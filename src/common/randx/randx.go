package randx

import (
	"math/rand"
	"strconv"
	"time"
)

// 生成自定义订单号 timestamp + 类用户ID + 随机数（可选）
func MustGenOrderNo(uid int64) string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(int(time.Now().Unix())) + strconv.Itoa(int(uid)) + strconv.Itoa(rand.Intn(1000))
}
