package sobjects

// Case as defined at
// https://developer.salesforce.com/docs/api-explorer/sobject/Case
type Case struct {
	Type            string
	Origin          string
	Reason          string
	Status          string
	OwnerId         string
	Subject         string
	ParentId        string
	Priority        string
	AccountId       string
	ContactId       string
	Description     string
	IsEscalated     bool
	SuppliedName    string
	SuppliedEmail   string
	SuppliedPhone   string
	SuppliedCompany string
}
