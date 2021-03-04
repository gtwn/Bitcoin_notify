package main

import (
	"time"

	"github.com/davecgh/go-spew/spew"
	config "github.com/gtwn/bitcoinnotify/cmd/server/Config"
	"github.com/gtwn/bitcoinnotify/pkg/svc"
	"github.com/sirupsen/logrus"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		logrus.WithError(err).Fatal("Read Config")
	}
	for {
		xrp,btc,jfin,err := svc.GetTicker24hr(cfg.BitAPI)
		if err != nil {
			logrus.WithError(err).Fatal("ticker fail")
		}
		spew.Dump(&xrp)
		spew.Dump(&btc)
		spew.Dump(&jfin)

		if err := svc.LineNotify(cfg.ChannelAccessToken, cfg.LineNoti ,xrp,btc,jfin) ; err != nil {
			logrus.WithError(err).Fatal("Notify Error")
		}
		
		time.Sleep(2 * time.Minute)
		
	}
}