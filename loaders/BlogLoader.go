package loaders

import (
	boxs "GoBlog/boxs"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsnMysql string = "mybatis:MyBatis@55998@tcp(127.0.0.1:3306)/" +
	"blog?charset=utf8mb4&parseTime=True&loc=Local"

/*
bug1 : 当查询操作没有找到结果的时候，仍然可以返回一个空数据。
	解决方式：与空数据进行比较，返回是否空数据
*/
/*
	通过 Blog 的 id 查找
*/
func LoadBlogByID(id uint64) (*boxs.Blog, bool) {
	dsn := dsnMysql
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接失败！！")
		fmt.Println(err.Error())
	}

	var blog boxs.Blog
	db.Table("t_blog").Find(&blog, id)
	return &blog, !(boxs.Blog{} == blog)
}

/*
	遍历所有 Blog 表中的 blog
*/
func LoadAllBlogs() (*[]boxs.Blog, bool) {
	dsn := dsnMysql
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接失败！！")
		fmt.Println(err.Error())
	}

	var blogs []boxs.Blog
	db.Table("t_blog").Find(&blogs)
	return &blogs, len(blogs) > 0
}

/*
	获取博客数量
*/
func GetNums() int64 {
	dsn := dsnMysql
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接失败！！")
		fmt.Println(err.Error())
	}

	var nums int64
	db.Table("t_blog").Count(&nums)
	return nums
}

/*
	通过年份获取相应的博客目录
*/
func LoadBlogsByYear(year int) (*[]boxs.Blog, bool) {
	dsn := dsnMysql
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接失败！！")
		fmt.Println(err.Error())
	}

	var blogs []boxs.Blog
	db.Table("t_blog").Find(&blogs, "DATE_FORMAT(t_blog.create_time,'%Y')=?", year)
	return &blogs, len(blogs) > 0
}

/*
	通过文章类型获取相应的博客目录
*/
func LoadBlogsByType(bType int) (*[]boxs.Blog, bool) {
	dsn := dsnMysql
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接失败！！")
		fmt.Println(err.Error())
	}

	var blogs []boxs.Blog
	db.Table("t_blog").Find(&blogs, "type_id=?", bType)
	return &blogs, len(blogs) > 0
}

/*
	通过组合条件对文章进行搜索，其中条件并不是必须的，如果条件都是为空，
	那么就显示所有文章。
*/
func LoadBlogsByConditions(bType *int, title *string, recommend *int) (*[]boxs.Blog, bool) {
	dsn := dsnMysql
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接失败！！")
		fmt.Println(err.Error())
	}

	conditions := make(map[string](interface{}))
	if bType != nil {
		conditions["type_id"] = *bType
	}
	if recommend != nil {
		conditions["recommend"] = *recommend
	}
	var blogs []boxs.Blog
	if title != nil {
		*title = "%" + *title + "%"
		db.Table("t_blog").Where(conditions).Find(&blogs, "title like ?", title)
	} else {
		db.Table("t_blog").Where(conditions).Find(&blogs)
	}
	return &blogs, len(blogs) > 0
}

/*
	通过文章 ID 删除文章
*/
func DelByID(id uint64) bool {
	dsn := dsnMysql
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接失败！！")
		fmt.Println(err.Error())
	}

	if err = db.Table("t_blog").Delete(boxs.Blog{}, id).Error; err != nil {
		fmt.Println("删除失败！！")
		fmt.Println(err.Error())
	}
	return true
}
