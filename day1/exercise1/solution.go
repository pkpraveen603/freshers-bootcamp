package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	Row int `json:"NumRows"`
	Col int `json:"NumCols"`
	Elements [][] int `json:"Matrix"`
}
func InitialiseMatrix(rows,cols int) *Matrix{
	matrix:= Matrix{Row: rows,Col: cols,Elements: make([][]int,rows)}
	for r := 0 ; r < rows ; r++ {
		matrix.Elements[r] = make([]int, cols)
		for c := 0; c < cols; c++ {
			matrix.Elements[r][c] = 0
		}
	}
	return &matrix
}
func (m Matrix) rows() int {
	return m.Row
}
func (m Matrix) cols() int {
	return m.Col
}
func (m *Matrix) setval(i, j, element int)  {
	m.Elements[i][j] = element
}
func (m *Matrix) add(n *Matrix) *Matrix {
	res := InitialiseMatrix(m.Row,m.Col)
	if m.Row == n.Row && m.Col == n.Col {
		for i := 0; i < m.Row; i++ {
			for j := 0; j < m.Col; j++ {
				res.Elements[i][j] = m.Elements[i][j] + n.Elements[i][j]
			}
		}
	}
	return res
}

func (m *Matrix)printJSON(){
	matrix,_ := json.Marshal(&m)
	fmt.Println(string(matrix))
}

func main() {
	a := InitialiseMatrix(2,2)
	a.printJSON()
	fmt.Println()
	b := InitialiseMatrix(2,2)
	b.printJSON()
	fmt.Println()
	a.setval(0, 1, 20)
	b.setval(0, 0, 10)
	c := a.add(b)
	c.printJSON()
}