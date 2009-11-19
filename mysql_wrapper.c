#include <stdio.h>
#include <string.h>
#include <mysql.h>
#include "mysql_wrapper.h"

void *MySqlConnect(const char *host, const char *user, 
    const char *passwd, const char *db, unsigned int port, 
    const char *unix_socket, unsigned long client_flag)
{

    MYSQL *conn = mysql_init(NULL);
    
    mysql_real_connect(conn, strlen(host) == 0 ? NULL : host, 
        strlen(user) == 0 ? NULL : user, 
        strlen(passwd) == 0 ? NULL : passwd, 
        strlen(db) == 0 ? NULL : db, port, 
        strlen(unix_socket) == 0 ? NULL : unix_socket, client_flag);

    return conn;
}

void MySqlClose(void *conn)
{
    mysql_close(conn);
}

int MySqlErrno(void *conn)
{
    return mysql_errno(conn);
}

const char *MySqlError(void *conn)
{
    return mysql_error(conn);
}
