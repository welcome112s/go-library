// Code generated by liverpcgen, DO NOT EDIT.
// source: *.proto files under this directory
// If you want to change this file, Please see README in go-common/app/tool/liverpc/protoc-gen-liverpc/
package testdata

import (
	"go-library/pkg/net/rpc/liverpc"
	"go-library/pkg/net/rpc/liverpc/testdata/v1"
	"go-library/pkg/net/rpc/liverpc/testdata/v2"
)

// Client that represents a liverpc room service api
type Client struct {
	cli *liverpc.Client
	// V1Room presents the controller in liverpc
	V1Room v1.RoomRPCClient
	// V2Room presents the controller in liverpc
	V2Room v2.RoomRPCClient
}

// DiscoveryAppId the discovery id is not the tree name
var DiscoveryAppId = "live.room"

// New a Client that represents a liverpc live.room service api
// conf can be empty, and it will use discovery to find service by default
// conf.AppID will be overwrite by a fixed value DiscoveryAppId
// therefore is no need to set
func New(conf *liverpc.ClientConfig) *Client {
	if conf == nil {
		conf = &liverpc.ClientConfig{}
	}
	conf.AppID = DiscoveryAppId
	var realCli = liverpc.NewClient(conf)
	cli := &Client{cli: realCli}
	cli.clientInit(realCli)
	return cli
}

func (cli *Client) GetRawCli() *liverpc.Client {
	return cli.cli
}

func (cli *Client) clientInit(realCli *liverpc.Client) {
	cli.V1Room = v1.NewRoomRPCClient(realCli)
	cli.V2Room = v2.NewRoomRPCClient(realCli)
}
