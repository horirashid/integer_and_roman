package main

import (
	"testing"
)

func TestRoman2Int01(t *testing.T) {
	number := RomanString2Int("MCMXC")
	if number != 1990 {
		t.Errorf("Test01 is failed, error output : %d", number)
	}
}

func TestRoman2Int02(t *testing.T) {
	number := RomanString2Int("MMMMCMXCIX")
	if number != 4999 {
		t.Errorf("Test02 is failed, error output : %d", number)
	}
}

func TestRoman2Int03(t *testing.T) {
	number := RomanString2Int("DCCCLXXXVIII")
	if number != 888 {
		t.Errorf("Test02 is failed, error output : %d", number)
	}
}
