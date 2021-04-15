package svc

import (
	"fmt"

	resty "github.com/go-resty/resty/v2"
	"github.com/gtwn/bitcoinnotify/pkg/model"
)


func LineNotify(chs string,api string,xrp *model.Ticker, btc *model.Ticker, jfin *model.Ticker, xzc *model.Ticker) error {
	
	message := fmt.Sprintf("\nlast XRP: %s\nhigh XRP: %s\nlow XRP: %s\n------------------\nlast BTC: %s\nhigh BTC: %s\nlow BTC %s\n------------------\nlast JFIN: %s\nhigh JFIN: %s\nlow JFIN: %s\n------------------\nlast XZC: %s\nhigh XZC: %s\nlow XZC: %s",
	xrp.LastPrice,xrp.HighPrice,xrp.LowPrice,btc.LastPrice,btc.HighPrice,btc.LowPrice,jfin.LastPrice,jfin.HighPrice,jfin.LowPrice,xzc.LastPrice,xzc.HighPrice,xzc.LowPrice)

	client := resty.New()
	auth := fmt.Sprintf("Bearer %s",chs)
	if _,err := client.R().
	SetHeaders(map[string]string{
		"Content-Type" :"application/x-www-form-urlencoded", 
		"Authorization": auth,
	}).
	SetFormData(map[string]string{
		"message": message,
	}).Post(api) ; err != nil {
		return err
	}
	
	return nil

}