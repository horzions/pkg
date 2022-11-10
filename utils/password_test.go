package utils_test

import (
	"testing"

	"github.com/horzions/pkg/utils"
)

func TestPassword(t *testing.T) {
	test, err := utils.Password("password")
	if err != nil {
		t.Errorf("HashPassword  err：%v", err)
	}
	t.Logf("输出哈希值为：%s", test)
}

func TestCheckPassword(t *testing.T) {
	bools := utils.CheckPassword("password", `$2a$14$QYRa2RMp7gn.1QwlxGXKJu6brn.LQYxOm4foTWhCeH8REr5PO81YO`)
	t.Logf("aaa 是否哈希：%v", bools)
}
