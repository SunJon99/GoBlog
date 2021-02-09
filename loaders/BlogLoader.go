package loaders

import (
	. "GoBlog/boxs"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsm_mysql string = "mybatis:MyBatis@55998@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"

/*
	通过 Blog 的 id 查找
*/
func LoadOneBlogByID(id uint64) *Blog {
	dsn := dsm_mysql
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接失败！！")
		fmt.Println(err.Error())
	}

	var blog Blog
	db.Table("t_blog").Find(&blog, "1")
	return &blog
}

/*
	遍历所有 Blog 表中的 blog
*/
func LoadAllBlogs() *[]Blog {
	dsn := dsm_mysql
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接失败！！")
		fmt.Println(err.Error())
	}

	var blogs []Blog
	db.Table("t_blog").Find(&blogs)
	return &blogs
}
