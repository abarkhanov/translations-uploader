// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ApiTranslationsClient is an autogenerated mock type for the ApiTranslationsClient type
type ApiTranslationsClient struct {
	mock.Mock
}

// AddToken provides a mock function with given fields: orgID, emailType, token, translations
func (_m *ApiTranslationsClient) AddToken(orgID string, emailType string, token string, translations []map[string]string) error {
	ret := _m.Called(orgID, emailType, token, translations)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, []map[string]string) error); ok {
		r0 = rf(orgID, emailType, token, translations)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
