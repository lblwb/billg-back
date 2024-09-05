package currency

import (
	"backend/internal/database"
	"errors"
	"gorm.io/gorm"
	//"backend/app/service/workers/currency_rates"
)

type CurrRates struct {
	Id          uint    `json:"id" gorm:"id"`
	Currency    string  `json:"currency" gorm:"currency"`
	DirCurrency string  `json:"dir_currency" gorm:"dir_currency"`
	Value       float64 `json:"value" gorm:"value"`
	Crypto      bool    `json:"cr" gorm:"crypto"`
}

func (CurrRates) TableName() string {
	return "currency_rates"
}

type RatesEntity struct {
	db            *database.StorageDb
	CurrencyRates []CurrRates
}

func NewRatesEntity(db *database.StorageDb) *RatesEntity {
	return &RatesEntity{
		db: db,
	}
}

func (re RatesEntity) GetAllExcRates() ([]CurrRates, error) {
	var currencyRates []CurrRates
	db, err := re.db.GetDB()
	if err != nil {
		return []CurrRates{}, err
	}

	currency := db.
		Select("*").
		//Where("currency = ?", "RUB").
		Find(&currencyRates, CurrRates{
			Currency: "RUB",
		})

	return currencyRates, currency.Error
}

func (re RatesEntity) GetExcRateByDirCurrency(dirCurrency string) (CurrRates, string, error) {
	var currencyRates CurrRates
	db, err := re.db.GetDB()
	if err != nil {
		return CurrRates{}, "", err
	}

	currency := db.
		Select("*").
		Where("currency = ? AND dir_currency = ?", "RUB", dirCurrency).
		First(&currencyRates)
	return currencyRates, dirCurrency, currency.Error
}

// SaveExchangeRates
func (re *RatesEntity) SaveExchangesRates(currencyRates []CurrRates) error {
	// Assuming you have access to the database instance in re.db
	db, err := re.db.GetDB()
	if err != nil {
		return err
	}

	for _, rate := range currencyRates {
		var existingRate CurrRates
		if err := db.Where("dir_currency = ?", rate.DirCurrency).First(&existingRate).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// dirCurrency не найден, выполняем вставку новой записи
				if err := db.Create(&rate).Error; err != nil {
					return err
				}
			} else {
				// Произошла ошибка при поиске dirCurrency
				return err
			}
		} else {
			// dirCurrency найден, обновляем его значение
			if err := db.Model(&existingRate).Updates(rate).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
