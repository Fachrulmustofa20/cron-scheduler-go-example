package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Fachrulmustofa20/cron-scheduler-go/jobs"
	cron "github.com/robfig/cron/v3"
)

func main() {
	// set scheduler with timezone
	timeZone, _ := time.LoadLocation("Asia/Jakarta")
	schedule := cron.New(cron.WithLocation(timeZone))
	defer schedule.Stop()

	// task scheduler
	schedule.AddFunc("* * * * *", jobs.SendNotifHi)
	schedule.AddFunc("*/5 * * * *", jobs.SendEveryFiveMinutes)
	schedule.AddFunc("*/2 * * * *", jobs.SendEveryTwoMinutes)

	go schedule.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
