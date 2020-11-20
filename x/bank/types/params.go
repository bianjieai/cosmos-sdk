package types

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	// DefaultSendEnabled enabled
	DefaultSendEnabled = true

	// DefaultUnrestrictedTokenTransfer is set to false, which
	// means that either of two parties of the transfer must be the token owner
	DefaultUnrestrictedTokenTransfer = false
)

var (
	// KeySendEnabled is store's key for SendEnabled Params
	KeySendEnabled = []byte("SendEnabled")
	// KeyDefaultSendEnabled is store's key for the DefaultSendEnabled option
	KeyDefaultSendEnabled = []byte("DefaultSendEnabled")
	// KeyUnrestrictedTokenTransfer is store's key for the UnrestrictedTokenTransfer param
	KeyUnrestrictedTokenTransfer = []byte("UnrestrictedTokenTransfer")
)

// ParamKeyTable for bank module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new parameter configuration for the bank module
func NewParams(
	defaultSendEnabled bool,
	sendEnabledParams SendEnabledParams,
	unrestrictedTokenTransfer bool,
) Params {
	return Params{
		SendEnabled:               sendEnabledParams,
		DefaultSendEnabled:        defaultSendEnabled,
		UnrestrictedTokenTransfer: unrestrictedTokenTransfer,
	}
}

// DefaultParams is the default parameter configuration for the bank module
func DefaultParams() Params {
	return Params{
		SendEnabled: SendEnabledParams{},
		// The default send enabled value allows send transfers for all coin denoms
		DefaultSendEnabled:        true,
		UnrestrictedTokenTransfer: DefaultUnrestrictedTokenTransfer,
	}
}

// Validate all bank module parameters
func (p Params) Validate() error {
	if err := validateSendEnabledParams(p.SendEnabled); err != nil {
		return err
	}

	if err := validateIsBool(p.DefaultSendEnabled); err != nil {
		return err
	}

	return validateIsBool(p.UnrestrictedTokenTransfer)
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// SendEnabledDenom returns true if the given denom is enabled for sending
func (p Params) SendEnabledDenom(denom string) bool {
	for _, pse := range p.SendEnabled {
		if pse.Denom == denom {
			return pse.Enabled
		}
	}
	return p.DefaultSendEnabled
}

// SetSendEnabledParam returns an updated set of Parameters with the given denom
// send enabled flag set.
func (p Params) SetSendEnabledParam(denom string, sendEnabled bool) Params {
	var sendParams SendEnabledParams
	for _, p := range p.SendEnabled {
		if p.Denom != denom {
			sendParams = append(sendParams, NewSendEnabled(p.Denom, p.Enabled))
		}
	}
	sendParams = append(sendParams, NewSendEnabled(denom, sendEnabled))
	return NewParams(p.DefaultSendEnabled, sendParams, p.UnrestrictedTokenTransfer)
}

// ParamSetPairs implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeySendEnabled, &p.SendEnabled, validateSendEnabledParams),
		paramtypes.NewParamSetPair(KeyDefaultSendEnabled, &p.DefaultSendEnabled, validateIsBool),
		paramtypes.NewParamSetPair(KeyUnrestrictedTokenTransfer, &p.UnrestrictedTokenTransfer, validateIsBool),
	}
}

// SendEnabledParams is a collection of parameters indicating if a coin denom is enabled for sending
type SendEnabledParams []*SendEnabled

func validateSendEnabledParams(i interface{}) error {
	params, ok := i.([]*SendEnabled)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	// ensure each denom is only registered one time.
	registered := make(map[string]bool)
	for _, p := range params {
		if _, exists := registered[p.Denom]; exists {
			return fmt.Errorf("duplicate send enabled parameter found: '%s'", p.Denom)
		}
		if err := validateSendEnabled(*p); err != nil {
			return err
		}
		registered[p.Denom] = true
	}
	return nil
}

// NewSendEnabled creates a new SendEnabled object
// The denom may be left empty to control the global default setting of send_enabled
func NewSendEnabled(denom string, sendEnabled bool) *SendEnabled {
	return &SendEnabled{
		Denom:   denom,
		Enabled: sendEnabled,
	}
}

// String implements stringer insterface
func (se SendEnabled) String() string {
	out, _ := yaml.Marshal(se)
	return string(out)
}

func validateSendEnabled(i interface{}) error {
	param, ok := i.(SendEnabled)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return sdk.ValidateDenom(param.Denom)
}

func validateIsBool(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}
