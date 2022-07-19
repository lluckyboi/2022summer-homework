package main

import "2022summer-homework/day03/cron"

func main() {
	c := cron.NewCron()
	c.AddEntry("*/5 * * * * * *", func() {
		println("5s ")
	})

	c.Start()
}
