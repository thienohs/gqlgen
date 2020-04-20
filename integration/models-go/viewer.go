package models

import "github.com/thienohs/gqlgen/integration/remote_api"

type Viewer struct {
	User *remote_api.User
}
