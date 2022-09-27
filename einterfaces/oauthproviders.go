// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package einterfaces

import (
	"fmt"
	"io"

	"github.com/mattermost/mattermost-server/v6/model"
)

type OAuthProvider interface {
	GetUserFromJSON(data io.Reader, tokenUser *model.User) (*model.User, error)
	GetSSOSettings(config *model.Config, service string) (*model.SSOSettings, error)
	GetUserFromIdToken(idToken string) (*model.User, error)
	IsSameUser(dbUser, oAuthUser *model.User) bool
}

var oauthProviders = make(map[string]OAuthProvider)

func RegisterOAuthProvider(name string, newProvider OAuthProvider) {
	fmt.Println("******************************** RegisterOAuthProvider ********************************")
	fmt.Println("_____ newProvider: ", newProvider)

	oauthProviders[name] = newProvider
}

func GetOAuthProvider(name string) OAuthProvider {
	fmt.Println("******************************** GetOAuthProvider ********************************")
	fmt.Println("_____ oauthProviders: ", oauthProviders)
	
	provider, ok := oauthProviders[name]
	fmt.Println("provider: ", provider)

	if ok {
		return provider
	}
	return nil
}
