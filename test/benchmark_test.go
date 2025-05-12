package test

import (
	"encoding/json"
	"sync"
	"testing"
	"time"
)

func BenchmarkSyncPool(b *testing.B) {

	b.Run("with sync pool", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			Pool()
		}
	})

	//b.ReportAllocs()
	//b.ResetTimer()
	//for i := 0; i < b.N; i++ {
	//	Pool()
	//}
}

func BenchmarkWOSyncPool(b *testing.B) {
	b.Run("without sync pool", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			WoPool()
		}
	})

}

type syncPool struct {
	Number int    `json:"number"`
	String string `json:"string"`
}

type Object struct {
	s1 string
	s2 string
	t1 time.Time
	t2 time.Time
	i1 int64
	i2 int64
}

func Pool() {
	var arg = &sync.Pool{
		New: func() any {
			return new(Object)
		}}

	//var globalSink []*Object
	for j := 0; j < 10000; j++ {

		o := arg.Get().(*Object)

		o.i1 = int64(j)

		arg.Put(o)

		//globalSink = append(globalSink, o)
	}

	//fmt.Println(globalSink)

	//o := arg.Get().(*Object)
	//fmt.Println(o)
}

func WoPool() {
	//var globalSink []*Object

	for j := 0; j < 10000; j++ {
		o := new(Object)

		o.i1 = int64(j)

		//globalSink = append(globalSink, o)
	}
	//fmt.Println(globalSink)
	//fmt.Println(o)
}

func newObject() *Object {
	return &Object{}
}

func (s *syncPool) Reset() {
	s = nil

}

func SyncPool() {
	var arg = sync.Pool{New: func() any {
		return &syncPool{
			Number: 123123,
			String: "asdasd",
		}
	}}

	arg2 := arg.Get().(*syncPool)

	//fmt.Println(arg2)

	arg2.Reset()

	//fmt.Println(arg2)

	arg.Put(arg2)

	//fmt.Println(arg2)
}

func WOSyncPool() {
	arg := syncPool{
		Number: 123123,
		String: "asdasd",
	}

	marshall, err := json.Marshal(arg)
	if err != nil {
		return
	}
	arg2 := syncPool{}
	err = json.Unmarshal(marshall, &arg2)
	if err != nil {
		return
	}

}
