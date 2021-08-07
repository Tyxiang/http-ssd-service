package main

import (
	"fmt"
	"pool"
	"testing"
)

func init() {
	fmt.Println(pool.AddTable("1234"))
	//fmt.Println(pool.AddTable("1234", "1", 3))
	//fmt.Println(pool.CopyTable("1234", "1", []byte{31, 32, 33, 34}))
	fmt.Println(pool.Tables)
}

func BenchmarkAddTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool.AddTable("1234")
	}
}

func BenchmarkReadTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool.ReadTable("1234")
	}
}

func BenchmarkCopyTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool.CopyTable("1234", "5678")
	}
}

func BenchmarkDeleteTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool.DeleteTable("1234")
	}
}
