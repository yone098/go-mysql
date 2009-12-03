/*

    mysql wrapper 
    This source code was made referring to Eden's and Peter's source code.
    Eden http://github.com/eden/mysqlgo
    Peter http://github.com/phf/go-sqlite

*/
package mysql

/*
#include <stdlib.h>
#include "wrapper.h"
*/
import "C"

import "fmt"
import "unsafe"
import "os"
import "db"
import "sync"

var MaxFetchCount = 65535

type Connection struct {
    /* pointer to struct mysql */
    handle C.wmysql;
    queryLock sync.Mutex;
}

/* MYSQL cursors, will be renamed/refactored soon */
type Cursor struct {
    /* statement we were created for */
    statement *Statement;
    /* connection we were created on */
    connection *Connection;
    /* the last query yielded results */
    result bool;
}

type Statement struct {
    /* pointer to struct mysql statement */
    handle C.wres;
    connection *Connection;
    query string;
    nfields int;
}

func init() {
    // void
}

/* idiom to ensure that signatures are exactly as specified in db */
var Version db.VersionSignature;
var Open db.OpenSignature;
func init() {
    Version = version;
    Open = open;
}

func version() (data map[string]string, error os.Error)
{
    data = make(map[string]string);

    //v := int(C.wm_get_client_version());
    v := 123;
    cv := fmt.Sprintf("%d", v);
    data["version"] = cv;

    // mysql source id ?
    data["mysql.sourceid"] = "";

    // versionnumber
    data["mysql.versionnumber"] = "";

    return data, nil;
}

type Any interface{};
type ConnectionInfo map[string] Any;

// Unwrap function which returns an unsafe.Pointer for the given argument.
// This is used because cgo translates the typedefs from the wrapper into
// unsafe pointers instead of _C_typedefname.
func use(h interface {}) (rval unsafe.Pointer) {
    switch ptr := h.(type) {
    case C.wmysql:      rval = unsafe.Pointer(ptr)
    case C.wrow:   rval = unsafe.Pointer(ptr)
    case C.wres:   rval = unsafe.Pointer(ptr)
    case C.wfield: rval = unsafe.Pointer(ptr)
    default:        panic("Tried to use() unknown type\n")
    }
    return;
}

func LastError(mysql C.wmysql) os.Error {
    if err := C.wm_error(use(mysql)); *err != 0 {
        return os.NewError(C.GoString(err));
    }
    return nil;
}

func open(info ConnectionInfo) (connection db.Connection, error os.Error)
{
    // parse and require check
    host, port, uname, pass, dbname, error := parseConnInfo(info);
    args := []*C.char{
        C.CString(host), C.CString(uname), C.CString(pass), C.CString(dbname)};

    conn := new (Connection);
    conn.handle = C.wm_init(nil);
    C.wm_real_connect(
        use(conn.handle),
        args[0],
        args[1],
        args[2],
        args[3],
        C.int(port));

    for i, _ := range args {
        C.free(unsafe.Pointer(args[i]));
    }

    if error = LastError(conn.handle); error != nil {
        conn.handle = nil;
    }

    connection = conn;

    return;
}

/* === Connection === */

/*
    Fill in a DatabaseError with information about
    the last error from MySQL.
*/
func (self *Connection) error() (error os.Error) {
    e := new(DatabaseError);
    e.basic = int(C.wm_errno(use(self.handle)));
    // mysql is not extended error code
    e.extended = 0;
    e.message = string(C.GoString(C.wm_error(use(self.handle))));
    return e;
}

/*
    Precompile query into Statement.
*/
func (self *Connection) Prepare(query string) (statement db.Statement, error os.Error)
{
    s := new(Statement);
    s.query = query;
    s.connection = self;

    statement = s;

    return;
}

