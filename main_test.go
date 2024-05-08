package main

import "testing"

func BenchmarkInsertPath(b *testing.B) {
	b.StopTimer()

	dir := NewTrie()

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		dir.NewDir("user/public/main/")
	}
}

func BenchmarkPrintAllPath(b *testing.B) {
	b.StopTimer()

	dir := NewTrie()

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		dir.PrintDirTree()
	}
}
