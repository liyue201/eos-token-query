package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/liyue201/go-logger"
	"fmt"
)


func getAccounts(url string) []string  {
	acounts := []string{}

	doc, err := goquery.NewDocument(url)
	if err != nil{
		logger.Errorf("[getTotkenAccounts] %s",err)
	}

	fc := doc.Find(".common-lsit-data_table_fc")
	fc.Each(func(i int, content *goquery.Selection) {
		//logger.Debugf("i=%d , fc=%s", i, content.Text())

		span := content.Find("span")

		span.Each(func(i int, content *goquery.Selection) {
			account := content.Text()
			//logger.Debugf("span=%s", account)
			acounts = append(acounts, account)
		})
	})
	return acounts
}

func getTokenAccounts(contract, symbol string, limit int) []string {
	page := (limit + 19)/ 20

	accounts := []string{}
	for i := 1; i <= page ; i++{
		url := fmt.Sprintf("https://eosmonitor.io/rank_coin?page=%d&coin=%s&name=%s", i, symbol, contract)
		acs := getAccounts(url)
		if len(acs) == 0 {
			break
		}
		accounts = append(accounts, acs...)
	}
	//logger.Debugf("accounts=", accounts)
	return accounts
}

func test()  {
	accounts := getTokenAccounts("eosbuttonbtn", "BTN", 100)
	logger.Debugf("accounts=", accounts)
}

func main()  {
	test()
}
