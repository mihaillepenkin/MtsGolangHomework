package api

import (
	"encoding/json"
	"homework3/internal/application/dto"
	"homework3/internal/application/usecase"
	"net/http"
)


type HTTPHandler struct {
	service *usecase.UserService
}

func NewHandler(service *usecase.UserService) *HTTPHandler {
	controller := new(HTTPHandler)
	controller.service = service
	return controller
}

func (h *HTTPHandler) TopUpBalanceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input dto.TopUpBalanceInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var output dto.TopUpBalanceOutputDto
	err = h.service.TopUpBalance(input.PhoneNumber, input.Ruble, input.Kopeck)
	if (err != nil) {
		w.WriteHeader(http.StatusInternalServerError)
		output = dto.TopUpBalanceOutputDto{Status: "error", Msg: err.Error()}
	} else {
		output = dto.TopUpBalanceOutputDto{Status: "ok", Msg: ""}
	}
	json.NewEncoder(w).Encode(output)
}

func (h *HTTPHandler) TransferMoneyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input dto.TransferMoneyInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var output dto.TransferMoneyOutputDto
	err = h.service.TransferMoney(input.ToNumber, input.FromNumber, input.Ruble, input.Kopeck)
	if (err != nil) {
		w.WriteHeader(http.StatusInternalServerError)
		output = dto.TransferMoneyOutputDto{Status: "error", Msg: err.Error()}
	} else {
		output = dto.TransferMoneyOutputDto{Status: "ok", Msg: ""}
	}
	json.NewEncoder(w).Encode(output) 
}

func (h *HTTPHandler) GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input dto.GetBalanceInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var output dto.GetBalanceOutputDto
	ruble, kopeck, err := h.service.GetBalance(input.PhoneNumber)
	if (err != nil) {
		w.WriteHeader(http.StatusInternalServerError)
		output = dto.GetBalanceOutputDto{Status: "error", Msg: err.Error(), Ruble : "", Kopeck: ""}
	} else {
		output = dto.GetBalanceOutputDto{Status: "ok", Msg: "", Ruble : ruble, Kopeck: kopeck}
	}
	json.NewEncoder(w).Encode(output) 
}

func (h *HTTPHandler) NewUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input dto.NewUserInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var output dto.NewUserOutputDto
	err = h.service.NewUser(input.PhoneNumber, input.Name)
	if (err != nil) {
		w.WriteHeader(http.StatusInternalServerError)
		output = dto.NewUserOutputDto{Status: "error", Msg: err.Error()}
	} else {
		output = dto.NewUserOutputDto{Status: "ok", Msg: ""}
	}
	json.NewEncoder(w).Encode(output) 
}

func (h *HTTPHandler) SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/up_balance", h.TopUpBalanceHandler)
	mux.HandleFunc("/new_user", h.NewUser)
	mux.HandleFunc("/transfer", h.TransferMoneyHandler)
	mux.HandleFunc("/get_balance", h.GetBalanceHandler)

	return mux
}