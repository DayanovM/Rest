package http

import (
	"net/http"
	"restapi/banking"
)

type HTTPHandlers struct {
	bank *banking.Bank
}

func NewHTTPHandlers(bank *banking.Bank) *HTTPHandlers {
	return &HTTPHandlers{}
}

func (h *HTTPHandlers) HandleCreateAccount(w http.ResponseWriter, r *http.Request) {
	/*
	   	pattern: /tasks

	   method: POST
	   info: JSON HTTP request body

	   succeed:
	   - status code: 201 created
	   - response: JSON with Account

	   failed:
	   - status code: 400, 409, 500
	   - response: JSON with error + time

	*/
}

func (h *HTTPHandlers) HandleDeposit(w http.ResponseWriter, r *http.Request) {
	/*
	   	pattern: /tasks/{title}

	   method: PATCH
	   info: JSON HTTP request body

	   succeed:
	   - status code: 201 created
	   - response: JSON with Account

	   failed:
	   - status code: 400, 409, 500
	   - response: JSON with error + time

	*/
}
