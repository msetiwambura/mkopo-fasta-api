package utils

import (
	"github.com/gin-gonic/gin"
	"loanapi/responses"
)

func CreateSuccessResponse[T any](statusMessage, statusDesc string, data []T) responses.Response[T] {
	header := responses.ResponseHeader{
		StatusMessage: statusMessage,
		StatusCode:    1000,
		StatusDesc:    statusDesc,
	}

	return responses.Response[T]{
		ResponseHeader: header,
		ResponseBody:   data,
	}
}

func CreateMapResponse(statusMessage, statusDesc string, data gin.H) gin.H {
	header := responses.ResponseHeader{
		StatusMessage: statusMessage,
		StatusCode:    1000,
		StatusDesc:    statusDesc,
	}

	return gin.H{
		"ResponseHeader": header,
		"ResponseBody":   data,
	}
}

func CreateErrorResponse(message, statusDesc string) responses.Response[interface{}] {
	header := responses.ResponseHeader{
		StatusMessage: message,
		StatusCode:    1001,
		StatusDesc:    statusDesc,
	}

	return responses.Response[interface{}]{
		ResponseHeader: header,
		ResponseBody:   nil,
	}
}

//func CreateGenericSuccessResponse[T any](statusMessage, statusDesc string, responseBody T) map[string]interface{} {
//	response := map[string]interface{}{
//		"ResponseHeader": map[string]interface{}{
//			"StatusCode":    1000,
//			"StatusMessage": statusMessage,
//			"StatusDesc":    statusDesc,
//		},
//		"ResponseBody": responseBody,
//	}
//	return response
//}

type ResponseHeader struct {
	StatusCode    int    `json:"StatusCode"`
	StatusMessage string `json:"StatusMessage"`
	StatusDesc    string `json:"StatusDesc"`
}

// Response defines the complete structure of the response
type Response[T any] struct {
	ResponseHeader ResponseHeader `json:"ResponseHeader"`
	ResponseBody   T              `json:"ResponseBody"`
}

// CreateGenericSuccessResponse CreateSuccessResponse creates a standard response format with generic type for ResponseBody
func CreateGenericSuccessResponse[T any](statusMessage, statusDesc string, responseBody T) Response[T] {
	return Response[T]{
		ResponseHeader: ResponseHeader{
			StatusCode:    1000,
			StatusMessage: statusMessage,
			StatusDesc:    statusDesc,
		},
		ResponseBody: responseBody,
	}
}
