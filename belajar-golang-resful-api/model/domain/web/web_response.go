package web

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	DATA   interface{} `json:"data"`
}
