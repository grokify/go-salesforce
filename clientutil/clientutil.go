package clientutil

import (
	"fmt"
	"net/http"
)

type ClientUtil struct {
	HTTPClient *http.Client
	Instance   string
	Version    string
}

type Describe struct {
	Fields []Field `json:"fields,omitempty"`
}

type Field struct {
	Name           string          `json:"name,omitempty"`
	PicklistValues []PicklistValue `json:"picklistValues,omitempty"`
	Type           string          `json:"type,omitempty"`
}

type PicklistValue struct {
	Active       bool   `json:"active"`
	DefaultValue bool   `json:"defaultValue"`
	Label        string `json:"label,omitempty"`
	Value        string `json:"value,omitempty"`
}

// Describe fetches https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/resources_sobject_describe.htm
func (cu *ClientUtil) Describe(sobjectName string) (*http.Response, error) {
	url := DescribeURL(cu.Instance, cu.Version, sobjectName)

	return cu.HTTPClient.Get(url)
}

func DescribeURL(instance, version, sobjectName string) string {
	urlFmt := "https://%s.salesforce.com/services/data/%s/sobjects/%s/describe/"
	return fmt.Sprintf(urlFmt, instance, version, sobjectName)
}
