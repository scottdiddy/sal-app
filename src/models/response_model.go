package models

type ResponseMessage struct {
	Status bool //status returns if the API call was successful
	Message string
	Data    interface{}
}
