package uniques

import (
	"fmt"
	"testing"

	"github.com/rashintha/uniques/standards/gs1"
)

func TestEPCtoUPC(t *testing.T) {
	fmt.Println("Running test: EPC to UPC")
	epc := "30340bdf184af9995006bc04"
	result, err := gs1.EPCtoUPC(&epc)

	if result == 0 && err != nil || result != 194502767742 {
		t.Errorf("Error in passing test %v", err)
		return
	}

	fmt.Printf("Test Passed: EPC to UPC\n\n")
}

func TestEPCstoUPCs(t *testing.T) {
	fmt.Println("Running test: EPCs to UPCs")

	epcs := []string{
		"30340BDF183B02B5EC93A804",
		"30340BDF1849BD7CD9699804",
		"30340BDF1849BD7A9AD9C004",
		"30340BDF1458AFD93799D004",
		"30340BDF1458AFC0986E6804",
		"30340BDF1458AFE485D73C04",
		"30340BEAB035F3EBF2587404",
		"30340BF8E83C7289D4E58804",
		"30340BF8E83C73D679DEDC04",
		"30340C05C004E25A8ACDF804",
		"30340BF8E83C765B5F969404",
		"30340BF8E83C78F7EF41A004",
		"30340C00040CA2A523407804",
		"30340BC9E821A72B6ED47404",
		"30340C05A017BBA4F6F7FC04",
	}

	upcs := []int64{
		194502604269,
		194502755091,
		194502755091,
		194501908153,
		194501908153,
		194501908153,
		195244552474,
		196154618984,
		196154619035,
		196976050016,
		196154619134,
		196154619233,
		196609129386,
		193146344609,
		196968243020,
	}

	for index, epc := range epcs {
		upc, err := gs1.EPCtoUPC(&epc)

		if err != nil || upc != upcs[index] {
			t.Errorf("Error passing the test %v", err)
			return
		}
	}

	fmt.Printf("Test Passed: EPCs to UPCs\n\n")
}
