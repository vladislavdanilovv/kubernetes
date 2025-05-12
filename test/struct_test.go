package test

import (
	"testing"
)

type sortStruct struct {
	i1, i2, i3, i4 int
	b1, b2, b3, b4 bool
}

type randomlyStruct struct {
	i1 int
	b1 bool
	i2 int
	b2 bool
	i3 int
	b3 bool
	i4 int
	b4 bool
}

func BenchmarkStructSort(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sortStructFunc(i)
	}
}

func BenchmarkStructRandomly(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		randomlyStructFunc(i)
	}
}

type structI interface {
	set(i int)
}

func asd(inter structI, i int) {
	inter.set(i)
}

func sortStructFunc(i int) {
	var arg structI = sortStruct{}

	arg.set(i)

	//fmt.Println(i)
}

func randomlyStructFunc(i int) {
	var arg structI = randomlyStruct{}

	arg.set(i)

	//fmt.Println(i)
}

func (s sortStruct) set(i int) {
	s.i1 = i
	s.i2 = i
	s.i3 = i
	s.i4 = i

	s.b1 = true
	s.b2 = true
	s.b3 = true
	s.b4 = true

	//return sortStruct
}

func (s randomlyStruct) set(i int) {

	s.i1 = i
	s.i2 = i
	s.i3 = i
	s.i4 = i

	s.b1 = true
	s.b2 = true
	s.b3 = true
	s.b4 = true
	//return sortStruct
}
