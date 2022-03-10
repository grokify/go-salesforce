package salesforcefsdb

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/grokify/mogo/net/httputilmore"
	"github.com/grokify/mogo/net/urlutil"
)

type SalesforceTokenResponse struct {
	ID          string `json:"id"`
	IssuedAt    int64  `json:"issued_at"`
	TokenType   string `json:"token_type"`
	InstanceURL string `json:"instance_url"`
	AccessToken string `json:"access_token"`
	Signature   string `json:"signature"`
}

type RestClient struct {
	Config        SalesforceClientConfig
	HTTPHeaders   map[string]string
	TokenResponse SalesforceTokenResponse
}

func NewRestClient(cfg SalesforceClientConfig) RestClient {
	cl := RestClient{}
	cfg.ConfigToken.Inflate()
	cl.Config = cfg
	cl.TokenResponse = SalesforceTokenResponse{}
	cl.HTTPHeaders = map[string]string{}
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
	resp, err := http.PostForm(cl.Config.ConfigToken.TokenURL, cl.Config.ConfigToken.URLValues)
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

func (cl *RestClient) GetSobjectURLForSfidAndType(sSfid string, sType string) string {
	aURL := []string{"https:/",
		cl.Config.ConfigGeneral.APIFqdn,
		"services/data",
		"v" + cl.Config.ConfigGeneral.APIVersion,
		"sobjects",
		sType,
		sSfid}
	return strings.Join(aURL, "/")
}

func (cl *RestClient) GetSobjectResponseForSfidAndType(sSfid string, sType string) (*http.Response, error) {
	sURL := cl.GetSobjectURLForSfidAndType(sSfid, sType)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, sURL, nil)
	if err != nil {
		return nil, err
	}
	accessToken, err := cl.GetAccessToken()
	if err != nil {
		return nil, err
	}
	req.Header.Add(httputilmore.HeaderAuthorization, "Bearer "+accessToken)
	return client.Do(req)
}

func (cl *RestClient) GetSoqlResponse(sSoql string) (string, error) {
	dQry := map[string][]string{"q": {sSoql}}
	soqlURLSlice := []string{"https:/",
		cl.Config.ConfigGeneral.APIFqdn,
		"services/data",
		"v" + cl.Config.ConfigGeneral.APIVersion,
		"query"}
	soqlURL := strings.Join(soqlURLSlice, "/")
	soqlURLGo, err := urlutil.URLAddQueryString(soqlURL, dQry)
	if err != nil {
		return soqlURL, err
	}
	return soqlURLGo.String(), nil
}
