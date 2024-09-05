package exchange

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Valute struct {
	ID       string  `json:"ID"`
	NumCode  string  `json:"NumCode"`
	CharCode string  `json:"CharCode"`
	Nominal  int     `json:"Nominal"`
	Name     string  `json:"Name"`
	Value    float64 `json:"Value"`
	Previous float64 `json:"Previous"`
	Crypto   bool    `json:"-"`
}

type ExchangeRates struct {
	Date         time.Time         `json:"Date"`
	PreviousDate time.Time         `json:"PreviousDate"`
	PreviousURL  string            `json:"PreviousURL"`
	Timestamp    time.Time         `json:"Timestamp"`
	Valute       map[string]Valute `json:"Valute"`
}

func GetCbrCurrencies() ExchangeRates {
	req, err := http.NewRequest("GET", "https://www.cbr-xml-daily.ru/daily_json.js", nil)
	if err != nil {
		// handle error
		log.Fatal(err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		log.Fatal(err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return ExchangeRates{}
	}

	fmt.Println("Status: " + resp.Status)
	fmt.Println("StatusCode: " + strconv.Itoa(resp.StatusCode))

	if resp.StatusCode == 304 {
		return ExchangeRates{}
	}

	var exchangeRates ExchangeRates
	err = json.Unmarshal(body, &exchangeRates)
	if err != nil {
		fmt.Println("Error:", err)
		return ExchangeRates{}
	}

	fmt.Println("Date:", exchangeRates.Date)
	fmt.Println("PreviousDate:", exchangeRates.PreviousDate)
	fmt.Println("PreviousURL:", exchangeRates.PreviousURL)
	fmt.Println("Timestamp:", exchangeRates.Timestamp)

	for _, valute := range exchangeRates.Valute {
		valute.Crypto = false
	}

	//for currencyCode, valute := range exchangeRates.Valute {
	//	fmt.Printf("Currency: %s, Name: %s, Value: %f, Previous: %f\n", currencyCode, valute.Name, valute.Value, valute.Previous)
	//}

	return exchangeRates
}
