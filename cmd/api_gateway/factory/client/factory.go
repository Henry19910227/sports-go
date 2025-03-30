package client

import (
	userClient "api_gateway/client/user"
)

type factory struct {
	userClient userClient.Client
}

func New() (Factory, error) {
	userServ, err := userClient.New("127.0.0.1:50051")
	if err != nil {
		return nil, err
	}
	return &factory{userServ}, nil
}

func (f *factory) UserClient() userClient.Client {
	return f.userClient
}
