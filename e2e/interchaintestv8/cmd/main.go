package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	FeeDenom            = "ulom"
	Bech32PrefixAccAddr = "lom"
	Bech32PrefixAccPub  = "lompub"

	FlagEthRPC    = "eth-rpc"
	DefaultEthRPC = "http://localhost:8545"

	FlagIcs26Address    = "ics26-address"
	DefaultIcs26Address = "0xa513e6e4b8f2a923d98304ec87f64353c4d5c853"

	FlagIcs20Address    = "ics20-address"
	DefaultIcs20Address = "0x2279b7a0a67db372996a5fab50d91eaa73d2ebe6"

	FlagErc20Address    = "erc20-address"
	DefaultErc20Address = "0xb7f8bc63bbcad18155201308c8f3540b07f84f5e"

	FlagCosmosRPC    = "cosmos-rpc"
	DefaultCosmosRPC = "http://localhost:26657"

	FlagCosmosGRPC    = "cosmos-grpc"
	DefaultCosmosGRPC = "localhost:9090"

	FlagCosmosChainID    = "cosmos-chain-id"
	DefaultCosmosChainID = "localnet-1"

	FlagEthChainID    = "ethereum-chain-id"
	DefaultEthChainID = "11155111"

	FlagSourceClientID      = "source-client-id"
	FlagCosmosClientIDOnEth = "client-id-on-eth"
	FlagEthClientIDOnCosmos = "client-id-on-cosmos"

	// TODO: Add the non-mock versions of these
	MockTendermintClientID = "cosmoshub-1"
	MockEthClientID        = "08-wasm-0"

	EnvEthPrivateKey    = "ETH_PRIVATE_KEY"
	EnvCosmosPrivateKey = "COSMOS_PRIVATE_KEY"

	RelayerURL = "localhost:3000"

	EnvRelayerWallet = "RELAYER_WALLET"

	FlagVerbose = "verbose"

	FlagTransferWithCallbacksMemo = "transfer-with-callbacks-memo"
)

func main() {
	if err := RootCmd().Execute(); err != nil {
		fmt.Println("Something went wrong!")
		fmt.Printf("Error: %+v\n", err)
	}
}

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "eureka-cli",
		Short: "IBC Eureka CLI",
	}

	rootCmd.AddCommand(TransferFromEth())
	rootCmd.AddCommand(RelayTxCmd())
	rootCmd.AddCommand(BalanceCmd())
	rootCmd.AddCommand(TransferFromCosmos())

	rootCmd.PersistentFlags().BoolP(FlagVerbose, "v", false, "verbose output")

	return rootCmd
}

func AddEthFlags(cmd *cobra.Command) {
	cmd.Flags().String(FlagEthRPC, DefaultEthRPC, "Ethereum RPC URL")
	cmd.Flags().String(FlagIcs26Address, DefaultIcs26Address, "ICS26 contract address")
	cmd.Flags().String(FlagIcs20Address, DefaultIcs20Address, "ICS20 contract address")
	cmd.Flags().String(FlagErc20Address, DefaultErc20Address, "ERC20 contract address")
}

func AddCosmosFlags(cmd *cobra.Command) {
	cmd.Flags().String(FlagCosmosRPC, DefaultCosmosRPC, "Cosmos RPC URL")
	cmd.Flags().String(FlagCosmosGRPC, DefaultCosmosGRPC, "Cosmos gRPC URL")
	cmd.Flags().String(FlagCosmosChainID, DefaultCosmosChainID, "Cosmos Chain ID")
}

func IsVerbose(cmd *cobra.Command) bool {
	verbose, _ := cmd.Flags().GetBool(FlagVerbose)
	return verbose
}
