package v0

import "C"
import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/vmmgr/imacon/pkg/api/core/storage"
	"github.com/vmmgr/imacon/pkg/api/store"
	"log"
	"time"
)

func Create(storage *storage.Storage) (*storage.Storage, error) {
	db, err := store.ConnectDB()
	if err != nil {
		log.Println("database connection error")
		return storage, fmt.Errorf("(%s)error: %s\n", time.Now(), err.Error())
	}
	defer db.Close()

	err = db.Create(&storage).Error
	return storage, err
}

func Delete(storage *storage.Storage) error {
	db, err := store.ConnectDB()
	if err != nil {
		log.Println("database connection error")
		return fmt.Errorf("(%s)error: %s\n", time.Now(), err.Error())
	}
	defer db.Close()

	return db.Delete(storage).Error
}

func Update(flags uint, data *storage.Storage) error {
	db, err := store.ConnectDB()
	if err != nil {
		log.Println("database connection error")
		return fmt.Errorf("(%s)error: %s\n", time.Now(), err.Error())
	}
	defer db.Close()

	var result *gorm.DB
	if storage.UpdatePath == flags {
		result = db.Model(&storage.Storage{Model: gorm.Model{ID: data.ID}}).Update(storage.Storage{Path: data.Path})
	} else if storage.UpdateGroup == flags {
		result = db.Model(&storage.Storage{Model: gorm.Model{ID: data.ID}}).Update(storage.Storage{GroupID: data.GroupID})
	} else if storage.UpdateAll == flags {
		result = db.Model(&storage.Storage{Model: gorm.Model{ID: data.ID}}).Update(storage.Storage{
			GroupID: data.GroupID, Type: data.Type, Path: data.Path, CloudInit: data.CloudInit, MinCPU: data.MinCPU,
			MinMem: data.MinMem, OS: data.OS, Lock: data.Lock, Name: data.Name, UUID: data.UUID})
	} else {
		log.Println("select error")
		return fmt.Errorf("(%s)error: select\n", time.Now())
	}
	return result.Error
}

func Get(flags int, data *storage.Storage) storage.ResultDatabase {
	db, err := store.ConnectDB()
	if err != nil {
		log.Println("database connection error")
		return storage.ResultDatabase{Err: fmt.Errorf("(%s)error: %s\n", time.Now(), err.Error())}
	}
	defer db.Close()

	var storageStruct []storage.Storage

	if storage.SearchID == flags { //ID
		err = db.First(&storageStruct, data.ID).Error
	} else if storage.SearchGroupID == flags { //NodeStorage内の全StorageからGIDを検索
		err = db.Where("group_id = ?", data.GroupID).Find(&storageStruct).Error
	} else if storage.SearchType == flags {
		err = db.Where("type = ?", data.Type).Find(&storageStruct).Error
	} else if storage.SearchAdmin == flags {
		err = db.Where("admin = ?", data.Admin).Find(&storageStruct).Error
	} else if storage.SearchUUID == flags {
		err = db.Where("uuid = ?", data.UUID).Find(&storageStruct).Error
	} else if storage.SearchName == flags {
		err = db.Where(" name = ?", data.Name).Find(&storageStruct).Error
	} else {
		log.Println("select error")
		return storage.ResultDatabase{Err: fmt.Errorf("(%s)error: select\n", time.Now())}
	}
	return storage.ResultDatabase{Storage: storageStruct, Err: err}
}

func GetAll() storage.ResultDatabase {
	db, err := store.ConnectDB()
	if err != nil {
		log.Println("database connection error")
		return storage.ResultDatabase{Err: fmt.Errorf("(%s)error: %s\n", time.Now(), err.Error())}
	}
	defer db.Close()

	var storages []storage.Storage
	err = db.Find(&storages).Error
	return storage.ResultDatabase{Storage: storages, Err: err}
}
