package util

import (
	"context"
	"math/rand"
)

// 获取单个随机数
func GetRandomNum(ctx context.Context, nums []int64) int64 {
	if len(nums) <= 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return nums[rand.Intn(len(nums)-1)]
}
