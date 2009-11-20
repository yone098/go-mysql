#ifndef __MYSQL_WRAPPER_H__
#define __MYSQL_WRAPPER_H__

/*
    wrapper for mysql
*/

/* really mysql */
typedef void *wm_mysql;
typedef void *wm_res;
typedef void *wm_row;
typedef void *wm_field;

wm_mysql wm_init(wm_mysql mysql);
const char *wm_error(wm_mysql mysql);
wm_mysql wm_real_connect(wm_mysql mysql, const char *host, const char *user,
    const char* passwd, const char *db, int port);
void wm_close(wm_mysql mysql);
unsigned long wm_get_client_version(void);
int wm_errno(wm_mysql mysql);
const char *wm_error(wm_mysql mysql);
void mw_free_result(wm_res res);
int mw_query(wm_mysql mysql, const char *q);
wm_res mw_store_result(wm_mysql mysql);
char *mw_row(wm_row row, int i);
const char *mw_field_name_at(wm_field field, int i);
int mw_field_type_at(wm_field field, int i);
int mw_field_count(wm_mysql mysql);
int mw_num_fields(wm_res res);
wm_field mw_fetch_fields(wm_res res);
wm_row mw_fetch_row(wm_res res);
unsigned long long mw_num_rows(wm_res res);

void mw_thread_init(void);
void mw_thread_end(void);
 
 
#endif /* __MYSQL_WRAPPER_H__ */
