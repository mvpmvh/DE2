// Autogenerated by Thrift Compiler (0.9.1)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package user

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"source.discoveryeducation.com/users/mhatch/repos/de2/thrift/gen-go/generated/shared"
)

// (needed to ensure safety because of naive import list construction.)
var _ = math.MinInt32
var _ = thrift.ZERO
var _ = fmt.Printf

var _ = shared.GoUnusedProtection__

type UserServicer interface {
	shared.BinaryMessenger
}

type UserServicerClient struct {
	*shared.BinaryMessengerClient
}

func NewUserServicerClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *UserServicerClient {
	return &UserServicerClient{BinaryMessengerClient: shared.NewBinaryMessengerClientFactory(t, f)}
}

func NewUserServicerClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *UserServicerClient {
	return &UserServicerClient{BinaryMessengerClient: shared.NewBinaryMessengerClientProtocol(t, iprot, oprot)}
}

type UserServicerProcessor struct {
	*shared.BinaryMessengerProcessor
}

func NewUserServicerProcessor(handler UserServicer) *UserServicerProcessor {
	self0 := &UserServicerProcessor{shared.NewBinaryMessengerProcessor(handler)}
	return self0
}

// HELPER FUNCTIONS AND STRUCTURES
