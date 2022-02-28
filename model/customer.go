package model

type Customer struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	MobileNumber string `json:"mobile_number"`
	EmailId      string `json:"email_id"`
	Address      string `json:"address"`
}
type Response struct {
	Status int         `json:"status"`
	Error  bool        `json:"error"`
	Data   interface{} `json:"data"`
}
