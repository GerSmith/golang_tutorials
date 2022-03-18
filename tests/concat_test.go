/*
	Тесты производительности при конкатенации строк
*/

package main

import (
	"bytes"
	"strings"
	"testing"
)

const ConcatSteps = 1_000_000

func BenchmarkPlus(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var str string
		for i := 0; i < ConcatSteps; i++ {
			str += "x"
		}
	}
}

func BenchmarkBytes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buffer bytes.Buffer
		for i := 0; i < ConcatSteps; i++ {
			buffer.WriteString("x")
		}
	}
}

func BenchmarkStrings(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var str strings.Builder
		for i := 0; i < ConcatSteps; i++ {
			str.WriteString("x")
		}
	}
}
