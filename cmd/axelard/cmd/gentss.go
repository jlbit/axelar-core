package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/spf13/cobra"

	tssTypes "github.com/axelarnetwork/axelar-core/x/tss/types"
)

const (
	flagHeartbeatPeriod       = "heartbeat-period"
	flagSignedBlocksPerWindow = "signed-blocks-window"
)

// SetGenesisTSSCmd returns set-genesis-chain-params cobra Command.
func SetGenesisTSSCmd(defaultNodeHome string,
) *cobra.Command {
	var (
		heartbeatPeriod       int64
		signedBlocksPerWindow string
	)

	cmd := &cobra.Command{
		Use:   "set-genesis-tss",
		Short: "Set the genesis parameters for the tss module",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			cdc := clientCtx.Codec

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			config.SetRoot(clientCtx.HomeDir)

			genFile := config.GenesisFile()
			appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
			if err != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}
			genesisTSS := tssTypes.GetGenesisStateFromAppState(cdc, appState)

			if heartbeatPeriod > 0 {
				genesisTSS.Params.HeartbeatPeriodInBlocks = heartbeatPeriod
			}

			if signedBlocksPerWindow != "" {
				threshold, err := parseThreshold(signedBlocksPerWindow)
				if err != nil {
					return err
				}

				genesisTSS.Params.MaxMissedBlocksPerWindow = threshold
			}

			genesisTSSBz, err := cdc.MarshalJSON(&genesisTSS)
			if err != nil {
				return fmt.Errorf("failed to marshal tss genesis state: %w", err)
			}

			appState[tssTypes.ModuleName] = genesisTSSBz

			appStateJSON, err := json.Marshal(appState)
			if err != nil {
				return fmt.Errorf("failed to marshal application genesis state: %w", err)
			}
			genDoc.AppState = appStateJSON

			return genutil.ExportGenesisFile(genDoc, genFile)
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "node's home directory")
	cmd.Flags().Int64Var(&heartbeatPeriod, flagHeartbeatPeriod, 0, "time period in blocks for tss to emit the event asking validators to send their heartbeats")
	cmd.Flags().StringVar(&signedBlocksPerWindow, flagSignedBlocksPerWindow, "", "the signed blocks window to be considered when calculating the missed blocks percentage")
	return cmd
}
