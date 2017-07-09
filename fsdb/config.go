package salesforcefsdb

import (
	"net/url"
)

type SalesforceClientConfig struct {
	ConfigGeneral SalesforceClientConfigGeneral
	ConfigToken   SalesforceClientConfigToken
}

type SalesforceClientConfigGeneral struct {
	ApiFqdn           string
	ApiVersion        string
	DataDir           string
	MaxAgeSec         int64
	FlagDisableRemote bool
	FlagVerbose       bool
	FlagSaveFs        bool
}

type SalesforceClientConfigToken struct {
	TokenUrl     string
	GrantType    string
	ClientId     string
	ClientSecret string
	Username     string
	Password     string
	UrlValues    url.Values
}

func (ct *SalesforceClientConfigToken) Inflate() {
	vals := url.Values{}
	vals.Add("grant_type", ct.GrantType)
	vals.Add("client_id", ct.ClientId)
	vals.Add("client_secret", ct.ClientSecret)
	vals.Add("username", ct.Username)
	vals.Add("password", ct.Password)
	ct.UrlValues = vals
}
