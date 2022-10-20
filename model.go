package gomatrix

import (
	"fmt"
	"math"
	"math/rand"
)

type Matrix struct {
	N int // matrix size
	M int // matrix size

	mat [][]float32
}

func Sigmoid(x float32)float32{
	return float32(1)/float32(1+math.Exp(float64(x)))
}

func NewMatrix(n, m int) *Matrix {
	mat := new(Matrix)
	mat.N = n
	mat.M = m
	mat.mat = make([][]float32, m)
	for i := 0; i < m; i++ {
		mat.mat[i] = make([]float32, n)
	}
	return mat
}

func (m *Matrix)Multiplication(a *Matrix)bool{
	if m.N != a.N || m.M != a.M{
		return false
	}
	for i,_:=range m.mat{
		for j,_:=range m.mat[i]{
			m.mat[i][j] *= a.mat[i][j]
		}
	}
	return true
}

func (m *Matrix) Scalar(k float32) {
	for i,_:=range m.mat {
		for j,_:=range m.mat[i] {
			m.mat[i][j] *= k
		}
	}
}

func (m *Matrix) ScalarPlus(k float32) {
	for i,_:=range m.mat {
		for j,_:=range m.mat[i] {
			m.mat[i][j] += k
		}
	}
}

func (m *Matrix) ScalarDivision(k float32) {
	for i,_:=range m.mat {
		for j,_:=range m.mat[i] {
			m.mat[i][j] /= k
		}
	}
}

func (m *Matrix) Plus(mm *Matrix) {
	if m.M == mm.M && m.N == mm.N {
	for i,_:=range m.mat {
		for j,_:=range m.mat[i] {
				m.mat[i][j] += mm.mat[i][j]
			}
		}
	}
}

func (m *Matrix) Minus(mm *Matrix) {
	if m.M == mm.M && m.N == mm.N {
		for i,_:=range m.mat {
			for j,_:=range m.mat[i] {
				m.mat[i][j] -= mm.mat[i][j]
			}
		}
	}
}

func (m *Matrix) Sigmoid() {
	for i,_:=range m.mat {
		for j,_:=range m.mat[i] {
			m.mat[i][j] = Sigmoid(m.mat[i][j])
		}
	}
}

func (m *Matrix) Print() {
	for i,_ := range m.mat {
		fmt.Println(m.mat[i])
	}
}

func (m *Matrix) Put(i,j int,x float32) {
	m.mat[j][i] = x
}

func (m *Matrix) Get(i,j int) float32 {
	return m.mat[j][i]
}

const RAND = 10000

func NewRandom05to05Matrix(n,m int)*Matrix{
	mat := NewMatrix(n,m)
	for i,_:=range mat.mat{
		for j,_:=range mat.mat[i]{
			mat.mat[i][j] = float32(rand.Intn(RAND)) / float32(RAND) - 0.5
		}
	}
	return mat
}

func NewDotMatrix (a , b *Matrix)*Matrix{
	m := NewMatrix(b.N,a.M)
	if a.N != b.M {
		fmt.Println("not",a.N,a.M,b.N,b.M)
		return NewMatrix(0,0)
	} else {
		fmt.Println("ok",a.N,a.M,b.N,b.M)
	}
	b.Print()
	for i:=0;i<a.M;i++ {
		for j:=0;j<b.N;j++ {
			for k:=0;k<a.N;k++{
				fmt.Println(j,k)
				fmt.Println(len(a.mat))
				aa := a.mat[k][i]
				bb:=b.mat[j][k]
				o:= aa * bb
				m.mat[j][i] = o
			}
		}
	}
	return m
}

func NewCopiedMatrix(mm *Matrix)*Matrix{
	m := NewMatrix(mm.N,mm.M)
	for i,_:=range m.mat{
		for j,_:=range m.mat[i]{
			m.mat[i][j] = mm.mat[i][j]
		}
	}
	return m
}

func NewCopyMatrix(a *Matrix)*Matrix{
	m := NewMatrix(a.N,a.M)
	m.Plus(a)
	return m
}

func (m *Matrix)GetT()*Matrix{
	t := NewMatrix(m.M,m.N)
	for i,_:=range m.mat{
		for j,_:=range m.mat[i]{
			t.mat[j][i] = m.mat[i][j]
		}
	}
	return t
}