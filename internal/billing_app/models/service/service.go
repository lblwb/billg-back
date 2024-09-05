package service

import (
	"backend/internal/billing_app/models/tariff"
	"backend/internal/database"
	"database/sql"
	"gorm.io/gorm"
)

type Services struct {
	//ID uint `gorm:"primarykey"`
	Id           uint   `json:"-" gorm:"id, primaryKey"`
	Slug         string `json:"slug"`
	Name         string `json:"name"`
	FullName     string `json:"full_name"`
	FullNameEn   string `json:"full_name_en"`
	DeviceName   string `json:"device_name"`
	DeviceSlug   string `json:"device_slug"`
	BannerDesc   string `json:"banner_desc"`
	BannerDescEn string `json:"banner_desc_en"`
	//Tariffs    []Tariffs `json:"tariffs"`
	Tariffs []tariff.TariffsServices `json:"tariffs" gorm:"foreignkey:ServiceID;association_foreignkey:Id"`
}

type ServicesEntity struct {
	db *database.StorageDb
}

func NewServicesEntity(db *database.StorageDb) *ServicesEntity {
	return &ServicesEntity{
		db: db,
	}
}

// GetServiceBySlug Получение сервиса по Slug
func (se ServicesEntity) GetServiceBySlug(slugName string) (Services, error) {
	var service Services
	db, err := se.db.GetDB()
	if err != nil {
		return Services{}, err
	}
	dbData := db.Select("*").
		Where("slug = @slug", sql.Named("slug", slugName)).
		Preload("Tariffs", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at asc") // Сортировка через связь "Tariffs"
		}).
		First(&service)

	if dbData.Error != nil {
		return service, dbData.Error
	} else {
		return service, nil
	}
}

// GetAllServices - Получение всего спектра услуг
func (se ServicesEntity) GetAllServices() []Services {
	var serviceList []Services
	db, err := se.db.GetDB()
	if err != nil {
		return []Services{}
	}
	// Fetch services
	if db.Select("*").
		Preload("Tariffs").
		Find(&serviceList).Error != nil {
		return serviceList
	} else {
		return serviceList
	}
}
