#ifndef __WRAPPER_H__
#define __WRAPPER_H__

/*
    wrapper for mysql
*/

/* really mysql */
typedef void *wmysql;
typedef void *wres;
typedef void *wrow;
typedef void *wfield;

wmysql wm_init(wmysql mysql);
const char *wm_error(wmysql mysql);
wmysql wm_real_connect(wmysql mysql, const char *host, const char *user,
    const char* passwd, const char *db, int port);
void wm_close(wmysql mysql);
unsigned long wm_get_client_version(void);
int wm_errno(wmysql mysql);
const char *wm_error(wmysql mysql);
void wm_free_result(wres res);
int wms_query(wmysql mysql, const char *q);
wres wm_store_result(wmysql mysql);
char *wm_row(wrow row, int i);
const char *wfield_name_at(wfield field, int i);
int wfield_type_at(wfield field, int i);
int wfield_count(wmysql mysql);
int wm_num_fields(wres res);
wfield wm_fetch_fields(wres res);
wrow wm_fetch_row(wres res);
unsigned long long wm_num_rows(wres res);

void wm_thread_init(void);
void wm_thread_end(void);
 
 
#endif /* __WRAPPER_H__ */
