package dao

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type EmpDao interface {
	Insert(emp Emp)
	delete(emp Emp)
	update(emp Emp)
	One(empNo string) Emp
}

type EmpDaoImpl struct {
	Db *sql.DB
}

func (empDaoImpl EmpDaoImpl) Insert(emp Emp) {
	db := empDaoImpl.Db
	begin, _ := db.Begin()
	exec, err := begin.Exec("INSERT INTO `emp` (`EMPNO`, `ENAME`, `JOB`, `MGR`, `HIREDATE`, `SAL`, `COMM`, `DEPTNO`) "+
		"VALUES"+
		"(?, ?, ?, ?, ?, ?, ?, ?);",
		emp.Empno, emp.Ename, emp.Job, emp.Mgr, emp.HireDate.Format("2006-01-02"), emp.Sal, emp.Comm, emp.DeptNo)
	if err != nil {
		fmt.Println(err)
	}
	_ = begin.Commit()
	fmt.Println(exec.RowsAffected())
}
func (empDaoImpl EmpDaoImpl) delete(emp Emp) {

}
func (empDaoImpl EmpDaoImpl) update(emp Emp) {

}

func (empDaoImpl EmpDaoImpl) One(empNo int) Emp {
	var emp Emp
	query, err := empDaoImpl.Db.Prepare("select empno,ENAME,sal,HireDate,Comm,DeptNo from emp where empno=?")
	if err != nil {
		log.Println(err)
	}
	row := query.QueryRow(empNo)
	err = row.Scan(&emp.Empno, &emp.Ename, &emp.Sal, &emp.HireDate, &emp.Comm, &emp.DeptNo)
	if err != nil {
		log.Println(err)
	}
	return emp
}

type Emp struct {
	Empno    int `uri:"empno"`
	Ename    string
	Job      string
	Mgr      int8
	HireDate time.Time
	Sal      float64
	Comm     float64
	DeptNo   string
}
