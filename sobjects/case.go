package sobjects

// Case as defined at
// https://developer.salesforce.com/docs/api-explorer/sobject/Case
type Case struct {
	Type            string
	Origin          string
	Reason          string
	Status          string
	OwnerID         string
	Subject         string
	ParentID        string
	Priority        string
	AccountID       string
	ContactID       string
	Description     string
	IsEscalated     bool
	SuppliedName    string
	SuppliedEmail   string
	SuppliedPhone   string
	SuppliedCompany string
}
