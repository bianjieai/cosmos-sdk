package keyring

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
)

type SignatureAlgo interface {
	Name() hd.PubKeyType
	Derive() hd.DeriveFn
	Generate() hd.GenerateFn
}

func NewSigningAlgoFromString(str string) (SignatureAlgo, error) {
	switch str {
	case string(hd.Secp256k1.Name()):
		return hd.Secp256k1, nil
	case string(hd.Sm2.Name()):
		return hd.Sm2, nil
	default:
		return nil, fmt.Errorf("provided algorithm `%s` is not supported", str)
	}
}

type SigningAlgoList []SignatureAlgo

func (l SigningAlgoList) Contains(algo SignatureAlgo) bool {
	for _, cAlgo := range l {
		if cAlgo.Name() == algo.Name() {
			return true
		}
	}

	return false
}
