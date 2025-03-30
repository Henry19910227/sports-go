package client

import userClient "api_gateway/client/user"

type Factory interface {
	UserClient() userClient.Client
}
