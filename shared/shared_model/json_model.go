package shared_model

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SingleResponse struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
}

type PagedResponse struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
	Paging Paging      `json:"paging"`
}

type StatusSucces struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
