package promos

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "promos/api/promos/promos"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "LastHeader",
					Use:            "last-header [height]",
					Short:          "Query last-header",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "height"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "InitLightClient",
					Use:            "init-light-client [height] [light-block]",
					Short:          "Send a init-light-client tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "height"}, {ProtoField: "lightBlock"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
