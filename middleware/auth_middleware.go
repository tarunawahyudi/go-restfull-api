package middleware

import (
	"belajar-golang-restfull-api/helper"
	"belajar-golang-restfull-api/model/dto"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "secret" == request.Header.Get("X-API-Key") {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		response := dto.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.WriteToResponseBody(writer, response)
	}
}
