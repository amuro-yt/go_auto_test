package main

import "testing"

// 偶数かどうかを判定
func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "even" {
		t.Errorf("expected: even, actual: %s", result)
	}
}
