package svc

import (
	"fmt"

	"github.com/gtwn/bitcoinnotify/pkg/model"
	resty "github.com/go-resty/resty/v2"
)

func GetTicker24hr(APITDAX string) (*model.Ticker,*model.Ticker,*model.Ticker, error) {

	client := resty.New()
	var xrp model.Ticker
	var btc model.Ticker
	var jfin model.Ticker
	url := fmt.Sprintf("%s/ticker/24hr",APITDAX)
	if _,err := client.R().SetResult(&xrp).Get(url+"?symbol=xrp_thb"); err != nil {
		return nil,nil,nil,err
	}

	if _,err := client.R().SetResult(&btc).Get(url+"?symbol=btc_thb"); err != nil {
		return nil,nil,nil,err
	}

	if _,err := client.R().SetResult(&jfin).Get(url+"?symbol=jfin_thb"); err != nil {
		return nil,nil,nil,err
	}

	return &xrp,&btc,&jfin, nil
}