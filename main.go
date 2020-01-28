package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DataSet = []map[string]string

func GetDataSet(rows *sql.Rows) DataSet {
	cols, _ := rows.Columns()
	colsize := len(cols)
	dataset := DataSet{}

	for rows.Next() {
		colmap := make(map[string]string)
		// colmap := make(map[string]string)
		coldata := make([]interface{}, colsize)

		for i := 0; i < colsize; i++ {
			coldata[i] = new(interface{})
		}
		rows.Scan(coldata...)

		for i, m := range cols {
			v := coldata[i].(*interface{})
			if *v == nil {
				colmap[m] = ""
			} else {
				colmap[m] = string((*v).([]byte))
			}
			// colmap[m] = fmt.Sprintf("%s", *v)
		}
		dataset = append(dataset, colmap)
	}

	return dataset
}

type flags struct {
	Driver *string
	Source *string
	Query  *string
}

func getArgs() (flags, bool) {
	args := flags{}

	args.Driver = flag.String("d", "", "driver name  (e.g mysql)")
	args.Source = flag.String("s", "", "source (e.g user:passwd@tcp(host:3306)/database )")
	args.Query = flag.String("q", "", "query ")
	flag.Parse()

	isFlagPassed := func(name string) bool {
		found := false
		flag.Visit(func(f *flag.Flag) {
			if f.Name == name {
				found = true
			}
		})
		return found
	}

	found := isFlagPassed("d")
	found = found && isFlagPassed("s")
	found = found && isFlagPassed("q")

	if !found {
		flag.Usage()
	}
	return args, found
}

func main() {
	args, found := getArgs()
	if !found {
		return
	}

	db, err := sql.Open(*args.Driver, *args.Source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(*args.Query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	ds := GetDataSet(rows)

	for _, r := range ds {
		bs, _ := json.Marshal(r)
		fmt.Println(string(bs))
	}
}
