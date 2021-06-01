// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.20
// source: satellite.proto

package pb

import (
	bytes "bytes"
	context "context"
	errors "errors"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"

	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_satellite_proto struct{}

func (drpcEncoding_File_satellite_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_satellite_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return proto.Unmarshal(buf, msg.(proto.Message))
}

func (drpcEncoding_File_satellite_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	var buf bytes.Buffer
	err := new(jsonpb.Marshaler).Marshal(&buf, msg.(proto.Message))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (drpcEncoding_File_satellite_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return jsonpb.Unmarshal(bytes.NewReader(buf), msg.(proto.Message))
}

type DRPCConfigClient interface {
	DRPCConn() drpc.Conn

	MinimumStorageNode(ctx context.Context, in *MinimumStorageNodeRequest) (*MinimumStorageNodeResponse, error)
}

type drpcConfigClient struct {
	cc drpc.Conn
}

func NewDRPCConfigClient(cc drpc.Conn) DRPCConfigClient {
	return &drpcConfigClient{cc}
}

func (c *drpcConfigClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcConfigClient) MinimumStorageNode(ctx context.Context, in *MinimumStorageNodeRequest) (*MinimumStorageNodeResponse, error) {
	out := new(MinimumStorageNodeResponse)
	err := c.cc.Invoke(ctx, "/satellite.Config/MinimumStorageNode", drpcEncoding_File_satellite_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCConfigServer interface {
	MinimumStorageNode(context.Context, *MinimumStorageNodeRequest) (*MinimumStorageNodeResponse, error)
}

type DRPCConfigUnimplementedServer struct{}

func (s *DRPCConfigUnimplementedServer) MinimumStorageNode(context.Context, *MinimumStorageNodeRequest) (*MinimumStorageNodeResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), 12)
}

type DRPCConfigDescription struct{}

func (DRPCConfigDescription) NumMethods() int { return 1 }

func (DRPCConfigDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/satellite.Config/MinimumStorageNode", drpcEncoding_File_satellite_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCConfigServer).
					MinimumStorageNode(
						ctx,
						in1.(*MinimumStorageNodeRequest),
					)
			}, DRPCConfigServer.MinimumStorageNode, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterConfig(mux drpc.Mux, impl DRPCConfigServer) error {
	return mux.Register(impl, DRPCConfigDescription{})
}

type DRPCConfig_MinimumStorageNodeStream interface {
	drpc.Stream
	SendAndClose(*MinimumStorageNodeResponse) error
}

type drpcConfig_MinimumStorageNodeStream struct {
	drpc.Stream
}

func (x *drpcConfig_MinimumStorageNodeStream) SendAndClose(m *MinimumStorageNodeResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_satellite_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}
