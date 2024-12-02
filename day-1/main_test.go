package main

import "testing"

func BenchmarkFindScore(b *testing.B) {
	fileName := "input.txt"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {

		_ = findTotalScore(fileName)
	}
}
