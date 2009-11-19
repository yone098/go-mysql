// yone098
package mysql

// #include "mysql_wrapper.h"
import "C"

import (
    "unsafe";
    "fmt";
)

func Connect(host string, user string, passwd string, db string, port uint8, unix_socket string, client_flag uint32) (unsafe.Pointer, map[string] string) {

    conn := C.MySqlConnect(C.CString(host), C.CString(user), 
        C.CString(passwd), C.CString(db), C.uint(port), 
        C.CString(unix_socket), C.ulong(client_flag));
    errno := C.MySqlErrno(conn);
    if (errno == 0) {
        return conn, nil;
    }

    return conn, map[string] string {
        "errno": fmt.Sprintf("%d", errno),
        "error": C.GoString(C.MySqlError(conn)),
    };
}


func Close(conn unsafe.Pointer) {
    C.MySqlClose(conn);
}
