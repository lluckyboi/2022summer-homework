package cron

import "time"

type ComCycle struct {
	Minute, Hour, Dom, Month, Dow uint64
	Location                      *time.Location
}

// Next 实现Cycle接口
//func (s *ComCycle) Next(t time.Time) time.Time {
//
//}
