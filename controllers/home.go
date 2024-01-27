package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/model"
	"log"
	"net/http"
	"strconv"
	"time"
)

// gorm document: https://gorm.io/zh_CN/docs/index.html

func GetUserInfo(c *gin.Context) {
	db := model.GetDB()

	var userItem []model.User
	if err := db.Find(&userItem).Error; err != nil {
		fmt.Println("查询失败", err)
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
		return
	}
	fmt.Println(userItem)

	c.JSON(200, gin.H{"success": true, "data": userItem})
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

func UpdateUserEmail(c *gin.Context) {
	userId := c.Param("id")
	var newInfo model.PostUser
	if err := c.ShouldBindJSON(&newInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	u := model.User{}
	db := model.GetDB()
	db.Find(&u, userId)
	u.Email = newInfo.Email
	fmt.Println(newInfo.Email, newInfo.Name, newInfo.Address)

	if err := db.Save(&u).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": newInfo})
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

func GetArticleInfoById(c *gin.Context) {
	articleId := c.Param("id")

	articleIdNum, err := strconv.Atoi(articleId)
	if err != nil {
		c.JSON(400, gin.H{"success": false, "error": "error format article id"})
		return
	}

	var article model.Article

	db := model.GetDB()
	if result := db.Find(&article, "id = ?", articleIdNum); result.Error != nil {
		c.JSON(500, gin.H{"success": false, "error": result.Error.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "data": article})
}

func PostArticle(c *gin.Context) {
	article := model.Article{}
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"success": false, "error": err.Error()})
		return
	}
	fmt.Println(article)

	db := model.GetDB()
	result := db.Create(&article)
	if result.Error != nil {
		c.JSON(400, gin.H{"success": false, "error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": article})
}

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	db := model.GetDB()
	if result := db.Delete(&model.Article{}, id); result.Error != nil {
		c.JSON(500, gin.H{"success": false, "error": result.Error.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true})
}
