/* package for lazy-loading services */

package factories

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/mvpmvh/beego"
	"source.discoveryeducation.com/users/mhatch/repos/de2/services"
	"source.discoveryeducation.com/users/mhatch/repos/de2/thrift/gen-go/generated/user"
)

func NewUserService() *services.UserService {
	var transport thrift.TTransport
	var err error

	addr := beego.AppConfig.String("socketAddress")

	if transport, err = thrift.NewTSocket(addr); err != nil {
		panic("Unable to initialize socket!")
	}

	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport = transportFactory.GetTransport(transport)

	if !transport.IsOpen() {
		if err = transport.Open(); err != nil {
			panic("Unable to open socket!")
		}
	}

	return &services.UserService{user.NewUserServicerClientFactory(transport, protocolFactory)}
}
