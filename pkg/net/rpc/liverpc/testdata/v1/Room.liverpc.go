// Code generated by protoc-gen-liverpc v0.1, DO NOT EDIT.
// source: v1/Room.proto

/*
Package v1 is a generated liverpc stub package.
This code was generated with go-common/app/tool/liverpc/protoc-gen-liverpc v0.1.

It is generated from these files:
	v1/Room.proto
*/
package v1

import context "context"

import proto "github.com/golang/protobuf/proto"
import "github.com/welcome112s/go-library/pkg/pkg/net/rpc/liverpc"

var _ proto.Message // generate to suppress unused imports
// Imports only used by utility functions:

// ==============
// Room Interface
// ==============

type RoomRPCClient interface {
	// * 根据房间id获取房间信息
	GetInfoById(ctx context.Context, req *RoomGetInfoByIdReq, opts ...liverpc.CallOption) (resp *RoomGetInfoByIdResp, err error)

	// * 获取房间基本信息接口，前端/移动端房间页使用
	GetInfo(ctx context.Context, req *RoomGetInfoReq, opts ...liverpc.CallOption) (resp *RoomGetInfoResp, err error)
}

// ====================
// Room Live Rpc Client
// ====================

type roomRPCClient struct {
	client *liverpc.Client
}

// NewRoomRPCClient creates a client that implements the RoomRPCClient interface.
func NewRoomRPCClient(client *liverpc.Client) RoomRPCClient {
	return &roomRPCClient{
		client: client,
	}
}

func (c *roomRPCClient) GetInfoById(ctx context.Context, in *RoomGetInfoByIdReq, opts ...liverpc.CallOption) (*RoomGetInfoByIdResp, error) {
	out := new(RoomGetInfoByIdResp)
	err := doRPCRequest(ctx, c.client, 1, "Room.get_info_by_id", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomRPCClient) GetInfo(ctx context.Context, in *RoomGetInfoReq, opts ...liverpc.CallOption) (*RoomGetInfoResp, error) {
	out := new(RoomGetInfoResp)
	err := doRPCRequest(ctx, c.client, 1, "Room.get_info", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// =====
// Utils
// =====

func doRPCRequest(ctx context.Context, client *liverpc.Client, version int, method string, in, out proto.Message, opts []liverpc.CallOption) (err error) {
	err = client.Call(ctx, version, method, in, out, opts...)
	return
}
