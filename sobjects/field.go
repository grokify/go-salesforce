package sobjects

import "golang.org/x/net/html"

type FieldDefinition struct {
	Aggregatable        bool   `json:"aggregatable,omitempty"`
	AutoNumber          bool   `json:"autoNumber,omitempty"`
	ByteLength          int    `json:"byteLength,omitempty"`
	Calculated          bool   `json:"calculated,omitempty"`
	CaseSensitive       bool   `json:"caseSensitive,omitempty"`
	Createable          bool   `json:"createable,omitempty"`
	Custom              bool   `json:"custom,omitempty"`
	DefaultedOnCreate   bool   `json:"defaultedOnCreate,omitempty"`
	DeprecatedAndHidden bool   `json:"deprecatedAndHidden,omitempty"`
	Digits              int    `json:"digits,omitempty"`
	Filterable          bool   `json:"filterable,omitempty"`
	Groupable           bool   `json:"groupable,omitempty"`
	IdLookup            bool   `json:"idLookup,omitempty"`
	Label               string `json:"label,omitempty"`
	Length              int    `json:"length,omitempty"`
	Name                string `json:"name,omitempty"`
	NameField           bool   `json:"nameField,omitempty"`
	NamePointing        bool   `json:"namePointing,omitempty"`
	Nillable            bool   `json:"nillable,omitempty"`
	Permissionable      bool   `json:"permissionable,omitempty"`
	Precision           int    `json:"precision,omitempty"`
	QueryByDistance     bool   `json:"queryByDistance,omitempty"`
	RestrictedPicklist  bool   `json:"restrictedPicklist,omitempty"`
	Scale               int    `json:"scale,omitempty"`
	SoapType            string `json:"soapType,omitempty"`
	Sortable            bool   `json:"sortable,omitempty"`
	Type                string `json:"type,omitempty"`
	Unique              bool   `json:"unique,omitempty"`
	Updateable          bool   `json:"updateable,omitempty"`
}

/*
type ObjectDefinitionParserHTML struct {
	Fields []FieldDefinition
}

func (parser *ObjectDefinitionParserHTML) Parse(wbHTML string) {
	fields := []FieldDefinition{}
}
*/
