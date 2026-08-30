package core

import "time"

// defaultSeed 提供一个基于当前时间的随机种子，供未显式指定 Rand 时使用。
func defaultSeed() int64 {
	return time.Now().UnixNano()
}
