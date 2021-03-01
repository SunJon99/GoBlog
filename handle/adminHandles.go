package handle

import (
	boxs "GoBlog/boxs"
	loaders "GoBlog/loaders"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func AdminHandle(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseFiles("templates/admin/bolgs.html"))

	if err := templates.ExecuteTemplate(w, "bolgs.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//@Description 处理页面的blogs请求，返回数据库中所有的blog信息
// 这里页面需要的信息只是 blog 中的 id，title，type，recommend，updatetime

//@TODO 将数据查询出来的blog数据进行封装，并转换为JSON的数据形式，返回给客户端
func GetBlogsHandle(w http.ResponseWriter, r *http.Request) {
	blogs, isLoad := loaders.LoadAllBlogs()
	if isLoad == false {
		log.Println("数据库数据加载出错！！")

		w.WriteHeader(505)
		fmt.Fprintln(w, "服务器内部出现错误！！")
		return
	}
	blogsInfo := make([]boxs.BlogInfo, len(*blogs))
	for i, blog := range *blogs {
		timeString := blog.UpdatedAt.Format("2006-01-02 15:04:05")
		blogsInfo[i] = boxs.BlogInfo{
			ID:         blog.ID,
			Title:      blog.Title,
			Type:       fmt.Sprint(blog.Type),
			Recommend:  fmt.Sprint(blog.Recommend),
			UpdateTime: timeString}
	}
	data, err := json.MarshalIndent(blogsInfo, "", "\t\t")
	if err != nil {
		log.Println("json转换出现问题！！")
		log.Println(err)

		w.WriteHeader(505)
		fmt.Fprintln(w, "服务器内部出现错误！！")
		return
	}

	jsonString := string(data)
	fmt.Println(jsonString)
	w.Write([]byte(jsonString))

}
