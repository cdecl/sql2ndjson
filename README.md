
## sql2ndjson
sql(db) to ndjson format output

### Build 

```
make build
```

### Example 

```
$ bin\sql2ndjson.exe
Usage of bin\sql2ndjson.exe:
  -d string
        driver name  (e.g mysql)
  -q string
        query
  -s string
        source (e.g user:passwd@tcp(host:3306)/database )

```

```
sql2ndjson -d "mysql" -s "root:passwd@tcp(host:3306)/database" -q "select * from tablename"
```
