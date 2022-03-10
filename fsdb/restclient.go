package salesforcefsdb

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/grokify/mogo/net/urlutil"
)

type SalesforceTokenResponse struct {
	Id          string `json:"id"`
	IssuedAt    int64  `json:"issued_at"`
	TokenType   string `json:"token_type"`
	InstanceUrl string `json:"instance_url"`
	AccessToken string `json:"access_token"`
	Signature   string `json:"signature"`
}

type RestClient struct {
	Config        SalesforceClientConfig
	HttpHeaders   map[string]string
	TokenResponse SalesforceTokenResponse
}

func NewRestClient(cfg SalesforceClientConfig) RestClient {
	cl := RestClient{}
	cfg.ConfigToken.Inflate()
	cl.Config = cfg
	cl.TokenResponse = SalesforceTokenResponse{}
	cl.HttpHeaders = map[string]string{}
	return cl
}

func (cl *RestClient) GetAccessToken() (string, error) {
	if len(cl.TokenResponse.AccessToken) > 0 {
		return cl.TokenResponse.AccessToken, nil
	}
	err := cl.LoadToken()
	return cl.TokenResponse.AccessToken, err
}

func (cl *RestClient) LoadToken() error {
	resp, err := http.PostForm(cl.Config.ConfigToken.TokenUrl, cl.Config.ConfigToken.UrlValues)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	tokRes := SalesforceTokenResponse{}
	err = json.Unmarshal(body, &tokRes)
	if err != nil {
		return err
	}
	cl.TokenResponse = tokRes
	return nil
}

func (cl *RestClient) GetSobjectUrlForSfidAndType(sSfid string, sType string) string {
	aUrl := []string{"https:/",
		cl.Config.ConfigGeneral.ApiFqdn,
		"services/data",
		"v" + cl.Config.ConfigGeneral.ApiVersion,
		"sobjects",
		sType,
		sSfid}
	sUrl := strings.Join(aUrl, "/")
	return sUrl
}

func (cl *RestClient) GetSobjectResponseForSfidAndType(sSfid string, sType string) (*http.Response, error) {
	sUrl := cl.GetSobjectUrlForSfidAndType(sSfid, sType)
	client := &http.Client{}
	req, err := http.NewRequest("GET", sUrl, nil)
	if err != nil {
		return &http.Response{}, err
	}
	accessToken, err := cl.GetAccessToken()
	if err != nil {
		return &http.Response{}, err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)
	res, err := client.Do(req)
	return res, err
}

func (cl *RestClient) GetSoqlResponse(sSoql string) (string, error) {
	dQry := map[string][]string{"q": {sSoql}}
	soqlUrlSlice := []string{"https:/",
		cl.Config.ConfigGeneral.ApiFqdn,
		"services/data",
		"v" + cl.Config.ConfigGeneral.ApiVersion,
		"query"}
	soqlUrl := strings.Join(soqlUrlSlice, "/")
	soqlUrlGo, err := urlutil.URLAddQueryString(soqlUrl, dQry)
	if err != nil {
		return soqlUrl, err
	}
	return soqlUrlGo.String(), nil
}
