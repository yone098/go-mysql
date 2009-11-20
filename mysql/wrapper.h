#ifndef __MYSQL_WRAPPER_H__
#define __MYSQL_WRAPPER_H__

/*
    wrapper for mysql
*/

/* really mysql */
typedef void *wm_mysql;

typedef void *wm_res;

wm_mysql wm_init(wm_mysql mysql);

const char *wm_error(wm_mysql mysql);

wm_mysql wm_real_connect(wm_mysql mysql, const char *host, const char *user,
    const char* passwd, const char *db, int port);

void wm_close(wm_mysql mysql);

unsigned long wm_get_client_version();

int wm_errno(wm_mysql mysql);

const char *wm_error(wm_mysql mysql);
 
 
#endif /* __MYSQL_WRAPPER_H__ */