/*
    Execute precompiled Statement with given parameters (if any).
*/
func (self *Connection) Execute(statement db.Statement, parameters ...) (cursor db.Cursor, error os.Error)
{
    self.lock();

    s, ok := statement.(*Statement);
    if !ok {
        error = &InterfaceError{"Execute: Not an mysql statement!"};
        return;
    }
    // TODO bind parameter
    query := fmt.Sprintf(s.query, parameters);
    q := C.CString(query);
    rcode := C.wms_query(use(self.handle), q);
    C.free(unsafe.Pointer(q));

    if error = LastError(self.handle); error != nil || rcode != 0 {
        if error == nil { 
            error = os.NewError("Query failed.") 
        }
        goto UnlockAndReturn;
    }
    s.nfields = int(C.wfield_count(use(self.handle)));
    s.handle = C.wm_store_result(use(self.handle));
    error = LastError(self.handle);
    if error != nil || (s.handle == nil && s.nfields > 0) {
        if error == nil {
            error = os.NewError("No results returned.");
            s.cleanup();
        }
        goto UnlockAndReturn;
    }
    c := new(Cursor);
    c.statement = s;
    c.connection = self;
    c.result = true;
    cursor = c;

UnlockAndReturn:
    self.unlock();
    return;
}


func (self *Connection) Close() (error os.Error) {
    C.wm_close(use(self.handle));
    self.handle = nil;
    return;
}

func (self *Connection) lock() { self.queryLock.Lock() }
func (self *Connection) unlock() { self.queryLock.Unlock() }

/* === Statement === */


func (self *Statement) cleanup() {
    if self.handle != nil {
        C.wm_free_result(use(self.handle));
        self.handle = nil;
        self.nfields = 0;
    }
}

/* === Cursor === */


func (self *Cursor) FetchOne() (data []interface {}, error os.Error)
{
    if !self.result {
        error = &InterfaceError{"FetchOne: No results to fetch!"};
        return;
    }
    
    row := C.wm_fetch_row(use(self.statement.handle));
    error = LastError(self.connection.handle);

    if row != nil && error == nil {
        data = make([]interface {}, self.statement.nfields);
        for i := 0; i < self.statement.nfields; i += 1 {
            data[i] = C.GoString(C.wm_row(use(row), C.int(i)));
        }
    }

    return;
}

func (self *Cursor) FetchMany(count int) (data [][]interface {}, error os.Error)
{
    if count == 0 { return nil, os.NewError("Invalid count") }

    data = make([][]interface {}, count);
    i := 0;
    row, err := self.FetchOne();
    for i < count && row != nil && err == nil {
        data[i] = row;
        i += 1;
        row, err = self.FetchOne();
    }
    if err != nil { data = nil }

    return;
}

func (self *Cursor) FetchAll() ([][]interface {}, os.Error)
{
    count := self.RowCount();
    if uint64(MaxFetchCount) <= count {
        return nil, os.NewError("Too many rows in result set. Use Fetch One or FetchMany insted");
    }
    return self.FetchMany(int(count));
}

// Returns the number of rows returned from the current result set.
func (self *Cursor) RowCount() uint64 {
    if !self.result { return 0 }
    return uint64(C.wm_num_rows(use(self.statement.handle)));
}

func (self *Cursor) Close() os.Error
{
    self.statement.cleanup();
    return nil;
}



func parseConnInfo(info ConnectionInfo) (host string, port int, uname string,
    pass string, dbname string, error os.Error)
{
    if host, error = getMapStringValue("host", info); error != nil {
        return;
    }
    if port, error = getMapIntValue("port", info); error != nil {
        return;
    }
    if uname, error = getMapStringValue("uname", info); error != nil {
        return;
    }
    if pass, error = getMapStringValue("pass", info); error != nil {
        return;
    }
    if dbname, error = getMapStringValue("dbname", info); error != nil {
        return;
    }
    
    return;
}

func getMapStringValue(key string, m map[string] Any) (val string, error os.Error)
{
    ok := false;
    any := Any(nil);

    any, ok = m[key];
    if !ok {
        error = &InterfaceError{
            fmt.Sprintf("Open: No \"%s\" in arguments map.", key)};
        return;
    }
    val, ok = any.(string);
    if !ok {
        error = &InterfaceError{
            fmt.Sprintf("Open: \"%s\" argument not a string.", key)};
        return;
    }
    
    return;
}

func getMapIntValue(key string, m map[string] Any) (val int, error os.Error)
{
    ok := false;
    any := Any(nil);

    any, ok = m[key];
    if !ok {
        error = &InterfaceError{
            fmt.Sprintf("Open: No \"%s\" in arguments map.", key)};
        return;
    }
    val, ok = any.(int);
    if !ok {
        error = &InterfaceError{
            fmt.Sprintf("Open: \"%s\" argument not a int.", key)};
        return;
    }
    
    return;
}
