package main

import (
	"testing"
)

func TestInt2Roman01(t *testing.T) {
	roman := int_to_roman(1990)
	if roman != "MCMXC" {
		t.Errorf("Test01 is failed, error output : %s", roman)
	}
}

func TestInt2Roman02(t *testing.T) {
	roman := int_to_roman(4999)
	if roman != "MMMMCMXCIX" {
		t.Errorf("Test02 is failed, error output : %s", roman)
	}
}

func TestInt2Roman03(t *testing.T) {
	number := int_to_roman(888)
	if number != "DCCCLXXXVIII" {
		t.Errorf("Test03 is failed, error output : %s", number)
	}
}
