package common

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Abc() string {
	return "abc"
}

func JsonRetorno(mensagem string, codigo int) gin.H {
	uuidGerado := uuid.New()
	var errorCode string

	switch codigo {
	case 400:
		errorCode = "PATRULHAAGRICOLA_CADASTROS_ERROR_BAD_REQUEST"
	case 401:
		errorCode = "PATRULHAAGRICOLA_CADASTROS_ERROR_UNAUTHORIZED"
	case 403:
		errorCode = "PATRULHAAGRICOLA_CADASTROS_ERROR_FORBIDDEN"
	case 404:
		errorCode = "PATRULHAAGRICOLA_CADASTROS_ERROR_NOT_FOUND"
	case 409:
		errorCode = "PATRULHAAGRICOLA_CADASTROS_ERROR_CONFLICT"
	case 500:
		errorCode = "PATRULHAAGRICOLA_CADASTROS_ERROR_INTERNAL_SERVER_ERROR"
	default:
		errorCode = "PATRULHAAGRICOLA_CADASTROS_ERROR_UNKNOWN"
	}

	json := gin.H{
		"message":        mensagem,
		"debugID":        uuidGerado,
		"errorCode":      errorCode,
		"httpStatusCode": codigo,
	}
	return json
}
