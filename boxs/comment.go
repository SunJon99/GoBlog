package boxs

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID              uint64
	NickName        string
	Email           string
	CreatedAt       time.Time
	BlogID          uint64
	ParentCommentID uint64
	RootComment     uint64
}
