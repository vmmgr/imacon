package v2

import (
	"bytes"
	"encoding/json"
	"github.com/vmmgr/imacon/pkg/api/core/controller"
	"github.com/vmmgr/imacon/pkg/api/core/tool/config"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func SendController(url string, data controller.Controller) error {
	client := &http.Client{}
	client.Timeout = time.Second * 5

	body, _ := json.Marshal(data)

	//Header部分
	header := http.Header{}
	header.Set("Content-Length", "10000")
	header.Add("Content-Type", "application/json")
	header.Add("TOKEN_1", config.Conf.Controller.Auth.Token1)
	header.Add("TOKEN_2", config.Conf.Controller.Auth.Token2)

	//リクエストの作成
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header = header

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
