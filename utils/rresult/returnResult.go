package rResult

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Totoal  int64       `json:"totoal"`
}
