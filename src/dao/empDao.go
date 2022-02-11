package dao

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type EmpDao interface {
	Insert(emp Emp)
	delete(emp Emp)
	update(emp Emp)
	One(empNo string) Emp
}

type EmpDaoImpl struct {
	Db *sqlx.DB
}

func (empDaoImpl EmpDaoImpl) Insert(emp Emp) {
	db := empDaoImpl.Db
	begin, _ := db.Begin()
	exec, err := begin.Exec("INSERT INTO `emp` (`EMPNO`, `ENAME`, `JOB`, `MGR`, `HIREDATE`, `SAL`, `COMM`, `DEPTNO`) "+
		"VALUES"+
		"(?, ?, ?, ?, ?, ?, ?, ?);",
		emp.Empno, emp.Ename, emp.Job, emp.Mgr, "2006-01-02", emp.Sal, emp.Comm, emp.DeptNo)
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
	err := empDaoImpl.Db.Get(&emp, "select * from emp where empno=?", empNo)
	if err != nil {
		log.Println(err)
	}
	return emp
}

type Emp struct {
	Empno    int     `uri:"empno" db:"EMPNO"`
	Ename    string  `db:"ENAME"`
	Job      string  `db:"JOB"`
	Mgr      int     `db:"MGR"`
	HireDate string  `db:"HIREDATE"`
	Sal      float64 `db:"SAL"`
	Comm     float64 `db:"COMM"`
	DeptNo   string  `db:"DEPTNO"`
}
