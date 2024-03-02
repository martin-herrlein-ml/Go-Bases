package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalcTax_lowerBasicSalary(t *testing.T) {
	//arrange
	var salary float64 = 30000
	var expectedResult float64 = 0

	//act
	obtainResult := CalcTax(salary)

	//assert
	require.Equal(t, expectedResult, obtainResult)
}

func TestCalcTax_middleSalary(t *testing.T) {
	//arrange
	var salary float64 = 55000
	var expectedResult float64 = 9350

	//act
	obtainResult := CalcTax(salary)

	//assert
	require.Equal(t, expectedResult, obtainResult)
}

func TestCalcTax_higherSalary(t *testing.T) {
	salary := float64(150000)
	expectedResult := float64(40500)

	obtainResult := CalcTax(salary)

	require.Equal(t, expectedResult, obtainResult)
}
