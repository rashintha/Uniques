package gs1

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

type SGTIN96 struct {
	Filter                      int64
	Partition                   int64
	Gs1CompanyPrefix            int64
	IndicatorDigitItemRefNumber int64
	SerialNumber                int64
	UPC                         int64
}

func sgtin96Decode(epc *string) (*SGTIN96, error) {
	val, success := hexToBigInt(*epc)
	if !success {
		return &SGTIN96{}, errors.New("failed to convert hex string to big.Int")
	}

	// Decoding the filter
	filter := new(big.Int).Set(val)
	shift := 85
	filter.Rsh(filter, uint(shift))
	mask := new(big.Int)
	mask.SetString("7", 16)
	filter.And(filter, mask)

	// Decoding the partition
	partition := new(big.Int).Set(val)
	shift = 82
	partition.Rsh(partition, uint(shift))
	mask = new(big.Int)
	mask.SetString("7", 16)
	partition.And(partition, mask)

	companyPrefixDigitsCount := 40
	indicatorDigitsCount := 4
	maskValCompanyPrefix := "FFFFFFFFFF"
	maskValIndicator := "F"

	switch partition.Int64() {
	case 1:
		companyPrefixDigitsCount = 37
		indicatorDigitsCount = 7
		maskValCompanyPrefix = "1FFFFFFFFF"
		maskValIndicator = "7F"

	case 2:
		companyPrefixDigitsCount = 34
		indicatorDigitsCount = 10
		maskValCompanyPrefix = "3FFFFFFFF"
		maskValIndicator = "3FF"

	case 3:
		companyPrefixDigitsCount = 30
		indicatorDigitsCount = 14
		maskValCompanyPrefix = "3FFFFFFF"
		maskValIndicator = "3FFF"

	case 4:
		companyPrefixDigitsCount = 27
		indicatorDigitsCount = 17
		maskValCompanyPrefix = "7FFFFFF"
		maskValIndicator = "1FFFF"

	case 5:
		companyPrefixDigitsCount = 24
		indicatorDigitsCount = 20
		maskValCompanyPrefix = "FFFFFF"
		maskValIndicator = "FFFFF"

	case 6:
		companyPrefixDigitsCount = 20
		indicatorDigitsCount = 24
		maskValCompanyPrefix = "FFFFF"
		maskValIndicator = "FFFFFF"

	}

	// Decoding GS1 Company Prefix
	gs1CompanyPrefix := new(big.Int).Set(val)
	shift = 82 - companyPrefixDigitsCount
	gs1CompanyPrefix.Rsh(gs1CompanyPrefix, uint(shift))
	mask = new(big.Int)
	mask.SetString(maskValCompanyPrefix, 16)
	gs1CompanyPrefix.And(gs1CompanyPrefix, mask)

	// Decoding Indicator Digit + Item Reference Number
	indicatorItemRefNo := new(big.Int).Set(val)
	shift = 82 - companyPrefixDigitsCount - indicatorDigitsCount
	indicatorItemRefNo.Rsh(indicatorItemRefNo, uint(shift))
	mask = new(big.Int)
	mask.SetString(maskValIndicator, 16)
	indicatorItemRefNo.And(indicatorItemRefNo, mask)

	// Decoding serial number
	mask = new(big.Int)
	mask.SetString("3FFFFFFFFF", 16)
	serialNo := new(big.Int).Set(val)
	serialNo.And(serialNo, mask)

	// Decoding the UPC number
	itemRefNo := fmt.Sprintf("%d", indicatorItemRefNo.Int64())

	if len(itemRefNo) == 6 {
		itemRefNo = itemRefNo[1:]
	}

	itemRefNoInt, err := strconv.Atoi(itemRefNo)
	if err != nil {
		return &SGTIN96{}, err
	}

	upcStr := fmt.Sprintf("%d%05d", gs1CompanyPrefix.Int64(), itemRefNoInt)

	checkDigit, err := calculateUPCCheckDigit(&upcStr)
	if err != nil {
		return &SGTIN96{}, err
	}

	upcStrUpdated := fmt.Sprintf("%s%d", upcStr, checkDigit)

	upc, err := strconv.ParseInt(upcStrUpdated, 10, 64)
	if err != nil {
		return &SGTIN96{}, err
	}

	return &SGTIN96{
		Filter:                      filter.Int64(),
		Partition:                   partition.Int64(),
		Gs1CompanyPrefix:            gs1CompanyPrefix.Int64(),
		IndicatorDigitItemRefNumber: indicatorItemRefNo.Int64(),
		SerialNumber:                serialNo.Int64(),
		UPC:                         int64(upc),
	}, nil
}

func calculateUPCCheckDigit(upc *string) (int, error) {
	if len(*upc) != 11 {
		return 0, fmt.Errorf("UPC must be 11 digits %v", *upc)
	}

	sum := 0
	for i := 0; i < 11; i++ {
		digit, err := strconv.Atoi(string((*upc)[i]))
		if err != nil {
			return 0, err
		}

		if i%2 == 0 {
			sum += digit * 3
		} else {
			sum += digit
		}
	}

	checkDigit := 10 - (sum % 10)
	if checkDigit == 10 {
		checkDigit = 0
	}

	return checkDigit, nil
}
