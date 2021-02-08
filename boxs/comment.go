package boxs

import (
	"time"
)

type Comment struct {
	ID              uint64
	NickName        string
	Email           string
	CreatedAt       time.Time
	BlogID          uint64
	ParentCommentID uint64
	RootComment     uint64
}
