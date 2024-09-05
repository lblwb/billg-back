package user

import (
	"backend/internal/billing_app/models/order"
	"backend/internal/database"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type AuthUserHiddenField struct {
	//ID uint `gorm:"primarykey"`
	Id       uint   `json:"id" gorm:"primaryKey"`
	Password []byte `json:"password" `
}

type AuthUser struct {
	Username string `json:"username"`
}

type Users struct {
	Id       uint   `json:"-" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"-" gorm:"password"`
	//UserAuthHidden
	TelegramId string    `json:"t_id"`
	Birthday   time.Time `json:"birthday"`
	//
	Balance UsersBalance              `json:"balance" gorm:"foreignkey:UserID"`
	Orders  []order.UserOrderServices `json:"orders" gorm:"foreignkey:UserID"`
	//BalanceUSD UsersBalance `json:"balance_usd" gorm:"foreignkey:UserID;association_foreignkey:Id"`
}

type UsersEntity struct {
	db                 *database.StorageDb
	usersBalanceEntity *UsersBalanceEntity
}

func NewUsersEntity(db *database.StorageDb) *UsersEntity {
	usersEntity := &UsersEntity{
		db:                 db,
		usersBalanceEntity: NewUserBalanceEntity(db),
	}
	//usersEntity.usersBalanceEntity = NewUserBalanceEntity(db)
	return usersEntity
}

// CalcExchangesRates Метод добавления USD баланса на основе курса обмена
//func (u *Users) CalcExchangesRates(dirCurrency string) *Users {
//var exchangeRate float64 = 0
//
//if dirCurrency == "" {
//	dirCurrency = "USD"
//}
//
//cacheManager := CacheManagerSync.NewCacheManagerSync()
//byDirCurrencyUsd := cacheManager.Remember(fmt.Sprintf("calc_currency_rate_%s", dirCurrency), 60*time.Second, func() interface{} {
//	byDirCurrencyDb, dirCur, err := currency.GetExcRateByDirCurrency(dirCurrency)
//	if err != nil || byDirCurrencyDb.Value == 0 {
//		fmt.Println("Calc Exh error: ", dirCur, err.Error())
//	} else {
//		return byDirCurrencyDb
//	}
//	return currency_rates.CurrencyRate{}
//})
////
//currencyRateUSDValue := byDirCurrencyUsd.(currency_rates.CurrencyRate).Value
////
//if byDirCurrencyUsd == nil || currencyRateUSDValue == 0 {
//	u.Balance.AmountUSD = 0
//	fmt.Println("Calc Exh error: ", byDirCurrencyUsd)
//	return u
//}
//
//fmt.Println("currencyRates->Value", u.Balance.Amount, currencyRateUSDValue, "=", u.Balance.Amount/currencyRateUSDValue)
//roundedValue := fmt.Sprintf("%.1f", u.Balance.Amount/currencyRateUSDValue)
//roundedFloat, _ := strconv.ParseFloat(roundedValue, 64)
//u.Balance.AmountUSD = roundedFloat
//	return u
//}

func (ue UsersEntity) GetAllUsers() ([]Users, error) {
	var users []Users
	db, err := ue.db.GetDB()
	if err != nil {
		return nil, err
	}
	db.Select("*").
		//Where("id = @id", sql.Named("id", id)).
		Order("username asc").
		Preload("Balance").
		Preload("Orders").
		Find(&users)

	//database.GetDB().Preload("Balance").First(&user)
	return users, nil
}

func (ue UsersEntity) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (ue UsersEntity) GetUserByLogin(authUserData AuthUser) (Users, error) {
	user := Users{}
	db, err := ue.db.GetDB()
	if err != nil {
		return Users{}, err
	}

	db.Select("*").
		Where("username = @name", sql.Named("name", authUserData.Username)).
		Or("email = @email", sql.Named("email", authUserData.Username)).
		Or("telegram_id = @telegram", sql.Named("telegram", authUserData.Username)).
		//
		Preload("Balance").
		First(&user)

	//database.GetDB().Preload("Balance").First(&user)

	//fmt.Println(user)
	if db.Error != nil {
		return user, db.Error
	} else {
		return user, nil
	}
}

// GetUserByUsername - получение определенного пользователя по username
func (ue UsersEntity) GetUserByUsername(username string) (Users, error) {
	db, err := ue.db.GetDB()
	user := Users{}
	if err != nil {
		return Users{}, err
	}
	//Where("username = @name", sql.Named("name", username)).
	db.Select("*").
		Preload("Balance").
		First(&user, &Users{Username: username})

	amountExchanges, err := ue.usersBalanceEntity.CalculateAmountExchanges(user.Balance)
	if err != nil {
		return Users{}, err
	}
	//
	user.Balance.AmountExchanges = amountExchanges

	//database.GetDB().Preload("Balance").First(&user)

	//fmt.Println(user)
	if db.Error != nil {
		return user, db.Error
	} else {
		return user, nil
	}
}

// GetUserById - получение определенного пользователя по id
func (ue UsersEntity) GetUserById(id uint) Users {
	user := Users{}
	db, err := ue.db.GetDB()
	if err != nil {
		return Users{}
	}
	db.Select("*").
		Where("id = @id", sql.Named("id", id)).
		Preload("Balance").
		First(&user)

	//database.GetDB().Preload("Balance").First(&user)
	return user
}

// CreateUser - создание нового пользователя
func (ue UsersEntity) CreateUser(username string, password string) (*gorm.DB, error) {
	db, err := ue.db.GetDB()
	if err != nil {
		return db, err
	}
	password, err = ue.hashPassword(password)
	if err != nil {
		fmt.Println("Ошибка формирования шифра!пароля")
	}
	user := Users{
		Id:         1,
		Username:   username,
		Password:   password,
		TelegramId: "",
		Birthday:   time.Now(),
	}
	tx := db.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		return tx, nil
	}
}
