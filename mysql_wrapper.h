/*
 * mysql_wrapper.h
 */
#ifndef __MYSQL_WRAPPER_H__
#define __MYSQL_WRAPPER_H__
 
// mysql_real_connect wrapper
void *MySqlConnect(const char *host, const char *user, 
    const char *passwd, const char *db, unsigned int port, 
    const char *unix_socket, unsigned long client_flag);

// mysql_close wrapper
void MySqlClose(void *conn);

// mysql_error wrapper
const char *MySqlError(void *conn);

// mysql_errno wrapper
int MySqlErrno(void *conn);
 
#endif /* __MYSQL_WRAPPER_H__ */
 
 
