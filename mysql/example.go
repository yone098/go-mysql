package main

import (
    "db/mysql";
    "fmt";
    "os";
)

func main() {
    version, e := mysql.Version();
    if e != nil {
        fmt.Printf("error: %s\n", e.String());
    }
    for k, v := range version {
        fmt.Printf("version[%s] == %s\n", k, v);
    }

    info := mysql.ConnectionInfo{ 
        "host": "localhost",
        "port": 0,
        "uname": "username",
        "pass": "password",
        "dbname": "userdbname"
     };
    c, e := mysql.Open(info);
    if e != nil {
        fmt.Printf("open error: %s\n", e.String());
        os.Exit(1);
    }
    fmt.Printf("connection: %s\n", c);

    e = c.Close();
    if e != nil {
        fmt.Printf("close error: %s\n", e.String());
        os.Exit(1);
    }
    
}
