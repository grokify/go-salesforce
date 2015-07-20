package sobjects

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
)

type CampaignMemberSet struct {
	IdSet   IdSet
	Records []CampaignMember `xml:"records"`
}

func NewCampaignMemberSetFromXml(bytes []byte) (CampaignMemberSet, error) {
	set := CampaignMemberSet{IdSet: NewIdSet()}
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
		set.IdSet.AddId(record.Id)
		set.IdSet.AddId(record.ContactId)
		set.IdSet.AddId(record.LeadId)
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
	RecordTypeId       string
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
