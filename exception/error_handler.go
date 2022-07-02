package exception

import (
	"my-gram/helper"
	"my-gram/model/response"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	
	if notFoundError(writer, request, err) {
		return
	}

	if badRequestError(writer, request, err) {
		return
	}
	
	internalServerError(writer, request, err)
}

func badRequestError(writer http.ResponseWriter, request *http.Request, err interface{}) bool{
	exception, ok := err.(BadRequestError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
	
		webResponse := response.WebResponse {
			Code: http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: exception.Error,
		}
	
		helper.WriteToResponseBody(writer, http.StatusBadRequest ,webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
	
		webResponse := response.WebResponse {
			Code: http.StatusNotFound,
			Status: "NOT FOUND",
			Data: exception.Error,
		}
	
		helper.WriteToResponseBody(writer, http.StatusNotFound, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := response.WebResponse {
		Code: http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data: err,
	}

	helper.WriteToResponseBody(writer, http.StatusInternalServerError, webResponse)
}