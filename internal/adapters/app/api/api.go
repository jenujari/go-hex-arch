package api

import (
	"go-hex-arch/internal/ports"
)

type Adapter struct {
	db    ports.DbPort
	arith ports.ArithmaticPort
}

func NewAdapter(db ports.DbPort, arith ports.ArithmaticPort) *Adapter {
	return &Adapter{arith: arith, db: db}
}

func (apiA Adapter) GetAddition(a, b int32) (int32, error) {
	ans, err := apiA.arith.Addition(a, b)

	if err != nil {
		return 0, err
	}

	err = apiA.db.AddToHistory(ans, "addition")
	if err != nil {
		return 0, err
	}

	return ans, nil
}

func (apiA Adapter) GetSubtraction(a, b int32) (int32, error) {
	ans, err := apiA.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}

	err = apiA.db.AddToHistory(ans, "subtraction")
	if err != nil {
		return 0, err
	}

	return ans, nil
}

func (apiA Adapter) GetMultiplication(a, b int32) (int32, error) {
	ans, err := apiA.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}

	err = apiA.db.AddToHistory(ans, "multiplication")
	if err != nil {
		return 0, err
	}

	return ans, nil
}

func (apiA Adapter) GetDivision(a, b int32) (int32, error) {
	ans, err := apiA.arith.Division(a, b)
	if err != nil {
		return 0, err
	}

	err = apiA.db.AddToHistory(ans, "division")
	if err != nil {
		return 0, err
	}

	return ans, nil
}
