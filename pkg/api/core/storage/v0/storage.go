package v0

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/vmmgr/imacon/pkg/api/core/storage"
	"github.com/vmmgr/imacon/pkg/api/core/tool/config"
	"github.com/vmmgr/imacon/pkg/api/meta/json"
	store "github.com/vmmgr/imacon/pkg/api/store/storage/v0"
	"net/http"
	"strconv"
)

func Add(c *gin.Context) {
	var input storage.Add

	err := c.BindJSON(&input)
	if err != nil {
		json.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	// FileNameがない場合はErrorを返す
	if input.FileName != "" {
		json.ResponseError(c, http.StatusBadRequest, fmt.Errorf("Error: FileName is blank... "))
		return
	}

	// pathの定義
	path := ""

	for _, tmpConf := range config.Conf.Storage {
		if tmpConf.Type == input.Type {
			path = tmpConf.Path + "/" + input.FileName
		}
	}

	// typeよりpathが見つからない場合はErrorを返す
	if path == "" {
		json.ResponseError(c, http.StatusBadRequest, fmt.Errorf("Error: Not found... "))
		return
	}

	if fileExistsCheck(path) {
		json.ResponseError(c, http.StatusNotFound, fmt.Errorf("Error: file already exists... "))
		return
	}

	// DBに追加
	if result, err := store.Create(&storage.Storage{
		GroupID: input.GroupID, Type: input.Type, Path: input.FileName, CloudInit: &[]bool{input.CloudInit}[0],
		Admin: &[]bool{input.Admin}[0], MinCPU: input.MinCPU, MinMem: input.MinMem, OS: input.OS,
		Lock: &[]bool{false}[0]}); err != nil {
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
		json.ResponseOK(c, result.Storage)
	}
}

func GetAll(c *gin.Context) {
	if result := store.GetAll(); result.Err != nil {
		json.ResponseError(c, http.StatusInternalServerError, result.Err)
	} else {
		json.ResponseOK(c, result.Storage)
	}
}
