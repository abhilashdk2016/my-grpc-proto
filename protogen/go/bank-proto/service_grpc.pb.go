// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: proto/bank/service.proto

package bank_proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	BankService_GetCurrentBalance_FullMethodName     = "/bank.BankService/GetCurrentBalance"
	BankService_FetchExchangeRates_FullMethodName    = "/bank.BankService/FetchExchangeRates"
	BankService_SummarizeTransactions_FullMethodName = "/bank.BankService/SummarizeTransactions"
)

// BankServiceClient is the client API for BankService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BankServiceClient interface {
	GetCurrentBalance(ctx context.Context, in *CurrentBalanceRequest, opts ...grpc.CallOption) (*CurrentBalanceResponse, error)
	FetchExchangeRates(ctx context.Context, in *ExchangeRateRequest, opts ...grpc.CallOption) (BankService_FetchExchangeRatesClient, error)
	SummarizeTransactions(ctx context.Context, opts ...grpc.CallOption) (BankService_SummarizeTransactionsClient, error)
}

type bankServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBankServiceClient(cc grpc.ClientConnInterface) BankServiceClient {
	return &bankServiceClient{cc}
}

func (c *bankServiceClient) GetCurrentBalance(ctx context.Context, in *CurrentBalanceRequest, opts ...grpc.CallOption) (*CurrentBalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CurrentBalanceResponse)
	err := c.cc.Invoke(ctx, BankService_GetCurrentBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) FetchExchangeRates(ctx context.Context, in *ExchangeRateRequest, opts ...grpc.CallOption) (BankService_FetchExchangeRatesClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &BankService_ServiceDesc.Streams[0], BankService_FetchExchangeRates_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &bankServiceFetchExchangeRatesClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BankService_FetchExchangeRatesClient interface {
	Recv() (*ExchangeRateResponse, error)
	grpc.ClientStream
}

type bankServiceFetchExchangeRatesClient struct {
	grpc.ClientStream
}

func (x *bankServiceFetchExchangeRatesClient) Recv() (*ExchangeRateResponse, error) {
	m := new(ExchangeRateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bankServiceClient) SummarizeTransactions(ctx context.Context, opts ...grpc.CallOption) (BankService_SummarizeTransactionsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &BankService_ServiceDesc.Streams[1], BankService_SummarizeTransactions_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &bankServiceSummarizeTransactionsClient{ClientStream: stream}
	return x, nil
}

type BankService_SummarizeTransactionsClient interface {
	Send(*Transaction) error
	CloseAndRecv() (*TransactionSummary, error)
	grpc.ClientStream
}

type bankServiceSummarizeTransactionsClient struct {
	grpc.ClientStream
}

func (x *bankServiceSummarizeTransactionsClient) Send(m *Transaction) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bankServiceSummarizeTransactionsClient) CloseAndRecv() (*TransactionSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(TransactionSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BankServiceServer is the server API for BankService service.
// All implementations must embed UnimplementedBankServiceServer
// for forward compatibility
type BankServiceServer interface {
	GetCurrentBalance(context.Context, *CurrentBalanceRequest) (*CurrentBalanceResponse, error)
	FetchExchangeRates(*ExchangeRateRequest, BankService_FetchExchangeRatesServer) error
	SummarizeTransactions(BankService_SummarizeTransactionsServer) error
	mustEmbedUnimplementedBankServiceServer()
}

// UnimplementedBankServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBankServiceServer struct {
}

func (UnimplementedBankServiceServer) GetCurrentBalance(context.Context, *CurrentBalanceRequest) (*CurrentBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentBalance not implemented")
}
func (UnimplementedBankServiceServer) FetchExchangeRates(*ExchangeRateRequest, BankService_FetchExchangeRatesServer) error {
	return status.Errorf(codes.Unimplemented, "method FetchExchangeRates not implemented")
}
func (UnimplementedBankServiceServer) SummarizeTransactions(BankService_SummarizeTransactionsServer) error {
	return status.Errorf(codes.Unimplemented, "method SummarizeTransactions not implemented")
}
func (UnimplementedBankServiceServer) mustEmbedUnimplementedBankServiceServer() {}

// UnsafeBankServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BankServiceServer will
// result in compilation errors.
type UnsafeBankServiceServer interface {
	mustEmbedUnimplementedBankServiceServer()
}

func RegisterBankServiceServer(s grpc.ServiceRegistrar, srv BankServiceServer) {
	s.RegisterService(&BankService_ServiceDesc, srv)
}

func _BankService_GetCurrentBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrentBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).GetCurrentBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BankService_GetCurrentBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).GetCurrentBalance(ctx, req.(*CurrentBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankService_FetchExchangeRates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExchangeRateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BankServiceServer).FetchExchangeRates(m, &bankServiceFetchExchangeRatesServer{ServerStream: stream})
}

type BankService_FetchExchangeRatesServer interface {
	Send(*ExchangeRateResponse) error
	grpc.ServerStream
}

type bankServiceFetchExchangeRatesServer struct {
	grpc.ServerStream
}

func (x *bankServiceFetchExchangeRatesServer) Send(m *ExchangeRateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BankService_SummarizeTransactions_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BankServiceServer).SummarizeTransactions(&bankServiceSummarizeTransactionsServer{ServerStream: stream})
}

type BankService_SummarizeTransactionsServer interface {
	SendAndClose(*TransactionSummary) error
	Recv() (*Transaction, error)
	grpc.ServerStream
}

type bankServiceSummarizeTransactionsServer struct {
	grpc.ServerStream
}

func (x *bankServiceSummarizeTransactionsServer) SendAndClose(m *TransactionSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bankServiceSummarizeTransactionsServer) Recv() (*Transaction, error) {
	m := new(Transaction)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BankService_ServiceDesc is the grpc.ServiceDesc for BankService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BankService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bank.BankService",
	HandlerType: (*BankServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrentBalance",
			Handler:    _BankService_GetCurrentBalance_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FetchExchangeRates",
			Handler:       _BankService_FetchExchangeRates_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SummarizeTransactions",
			Handler:       _BankService_SummarizeTransactions_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/bank/service.proto",
}
