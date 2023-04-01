package http

import (
	"fmt"
	"go-hex-arch/internal/ports"
	"net/http"
)

type Adapter struct {
	api ports.APIPort
	mux *http.ServeMux
}

func NewAdapter(api ports.APIPort) *Adapter {
	adapter := &Adapter{api: api}
	mux := http.NewServeMux()

	mux.HandleFunc("/add", adapter.GetAddition)
	mux.HandleFunc("/sub", adapter.GetSubtraction)
	mux.HandleFunc("/mul", adapter.GetMultiplication)
	mux.HandleFunc("/sub", adapter.GetDivision)

	adapter.mux = mux

	return adapter
}

func (httpa Adapter) Run() error {
	if err := http.ListenAndServe(":8080", httpa.mux); err != nil {
		return fmt.Errorf("failed to serve http server over port 8080: %v", err)
	}
	return nil
}
