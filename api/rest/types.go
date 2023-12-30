package rest

type Response struct {
	Status int `json:"status"`
	Result any `json:"result"`
}

type ResponseCount struct {
	Count int `json:"count"`
}
