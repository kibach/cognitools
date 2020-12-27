package cognitoclient

import (
	cognitoIDP "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type CognitoPool struct {
	IDP *cognitoIDP.CognitoIdentityProvider
	ID  string
}

func NewPool(idp *cognitoIDP.CognitoIdentityProvider, id string) *CognitoPool {
	p := new(CognitoPool)
	p.IDP = idp
	p.ID = id
	return p
}

func (p *CognitoPool) GetInfo() (*cognitoIDP.UserPoolType, error) {
	opInput := &cognitoIDP.DescribeUserPoolInput{}
	opInput.SetUserPoolId(p.ID)

	response, err := p.IDP.DescribeUserPool(opInput)
	if err != nil {
		return nil, err
	}

	return response.UserPool, nil
}

func (p *CognitoPool) GetUsers(paginationToken string) ([]*cognitoIDP.UserType, *string, error) {
	opInput := &cognitoIDP.ListUsersInput{}
	opInput.SetUserPoolId(p.ID).
		SetLimit(60)
	if paginationToken != "" {
		opInput.SetPaginationToken(paginationToken)
	}

	poolUsers, err := p.IDP.ListUsers(opInput)
	if err != nil {
		return nil, nil, err
	}

	return poolUsers.Users, poolUsers.PaginationToken, nil
}

func (p *CognitoPool) CreateUser(request *CreateUserRequest) (*cognitoIDP.UserType, error) {
	opInput := cognitoIDP.AdminCreateUserInput{}
	opInput.SetUserAttributes(request.Attributes).
		SetUsername(request.Username).
		SetTemporaryPassword(request.Password).
		SetUserPoolId(p.ID)

	result, err := p.IDP.AdminCreateUser(&opInput)
	if err != nil {
		return nil, err
	}

	return result.User, err
}

func (p *CognitoPool) GetUser(username string) *CognitoUser {
	return NewUser(p.IDP, p.ID, username)
}
