package sobjects

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
)

type CampaignMemberSet struct {
	IdSet      IdSet                     `xml:"-"`
	Records    []CampaignMember          `xml:"records"`
	RecordsMap map[string]CampaignMember `xml:"-"`
}

func NewCampaignMemberSet() CampaignMemberSet {
	set := CampaignMemberSet{
		IdSet:      NewIdSet(),
		Records:    []CampaignMember{},
		RecordsMap: map[string]CampaignMember{}}
	return set
}

func NewCampaignMemberSetFromXml(bytes []byte) (CampaignMemberSet, error) {
	set := NewCampaignMemberSet()
	err := xml.Unmarshal(bytes, &set)
	set.Inflate()
	return set, err
}

func NewCampaignMemberSetFromXmlFile(filepath string) (CampaignMemberSet, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return CampaignMemberSet{}, err
	}
	return NewCampaignMemberSetFromXml(bytes)
}

func (set *CampaignMemberSet) Inflate() {
	for _, record := range set.Records {
		if len(record.Id) > 0 {
			set.IdSet.AddId(record.Id)
			set.RecordsMap[record.Id] = record
		}
		if len(record.ContactId) > 0 {
			set.IdSet.AddId(record.ContactId)
		}
		if len(record.LeadId) > 0 {
			set.IdSet.AddId(record.LeadId)
		}
	}
}

type CampaignMember struct {
	Id                 string
	CampaignId         string
	ContactId          string
	CreatedDate        string
	CurrencyIsoCode    string
	FirstRespondedDate string
	HasResponded       bool
	LeadId             string
	Name               string
	Status             string
}

func NewCampaignMemberFromJson(bytes []byte) (CampaignMember, error) {
	obj := CampaignMember{}
	err := json.Unmarshal(bytes, &obj)
	return obj, err
}

func NewCampaignMemberFromJsonFile(filepath string) (CampaignMember, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return CampaignMember{}, err
	}
	return NewCampaignMemberFromJson(bytes)
}
