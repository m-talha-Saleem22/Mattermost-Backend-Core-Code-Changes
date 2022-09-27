// Code generated by mockery v2.10.4. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	io "io"

	model "github.com/mattermost/mattermost-server/v6/model"
	mock "github.com/stretchr/testify/mock"
)

// OAuthProvider is an autogenerated mock type for the OAuthProvider type
type OAuthProvider struct {
	mock.Mock
}

// GetSSOSettings provides a mock function with given fields: config, service
func (_m *OAuthProvider) GetSSOSettings(config *model.Config, service string) (*model.SSOSettings, error) {
	ret := _m.Called(config, service)

	var r0 *model.SSOSettings
	if rf, ok := ret.Get(0).(func(*model.Config, string) *model.SSOSettings); ok {
		r0 = rf(config, service)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SSOSettings)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Config, string) error); ok {
		r1 = rf(config, service)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserFromIdToken provides a mock function with given fields: idToken
func (_m *OAuthProvider) GetUserFromIdToken(idToken string) (*model.User, error) {
	ret := _m.Called(idToken)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(idToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(idToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserFromJSON provides a mock function with given fields: data, tokenUser
func (_m *OAuthProvider) GetUserFromJSON(data io.Reader, tokenUser *model.User) (*model.User, error) {
	ret := _m.Called(data, tokenUser)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(io.Reader, *model.User) *model.User); ok {
		r0 = rf(data, tokenUser)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(io.Reader, *model.User) error); ok {
		r1 = rf(data, tokenUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsSameUser provides a mock function with given fields: dbUser, oAuthUser
func (_m *OAuthProvider) IsSameUser(dbUser *model.User, oAuthUser *model.User) bool {
	
	ret := _m.Called(dbUser, oAuthUser)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*model.User, *model.User) bool); ok {
		r0 = rf(dbUser, oAuthUser)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
