package balance

import (
	"encoding/json"
	"fmt"
	"github.com/fabiopsouza/balance/internal/core/port/balance"
	"net/http"
	"strings"
)

type WebHandler struct {
	UseCase balance.UseCase
}

func NewHandler(useCase balance.UseCase) balance.Handler {
	return &WebHandler{
		UseCase: useCase,
	}
}

func (h *WebHandler) List(w http.ResponseWriter, r *http.Request) {
	accountID := strings.TrimPrefix(r.URL.Path, "/balances/")
	if accountID == "" {
		fmt.Println("Account ID is required")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Account ID is required"))
		return
	}

	output, err := h.UseCase.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
