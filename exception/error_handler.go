package exception

import (
	"belajar-golang-restfull-api/helper"
	"belajar-golang-restfull-api/model/dto"
	"errors"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(writer, request, err) {
		return
	}

	if validationErrors(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationErrors(writer http.ResponseWriter, _ *http.Request, err interface{}) bool {
	var exception validator.ValidationErrors
	ok := errors.As(err.(error), &exception)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		response := dto.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(writer, response)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, _ *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		response := dto.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, response)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, _ *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	response := dto.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, response)
}
