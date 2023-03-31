package arithmatic

import "fmt"

type Ops struct {
	inputA int32
	inputB int32
	opName string
	result int32
}

type Adapter struct {
	history []Ops
}

func NewAdapter() *Adapter {
	return &Adapter{
		history: make([]Ops, 10),
	}
}

func (arith *Adapter) Addition(a, b int32) (int32, error) {
	v := a + b
	op := Ops{inputA: a, inputB: b, opName: "Addition", result: v}
	arith.history = append(arith.history, op)
	return v, nil
}

func (arith *Adapter) Subtraction(a, b int32) (int32, error) {
	v := a - b
	op := Ops{inputA: a, inputB: b, opName: "Subtraction", result: v}
	arith.history = append(arith.history, op)
	return v, nil
}

func (arith *Adapter) Multiplication(a, b int32) (int32, error) {
	v := a * b
	op := Ops{inputA: a, inputB: b, opName: "Multiplication", result: v}
	arith.history = append(arith.history, op)
	return v, nil
}

func (arith *Adapter) Division(a, b int32) (int32, error) {
	if b == 0 {
		return 0, fmt.Errorf("cant process as b is zero")
	}
	v := a / b
	op := Ops{inputA: a, inputB: b, opName: "Division", result: v}
	arith.history = append(arith.history, op)
	return v, nil
}
