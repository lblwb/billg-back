package tariff

import (
	"backend/internal/database"
	"database/sql"
	"encoding/json"
	"time"
)

/**
Модель услуги
*/

type TariffsServices struct {
	Id          uint            `json:"-" gorm:"id"`
	ServiceID   uint            `json:"-" gorm:"service_id"`
	Slug        string          `json:"slug"`
	FullName    string          `json:"full_name"`
	FullNameEn  string          `json:"full_name_en"`
	FirstAmount string          `json:"first_amount"`
	DescAlert   string          `json:"desc_alert"`
	DescAlertEn string          `json:"desc_alert_en"`
	CreatedAt   time.Time       `json:"-"`
	Params      json.RawMessage `json:"params"`
}

type TariffsServiceEntity struct {
	db             *database.StorageDb
	tariffServices *TariffsServices
}

func NewTariffsServiceEntity(db *database.StorageDb) *TariffsServiceEntity {
	return &TariffsServiceEntity{
		db: db,
	}
}

// GetTariffBySlug Получение тарифа по Slug
func (tse TariffsServiceEntity) GetTariffBySlug(slugName string) (TariffsServices, error) {
	var tariff TariffsServices

	dbData, err := tse.db.GetDB()
	if err != nil {
		return TariffsServices{}, err
	}

	db := dbData.
		Select("*").
		Where("slug = @slug", sql.Named("slug", slugName)).
		//Preload("Tariffs", func(db *gorm.DB) *gorm.DB {
		//	return db.Order("created_at asc") // Сортировка через связь "Tariffs"
		//}).
		First(&tariff)

	if db.Error != nil {
		return tariff, db.Error
	} else {
		return tariff, nil
	}
}
