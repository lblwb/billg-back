package order

import (
	"backend/internal/database"
)

type UserOrderSvcEntity struct {
	db *database.StorageDb
}

func NewUserOrderSvcEntity(db *database.StorageDb) *UserOrderSvcEntity {
	return &UserOrderSvcEntity{
		db: db,
	}
}

func (uos UserOrderSvcEntity) GetAllUserOrdersSvc() ([]UserOrderServices, error) {

	var userOrdSvc []UserOrderServices
	//user

	db, err := uos.db.GetDB()
	if err != nil {
		return nil, err
	}

	err = db.
		//
		Preload("Services").
		Preload("Services.Order").
		Preload("Services.Service").
		Order("updated_at desc").
		//
		Find(&userOrdSvc).Error

	return userOrdSvc, err
}

// InsertUserOrderServices - Создание заказанных услуг пользователя
func (uos UserOrderSvcEntity) InsertUserOrderServices(order *UserOrderServices) error {
	//order := UserOrderServices{
	//	UserID:      1,
	//	TotalAmount: 1500.00,
	//	PromoCode:   "PROMO123",
	//	//OrderParams: map[string]interface{}{
	//	//	"domain name": map[string]interface{}{"price": 1200.00, "value": "domain.com"},
	//	//	"speed":       map[string]interface{}{"price": 300.00, "value": 240},
	//	//},
	//	UpdatedAt: time.Now(),
	//	Status:    "pending", // Начальный статус заявки
	//}

	db, err := uos.db.GetDB()
	if err != nil {
		return err
	}
	//Сохранение
	return db.
		Preload("Services").
		Preload("Order").
		Preload("TariffsServices").
		Create(&order).Error
}

// GetUserOrderSvcBySlug - Информация о заказе
func (uos UserOrderSvcEntity) GetUserOrderSvcBySlug(slugName string) (UserOrderServices, error) {

	userOrdSvc := UserOrderServices{}
	//user

	db, err := uos.db.GetDB()
	if err != nil {
		return UserOrderServices{}, err
	}

	err = db.
		//
		Preload("Services").
		Preload("Services.Order").
		Preload("Services.Service").
		//
		First(&userOrdSvc, &UserOrderServices{
			Slug: slugName,
		}).Error

	return userOrdSvc, err
}
