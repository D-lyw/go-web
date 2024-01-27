package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/model"
	"log"
	"net/http"
	"time"
)

func GetUserInfo(c *gin.Context) {
	db := model.GetDB()

	var userItem []model.User
	if err := db.Find(&userItem).Error; err != nil {
		fmt.Println("查询失败", err)
	}
	fmt.Println(userItem)

	c.JSON(200, userItem)
}

func AddUserInfo(c *gin.Context) {
	u := model.User{}
	if c.ShouldBind(&u) == nil {
		log.Println(u.Name, u.Email)
	}
	fmt.Println(u)

	// 存入数据库
	db := model.GetDB()
	res := db.Create(&u)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, u)
	}
	c.JSON(http.StatusOK, u)
}

func DeleteUser(c *gin.Context) {
	userId, _ := c.GetQuery("name")
	fmt.Println(userId)

	db := model.GetDB()
	if err := db.Where("name = ?", userId).Delete(&model.User{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		panic(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func GetHotList(c *gin.Context) {
	go func() {
		time.Sleep(5 * time.Second)
		log.Println("request: ", c.Request.URL.Path)
	}()
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func GetArticleById(c *gin.Context) {
	articleId, ok := c.GetQuery("id")
	if !ok {
		c.String(400, "No article id")
	}
	c.String(200, `%s`, articleId)
}

func GetBookInfoById(c *gin.Context) {
	bookId := c.Param("id")

	c.String(200, `书籍 ID: %s`, bookId)
}
