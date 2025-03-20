package encoding

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/hetu-project/hetu-checkpoint/proto/types"
)

// RegisterInterfaces registers the tendermint concrete client-related
// implementations and interfaces.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.AccountI)(nil),
		&types.EthAccount{},
	)
	registry.RegisterImplementations(
		(*authtypes.GenesisAccount)(nil),
		&types.EthAccount{},
	)
	registry.RegisterImplementations(
		(*tx.TxExtensionOptionI)(nil),
		&types.ExtensionOptionsWeb3Tx{},
		&types.ExtensionOptionDynamicFeeTx{},
	)
	// regist checkpointing interfaces
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&types.MsgRegistValidator{},
	)
}
