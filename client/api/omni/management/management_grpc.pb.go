// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: omni/management/management.proto

package management

import (
	context "context"

	common "github.com/siderolabs/talos/pkg/machinery/api/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ManagementService_Kubeconfig_FullMethodName                 = "/management.ManagementService/Kubeconfig"
	ManagementService_Talosconfig_FullMethodName                = "/management.ManagementService/Talosconfig"
	ManagementService_Omniconfig_FullMethodName                 = "/management.ManagementService/Omniconfig"
	ManagementService_MachineLogs_FullMethodName                = "/management.ManagementService/MachineLogs"
	ManagementService_ValidateConfig_FullMethodName             = "/management.ManagementService/ValidateConfig"
	ManagementService_CreateServiceAccount_FullMethodName       = "/management.ManagementService/CreateServiceAccount"
	ManagementService_RenewServiceAccount_FullMethodName        = "/management.ManagementService/RenewServiceAccount"
	ManagementService_ListServiceAccounts_FullMethodName        = "/management.ManagementService/ListServiceAccounts"
	ManagementService_DestroyServiceAccount_FullMethodName      = "/management.ManagementService/DestroyServiceAccount"
	ManagementService_KubernetesUpgradePreChecks_FullMethodName = "/management.ManagementService/KubernetesUpgradePreChecks"
	ManagementService_KubernetesSyncManifests_FullMethodName    = "/management.ManagementService/KubernetesSyncManifests"
	ManagementService_CreateSchematic_FullMethodName            = "/management.ManagementService/CreateSchematic"
)

// ManagementServiceClient is the client API for ManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagementServiceClient interface {
	Kubeconfig(ctx context.Context, in *KubeconfigRequest, opts ...grpc.CallOption) (*KubeconfigResponse, error)
	Talosconfig(ctx context.Context, in *TalosconfigRequest, opts ...grpc.CallOption) (*TalosconfigResponse, error)
	Omniconfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*OmniconfigResponse, error)
	MachineLogs(ctx context.Context, in *MachineLogsRequest, opts ...grpc.CallOption) (ManagementService_MachineLogsClient, error)
	ValidateConfig(ctx context.Context, in *ValidateConfigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateServiceAccount(ctx context.Context, in *CreateServiceAccountRequest, opts ...grpc.CallOption) (*CreateServiceAccountResponse, error)
	RenewServiceAccount(ctx context.Context, in *RenewServiceAccountRequest, opts ...grpc.CallOption) (*RenewServiceAccountResponse, error)
	ListServiceAccounts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListServiceAccountsResponse, error)
	DestroyServiceAccount(ctx context.Context, in *DestroyServiceAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	KubernetesUpgradePreChecks(ctx context.Context, in *KubernetesUpgradePreChecksRequest, opts ...grpc.CallOption) (*KubernetesUpgradePreChecksResponse, error)
	KubernetesSyncManifests(ctx context.Context, in *KubernetesSyncManifestRequest, opts ...grpc.CallOption) (ManagementService_KubernetesSyncManifestsClient, error)
	CreateSchematic(ctx context.Context, in *CreateSchematicRequest, opts ...grpc.CallOption) (*CreateSchematicResponse, error)
}

type managementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewManagementServiceClient(cc grpc.ClientConnInterface) ManagementServiceClient {
	return &managementServiceClient{cc}
}

