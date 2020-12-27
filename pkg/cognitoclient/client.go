package cognitoclient

import (
	"github.com/aws/aws-sdk-go/aws/session"
	cognitoIDP "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type CognitoClient struct {
	IDP     *cognitoIDP.CognitoIdentityProvider
	Session *session.Session
}

func GetAWSSession() *session.Session {
	return session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
}

func NewClient() *CognitoClient {
	c := new(CognitoClient)
	c.Session = GetAWSSession()
	c.IDP = cognitoIDP.New(c.Session)
	return c
}

func (c *CognitoClient) ListPools() ([]*cognitoIDP.UserPoolDescriptionType, error) {
	opInput := &cognitoIDP.ListUserPoolsInput{}
	opInput.SetMaxResults(50)

	response, err := c.IDP.ListUserPools(opInput)
	if err != nil {
		return nil, err
	}

	return response.UserPools, err
}

func (c *CognitoClient) GetPool(poolId string) *CognitoPool {
	return NewPool(c.IDP, poolId)
}
