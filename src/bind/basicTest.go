package main

import (
	"fmt"
	"go-web/src/dao"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	engine := gin.Default()
	engine.GET("/:empno", func(context *gin.Context) {
		context.Next()
		var emp dao.Emp
		err := context.ShouldBindUri(&emp)
		log.Println(emp.Empno)
		if err != nil {
			fmt.Println(err)
			return
		}
		connect, err := sqlx.Connect("mysql", "root:@tcp(localhost:3306)/test")
		if err != nil {
			log.Println(err)
			return
		}
		if err != nil {
			log.Println(err)
			return
		}

		daoImpl := dao.EmpDaoImpl{Db: connect}
		emp = daoImpl.One(emp.Empno)

		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": emp,
		})
	})
	err := engine.Run()
	if err != nil {
		return
	}
}
