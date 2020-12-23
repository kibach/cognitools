package cognito_client

import (
	cognitoIDP "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type CognitoUser struct {
	IDP      *cognitoIDP.CognitoIdentityProvider
	PoolID   string
	Username string
}

type UserAttributeList []*cognitoIDP.AttributeType

func NewUser(idp *cognitoIDP.CognitoIdentityProvider, poolId string, username string) *CognitoUser {
	u := new(CognitoUser)
	u.IDP = idp
	u.PoolID = poolId
	u.Username = username
	return u
}

func (u *CognitoUser) GetInfo() (*cognitoIDP.UserType, *UserMFAPreferences, error) {
	opInput := &cognitoIDP.AdminGetUserInput{}
	opInput.SetUserPoolId(u.PoolID).
		SetUsername(u.Username)

	userInfo, err := u.IDP.AdminGetUser(opInput)
	if err != nil {
		return nil, nil, err
	}

	user := &cognitoIDP.UserType{}
	user.SetAttributes(userInfo.UserAttributes).
		SetEnabled(*userInfo.Enabled).
		SetMFAOptions(userInfo.MFAOptions).
		SetUserCreateDate(*userInfo.UserCreateDate).
		SetUserLastModifiedDate(*userInfo.UserLastModifiedDate).
		SetUserStatus(*userInfo.UserStatus).
		SetUsername(*userInfo.Username)

	mfaPreferences := &UserMFAPreferences{
		EnabledOptions: userInfo.UserMFASettingList,
		PreferedOption: userInfo.PreferredMfaSetting,
	}

	return user, mfaPreferences, nil
}

func (u *CognitoUser) UpdateAttributes(attributes []*cognitoIDP.AttributeType) error {
	opInput := &cognitoIDP.AdminUpdateUserAttributesInput{}
	opInput.SetUserPoolId(u.PoolID).
		SetUsername(u.Username).
		SetUserAttributes(attributes)

	_, err := u.IDP.AdminUpdateUserAttributes(opInput)

	return err
}

func (u *CognitoUser) ChangePassword(newPassword string) error {
	opInput := &cognitoIDP.AdminSetUserPasswordInput{}
	opInput.SetUserPoolId(u.PoolID).
		SetUsername(u.Username).
		SetPassword(newPassword).
		SetPermanent(true)

	_, err := u.IDP.AdminSetUserPassword(opInput)

	return err
}

func (u *CognitoUser) ConfirmSignup() error {
	opInput := &cognitoIDP.AdminConfirmSignUpInput{}
	opInput.SetUserPoolId(u.PoolID).
		SetUsername(u.Username)

	_, err := u.IDP.AdminConfirmSignUp(opInput)

	return err
}

func (u *CognitoUser) Delete() error {
	opInput := &cognitoIDP.AdminDeleteUserInput{}
	opInput.SetUserPoolId(u.PoolID).
		SetUsername(u.Username)

	_, err := u.IDP.AdminDeleteUser(opInput)

	return err
}

func (u *CognitoUser) GetGroups(paginationToken string) ([]*cognitoIDP.GroupType, *string, error) {
	opInput := &cognitoIDP.AdminListGroupsForUserInput{}
	opInput.SetUserPoolId(u.PoolID).
		SetUsername(u.Username).
		SetLimit(60)
	if paginationToken != "" {
		opInput.SetNextToken(paginationToken)
	}

	response, err := u.IDP.AdminListGroupsForUser(opInput)
	if err != nil {
		return nil, nil, err
	}

	return response.Groups, response.NextToken, nil
}

// func (u *CognitoUser) InitiateAuth() {
// 	opInput := &cognitoIDP.AdminInitiateAuthInput{}
// 	opInput.SetUserPoolId(u.PoolID).
// 		Set
// }
