package loaders

import (
	. "GoBlog/boxs"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LoadOneBlogByID(id uint64) *Blog {
	dsn := "mybatis:MyBatis@55998@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接失败！！")
		fmt.Println(err.Error())
	}

	var blog Blog
	db.Table("t_blog").Find(&blog, "1")
	return &blog
}
