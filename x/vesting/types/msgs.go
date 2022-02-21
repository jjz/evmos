package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdkvesting "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

var (
	_ sdk.Msg = &MsgCreateClawbackVestingAccount{}
	_ sdk.Msg = &MsgClawback{}
)

const (
	TypeMsgCreateClawbackVestingAccount = "create_clawback_vesting_account"
	TypeMsgClawback                     = "clawback"
)

// NewMsgCreateClawbackVestingAccount creates new instance of MsgCreateClawbackVestingAccount
func NewMsgCreateClawbackVestingAccount(
	fromAddr, toAddr sdk.AccAddress,
	startTime int64,
	lockupPeriods,
	vestingPeriods []sdkvesting.Period,
	merge bool,
) *MsgCreateClawbackVestingAccount {
	return &MsgCreateClawbackVestingAccount{
		FromAddress:    fromAddr.String(),
		ToAddress:      toAddr.String(),
		StartTime:      startTime,
		LockupPeriods:  lockupPeriods,
		VestingPeriods: vestingPeriods,
		Merge:          merge,
	}
}

// Route returns the name of the module
func (msg MsgCreateClawbackVestingAccount) Route() string { return RouterKey }

// Type returns the the action
func (msg MsgCreateClawbackVestingAccount) Type() string { return TypeMsgCreateClawbackVestingAccount }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateClawbackVestingAccount) ValidateBasic() error {
	from, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return sdkerrors.Wrapf(err, "invalid from address %s", from)
	}

	to, err := sdk.AccAddressFromBech32(msg.ToAddress)
	if err != nil {
		return sdkerrors.Wrapf(err, "invalid to address %s", to)
	}

	lockupCoins := sdk.NewCoins()
	for i, period := range msg.LockupPeriods {
		if period.Length < 1 {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid period length of %d in period %d, length must be greater than 0", period.Length, i)
		}
		lockupCoins = lockupCoins.Add(period.Amount...)
	}

	vestingCoins := sdk.NewCoins()
	for i, period := range msg.VestingPeriods {
		if period.Length < 1 {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid period length of %d in period %d, length must be greater than 0", period.Length, i)
		}
		vestingCoins = vestingCoins.Add(period.Amount...)
	}

	// If both schedules are present, the must describe the same total amount.
	// IsEqual can panic, so use (a == b) <=> (a <= b && b <= a).
	if len(msg.LockupPeriods) > 0 && len(msg.VestingPeriods) > 0 &&
		!(lockupCoins.IsAllLTE(vestingCoins) && vestingCoins.IsAllLTE(lockupCoins)) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "vesting and lockup schedules must have same total coins")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateClawbackVestingAccount) GetSignBytes() []byte {
	return sdk.MustSortJSON(amino.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateClawbackVestingAccount) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil
	}
	return []sdk.AccAddress{from}
}

// NewMsgClawbackcreates new instance of MsgClawback. The dest_address may be
// nil - defaulting to the funder.
func NewMsgClawback(funder, addr, dest sdk.AccAddress) *MsgClawback {
	destString := ""
	if dest != nil {
		destString = dest.String()
	}
	return &MsgClawback{
		FunderAddress: funder.String(),
		Address:       addr.String(),
		DestAddress:   destString,
	}
}

// Route returns the message route for a MsgClawback.
func (msg MsgClawback) Route() string { return RouterKey }

// Type returns the message type for a MsgClawback.
func (msg MsgClawback) Type() string { return TypeMsgClawback }

// ValidateBasic runs stateless checks on the message
func (msg MsgClawback) ValidateBasic() error {
	funder, err := sdk.AccAddressFromBech32(msg.GetFunderAddress())
	if err != nil {
		return sdkerrors.Wrapf(err, "invalid funder address %s", funder)
	}

	addr, err := sdk.AccAddressFromBech32(msg.GetAddress())
	if err != nil {
		return sdkerrors.Wrapf(err, "invalid addr address %s", addr)
	}

	if msg.GetDestAddress() != "" {
		dest, err := sdk.AccAddressFromBech32(msg.GetDestAddress())
		if err != nil {
			return sdkerrors.Wrapf(err, "invalid dest address %s", dest)
		}
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgClawback) GetSignBytes() []byte {
	return sdk.MustSortJSON(amino.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgClawback) GetSigners() []sdk.AccAddress {
	funder, err := sdk.AccAddressFromBech32(msg.FunderAddress)
	if err != nil {
		return nil
	}
	return []sdk.AccAddress{funder}
}