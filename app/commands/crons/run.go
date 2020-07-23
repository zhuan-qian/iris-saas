package crons

import (
	"github.com/robfig/cron/v3"
)

func Run() {
	c := cron.New()
	defer c.Start()
	c.AddFunc("*/30 * * * *", EsLostBack)
}
