package main

import (
    "mysql";
    "log";
)

func main() {
    log.Stdoutf("this is mysql connect test\n");

    conn, err := mysql.Connect("localhost", "baduser", "badpasswd", "baddb", 0, "", 0); 
    if (err != nil) {
        log.Exitf("connect faled! errno:[%s] error:[%s]\n", err["errno"], err["error"]);
    }
    log.Stdoutf("connect success!\n");

    mysql.Close(conn);
}
