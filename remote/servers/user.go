/* remote servers that implement the services defined in the thrift files */

package main

import (
	"encoding/json"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	//"github.com/astaxie/beego"
	"source.discoveryeducation.com/users/mhatch/repos/de2/models"
	"source.discoveryeducation.com/users/mhatch/repos/de2/thrift/gen-go/generated/shared"
)

type UserServicerHandler struct {
}

func (UserServicerHandler) FindBy(marshalledCriteria []byte) (userResult []byte, err error) {
	var criteria models.Criteria
	var user = new(models.User)

	json.Unmarshal(marshalledCriteria, &criteria)

	userCriteria := criteria["user"].(map[string]interface{})
	fmt.Printf("userCriteria: %#v", userCriteria)
	if userCriteria["Email"] == "michael_hatch@discovery.com" {
		if userCriteria["Password"] == "Password" {
			var id int32 = 1
			user = &models.User{
				Id:        &id,
				Username:  "mhatch",
				Email:     "michael_hatch@discovery.com",
				Password:  "Password",
				FirstName: "Michael",
				LastName:  "Hatch",
			}
		}
	}

	userResult, err = json.Marshal(user)

	return userResult, err
}

func main() {
	handler := UserServicerHandler{}
	processor := shared.NewBinaryMessengerProcessor(handler)
	//addr := beego.AppConfig.String("socketAddress")
	addr := "localhost:3001"

	var transport thrift.TServerTransport
	var err error

	if transport, err = thrift.NewTServerSocket(addr); err != nil {
		panic(err)
		//panic("Unable to initialize socket!")
	}

	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	fmt.Printf("server started on %v\n", addr)
	server.Serve()
}
