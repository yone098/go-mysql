package main

import (
    //"db/mysql";
    "mysql";
    "fmt";
    "os";
    "rand";
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
        "port": 3306,
        "uname": "yone098",
        "pass": "yone098",
        "dbname": "golang"
     };
    c, e := mysql.Open(info);
    if e != nil {
        fmt.Printf("open error: %s\n", e.String());
        os.Exit(1);
    }
    fmt.Printf("connection: %s\n", c);
    fmt.Printf("About to prepare statement\n");
    s, e := c.Prepare("CREATE TEMPORARY TABLE __hello (i INT)");
    if e != nil {
        fmt.Printf("error: %s\n", e.String());
        os.Exit(1);
    }
    fmt.Printf("statement: %s\n", s);

    fmt.Printf("About to execute query\n");
    cur, e := c.Execute(s);
    if e != nil {
        fmt.Printf("error: %s\n", e.String());
        os.Exit(1);
    }
    fmt.Printf("corsor: %s\n", cur);

    fmt.Printf("Inserting 30 random ints");
    for i := 0; i < 30; i+=1 {
        s, e = c.Prepare("INSERT INTO __hello (i) values(%d)");
        cur, e = c.Execute(s, rand.Int());
        if e != nil {
            fmt.Printf("insert error: %s\n", e.String());
            os.Exit(1);
        }
    }

    s, e = c.Prepare("SELECT * FROM __hello ORDER BY i ASC");
    if e != nil {
        fmt.Printf("error: %s\n", e.String());
        os.Exit(1);
    }
    fmt.Printf("statement: %s\n", s);

    fmt.Printf("About to execute query\n");
    cur, e = c.Execute(s);
    if e != nil {
        fmt.Printf("error: %s\n", e.String());
        os.Exit(1);
    }
    fmt.Printf("corsor: %s\n", cur);

    fmt.Printf("About to fech one row\n");
    tuple, e := cur.FetchOne();
    for i := 0; tuple != nil; tuple, e = cur.FetchOne() {
        i++;
        fmt.Printf("row[%d]: %s\n", i, tuple[0]);
    }
    cur.Close();

    fmt.Printf("=====\n\nAbout to fetch many row\n");
    s, e = c.Prepare("SELECT * FROM __hello ORDER BY i ASC");
    cur, e = c.Execute(s);
    rows, e := cur.FetchMany(10);
    fmt.Printf("%s\n", rows);
    for _, y := range rows {
        fmt.Printf("data: %s\n", y);
    }
    cur.Close();

    fmt.Printf("=====\n\nAbout to fetch all row\n");
    s, e = c.Prepare("SELECT * FROM __hello ORDER BY i ASC");
    cur, e = c.Execute(s);
    rows, e = cur.FetchAll();
    fmt.Printf("%s\n", rows);
    for _, y := range rows {
        fmt.Printf("data: %s\n", y);
    }
    cur.Close();

    e = c.Close();
    if e != nil {
        fmt.Printf("close error: %s\n", e.String());
        os.Exit(1);
    }
}
