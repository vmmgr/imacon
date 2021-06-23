package controller

type Controller struct {
	UUID     string `json:"uuid"`
	Error    string `json:"error"`
	Progress uint   `json:"progress"`
	Finish   bool   `json:"finish"`
}
