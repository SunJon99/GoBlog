package boxs

import (
	"time"
)

type Blog struct {
	ID           uint64
	Title        string    //博客文章标题
	Content      string    //博客内容
	FirstPicture string    //博客首图地址
	Flag         string    //转载，原创，翻译声明
	Views        uint64    //查看人数
	Appreciation uint64    //赞赏人数
	Share        uint8     //是否开启share 0、1
	Commentabled uint8     //是否开启评论
	Published    uint8     //是否发表
	Recommend    uint8     //是否推荐
	CreatedAt    time.Time `gorm:"column:create_time"` //创建时间
	UpdatedAt    time.Time `gorm:"column:update_time"` //更新时间
	Type         uint64    `gorm:"column:type_id"`     //文章类型
	Author       uint64    `gorm:"column:user_id"`     //文章作者
	Description  string    //文章描述
	Tags         string    `gorm:"column:tag_ids"` //文章标签的字符串表现形式
}
