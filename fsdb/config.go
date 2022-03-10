package salesforcefsdb

import (
	"net/url"
)

type SalesforceClientConfig struct {
	ConfigGeneral SalesforceClientConfigGeneral
	ConfigToken   SalesforceClientConfigToken
}

type SalesforceClientConfigGeneral struct {
	APIFqdn           string
	APIVersion        string
	DataDir           string
	MaxAgeSec         int64
	FlagDisableRemote bool
	FlagVerbose       bool
	FlagSaveFs        bool
}

type SalesforceClientConfigToken struct {
	TokenURL     string
	GrantType    string
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
	URLValues    url.Values
}

func (ct *SalesforceClientConfigToken) Inflate() {
	vals := url.Values{}
	vals.Add("grant_type", ct.GrantType)
	vals.Add("client_id", ct.ClientID)
	vals.Add("client_secret", ct.ClientSecret)
	vals.Add("username", ct.Username)
	vals.Add("password", ct.Password)
	ct.URLValues = vals
}
