package describe

/*
https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/resources_sobject_describe.htm
*/

type Description struct {
	Activateable        bool                `json:"activateable"`
	ChildRelationships  []ChildRelationship `json:"childRelationships"`
	CompactLayoutable   bool                `json:"compactLayoutable"`
	Createable          bool                `json:"createable"`
	Custom              bool                `json:"custom"`
	CustomSetting       bool                `json:"customSetting"`
	Deletable           bool                `json:"deletable"`
	DeprecatedAndHidden bool                `json:"deprecatedAndHidden"`
	FeedEnabled         bool                `json:"feedEnabled"`
	Fields              []Field             `json:"fields"`
}

type ChildRelationship struct {
	CascadeDelete       bool          `json:"cascadeDelete"`
	ChildSObject        string        `json:"childSObject"`
	DeprecatedAndHidden bool          `json:"deprecatedAndHidden"`
	Field               string        `json:"field"`
	JunctionIdListNames []interface{} `json:"junctionIdListNames"`
	JunctionReferenceTo []interface{} `json:"junctionReferenceTo"`
	RelationshipName    string        `json:"relationshipName"`
	RestrictedDelete    bool          `json:"restrictedDelete"`
}

type Field struct {
	Aggregatable      bool `json:"aggregatable"`
	AiPredictionField bool `json:"aiPredictionField"`
	AutoNumber        bool `json:"autoNumber"`
	ByteLength        int  `json:"byteLength"`
	Calculated        bool `json:"calculated"`
	//CalculatedFormula" : null,
	CascadeDelete     bool   `json:"cascadeDelete"`
	CaseSensitive     bool   `json:"caseSensitive"`
	CompoundFieldName string `json:"compoundFieldName"`
	//ControllerName" : null,
	Createable bool `json:"createable"`
	Custom     bool `json:"custom"`
	//  "defaultValue" : null,
	//"defaultValueFormula" : null,
	DefaultedOnCreate        bool   `json:"defaultedOnCreate"`
	DependentPicklist        bool   `json:"dependentPicklist"`
	DeprecatedAndHidden      bool   `json:"deprecatedAndHidden"`
	Digits                   int    `json:"digits"`
	DisplayLocationInDecimal bool   `json:"displayLocationInDecimal"`
	Encrypted                bool   `json:"encrypted"`
	ExternalId               bool   `json:"externalId"`
	ExtraTypeInfo            string `json:"extraTypeInfo"`
	Filterable               bool   `json:"filterable"`
	//FilteredLookupInfo" : null,
	FormulaTreatNullNumberAsZero bool   `json:"formulaTreatNullNumberAsZero"`
	Groupable                    bool   `json:"groupable"`
	HighScaleNumber              bool   `json:"highScaleNumber"`
	HtmlFormatted                bool   `json:"htmlFormatted"`
	IdLookup                     bool   `json:"idLookup"`
	InlineHelpText               string `json:"inlineHelpText"` // null,
	Label                        string `json:"label"`
	Length                       int    `json:"length"`
	//"mask" : null,
	//"maskType" : null,
	Name                  string          `json:"name"`
	NameField             bool            `json:"nameField"`
	NamePointing          bool            `json:"namePointing"`
	Nillable              bool            `json:"nillable"`
	Permissionable        bool            `json:"permissionable"`
	PicklistValues        []PicklistValue `json:"picklistValues"`
	PolymorphicForeignKey bool            `json:"polymorphicForeignKey"`
	Precision             int             `json:"precision"`
	QueryByDistance       bool            `json:"queryByDistance"`
	//"referenceTargetField" : null,
	ReferenceTo      []string `json:"referenceTo"`
	RelationshipName string   `json:"relationshipName"` // null,
	//"relationshipOrder" : null,
	RestrictedDelete        bool   `json:"restrictedDelete"`
	RestrictedPicklist      bool   `json:"restrictedPicklist"`
	Scale                   int    `json:"scale"`
	PearchPrefilterable     bool   `json:"searchPrefilterable"`
	SoapType                string `json:"soapType"`
	Sortable                bool   `json:"sortable"`
	Type                    string `json:"type"`
	Unique                  bool   `json:"unique"`
	Updateable              bool   `json:"updateable"`
	WriteRequiresMasterRead bool   `json:"writeRequiresMasterRead"`
}

type PicklistValue struct {
	Active       bool   `json:"active"`
	DefaultValue bool   `json:"defaultValue"`
	Label        string `json:"type"`
	//"validFor" : null,
	Value string `json:"type"`
}
