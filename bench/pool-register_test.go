package main

import (
	"fmt"
	"pool"
	"testing"
)

func init() {
	fmt.Println(pool.AddTable("1234"))
	fmt.Println(pool.AddRegister("1234", "1", 3))
	//fmt.Println(pool.WriteRegister("1234", "1", []byte{31, 32, 33, 34}))
	fmt.Println(pool.Registers)
}

func BenchmarkAddRegister(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool.AddRegister("1234", "2", 10)
	}
}

func BenchmarkReadRegister(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool.ReadRegister("1234", "1")
	}
}

func BenchmarkWriteRegister(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool.WriteRegister("1234", "1", []byte{31, 32, 33})
	}
}

func BenchmarkDeleteRegister(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool.DeleteRegister("1234", "2")
	}
}

func BenchmarkExportRegisterFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool.ExportRegisterFile("1234", "2", "./")
	}
}

func BenchmarkImportRegisterFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool.ImportRegisterFile("1234", "2", "./")
	}
}

func BenchmarkRemoveRegisterFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool.RemoveRegisterFile("1234", "2", "./")
	}
}

func main() {
	//fmt.Println(pool.Registers)
}
