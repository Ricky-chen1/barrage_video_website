package cache

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	RankDailyKey     = "rank:daily"
	UserBlacklistKey = "blacklist:user"
	SearchVideoKey   = "search:video"
)

func UserKey(uid uint) string {
	return fmt.Sprintf("user:%s", strconv.FormatUint(uint64(uid), 10))
}

func VideoViewKey(id uint) string {
	return fmt.Sprintf("view:video:%s", strconv.FormatUint(uint64(id), 10))
}

func CommentReplyKey(id uint) string {
	return fmt.Sprintf("reply:comment:%s", strconv.FormatUint(uint64(id), 10))
}

func TokenKey(uid uint) string {
	return strings.Join([]string{"token", ":", UserKey(uid)}, "")
}
