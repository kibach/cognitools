package cognitoclient

import (
	cognitoIDP "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type CreateUserRequest struct {
	Attributes []*cognitoIDP.AttributeType `json:"Attributes" required:"true"`
	Username   string                      `json:"Username" required:"true"`
	Password   string                      `json:"Password" required:"true"`
}

type NewPasswordRequest struct {
	NewPassword string `json:"NewPassword" required:"true"`
}

type UserMFAPreferences struct {
	EnabledOptions []*string
	PreferedOption *string
}
