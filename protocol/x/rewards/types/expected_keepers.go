package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	assets "github.com/dydxprotocol/v4-chain/protocol/x/assets/types"
	prices "github.com/dydxprotocol/v4-chain/protocol/x/prices/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}

type FeeTiersKeeper interface {
	GetLowestMakerFee(ctx sdk.Context) int32
}

type PricesKeeper interface {
	GetMarketPrice(
		ctx sdk.Context,
		id uint32,
	) (market prices.MarketPrice, err error)
}

type AssetsKeeper interface {
	GetAsset(
		ctx sdk.Context,
		id uint32,
	) (val assets.Asset, err error)
}
