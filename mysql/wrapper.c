// license that can be found in the LICENSE file.

// See the README for more information.

#include "wrapper.h"
#include <stdio.h>
#include <mysql.h>

wmysql wm_init(wmysql mysql) {
    return (wmysql)mysql_init((MYSQL *)mysql);
}


int wm_errno(wmysql mysql) {
    return mysql_errno((MYSQL *)mysql);
}

const char *wm_error(wmysql mysql) {
    return mysql_error((MYSQL *)mysql);
}

wmysql wm_real_connect(wmysql mysql, const char *host, const char *user,
    const char* passwd, const char *db, int port) {
    return (wmysql)mysql_real_connect((MYSQL *)mysql, host, user, passwd, db,
            port, NULL, 0);
}

void wm_close(wmysql mysql) {
    mysql_close((MYSQL *)mysql);
}

unsigned long wm_get_client_version(void) {
    return mysql_get_client_version();
}

void wm_free_result(wres res) {
    mysql_free_result((MYSQL_RES *)res);
}

int wms_query(wmysql mysql, const char *q) {
    return mysql_query((MYSQL *)mysql, q);
}

wres wm_store_result(wmysql mysql) {
    return (wres)mysql_store_result((MYSQL *)mysql);
}

char *wm_row(wrow row, int i) {
    return (char *)((MYSQL_ROW)row)[i];
}

const char *wfield_name_at(wfield field, int i) {
    return ((MYSQL_FIELD *)field)[i].name;
}

int wfield_type_at(wfield field, int i) {
    return ((MYSQL_FIELD *)field)[i].type;
}

int wfield_count(wmysql mysql) {
    return mysql_field_count((MYSQL *)mysql);
}

int wm_num_fields(wres res) {
    return mysql_num_fields((MYSQL_RES *)res);
}

wfield wm_fetch_fields(wres res) {
    return (wfield)mysql_fetch_fields((MYSQL_RES *)res);
}

wrow wm_fetch_row(wres res) {
    return mysql_fetch_row((MYSQL_RES *)res);
}

unsigned long long wm_num_rows(wres res) {
    return mysql_num_rows((MYSQL_RES *)res);
}

void wm_thread_init(void) {
    (void)mysql_thread_init();
}

void wm_thread_end(void) {
    mysql_thread_end();
}
