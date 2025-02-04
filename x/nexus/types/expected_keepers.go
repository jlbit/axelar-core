package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/tendermint/tendermint/libs/log"

	evm "github.com/axelarnetwork/axelar-core/x/evm/types"
	"github.com/axelarnetwork/axelar-core/x/nexus/exported"
	reward "github.com/axelarnetwork/axelar-core/x/reward/exported"
)

//go:generate moq -out ./mock/expected_keepers.go -pkg mock . Nexus Snapshotter AxelarnetKeeper EVMBaseKeeper RewardKeeper

// Nexus provides functionality to manage cross-chain transfers
type Nexus interface {
	Logger(ctx sdk.Context) log.Logger

	InitGenesis(ctx sdk.Context, genState *GenesisState)
	ExportGenesis(ctx sdk.Context) *GenesisState

	SetParams(ctx sdk.Context, p Params)
	GetParams(ctx sdk.Context) Params

	IsChainActivated(ctx sdk.Context, chain exported.Chain) bool
	ActivateChain(ctx sdk.Context, chain exported.Chain)
	GetChains(ctx sdk.Context) []exported.Chain
	GetChain(ctx sdk.Context, chain string) (exported.Chain, bool)
	IsChainMaintainer(ctx sdk.Context, chain exported.Chain, maintainer sdk.ValAddress) bool
	AddChainMaintainer(ctx sdk.Context, chain exported.Chain, validator sdk.ValAddress) error
	RemoveChainMaintainer(ctx sdk.Context, chain exported.Chain, validator sdk.ValAddress) error
	GetChainMaintainers(ctx sdk.Context, chain exported.Chain) []sdk.ValAddress
	GetChainMaintainerStates(ctx sdk.Context, chain exported.Chain) []MaintainerState
	LatestDepositAddress(c context.Context, req *LatestDepositAddressRequest) (*LatestDepositAddressResponse, error)
	LinkAddresses(ctx sdk.Context, sender exported.CrossChainAddress, recipient exported.CrossChainAddress) error
	DeactivateChain(ctx sdk.Context, chain exported.Chain)
	RegisterFee(ctx sdk.Context, chain exported.Chain, feeInfo exported.FeeInfo) error
	GetFeeInfo(ctx sdk.Context, chain exported.Chain, asset string) (feeInfo exported.FeeInfo, found bool)
}

// Snapshotter provides functionality to the snapshot module
type Snapshotter interface {
	GetOperator(ctx sdk.Context, proxy sdk.AccAddress) sdk.ValAddress
	GetProxy(ctx sdk.Context, operator sdk.ValAddress) (addr sdk.AccAddress, active bool)
}

// StakingKeeper provides functionality to the staking module
type StakingKeeper interface {
	Validator(ctx sdk.Context, addr sdk.ValAddress) stakingtypes.ValidatorI
	PowerReduction(sdk.Context) sdk.Int
	GetLastTotalPower(sdk.Context) sdk.Int
}

// AxelarnetKeeper provides functionality to the axelarnet module
type AxelarnetKeeper interface {
	IsCosmosChain(ctx sdk.Context, chain string) bool
}

// EVMBaseKeeper provides functionality to get evm chain keeper
type EVMBaseKeeper interface {
	ForChain(chain string) evm.ChainKeeper
}

// RewardKeeper provides functionality to get reward keeper
type RewardKeeper interface {
	GetPool(ctx sdk.Context, name string) reward.RewardPool
}
