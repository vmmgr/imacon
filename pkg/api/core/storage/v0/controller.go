package v0

func SendController(data interface{}) {
	//client := &http.Client{}
	//client.Timeout = time.Second * 5
	//
	//body, _ := json.Marshal(data)
	//
	////Header部分
	//header := http.Header{}
	//header.Set("Content-Length", "10000")
	//header.Add("Content-Type", "application/json")
	//header.Add("TOKEN_1", config.Conf.Controller.Auth.Token1)
	//header.Add("TOKEN_2", hash.Generate(config.Conf.Controller.Auth.Token2+config.Conf.Controller.Auth.Token3))
	//
	////リクエストの作成
	//req, err := http.NewRequest("POST", "http://"+config.Conf.Controller.User.IP+":"+
	//	strconv.Itoa(config.Conf.Controller.User.Port)+"/api/v1/controller/chat", bytes.NewBuffer(body))
	//if err != nil {
	//	return
	//}
	//
	//req.Header = header
	//
	//resp, err := client.Do(req)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//defer resp.Body.Close()
	//
	//_, err = ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
}
