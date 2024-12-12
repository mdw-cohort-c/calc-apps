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
	router.Handle("GET /sub", newHTTPHandler(logger, &calc.Subtraction{}))
	router.Handle("GET /mul", newHTTPHandler(logger, &calc.Multiplication{}))
	router.Handle("GET /div", newHTTPHandler(logger, &calc.Division{}))
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
	_, err = fmt.Fprint(response, c)
	if err != nil {
		this.logger.Println(err)
	}
}
