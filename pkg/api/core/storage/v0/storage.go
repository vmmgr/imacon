package v0

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/vmmgr/imacon/pkg/api/core/storage"
	"github.com/vmmgr/imacon/pkg/api/core/tool/config"
	"github.com/vmmgr/imacon/pkg/api/core/tool/gen"
	"github.com/vmmgr/imacon/pkg/api/meta/json"
	store "github.com/vmmgr/imacon/pkg/api/store/storage/v0"
	"log"
	"net/http"
	"strconv"
)

func Add(c *gin.Context) {
	var input storage.Storage

	err := c.BindJSON(&input)
	if err != nil {
		json.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	// File Pathがない場合はErrorを返す
	if input.Path == "" {
		json.ResponseError(c, http.StatusBadRequest, fmt.Errorf("Error: FileName is blank... "))
		return
	}

	// pathの定義
	var path string

	for _, tmpConf := range config.Conf.Storage {
		if tmpConf.Type == input.Type {
			path = tmpConf.Path + "/" + input.Path
		}
	}

	log.Println("Path: " + path)

	// typeよりpathが見つからない場合はErrorを返す
	if path == "" {
		json.ResponseError(c, http.StatusBadRequest, fmt.Errorf("Error: Not found... "))
		return
	}

	if fileExistsCheck(path) {
		json.ResponseError(c, http.StatusNotFound, fmt.Errorf("Error: file already exists... "))
		return
	}

	uuid, err := gen.GenerateUUID()
	if err != nil {
		json.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	// DBに追加
	if result, err := store.Create(&storage.Storage{Name: input.Name, UUID: uuid, GroupID: input.GroupID,
		Type: input.Type, Path: input.Path, CloudInit: input.CloudInit, Admin: input.Admin, MinCPU: input.MinCPU,
		MinMem: input.MinMem, OS: input.OS, Lock: &[]bool{false}[0]}); err != nil {
		json.ResponseError(c, http.StatusInternalServerError, err)
	} else {
		json.ResponseOK(c, result)
	}

}

func Update(c *gin.Context) {
	var input storage.Storage

	err := c.BindJSON(&input)
	if err != nil {
		json.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		json.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	getDB := store.Get(storage.SearchID, &storage.Storage{Model: gorm.Model{ID: uint(id)}})
	if getDB.Err != nil {
		json.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	update(&input, &getDB.Storage[0])

	if err := store.Update(storage.UpdateAll, &getDB.Storage[0]); err != nil {
		json.ResponseError(c, http.StatusInternalServerError, err)
	} else {
		json.ResponseOK(c, err)
	}
}

func Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		json.ResponseError(c, http.StatusBadRequest, err)
	}

	if result := store.Get(storage.SearchID, &storage.Storage{Model: gorm.Model{ID: uint(id)}}); result.Err != nil {
		json.ResponseError(c, http.StatusInternalServerError, result.Err)
	} else {
		response := result.Storage[0]
		for _, tmp := range config.Conf.Storage {
			if tmp.Type == response.Type {
				response := storage.Get{Path: tmp.Path + "/" + response.Path, CloudInit: *response.CloudInit}
				json.ResponseOK(c, response)
				return
			}
		}
		json.ResponseError(c, http.StatusInternalServerError, fmt.Errorf("Not Found: ID "))
	}
}

func GetName(c *gin.Context) {
	name := c.Param("name")

	if result := store.Get(storage.SearchName, &storage.Storage{Name: name}); result.Err != nil {
		json.ResponseError(c, http.StatusInternalServerError, result.Err)
	} else {
		response := result.Storage[0]
		for _, tmp := range config.Conf.Storage {
			if tmp.Type == response.Type {
				response := storage.Get{Path: tmp.Path + "/" + response.Path, CloudInit: *response.CloudInit}
				json.ResponseOK(c, response)
				return
			}
		}
		json.ResponseError(c, http.StatusInternalServerError, fmt.Errorf("Not Found: ID "))
	}
}

func GetUUID(c *gin.Context) {
	uuid := c.Param("uuid")

	if result := store.Get(storage.SearchUUID, &storage.Storage{UUID: uuid}); result.Err != nil {
		json.ResponseError(c, http.StatusInternalServerError, result.Err)
	} else {
		response := result.Storage[0]
		for _, tmp := range config.Conf.Storage {
			if tmp.Type == response.Type {
				response := storage.Get{Path: tmp.Path + "/" + response.Path, CloudInit: *response.CloudInit}
				json.ResponseOK(c, response)
				return
			}
		}
		json.ResponseError(c, http.StatusInternalServerError, fmt.Errorf("Not Found: ID "))
	}
}

func GetAll(c *gin.Context) {
	if result := store.GetAll(); result.Err != nil {
		json.ResponseError(c, http.StatusInternalServerError, result.Err)
	} else {
		json.ResponseOK(c, result.Storage)
	}
}
