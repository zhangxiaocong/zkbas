// Code generated by goctl. DO NOT EDIT!
// Source: globalRPC.proto

package globalrpc

import (
	"context"

	"github.com/zecrey-labs/zecrey-legend/service/rpc/globalRPC/globalRPCProto"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AssetResult                           = globalRPCProto.AssetResult
	ReqGetLatestAssetsListByAccountIndex  = globalRPCProto.ReqGetLatestAssetsListByAccountIndex
	ReqGetLatestPairInfo                  = globalRPCProto.ReqGetLatestPairInfo
	ReqGetLpValue                         = globalRPCProto.ReqGetLpValue
	ReqGetMaxOfferId                      = globalRPCProto.ReqGetMaxOfferId
	ReqGetNextNonce                       = globalRPCProto.ReqGetNextNonce
	ReqGetSwapAmount                      = globalRPCProto.ReqGetSwapAmount
	ReqSendTx                             = globalRPCProto.ReqSendTx
	RespGetLatestAssetsListByAccountIndex = globalRPCProto.RespGetLatestAssetsListByAccountIndex
	RespGetLatestPairInfo                 = globalRPCProto.RespGetLatestPairInfo
	RespGetLpValue                        = globalRPCProto.RespGetLpValue
	RespGetMaxOfferId                     = globalRPCProto.RespGetMaxOfferId
	RespGetNextNonce                      = globalRPCProto.RespGetNextNonce
	RespGetSwapAmount                     = globalRPCProto.RespGetSwapAmount
	RespSendTx                            = globalRPCProto.RespSendTx
	TxDetailInfo                          = globalRPCProto.TxDetailInfo
	TxInfo                                = globalRPCProto.TxInfo

	GlobalRPC interface {
		//  Asset
		GetLatestAssetsListByAccountIndex(ctx context.Context, in *ReqGetLatestAssetsListByAccountIndex, opts ...grpc.CallOption) (*RespGetLatestAssetsListByAccountIndex, error)
		//  Liquidity
		GetLatestPairInfo(ctx context.Context, in *ReqGetLatestPairInfo, opts ...grpc.CallOption) (*RespGetLatestPairInfo, error)
		GetSwapAmount(ctx context.Context, in *ReqGetSwapAmount, opts ...grpc.CallOption) (*RespGetSwapAmount, error)
		GetLpValue(ctx context.Context, in *ReqGetLpValue, opts ...grpc.CallOption) (*RespGetLpValue, error)
		//  Transaction
		SendTx(ctx context.Context, in *ReqSendTx, opts ...grpc.CallOption) (*RespSendTx, error)
		GetNextNonce(ctx context.Context, in *ReqGetNextNonce, opts ...grpc.CallOption) (*RespGetNextNonce, error)
		//  NFT
		GetMaxOfferId(ctx context.Context, in *ReqGetMaxOfferId, opts ...grpc.CallOption) (*RespGetMaxOfferId, error)
	}

	defaultGlobalRPC struct {
		cli zrpc.Client
	}
)

func NewGlobalRPC(cli zrpc.Client) GlobalRPC {
	return &defaultGlobalRPC{
		cli: cli,
	}
}

//  Asset
func (m *defaultGlobalRPC) GetLatestAssetsListByAccountIndex(ctx context.Context, in *ReqGetLatestAssetsListByAccountIndex, opts ...grpc.CallOption) (*RespGetLatestAssetsListByAccountIndex, error) {
	client := globalRPCProto.NewGlobalRPCClient(m.cli.Conn())
	return client.GetLatestAssetsListByAccountIndex(ctx, in, opts...)
}

//  Liquidity
func (m *defaultGlobalRPC) GetLatestPairInfo(ctx context.Context, in *ReqGetLatestPairInfo, opts ...grpc.CallOption) (*RespGetLatestPairInfo, error) {
	client := globalRPCProto.NewGlobalRPCClient(m.cli.Conn())
	return client.GetLatestPairInfo(ctx, in, opts...)
}

func (m *defaultGlobalRPC) GetSwapAmount(ctx context.Context, in *ReqGetSwapAmount, opts ...grpc.CallOption) (*RespGetSwapAmount, error) {
	client := globalRPCProto.NewGlobalRPCClient(m.cli.Conn())
	return client.GetSwapAmount(ctx, in, opts...)
}

func (m *defaultGlobalRPC) GetLpValue(ctx context.Context, in *ReqGetLpValue, opts ...grpc.CallOption) (*RespGetLpValue, error) {
	client := globalRPCProto.NewGlobalRPCClient(m.cli.Conn())
	return client.GetLpValue(ctx, in, opts...)
}

//  Transaction
func (m *defaultGlobalRPC) SendTx(ctx context.Context, in *ReqSendTx, opts ...grpc.CallOption) (*RespSendTx, error) {
	client := globalRPCProto.NewGlobalRPCClient(m.cli.Conn())
	return client.SendTx(ctx, in, opts...)
}

func (m *defaultGlobalRPC) GetNextNonce(ctx context.Context, in *ReqGetNextNonce, opts ...grpc.CallOption) (*RespGetNextNonce, error) {
	client := globalRPCProto.NewGlobalRPCClient(m.cli.Conn())
	return client.GetNextNonce(ctx, in, opts...)
}

//  NFT
func (m *defaultGlobalRPC) GetMaxOfferId(ctx context.Context, in *ReqGetMaxOfferId, opts ...grpc.CallOption) (*RespGetMaxOfferId, error) {
	client := globalRPCProto.NewGlobalRPCClient(m.cli.Conn())
	return client.GetMaxOfferId(ctx, in, opts...)
}
