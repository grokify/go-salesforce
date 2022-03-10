package sobjects

import (
	"encoding/json"
	"encoding/xml"
	"os"
)

type CampaignMemberSet struct {
	IDSet      IDSet                     `xml:"-"`
	Records    []CampaignMember          `xml:"records"`
	RecordsMap map[string]CampaignMember `xml:"-"`
}

func NewCampaignMemberSet() CampaignMemberSet {
	set := CampaignMemberSet{
		IDSet:      NewIDSet(),
		Records:    []CampaignMember{},
		RecordsMap: map[string]CampaignMember{}}
	return set
}

func NewCampaignMemberSetFromXML(bytes []byte) (CampaignMemberSet, error) {
	set := NewCampaignMemberSet()
	err := xml.Unmarshal(bytes, &set)
	if err != nil {
		return set, err
	}
	set.Inflate()
	return set, err
}

func NewCampaignMemberSetFromXMLFile(filepath string) (CampaignMemberSet, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return CampaignMemberSet{}, err
	}
	return NewCampaignMemberSetFromXML(bytes)
}

func (set *CampaignMemberSet) Inflate() {
	for _, record := range set.Records {
		if len(record.ID) > 0 {
			set.IDSet.AddId(record.ID)
			set.RecordsMap[record.ID] = record
		}
		if len(record.ContactID) > 0 {
			set.IDSet.AddId(record.ContactID)
		}
		if len(record.LeadID) > 0 {
			set.IDSet.AddId(record.LeadID)
		}
	}
}

type CampaignMember struct {
	ID                 string
	CampaignID         string
	ContactID          string
	CreatedDate        string
	CurrencyIsoCode    string
	FirstRespondedDate string
	HasResponded       bool
	LeadID             string
	Name               string
	Status             string
}

func NewCampaignMemberFromJSON(bytes []byte) (CampaignMember, error) {
	obj := CampaignMember{}
	err := json.Unmarshal(bytes, &obj)
	return obj, err
}

func NewCampaignMemberFromJSONFile(filepath string) (CampaignMember, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return CampaignMember{}, err
	}
	return NewCampaignMemberFromJSON(bytes)
}
