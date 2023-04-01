package http

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func (httpa Adapter) GetAddition(w http.ResponseWriter, r *http.Request) {
	var err error

	A := r.URL.Query().Get("a")
	B := r.URL.Query().Get("b")

	if A == "" || B == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Invalid request")
		return
	}

	intA, err := strconv.Atoi(A)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error in parsing value of A : %v", err)
		return
	}

	intB, err := strconv.Atoi(B)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error in parsing value of B : %v", err)
		return
	}

	answer, err := httpa.api.GetAddition(int32(intA), int32(intB))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error in op : %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Value : %v", answer)
}

func (httpa Adapter) GetSubtraction(w http.ResponseWriter, r *http.Request) {
	var err error

	A := r.URL.Query().Get("a")
	B := r.URL.Query().Get("b")

	if A == "" || B == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Invalid request")
		return
	}

	intA, err := strconv.Atoi(A)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error in parsing value of A : %v", err)
		return
	}

	intB, err := strconv.Atoi(B)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error in parsing value of B : %v", err)
		return
	}

	answer, err := httpa.api.GetSubtraction(int32(intA), int32(intB))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error in op : %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Value : %v", answer)
}

func (httpa Adapter) GetMultiplication(w http.ResponseWriter, r *http.Request) {
	var err error

	A := r.URL.Query().Get("a")
	B := r.URL.Query().Get("b")

	if A == "" || B == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Invalid request")
		return
	}

	intA, err := strconv.Atoi(A)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error in parsing value of A : %v", err)
		return
	}

	intB, err := strconv.Atoi(B)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error in parsing value of B : %v", err)
		return
	}

	answer, err := httpa.api.GetMultiplication(int32(intA), int32(intB))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error in op : %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Value : %v", answer)
}

func (httpa Adapter) GetDivision(w http.ResponseWriter, r *http.Request) {
	var err error

	A := r.URL.Query().Get("a")
	B := r.URL.Query().Get("b")

	if A == "" || B == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Invalid request")
		return
	}

	intA, err := strconv.Atoi(A)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error in parsing value of A : %v", err)
		return
	}

	intB, err := strconv.Atoi(B)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error in parsing value of B : %v", err)
		return
	}

	answer, err := httpa.api.GetDivision(int32(intA), int32(intB))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error in op : %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Value : %v", answer)
}
