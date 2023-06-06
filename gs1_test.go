package uniques

import (
	"testing"

	"github.com/rashintha/uniques/standards/gs1"
)

func TestEPCtoUPC(t *testing.T) {
	result, err := gs1.EPCtoUPC("30340bdf184af9995006bc04")

	if result == 0 && err != nil || result != 194502767742 {
		t.Errorf("Error in passing test %v", err)
	}

}
