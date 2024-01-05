package dao

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mySQLClient = &gorm.DB{}

var mySQLConnPool = &sql.DB{}

func init() {
	dsn := "root:123456@tcp(localhost:3306)/jerry_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	mySQLConnPool, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// 设置连接池的大小
	mySQLConnPool.SetMaxIdleConns(10)
	mySQLConnPool.SetMaxOpenConns(100)
	mySQLConnPool.SetConnMaxLifetime(time.Hour)
}

func InitMySQL() {
	dsn := "root:123456@tcp(localhost:3306)/jerry_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	mySQLClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		ConnPool: mySQLConnPool,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = mySQLClient.Exec("CREATE DATABASE IF NOT EXISTS jerry_db").Error
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MySQL连接成功！")

	err = mySQLClient.AutoMigrate(&Person{})
	if err != nil {
		log.Fatal(err)
	}
}

type Person struct {
	gorm.Model // 加入这个可以自动添加updated_at和created_at字段，用于记录修改和创建时间
	Name       string
	Age        int
}

func GetFromMySQL(ctx *gin.Context) {
	var persons []Person

	result := mySQLClient.Where("age=?", 20).Find(&persons)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "MySQL 查询成功",
		"data":    persons,
	})
}

func WriteToMySQL(ctx *gin.Context) {
	person := Person{
		Name: "username",
		Age:  20,
	}

	err := mySQLClient.Create(&person).Error
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "MySQL 写入成功",
	})
}
