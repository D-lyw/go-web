package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go-web/controllers"
	"go-web/lib"
	"go-web/model"
)

func handleDB() {
	db, err := sql.Open("mysql", "root:20212021@/go_web")
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT user SET name=?,email=?,country=?,phone=?,address=?")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec("zhangsan", "dsfas@mail.com", "china", "24304989", "beijing china")
	if err != nil {
		panic(err)
	}
}

// More usage detail about gin https://gin-gonic.com/zh-cn/docs/examples/

func SetAndConfigRouter() *gin.Engine {
	//	init http server object
	r := gin.Default()

	// 注册自定义中间件
	r.Use(lib.MyLogger())

	// static file map
	r.Static("/static", "./resource")

	// handle request
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "OK")
	})

	r.GET("/userInfo", controllers.GetUserInfo)
	r.POST("/addUser", controllers.AddUserInfo)
	r.DELETE("/userInfo", controllers.DeleteUser)

	r.GET("/hotList", controllers.GetHotList)

	// 解析 URL 参数
	r.GET("/article", controllers.GetArticleById)

	// 解析路由参数
	r.GET("/article/:id", controllers.GetArticleInfoById)

	r.POST("/article", controllers.PostArticle)
	r.DELETE("/article/:id", controllers.DeleteArticle)

	return r
}

func main() {

	db := model.GetDB()
	fmt.Print(db)

	//handleDB()

	r := SetAndConfigRouter()

	// Listen Http Server at 0.0.0.0:8080
	r.Run()
}
