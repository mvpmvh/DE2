package services

import (
	"encoding/json"
	"source.discoveryeducation.com/users/mhatch/repos/de2/models"
	"source.discoveryeducation.com/users/mhatch/repos/de2/thrift/gen-go/generated/user"
)

type UserService struct {
	UserClient *user.UserServicerClient
}

func (this *UserService) FindUser(criteria models.Criteria) *models.User {
	var marshalledCriteria []byte
	var err error

	if marshalledCriteria, err = json.Marshal(criteria); err != nil {
		panic(err)
	}

	var response []byte
	if response, err = this.UserClient.FindBy(marshalledCriteria); err != nil {
		panic(err)
	}

	user := new(models.User)
	json.Unmarshal(response, user)

	return user
}
