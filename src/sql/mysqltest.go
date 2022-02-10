package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	open, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test")
	if err != nil {
		log.Println(err)
		return
	}

	defer func(open *sql.DB) {
		_ = open.Close()
	}(open)
	if err := open.Ping(); err != nil {
		panic(err)
	}

	insertRecord(open)

	deleteRecord(open)

	updateRecord(open)

	selectRecord(open)

}

func insertRecord(open *sql.DB) {
	begin, _ := open.Begin()
	prepare, err := begin.Prepare("INSERT INTO `emp` (`EMPNO`, `ENAME`, `JOB`, `MGR`, `HIREDATE`, `SAL`, `COMM`, `DEPTNO`) " +
		"VALUES" +
		"(?, ?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	exec, err := prepare.Exec(7934, "MILLER", "CLERK", 7782, "1982-01-23", 1300.00, nil, 10)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	_ = begin.Commit()
	affected, _ := exec.RowsAffected()
	fmt.Println(affected)
}

func selectRecord(open *sql.DB) {
	var name string
	var nameArr = make([]string, 0)
	query, err := open.Query("select ENAME from emp")
	for query.Next() {
		err = query.Scan(&name)
		if err != nil {
			panic(err)
		}
		nameArr = append(nameArr, name)
	}
	fmt.Println(nameArr)
}

func updateRecord(open *sql.DB) {
	begin, err := open.Begin()
	exec, err := begin.Prepare("update emp set sal=900 where empno=?")
	if err != nil {
		panic(err)
	}
	if _, err := exec.Exec("7369"); err != nil {
		panic(err)
	}
	if err := begin.Commit(); err != nil {
		panic(err)
	}
}

func deleteRecord(open *sql.DB) {
	begin, err := open.Begin()
	if err != nil {
		fmt.Println(err)
	}
	exec, err := begin.Prepare("delete from emp where empno=?")
	if err != nil {
		return
	}
	result, err := exec.Exec(7934)
	if err != nil {
		return
	}
	affected, _ := result.RowsAffected()
	fmt.Println(affected)
	if err := begin.Commit(); err != nil {
		return
	}
}
