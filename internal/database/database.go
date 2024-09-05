package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"sync"
)

type StorageDb struct {
	gormDb *gorm.DB
	appDsn string
	DbOnce sync.Once
	error  error
}

func NewStorageDb(dsn string) *StorageDb {
	var appDsn string
	if dsn != "" {
		return &StorageDb{
			appDsn: dsn,
		}
	}

	appDsn = os.Getenv("APP_PSQ_DSN")
	return &StorageDb{
		appDsn: appDsn,
	}
}

func (sd *StorageDb) Connect() *StorageDb {
	sd.DbOnce.Do(func() {
		// Инициализация соединения с БД
		sd.gormDb, sd.error = gorm.Open(postgres.Open(sd.appDsn), &gorm.Config{
			QueryFields:            true,
			SkipDefaultTransaction: true, //SYNC DATA
			PrepareStmt:            true, //INIT CACHE / FOR NEW TRANS
			Logger:                 logger.Default.LogMode(logger.Warn),
		})

		sd.gormDb.Session(&gorm.Session{})

		//if sd.error != nil {
		//	return sd
		//	//panic(sd.error) // Обработка ошибки - паника
		//}
	})
	return sd
}

//
//func (sd *StorageDb) Connect() *StorageDb {
//	//dsn := ""
//	var err error
//
//	sd.DbOnce.Do(func() {
//		sd.gormDb, sd.error = gorm.Open(postgres.Open(sd.appDsn), &gorm.Config{
//			QueryFields: true,
//			PrepareStmt: true,
//			Logger:      logger.Default.LogMode(logger.Info),
//			//DryRun:      true,
//			//Logger: logger.Default.LogMode(logger.Warn),
//		})
//		if err != nil {
//			panic(err)
//		}
//
//		//sd.gormDb = db
//	})
//	return sd
//}

//func (s *StoreDb) GetDb() (*gorm.DB, error) {
//	var err error
//	s.DbOnce.Do(func() {
//		s.gormDb, err = gorm.Open("mysql", s.appDsn)
//		if err != nil {
//			s.error = err
//			return
//		}
//	})
//	return s.gormDb, s.error
//}

func (sd *StorageDb) Ping() error {
	db, err := sd.gormDb.DB()
	if err != nil {
		return err
	}
	return db.Ping()
}

func (sd *StorageDb) GetDB() (*gorm.DB, error) {
	return sd.gormDb, sd.error
}
