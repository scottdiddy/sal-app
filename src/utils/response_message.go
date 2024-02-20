package utils

import "github.com/scottdiddy/sal-app/src/models"

func ResponseMessage(message string, data interface{}) models.ResponseMessage {
	return models.ResponseMessage{
		Message: message,
		Data: data,
	}
}