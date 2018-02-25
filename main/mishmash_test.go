package main

import (
	"testing"
)

func TestMishMash_Mish(t *testing.T) {
	if mishMash(2) != "Mish" {
		t.Error("Expected returned string to be 'Mish'")
	}
}
func TestMishMash_Mash(t *testing.T) {
	if mishMash(6) != "MishMash" {
		t.Error("Expected returned string to be 'MishMash'")
	}
}
func TestMishMash_MishMash(t *testing.T) {
	if mishMash(12) != "MishMash" {
		t.Error("Expected returned string to be 'MishMash'")
	}
}
func TestMishMash_Num(t *testing.T) {
	if mishMash(4) != "Mish" {
		t.Error("Expected returned string to be 'Mish'")
	}
}
