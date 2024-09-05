package exchange

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"log"
	"strconv"
	"strings"
	"time"
)

type BinancePairInfo struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

//type BinanceExcRates struct {}

var binanceAllowPairs = []string{
	"BTCRUB",
	"ETHRUB",
	"XRPRUB",
	"BNBRUB",
	"BUSDRUB",
	"USDTRUB",
	"LTCRUB",
	"ADARUB",
	"DOGERUB",
	"SHIBRUB",
	"MATICRUB",
	"DOTRUB",
	"SOLRUB",
	"ICPRUB",
	"TRURUB",
	"WAVESRUB",
	"ARPARUB",
	"FTMRUB",
	"NURUB",
	"ALGORUB",
	"NEORUB",
	"NEARRUB",
	"ARBRUB",
	"ARKMRUB",
	"WLDRUB",
}

//https://api.binance.com/api/v3/ticker/24hr

func GetBinanceTickByRub() ExchangeRates {
	var binanceBookList []BinancePairInfo
	var excRates ExchangeRates

	valute := make(map[string]Valute)

	client := resty.New()
	//
	resp, err := client.R().
		//EnableTrace().
		//SetResult(&htxList).
		Get("https://api.binance.com/api/v3/ticker/price")
	if err != nil {
		log.Printf("Error Binance parse, resp reason: %v", err)
	}

	if err := json.Unmarshal(resp.Body(), &binanceBookList); err != nil {
		log.Printf("Error Binance parse, json reason: %v", err)
	}

	//fmt.Println(&binanceBookList)

	for _, info := range binanceBookList {

		//log.Println(info.Symbol)

		if info.Symbol != "" {
			for _, pair := range binanceAllowPairs {
				if info.Symbol == pair {
					excRates.Date = time.Now()
					cryptoCharSymbol := strings.TrimSuffix(info.Symbol, "RUB")
					price, err := strconv.ParseFloat(info.Price, 64)
					if err != nil {
						log.Println(err)
					}
					//
					//fmt.Println(cryptoCharSymbol, info)
					valute[cryptoCharSymbol] = Valute{
						ID:       "",
						CharCode: cryptoCharSymbol,
						Value:    price,
						Name:     info.Symbol,
						Nominal:  1,
						Previous: price,
						Crypto:   true,
					}
					//fmt.Println(excRates.Valute)

				}
			}
		}

	}

	excRates.Valute = valute

	return excRates
}
