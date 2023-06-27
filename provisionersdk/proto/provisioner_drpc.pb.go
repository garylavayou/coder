// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.33
// source: provisionersdk/proto/provisioner.proto

package proto

import (
	context "context"
	errors "errors"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_provisionersdk_proto_provisioner_proto struct{}

func (drpcEncoding_File_provisionersdk_proto_provisioner_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_provisionersdk_proto_provisioner_proto) MarshalAppend(buf []byte, msg drpc.Message) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(buf, msg.(proto.Message))
}

func (drpcEncoding_File_provisionersdk_proto_provisioner_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return proto.Unmarshal(buf, msg.(proto.Message))
}

func (drpcEncoding_File_provisionersdk_proto_provisioner_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	return protojson.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_provisionersdk_proto_provisioner_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return protojson.Unmarshal(buf, msg.(proto.Message))
}

type DRPCProvisionerClient interface {
	DRPCConn() drpc.Conn

	Parse(ctx context.Context, in *Parse_Request) (DRPCProvisioner_ParseClient, error)
	Provision(ctx context.Context) (DRPCProvisioner_ProvisionClient, error)
}

type drpcProvisionerClient struct {
	cc drpc.Conn
}

func NewDRPCProvisionerClient(cc drpc.Conn) DRPCProvisionerClient {
	return &drpcProvisionerClient{cc}
}

func (c *drpcProvisionerClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcProvisionerClient) Parse(ctx context.Context, in *Parse_Request) (DRPCProvisioner_ParseClient, error) {
	stream, err := c.cc.NewStream(ctx, "/provisioner.Provisioner/Parse", drpcEncoding_File_provisionersdk_proto_provisioner_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcProvisioner_ParseClient{stream}
	if err := x.MsgSend(in, drpcEncoding_File_provisionersdk_proto_provisioner_proto{}); err != nil {
		return nil, err
	}
	if err := x.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DRPCProvisioner_ParseClient interface {
	drpc.Stream
	Recv() (*Parse_Response, error)
}

type drpcProvisioner_ParseClient struct {
	drpc.Stream
}

func (x *drpcProvisioner_ParseClient) GetStream() drpc.Stream {
	return x.Stream
}

func (x *drpcProvisioner_ParseClient) Recv() (*Parse_Response, error) {
	m := new(Parse_Response)
	if err := x.MsgRecv(m, drpcEncoding_File_provisionersdk_proto_provisioner_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcProvisioner_ProvisionClient) GetStream() drpc.Stream {
	return x.Stream
}

func (x *drpcProvisioner_ParseClient) RecvMsg(m *Parse_Response) error {
	return x.MsgRecv(m, drpcEncoding_File_provisionersdk_proto_provisioner_proto{})
}

func (c *drpcProvisionerClient) Provision(ctx context.Context) (DRPCProvisioner_ProvisionClient, error) {
	stream, err := c.cc.NewStream(ctx, "/provisioner.Provisioner/Provision", drpcEncoding_File_provisionersdk_proto_provisioner_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcProvisioner_ProvisionClient{stream}
	return x, nil
}

type DRPCProvisioner_ProvisionClient interface {
	drpc.Stream
	Send(*Provision_Request) error
	Recv() (*Provision_Response, error)
}

type drpcProvisioner_ProvisionClient struct {
	drpc.Stream
}

func (x *drpcProvisioner_ProvisionClient) Send(m *Provision_Request) error {
	return x.MsgSend(m, drpcEncoding_File_provisionersdk_proto_provisioner_proto{})
}

func (x *drpcProvisioner_ProvisionClient) Recv() (*Provision_Response, error) {
	m := new(Provision_Response)
	if err := x.MsgRecv(m, drpcEncoding_File_provisionersdk_proto_provisioner_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcProvisioner_ProvisionClient) RecvMsg(m *Provision_Response) error {
	return x.MsgRecv(m, drpcEncoding_File_provisionersdk_proto_provisioner_proto{})
}

type DRPCProvisionerServer interface {
	Parse(*Parse_Request, DRPCProvisioner_ParseStream) error
	Provision(DRPCProvisioner_ProvisionStream) error
}

type DRPCProvisionerUnimplementedServer struct{}

func (s *DRPCProvisionerUnimplementedServer) Parse(*Parse_Request, DRPCProvisioner_ParseStream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCProvisionerUnimplementedServer) Provision(DRPCProvisioner_ProvisionStream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCProvisionerDescription struct{}

func (DRPCProvisionerDescription) NumMethods() int { return 2 }

func (DRPCProvisionerDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/provisioner.Provisioner/Parse", drpcEncoding_File_provisionersdk_proto_provisioner_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCProvisionerServer).
					Parse(
						in1.(*Parse_Request),
						&drpcProvisioner_ParseStream{in2.(drpc.Stream)},
					)
			}, DRPCProvisionerServer.Parse, true
	case 1:
		return "/provisioner.Provisioner/Provision", drpcEncoding_File_provisionersdk_proto_provisioner_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCProvisionerServer).
					Provision(
						&drpcProvisioner_ProvisionStream{in1.(drpc.Stream)},
					)
			}, DRPCProvisionerServer.Provision, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterProvisioner(mux drpc.Mux, impl DRPCProvisionerServer) error {
	return mux.Register(impl, DRPCProvisionerDescription{})
}

type DRPCProvisioner_ParseStream interface {
	drpc.Stream
	Send(*Parse_Response) error
}

type drpcProvisioner_ParseStream struct {
	drpc.Stream
}

func (x *drpcProvisioner_ParseStream) Send(m *Parse_Response) error {
	return x.MsgSend(m, drpcEncoding_File_provisionersdk_proto_provisioner_proto{})
}

type DRPCProvisioner_ProvisionStream interface {
	drpc.Stream
	Send(*Provision_Response) error
	Recv() (*Provision_Request, error)
}

type drpcProvisioner_ProvisionStream struct {
	drpc.Stream
}

func (x *drpcProvisioner_ProvisionStream) Send(m *Provision_Response) error {
	return x.MsgSend(m, drpcEncoding_File_provisionersdk_proto_provisioner_proto{})
}

func (x *drpcProvisioner_ProvisionStream) Recv() (*Provision_Request, error) {
	m := new(Provision_Request)
	if err := x.MsgRecv(m, drpcEncoding_File_provisionersdk_proto_provisioner_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcProvisioner_ProvisionStream) RecvMsg(m *Provision_Request) error {
	return x.MsgRecv(m, drpcEncoding_File_provisionersdk_proto_provisioner_proto{})
}
