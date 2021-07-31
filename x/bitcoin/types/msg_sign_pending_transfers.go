package types

import (
	"fmt"

	"github.com/btcsuite/btcutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewSignPendingTransfersRequest - SignPendingTransfersRequest constructor
func NewSignPendingTransfersRequest(sender sdk.AccAddress, keyID string, masterKeyAmount btcutil.Amount) *SignPendingTransfersRequest {
	return &SignPendingTransfersRequest{
		Sender:          sender,
		KeyID:           keyID,
		MasterKeyAmount: masterKeyAmount,
	}
}

// Route returns the route for this message
func (m SignPendingTransfersRequest) Route() string {
	return RouterKey
}

// Type returns the type of the message
func (m SignPendingTransfersRequest) Type() string {
	return "SignPendingTransfers"
}

// ValidateBasic executes a stateless message validation
func (m SignPendingTransfersRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	if m.KeyID == "" {
		return fmt.Errorf("missing key ID")
	}

	if m.MasterKeyAmount < 0 {
		return fmt.Errorf("master key amount must be >= 0")
	}

	return nil
}

// GetSignBytes returns the message bytes that need to be signed
func (m SignPendingTransfersRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners returns the set of signers for this message
func (m SignPendingTransfersRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}