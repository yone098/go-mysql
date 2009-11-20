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

unsigned long wm_get_client_version(void) {
    return mysql_get_client_version();
}

void mw_free_result(wm_res res) {
    mysql_free_result((MYSQL_RES *)res);
}

int mw_query(wm_mysql mysql, const char *q) {
    return mysql_query((MYSQL *)mysql, q);
}

wm_res mw_store_result(wm_mysql mysql) {
    return (wm_res)mysql_store_result((MYSQL *)mysql);
}

char *mw_row(wm_row row, int i) {
    return (char *)((MYSQL_ROW)row)[i];
}

const char *mw_field_name_at(wm_field field, int i) {
    return ((MYSQL_FIELD *)field)[i].name;
}

int mw_field_type_at(wm_field field, int i) {
    return ((MYSQL_FIELD *)field)[i].type;
}

int mw_field_count(wm_mysql mysql) {
    return mysql_field_count((MYSQL *)mysql);
}

int mw_num_fields(wm_res res) {
    return mysql_num_fields((MYSQL_RES *)res);
}

wm_field mw_fetch_fields(wm_res res) {
    return (wm_field)mysql_fetch_fields((MYSQL_RES *)res);
}

wm_row mw_fetch_row(wm_res res) {
    return mysql_fetch_row((MYSQL_RES *)res);
}

unsigned long long mw_num_rows(wm_res res) {
    return mysql_num_rows((MYSQL_RES *)res);
}

void mw_thread_init(void) {
    (void)mysql_thread_init();
}

void mw_thread_end(void) {
    mysql_thread_end();
}
