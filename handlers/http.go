package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/mdw-cohort-c/calc-lib"
)

func NewRouter(logger *log.Logger) http.Handler {
	router := http.NewServeMux()
	router.Handle("GET /add", newHTTPHandler(logger, &calc.Addition{}))
	return router
}

type HTTPHandler struct {
	logger     *log.Logger
	calculator Calculator
}

func newHTTPHandler(logger *log.Logger, calculator Calculator) http.Handler {
	return &HTTPHandler{logger: logger, calculator: calculator}
}

func (this *HTTPHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	a, err := strconv.Atoi(query.Get("a"))
	if err != nil {
		http.Error(response, "a is invalid", http.StatusUnprocessableEntity)
		return
	}
	b, err := strconv.Atoi(query.Get("b"))
	if err != nil {
		http.Error(response, "b is invalid", http.StatusUnprocessableEntity)
		return
	}
	c := this.calculator.Calculate(a, b)
	_, _ = fmt.Fprint(response, c)
}
