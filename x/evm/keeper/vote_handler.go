package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/axelarnetwork/axelar-core/x/evm/types"
	nexus "github.com/axelarnetwork/axelar-core/x/nexus/exported"
	vote "github.com/axelarnetwork/axelar-core/x/vote/exported"
	"github.com/axelarnetwork/utils/slices"
)

// NewVoteHandler returns the handler for processing vote delivered by the vote module
func NewVoteHandler(cdc codec.Codec, keeper types.BaseKeeper, nexus types.Nexus) vote.VoteHandler {
	return func(ctx sdk.Context, result *vote.Vote) error {
		events, err := types.UnpackEvents(cdc, result.Results)
		if err != nil {
			return err
		}

		if len(events) == 0 {
			return fmt.Errorf("vote result has no events")
		}

		chainName := events[0].Chain
		if slices.Any(events, func(event types.Event) bool { return event.Chain != chainName }) {
			return fmt.Errorf("events are not from the same source chain")
		}

		chain, ok := nexus.GetChain(ctx, chainName)
		if !ok {
			return fmt.Errorf("%s is not a registered chain", chainName)
		}
		if !keeper.HasChain(ctx, chainName) {
			return fmt.Errorf("%s is not an evm chain", chainName)
		}

		chainK := keeper.ForChain(chain.Name)
		cacheCtx, writeCache := ctx.CacheContext()

		err = handleEvents(cacheCtx, chainK, events, chain)
		if err != nil {
			// set events to failed, we will deal with later
			for _, e := range events {
				chainK.SetFailedEvent(ctx, e)
			}
			return err
		}

		writeCache()
		ctx.EventManager().EmitEvents(cacheCtx.EventManager().Events())
		return nil
	}
}

func handleEvents(ctx sdk.Context, ck types.ChainKeeper, events []types.Event, chain nexus.Chain) error {
	for _, event := range events {
		var err error
		// validate event
		err = event.ValidateBasic()
		if err != nil {
			return fmt.Errorf("event %s: %s", event.GetID(), err.Error())
		}

		// check if event confirmed before
		eventID := event.GetID()
		if _, ok := ck.GetEvent(ctx, eventID); ok {
			return fmt.Errorf("event %s is already confirmed", eventID)
		}
		ck.SetConfirmedEvent(ctx, event)
		ck.Logger(ctx).Info(fmt.Sprintf("confirmed %s event %s in transaction %s", chain.Name, eventID, event.TxId.Hex()))

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(types.EventTypeEventConfirmation,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
				sdk.NewAttribute(types.AttributeKeyChain, event.Chain),
				sdk.NewAttribute(types.AttributeKeyTxID, event.TxId.Hex()),
				sdk.NewAttribute(types.AttributeKeyEventID, event.GetID()),
				sdk.NewAttribute(types.AttributeKeyEventType, event.GetEventType()),
				sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeValueConfirm)),
		)
	}

	return nil
}
