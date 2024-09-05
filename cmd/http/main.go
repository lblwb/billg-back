package main

import (
	"backend/cmd/bot/telegram/alerter"
	"backend/internal/billing_app/http/bootstrap"
	"backend/internal/billing_app/models/currency"
	"backend/internal/billing_app/service/api/auth/jwt_auth"
	"backend/internal/billing_app/service/workers/exchange"
	"backend/internal/database"
	"backend/pkg"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type ExchangeData struct {
	LastUpdated time.Time              `json:"last_updated"`
	Data        exchange.ExchangeRates `json:"data"`
}

var jwtPair string

func dbConnect() *database.StorageDb {
	return database.
		NewStorageDb(os.Getenv("DB_DSN")).
		Connect()
}

func setNewJwtPair() {
	jwt_auth.JwtPairKey = jwt_auth.GenerateNewPairs()
}

func updateExchangeData() ExchangeData {
	// Получение данных с биржи (пример)
	exchanges := exchange.GetCbrCurrencies()
	// Создание структуры данных для сохранения в файл
	exchangeData := ExchangeData{
		LastUpdated: time.Now(),
		Data:        exchanges,
	}

	exchangesCrypto := exchange.GetBinanceTickByRub()
	//log.Println(exchangesCrypto.Valute)

	for s, value := range exchangesCrypto.Valute {
		exchangeData.Data.Valute[s] = value
		//fmt.Println(s, value)
	}

	// Кодирование данных в формат JSON
	jsonData, err := json.Marshal(exchangeData)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return ExchangeData{}
	}

	// Сохранение данных в файл
	err = ioutil.WriteFile("exchange_data.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return ExchangeData{}
	}

	fmt.Println("Exchange data updated successfully.")

	return exchangeData
}

func excRatesFormat(currencies ExchangeData) []currency.CurrRates {
	var curRates []currency.CurrRates

	for _, val := range currencies.Data.Valute {
		curRates = append(curRates, currency.CurrRates{
			Currency:    "RUB",
			DirCurrency: val.CharCode,
			Value:       val.Value,
			Crypto:      val.Crypto,
		})
	}
	return curRates
}

func loadExcRates(db *database.StorageDb) {
	curExcGet := updateExchangeData()
	excRates := excRatesFormat(curExcGet)
	//log.Printf("%+v", excRates)

	err := currency.NewRatesEntity(db).SaveExchangesRates(excRates)
	if err != nil {
		//return
		log.Fatalln(err)
	}

	// Создаем канал для управления горутиной
	done := make(chan struct{})
	defer close(done)

	// Запускаем асинхронно горутину для выполнения запроса каждые 30 минут
	go func() {
		ticker := time.NewTicker(30 * time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				curExcGet := updateExchangeData()
				excRates := excRatesFormat(curExcGet)
				//log.Printf("%+v", excRates)

				err := currency.NewRatesEntity(db).SaveExchangesRates(excRates)
				if err != nil {
					//return
					log.Fatalln(err)
				}
				//
				//TODO: Записать курс в бд
			}
		}
	}()
}

//go run ./cmd/bot/telegram/alerter/main.go

//fun eventBroadcast(){
//go func() {
//	for {
//
//		//userOrder := events.NewPushEventToBroadcast("newUserOrder", userOrder[0])
//
//		orderService1 := order.GetServiceOrdersByUserId(1)
//		orderService2 := order.GetServiceOrdersByUserId(1)
//		//orderService3 := order.GetServiceOrdersByUserId(1)
//
//		//events.NewPushArrEventToBroadcast([]events.Event{
//		//	{orderService1[0], "newRegisterUser"},
//		//	{orderService2[0], "newUserOrder"},
//		//	//{orderService3[0], "newUserOrder"},
//		//})
//
//		time.Sleep(10 * time.Second)
//	}
//}()
//}

func main() {
	pkg.LoadEnv()
	//
	setNewJwtPair()
	db := dbConnect()
	go alerter.Run()

	//Парсинг курса TODO: расскомментить - для сбора фиата в рублях в другие валюты
	go loadExcRates(db)

	//Задача: Добавить swagger
	//ВажнаяЗадача: добавить swagger

	//Generate jwt secret key
	//jwt_auth.GenerateNewPairs()

	//Bootstrap Fiber App
	app := bootstrap.NewFiberAppBoot(db)

	log.Fatal(app.Listen(":7700"))
}
