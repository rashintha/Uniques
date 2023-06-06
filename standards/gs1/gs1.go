package gs1

import (
	"errors"
	"math/big"
)

func EPCtoUPC(epc string) (int64, error) {
	val, success := hexToBigInt(epc)
	if !success {
		return 0, errors.New("failed to convert hex string to big.Int")
	}

	totalBits := len(epc) * 4
	leadingZeros := totalBits - val.BitLen()

	// Decoding the header
	header := new(big.Int).Set(val)
	shift := header.BitLen() - 8 + leadingZeros
	header.Rsh(header, uint(shift))
	mask := new(big.Int)
	mask.SetString("FF", 16)
	header.And(header, mask)

	// If SGTIN-96
	if header.Cmp(new(big.Int).SetInt64(48)) == 0 {
		result, err := sgtin96Decode(epc)
		return result.UPC, err
	}

	return 0, errors.New("this decoder is not yet developed")
}

func hexToBigInt(hexStr string) (*big.Int, bool) {
	val := new(big.Int)
	_, success := val.SetString(hexStr, 16)
	return val, success
}
