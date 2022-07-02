package response

type WebResponse struct {
	Code   int         `josn:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}