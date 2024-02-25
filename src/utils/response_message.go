package utils

import "github.com/scottdiddy/sal-app/src/models"

func ResponseMessage(status bool, message string, data interface{}) models.ResponseMessage {
	return models.ResponseMessage{
		Status: status,
		Message: message,
		Data: data,
	}
}