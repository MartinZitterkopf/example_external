package matrix

import (
	"fmt"
	"math"
)

type Matrix struct {
	M         [][]float64
	H         int64
	W         int64
	quadratic bool
	maxValue  float64
}

func NewM(v ...[]float64) *Matrix {
	m := Matrix{}
	m.Set(v)
	return &m
}

func (m *Matrix) Set(v [][]float64) {
	m.H = int64(len(v))
	for i := 0; i < len(v); i++ {
		if i == 0 {
			m.W = int64(len(v[i]))
		} else {
			if m.W != int64(len(v[i])) {
				panic("error en el tamaño de la matrix")
			}
		}
		for j := 0; j < len(v[i]); j++ {
			value := math.Sqrt(v[i][j] * v[i][j])
			if value > m.maxValue {
				m.maxValue = value
			}
		}
	}
	m.M = v
	m.quadratic = m.H == m.W
}

func (m *Matrix) Print() {
	for i := 0; i < len((*m).M); i++ {
		fmt.Print("[ ")
		for j := 0; j < len((*m).M[i]); j++ {
			if j != 0 {
				fmt.Print(" ")
			}
			fmt.Print((*m).M[i][j])
		}
		fmt.Print(" ]")
		fmt.Println()
	}
	fmt.Println((*m).H, "x", (*m).W)
	fmt.Println()
}
