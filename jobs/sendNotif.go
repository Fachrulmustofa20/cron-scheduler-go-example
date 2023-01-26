package jobs

import (
	"fmt"
	"time"
)

func SendNotifHi() {
	fmt.Printf("[%s] Hi. This cron running every 1 minutes\n", time.Now().Format("2006-01-02 15:04:05"))
}

func SendEveryFiveMinutes() {
	fmt.Printf("[%s] Send every 5 minutes\n", time.Now().Format("2006-01-02 15:04:05"))
}

func SendEveryTwoMinutes() {
	fmt.Printf("[%s] Running every 2 minutes\n", time.Now().Format("2006-01-02 15:04:05"))
}
