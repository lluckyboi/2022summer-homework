package cron

import (
	"github.com/gorhill/cronexpr"
)

// Parse 解析 返回工作周期和err
// 尝试正则 没写出来
func (c *MyCron) Parse(spec string) (Cycle, error) {
	return cronexpr.Parse(spec)
}
