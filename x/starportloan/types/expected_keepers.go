package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type BankKeeper interface {
	// Methods imported from bank should be defined here
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}
