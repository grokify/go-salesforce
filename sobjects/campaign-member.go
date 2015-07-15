package sobjects

import (
	"encoding/xml"
	"io/ioutil"
)

type CampaignMemberSet struct {
	Records []CampaignMember `xml:"records"`
}

func NewCampaignMemberSetFromXml(bytes []byte) (CampaignMemberSet, error) {
	set := CampaignMemberSet{}
	err := xml.Unmarshal(bytes, &set)
	return set, err
}

func NewCampaignMemberSetFromXmlFile(filepath string) (CampaignMemberSet, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return CampaignMemberSet{}, err
	}
	return NewCampaignMemberSetFromXml(bytes)
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
	err := xml.Unmarshal(bytes, &obj)
	return obj, err
}

func NewCampaignMemberFromJsonFile(filepath string) (CampaignMember, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return CampaignMember{}, err
	}
	return NewCampaignMemberFromJson(bytes)
}
