/*

    mysql wrapper 

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

type Connection struct {
    /* pointer to struct mysql */
    handle C.wm_mysql;
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
    case C.wm_mysql:      rval = unsafe.Pointer(ptr)
    case C.wm_row:   rval = unsafe.Pointer(ptr)
    case C.wm_res:   rval = unsafe.Pointer(ptr)
    case C.wm_field: rval = unsafe.Pointer(ptr)
    default:        panic("Tried to use() unknown type\n")
    }
    return;
}

func LastError(mysql C.wm_mysql) os.Error {
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
    return nil, nil;
}

/*
    Execute precompiled Statement with given parameters (if any).
*/
func (self *Connection) Execute(statement db.Statement, parameters ...) (cursor db.Cursor, error os.Error)
{
    return nil, nil;
}

func (self *Connection) Close() (error os.Error) {
    C.wm_close(use(self.handle));
    self.handle = nil;
    return;
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