func (c *managementServiceClient) Kubeconfig(ctx context.Context, in *KubeconfigRequest, opts ...grpc.CallOption) (*KubeconfigResponse, error) {
	out := new(KubeconfigResponse)
	err := c.cc.Invoke(ctx, ManagementService_Kubeconfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) Talosconfig(ctx context.Context, in *TalosconfigRequest, opts ...grpc.CallOption) (*TalosconfigResponse, error) {
	out := new(TalosconfigResponse)
	err := c.cc.Invoke(ctx, ManagementService_Talosconfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) Omniconfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*OmniconfigResponse, error) {
	out := new(OmniconfigResponse)
	err := c.cc.Invoke(ctx, ManagementService_Omniconfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) MachineLogs(ctx context.Context, in *MachineLogsRequest, opts ...grpc.CallOption) (ManagementService_MachineLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ManagementService_ServiceDesc.Streams[0], ManagementService_MachineLogs_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &managementServiceMachineLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ManagementService_MachineLogsClient interface {
	Recv() (*common.Data, error)
	grpc.ClientStream
}

type managementServiceMachineLogsClient struct {
	grpc.ClientStream
}

func (x *managementServiceMachineLogsClient) Recv() (*common.Data, error) {
	m := new(common.Data)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *managementServiceClient) ValidateConfig(ctx context.Context, in *ValidateConfigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ManagementService_ValidateConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) CreateServiceAccount(ctx context.Context, in *CreateServiceAccountRequest, opts ...grpc.CallOption) (*CreateServiceAccountResponse, error) {
	out := new(CreateServiceAccountResponse)
	err := c.cc.Invoke(ctx, ManagementService_CreateServiceAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) RenewServiceAccount(ctx context.Context, in *RenewServiceAccountRequest, opts ...grpc.CallOption) (*RenewServiceAccountResponse, error) {
	out := new(RenewServiceAccountResponse)
	err := c.cc.Invoke(ctx, ManagementService_RenewServiceAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) ListServiceAccounts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListServiceAccountsResponse, error) {
	out := new(ListServiceAccountsResponse)
	err := c.cc.Invoke(ctx, ManagementService_ListServiceAccounts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) DestroyServiceAccount(ctx context.Context, in *DestroyServiceAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ManagementService_DestroyServiceAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) KubernetesUpgradePreChecks(ctx context.Context, in *KubernetesUpgradePreChecksRequest, opts ...grpc.CallOption) (*KubernetesUpgradePreChecksResponse, error) {
	out := new(KubernetesUpgradePreChecksResponse)
	err := c.cc.Invoke(ctx, ManagementService_KubernetesUpgradePreChecks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) KubernetesSyncManifests(ctx context.Context, in *KubernetesSyncManifestRequest, opts ...grpc.CallOption) (ManagementService_KubernetesSyncManifestsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ManagementService_ServiceDesc.Streams[1], ManagementService_KubernetesSyncManifests_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &managementServiceKubernetesSyncManifestsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ManagementService_KubernetesSyncManifestsClient interface {
	Recv() (*KubernetesSyncManifestResponse, error)
	grpc.ClientStream
}

type managementServiceKubernetesSyncManifestsClient struct {
	grpc.ClientStream
}

func (x *managementServiceKubernetesSyncManifestsClient) Recv() (*KubernetesSyncManifestResponse, error) {
	m := new(KubernetesSyncManifestResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *managementServiceClient) CreateSchematic(ctx context.Context, in *CreateSchematicRequest, opts ...grpc.CallOption) (*CreateSchematicResponse, error) {
	out := new(CreateSchematicResponse)
	err := c.cc.Invoke(ctx, ManagementService_CreateSchematic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagementServiceServer is the server API for ManagementService service.
// All implementations must embed UnimplementedManagementServiceServer
// for forward compatibility
type ManagementServiceServer interface {
	Kubeconfig(context.Context, *KubeconfigRequest) (*KubeconfigResponse, error)
	Talosconfig(context.Context, *TalosconfigRequest) (*TalosconfigResponse, error)
	Omniconfig(context.Context, *emptypb.Empty) (*OmniconfigResponse, error)
	MachineLogs(*MachineLogsRequest, ManagementService_MachineLogsServer) error
	ValidateConfig(context.Context, *ValidateConfigRequest) (*emptypb.Empty, error)
	CreateServiceAccount(context.Context, *CreateServiceAccountRequest) (*CreateServiceAccountResponse, error)
	RenewServiceAccount(context.Context, *RenewServiceAccountRequest) (*RenewServiceAccountResponse, error)
	ListServiceAccounts(context.Context, *emptypb.Empty) (*ListServiceAccountsResponse, error)
	DestroyServiceAccount(context.Context, *DestroyServiceAccountRequest) (*emptypb.Empty, error)
	KubernetesUpgradePreChecks(context.Context, *KubernetesUpgradePreChecksRequest) (*KubernetesUpgradePreChecksResponse, error)
	KubernetesSyncManifests(*KubernetesSyncManifestRequest, ManagementService_KubernetesSyncManifestsServer) error
	CreateSchematic(context.Context, *CreateSchematicRequest) (*CreateSchematicResponse, error)
	mustEmbedUnimplementedManagementServiceServer()
}

// UnimplementedManagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedManagementServiceServer struct {
}

func (UnimplementedManagementServiceServer) Kubeconfig(context.Context, *KubeconfigRequest) (*KubeconfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Kubeconfig not implemented")
}
func (UnimplementedManagementServiceServer) Talosconfig(context.Context, *TalosconfigRequest) (*TalosconfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Talosconfig not implemented")
}
func (UnimplementedManagementServiceServer) Omniconfig(context.Context, *emptypb.Empty) (*OmniconfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Omniconfig not implemented")
}
func (UnimplementedManagementServiceServer) MachineLogs(*MachineLogsRequest, ManagementService_MachineLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method MachineLogs not implemented")
}
func (UnimplementedManagementServiceServer) ValidateConfig(context.Context, *ValidateConfigRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateConfig not implemented")
}
func (UnimplementedManagementServiceServer) CreateServiceAccount(context.Context, *CreateServiceAccountRequest) (*CreateServiceAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateServiceAccount not implemented")
}
func (UnimplementedManagementServiceServer) RenewServiceAccount(context.Context, *RenewServiceAccountRequest) (*RenewServiceAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenewServiceAccount not implemented")
}
func (UnimplementedManagementServiceServer) ListServiceAccounts(context.Context, *emptypb.Empty) (*ListServiceAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServiceAccounts not implemented")
}
func (UnimplementedManagementServiceServer) DestroyServiceAccount(context.Context, *DestroyServiceAccountRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DestroyServiceAccount not implemented")
}
func (UnimplementedManagementServiceServer) KubernetesUpgradePreChecks(context.Context, *KubernetesUpgradePreChecksRequest) (*KubernetesUpgradePreChecksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KubernetesUpgradePreChecks not implemented")
}
func (UnimplementedManagementServiceServer) KubernetesSyncManifests(*KubernetesSyncManifestRequest, ManagementService_KubernetesSyncManifestsServer) error {
	return status.Errorf(codes.Unimplemented, "method KubernetesSyncManifests not implemented")
}
func (UnimplementedManagementServiceServer) CreateSchematic(context.Context, *CreateSchematicRequest) (*CreateSchematicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSchematic not implemented")
}
func (UnimplementedManagementServiceServer) mustEmbedUnimplementedManagementServiceServer() {}

// UnsafeManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagementServiceServer will
// result in compilation errors.
type UnsafeManagementServiceServer interface {
	mustEmbedUnimplementedManagementServiceServer()
}

func RegisterManagementServiceServer(s grpc.ServiceRegistrar, srv ManagementServiceServer) {
	s.RegisterService(&ManagementService_ServiceDesc, srv)
}

func _ManagementService_Kubeconfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KubeconfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).Kubeconfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementService_Kubeconfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).Kubeconfig(ctx, req.(*KubeconfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_Talosconfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TalosconfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).Talosconfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementService_Talosconfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).Talosconfig(ctx, req.(*TalosconfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_Omniconfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).Omniconfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementService_Omniconfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).Omniconfig(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_MachineLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MachineLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ManagementServiceServer).MachineLogs(m, &managementServiceMachineLogsServer{stream})
}

type ManagementService_MachineLogsServer interface {
	Send(*common.Data) error
	grpc.ServerStream
}

type managementServiceMachineLogsServer struct {
	grpc.ServerStream
}

func (x *managementServiceMachineLogsServer) Send(m *common.Data) error {
	return x.ServerStream.SendMsg(m)
}

func _ManagementService_ValidateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).ValidateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementService_ValidateConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).ValidateConfig(ctx, req.(*ValidateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_CreateServiceAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).CreateServiceAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementService_CreateServiceAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).CreateServiceAccount(ctx, req.(*CreateServiceAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_RenewServiceAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenewServiceAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).RenewServiceAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementService_RenewServiceAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).RenewServiceAccount(ctx, req.(*RenewServiceAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_ListServiceAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).ListServiceAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementService_ListServiceAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).ListServiceAccounts(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_DestroyServiceAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DestroyServiceAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).DestroyServiceAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementService_DestroyServiceAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).DestroyServiceAccount(ctx, req.(*DestroyServiceAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_KubernetesUpgradePreChecks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KubernetesUpgradePreChecksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).KubernetesUpgradePreChecks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementService_KubernetesUpgradePreChecks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).KubernetesUpgradePreChecks(ctx, req.(*KubernetesUpgradePreChecksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_KubernetesSyncManifests_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(KubernetesSyncManifestRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ManagementServiceServer).KubernetesSyncManifests(m, &managementServiceKubernetesSyncManifestsServer{stream})
}

type ManagementService_KubernetesSyncManifestsServer interface {
	Send(*KubernetesSyncManifestResponse) error
	grpc.ServerStream
}

type managementServiceKubernetesSyncManifestsServer struct {
	grpc.ServerStream
}

func (x *managementServiceKubernetesSyncManifestsServer) Send(m *KubernetesSyncManifestResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ManagementService_CreateSchematic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSchematicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).CreateSchematic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementService_CreateSchematic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).CreateSchematic(ctx, req.(*CreateSchematicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ManagementService_ServiceDesc is the grpc.ServiceDesc for ManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "management.ManagementService",
	HandlerType: (*ManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Kubeconfig",
			Handler:    _ManagementService_Kubeconfig_Handler,
		},
		{
			MethodName: "Talosconfig",
			Handler:    _ManagementService_Talosconfig_Handler,
		},
		{
			MethodName: "Omniconfig",
			Handler:    _ManagementService_Omniconfig_Handler,
		},
		{
			MethodName: "ValidateConfig",
			Handler:    _ManagementService_ValidateConfig_Handler,
		},
		{
			MethodName: "CreateServiceAccount",
			Handler:    _ManagementService_CreateServiceAccount_Handler,
		},
		{
			MethodName: "RenewServiceAccount",
			Handler:    _ManagementService_RenewServiceAccount_Handler,
		},
		{
			MethodName: "ListServiceAccounts",
			Handler:    _ManagementService_ListServiceAccounts_Handler,
		},
		{
			MethodName: "DestroyServiceAccount",
			Handler:    _ManagementService_DestroyServiceAccount_Handler,
		},
		{
			MethodName: "KubernetesUpgradePreChecks",
			Handler:    _ManagementService_KubernetesUpgradePreChecks_Handler,
		},
		{
			MethodName: "CreateSchematic",
			Handler:    _ManagementService_CreateSchematic_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MachineLogs",
			Handler:       _ManagementService_MachineLogs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "KubernetesSyncManifests",
			Handler:       _ManagementService_KubernetesSyncManifests_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "omni/management/management.proto",
}