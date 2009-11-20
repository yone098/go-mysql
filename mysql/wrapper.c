// license that can be found in the LICENSE file.

// See the README for more information.

#include "wrapper.h"
#include <stdio.h>
#include <mysql.h>

wm_mysql wm_init(wm_mysql mysql) {
    return (wm_mysql)mysql_init((MYSQL *)mysql);
}


int wm_errno(wm_mysql mysql) {
    return mysql_errno((MYSQL *)mysql);
}

const char *wm_error(wm_mysql mysql) {
    return mysql_error((MYSQL *)mysql);
}

wm_mysql wm_real_connect(wm_mysql mysql, const char *host, const char *user,
    const char* passwd, const char *db, int port) {
    return (wm_mysql)mysql_real_connect((MYSQL *)mysql, host, user, passwd, db,
            port, NULL, 0);
}

void wm_close(wm_mysql mysql) {
    mysql_close((MYSQL *)mysql);
}

unsigned long wm_get_client_version() {
    return mysql_get_client_version();
}
