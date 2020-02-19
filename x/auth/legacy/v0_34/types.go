// DONTCOVER
// nolint
package v0_34

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "auth"
)

type (
	Params struct {
		MaxMemoCharacters      uint64 `json:"max_memo_characters"`
		TxSigLimit             uint64 `json:"tx_sig_limit"`
		TxSizeCostPerByte      uint64 `json:"tx_size_cost_per_byte"`
		SigVerifyCostSm2       uint64 `json:"sig_verify_cost_sm2"`
		SigVerifyCostSecp256k1 uint64 `json:"sig_verify_cost_secp256k1"`
	}

	GenesisState struct {
		CollectedFees sdk.Coins `json:"collected_fees"`
		Params        Params    `json:"params"`
	}
)
