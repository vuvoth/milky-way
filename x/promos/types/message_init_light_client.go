package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgInitLightClient{}

func NewMsgInitLightClient(creator string, height uint64, lightBlock string) *MsgInitLightClient {
	return &MsgInitLightClient{
		Creator:    creator,
		Height:     height,
		LightBlock: lightBlock,
	}
}

func (msg *MsgInitLightClient) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
