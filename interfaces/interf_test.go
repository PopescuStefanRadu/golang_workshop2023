package interfaces_test

import (
	"errors"
	"fmt"
	"testing"
)

type Tester interface {
	Test() error
}

type TesterImpl struct {
	ShouldReturnError bool
}

func (t *TesterImpl) Test() error {
	if t == nil {
		return errors.New("the test has not yet been defined")
	}
	if t.ShouldReturnError {
		return nil
	}
	return errors.New("there was an error")
}

func TestInterfaces(t *testing.T) {
	var tester Tester
	if tester == nil {
		fmt.Println("tester is nil")
	}
	var testerImpl *TesterImpl
	tester = testerImpl
	if tester == nil {
		fmt.Println("tester is nil after assignment")
	} else {
		fmt.Println("tester is NOT nil after assignment")
	}
	_ = tester.Test()
}

func TestInterfaceCanHoldAnything(t *testing.T) {
	m := map[string]interface{}{}
	m["Test"] = false
	m["Test2"] = 1
	m["Test3"] = func() {}

	m2 := map[string]any{}
	fmt.Printf("m has %T, m2 has %T\n", m, m2)
}
