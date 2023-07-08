package models

type Tabler interface {
	TableName() string
}

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
