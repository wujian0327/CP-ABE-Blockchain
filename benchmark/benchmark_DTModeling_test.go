package benchmark

import (
	"CP-ABE-Blockchain/core"
	"testing"
)

const AccessStructure = "((company_A AND repairman) OR (company_B AND consumer) OR admin)"

//24          42031804 ns/op
func BenchmarkCreateDTObject_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		core.CreateDTObject("td_files/td_sample_1kb.json", AccessStructure)
	}
}

//24          47919967 ns/op
func BenchmarkCreateDTObject_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		core.CreateDTObject("td_files/td_sample_2kb.json", AccessStructure)
	}
}

//24          52335821 ns/op
func BenchmarkCreateDTObject_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		core.CreateDTObject("td_files/td_sample_4kb.json", AccessStructure)
	}
}

//21          48832114 ns/op
func BenchmarkCreateDTObject_8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		core.CreateDTObject("td_files/td_sample_8kb.json", AccessStructure)
	}
}

//19          58082558 ns/op
func BenchmarkCreateDTObject_16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		core.CreateDTObject("td_files/td_sample_16kb.json", AccessStructure)
	}
}

//28          44312732 ns/op
func BenchmarkUpdateDTObject_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		core.UpdateDTObject("td_files/td_sample_1kb.json")
	}
}

//21          49022900 ns/op
func BenchmarkUpdateDTObject_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		core.UpdateDTObject("td_files/td_sample_2kb.json")
	}
}

//28          45135311 ns/op
func BenchmarkUpdateDTObject_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		core.UpdateDTObject("td_files/td_sample_4kb.json")
	}
}

//24          41942725 ns/op
func BenchmarkUpdateDTObject_8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		core.UpdateDTObject("td_files/td_sample_8kb.json")
	}
}

//22          52529527 ns/op
func BenchmarkUpdateDTObject_16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		core.UpdateDTObject("td_files/td_sample_16kb.json")
	}
}
