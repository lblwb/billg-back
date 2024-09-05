package user

import (
	"backend/internal/billing_app/models/currency"
	"backend/internal/database"
	"database/sql"
	"errors"
)

type UsersBalance struct {
	Id              uint               `json:"-"  gorm:"-"`
	UserID          uint               `json:"-" gorm:"user_id"`
	AmountUSD       float64            `json:"amount_usd" gorm:"-"`
	Amount          float64            `json:"amount"  gorm:"amount"`
	FreezeBalance   bool               `json:"-"  gorm:"freeze_balance"`
	AmountExchanges map[string]float64 `json:"amount_exc" gorm:"-"`
}

type AmountExchanges struct {
	Currency string
	Value    float64
}

type UsersBalanceEntity struct {
	db          *database.StorageDb
	usersEntity *UsersEntity
	ratesEntity *currency.RatesEntity
}

func NewUserBalanceEntity(db *database.StorageDb) *UsersBalanceEntity {
	usersBalanceEntity := &UsersBalanceEntity{
		db: db,
		//usersEntity: NewUsersEntity(db),
		ratesEntity: currency.NewRatesEntity(db),
	}
	//usersBalanceEntity.usersEntity = usersEntity

	return usersBalanceEntity
}

// BeforeFind hook для автоматического обновления AmountExchanges перед каждым запросом баланса
//func (ub *UsersBalance) BeforeFirst(tx *gorm.DB) error {
//	amountExchanges, err := ub.CalculateAmountExchanges(tx)
//	if err != nil {
//		return err
//	}
//	// Записываем результаты в поле AmountExchanges структуры
//	ub.AmountExchanges = amountExchanges
//	return nil
//}

func (ube UsersBalanceEntity) CalculateAmountExchanges(userBalance UsersBalance) (map[string]float64, error) {
	rates, err := ube.ratesEntity.GetAllExcRates()
	if err != nil {
		return nil, err
	}

	// Создаем карту для хранения суммы в каждой валюте
	amountExchanges := map[string]float64{}

	//log.Println("rates", rates)

	// Рассчитываем баланс для каждой валюты
	for _, rate := range rates {
		amountInCurrency := userBalance.Amount / rate.Value
		//amountExchanges[rate.DirCurrency] = amountInCurrency
		amountExchanges[rate.DirCurrency] = amountInCurrency
		//log.Println(amountExchanges[rate.DirCurrency], amountInCurrency)
	}

	//fmt.Println(amountExchanges)

	// Записываем результаты в поле AmountExchanges структуры
	//ub.AmountExchanges = append(ub.AmountExchanges, amountExchanges)

	//return nil
	return amountExchanges, nil
}

// GetAvailBalance - Проверка на доступность баланса пользователя
func (ube UsersBalanceEntity) GetAvailBalance(UserId uint, availAmountBalance float64) (bool, float64, error) {
	user := ube.usersEntity.GetUserById(UserId)
	//userNotFound := Users{}
	//
	if user.Id == 0 {
		return false, 0, errors.New("not found user")
	} else if user.Balance.Amount >= availAmountBalance {
		return true, user.Balance.Amount, nil
	} else {
		return false, user.Balance.Amount, errors.New("avail Balance not amount sum")
	}
}

// WithdrawUserBalance - Списание суммы с баланса пользователя
func (ube UsersBalanceEntity) WithdrawUserBalance(UserId uint, amountAvail float64) (bool, string, error) {
	userBalance := UsersBalance{}
	db, err := ube.db.GetDB()
	if err != nil {
		return false, "not_success", errors.New("error, not avail")
	}

	err = db.
		Where("user_id = @user_id", sql.Named("user_id", UserId)).
		First(&userBalance).
		Error

	// Find the user and preload the balance association
	if err != nil {
		return false, "user_not_found", err
	}

	// Check if the balance is sufficient for the withdrawal
	if userBalance.Amount >= amountAvail {
		// Deduct the amount from the user's balance
		userBalance.Amount -= amountAvail

		// Save the updated user, including the balance, with "ON CONFLICT"
		if err := db.
			Where("user_id = @user_id", sql.Named("user_id", UserId)).
			Save(&userBalance).Error; err != nil {
			return false, "update_failed", err
		}

		return true, "success", nil
	} else {
		// Insufficient balance
		return false, "no_withdraw", errors.New("Баланс пользователя меньше запрашиваемой суммы")
	}

}

// DrawUserBalance - Списание суммы с баланса пользователя
func DrawUserBalance(UserId uint, amount float64) {

}
