package sobjects

import (
	"errors"
	"regexp"
	"strings"
)

type SObjectsInfo struct {
	Type2Prefix map[string]string
	Prefix2Type map[string]string
	rxChar3     *regexp.Regexp
	rxChar15    *regexp.Regexp
}

func NewSObjectsInfo() SObjectsInfo {
	types := SObjectsInfo{}
	types.rxChar3 = regexp.MustCompile(`^([0-9A-Za-z]{3})`)
	types.rxChar15 = regexp.MustCompile(`^([0-9A-Za-z]{15})`)
	types.loadMaps()
	return types
}

func (types *SObjectsInfo) GetID15ForID(id string) (string, error) {
	if len(id) == 15 {
		return id, nil
	} else if len(id) == 18 {
		rs15 := types.rxChar15.FindStringSubmatch(id)
		if len(rs15) > 1 {
			return rs15[1], nil
		}
	}
	return "", errors.New("sfdc id 15 not found")
}

func (types *SObjectsInfo) GetTypeForID(id string) (string, error) {
	prefix, err := types.GetPrefixForID(id)
	if err != nil {
		return "", err
	}
	return types.GetTypeForPrefix(prefix)
}

func (types *SObjectsInfo) GetPrefixForID(id string) (string, error) {
	if len(id) < 3 {
		return "", errors.New("sfdc id not provided")
	}
	rs3 := types.rxChar3.FindStringSubmatch(id)
	if len(rs3) > 0 {
		return rs3[1], nil
	}
	return "", errors.New("sobject prefix not found in id")
}

func (types *SObjectsInfo) GetTypeForPrefix(prefix string) (string, error) {
	if sobjectType, ok := types.Prefix2Type[prefix]; ok {
		return sobjectType, nil
	}
	return "", errors.New("sobject type not found for prefix")
}

func (types *SObjectsInfo) GetPrefixForType(sobjectType string) (string, error) {
	sobjectType = strings.ToUpper(sobjectType)
	if prefix, ok := types.Type2Prefix[sobjectType]; ok {
		return prefix, nil
	}
	return "", errors.New("sobject prefix not found for type")
}

func (types *SObjectsInfo) loadMaps() {
	type2prefix := map[string]string{}
	prefix2type := map[string]string{}
	csv := types.getTypesCsv()
	lines := strings.Split(csv, "\n")
	for i, line := range lines {
		if i == 0 {
			continue
		}
		parts := strings.Split(line, ", ")
		if len(parts) == 2 {
			sobjectType := parts[0]
			pref := parts[1]
			type2prefix[sobjectType] = pref
			prefix2type[pref] = sobjectType
		}
	}
	types.Type2Prefix = type2prefix
	types.Prefix2Type = prefix2type
}

func (types *SObjectsInfo) getTypesCsv() string {
	raw := `Entity, Prefix
ACCOUNT, 001
QUOTE, 0Q0
NOTE, 002
CONTACT, 003
USERS, 005
OPPORTUNITY, 006
ACTIVITY, 007
OPPORTUNITY_HISTORY, 008
FORECAST_ITEM, 00A
FILTER, 00B
DELETE_EVENT, 00C
ORGANIZATION, 00D
USER_ROLE, 00E
QUEUE, 00G
GROUPS, 00G
PARTNER, 00I
OPPORTUNITY_COMPETITOR, 00J
OPPORTUNITY_CONTACT_ROLE, 00K
CUSTOM_FIELD_DEFINITION, 00N
REPORT, 00O
ATTACHMENT, 00P
LEAD, 00Q
IMPORT_QUEUE, 00S
TASK, 00T
EVENT, 00U
EMAIL_TEMPLATE, 00X
EMAIL_TEMP, 00Y
COMMENTS, 00a
CUSTOM_RESOURCE_LINK, 00b
TRAINING, 00c
PROFILE, 00e
MH_BLUESHEET, 00f
MH_GOLDSHEET, 00g
LAYOUT, 00h
PRICEBOOK_MAPPING, 00i
PRICEBOOK_ENTRY_MAPPING, 00j
OPPORTUNITY_LINEITEM, 00k
FOLDER, 00l
EMAIL_ATTACHMENT_LOOKUP, 00m
EMAIL_ATTACHMENT_ARCHIVE, 00n
LINEITEM_SCHEDULE, 00o
USER_TEAM_MEMBER, 00p
OPP_TEAM_MEMBER, 00q
ACC_SHARE, 00r
ACC_SHARE_DEFAULT, 00s
OPP_SHARE, 00t
OPP_SHARE_DEFAULT, 00u
CAMPAIGN_MEMBER, 00v
PAYMENT_APPLICATION, 00w
BILLED_PRODUCT, 00x
PURCHASE_RULE, 00y
PURCHASE_RULE_ENTRY, 00z
CASE_SOLUTION, 010
GROUP_MEMBER, 011
RECORD_TYPE, 012
RECORD_TYPE_PICKLIST, 013
PROFILE_RECORD_TYPE, 014
DOCUMENT, 015
BRAND_TEMPLATE, 016
ENTITY_HISTORY, 017
EMAIL_STATUS, 018
BUSINESS_PROCESS, 019
BUSINESS_PROCESS_PICKLIST, 01A
LAYOUT_SECTION, 01B
LAYOUT_ITEM, 01C
PROFILE_LAYOUT, 01G
MAILMERGE_TEMPLATE, 01H
CUSTOM_ENTITY_DEFINITION, 01I
PICKLIST_MASTER, 01J
CURRENCY_TYPE, 01L
ACC_TEAM_MEMBER, 01M
ACTIVE_CONTENT, 01N
USER_UI_CONFIGURATION, 01O
PROFILE_TAB_CONFIGURATION, 01P
WORKFLOW_RULE, 01Q
RULE_FILTER, 01R
RULE_FILTER_ITEM, 01S
RULE_FILTER_ACTION, 01T
ACTION_ASSIGN_ESCALATE, 01U
ACTION_TASK, 01V
ACTION_EMAIL, 01W
ACTION_EMAIL_RECIPIENT, 01X
CAMPAIGN_MEMBER_STATUS, 01Y
DASHBOARD, 01Z
DASHBOARD_COMPONENT, 01a
FILTER_ITEM, 01b
FILTER_COLUMN, 01c
FOLDER_GROUPS, 01d
PICKLIST_SET, 01e
WEBEX_MEETING, 01f
API_QUERY, 01g
TRANSLATION, 01h
TRANSLATION_USER, 01i
TRANSLATION_VALUE, 01j
PROFILE_FLS_ITEM, 01k
ACTION_RESPONSE, 01l
BUSINESS_HOURS, 01m
CASE_SHARE, 01n
LEAD_SHARE, 01o
CUSTOM_TAB_DEFINITION, 01r
PRICEBOOK2, 01s
PRODUCT2, 01t
PRICEBOOK_ENTRY, 01u
PRICEBOOK_SHARE, 01v
OPP_UPDATE_REMINDER, 01w
OPP_UPDATE_REMINDER_STATS, 01x
CASE_SHARE_DEFAULT, 01y
CASE_ESCALATION, 01z
EVENT_ATTENDEE, 020
QUANTITY_FORECAST, 021
FISCAL_YEAR_SETTINGS, 022
APP_CALENDAR, 023
APP_CALENDAR_SHARING, 024
LIST_LAYOUT_ITEM, 025
PERIOD, 026
REVENUE_FORECAST, 027
OPPORTUNITY_OVERRIDE, 028
LINEITEM_OVERRIDE, 029
LEAD_SHARE_DEFAULT, 02A
LABEL_DEFINITION, 02B
LABEL_DATA, 02C
CASES_HISTORY2, 02D
HELP_SETTING, 02E
CUSTOM_FIELD_MAP, 02F
MH_GOLD_PROGRAM, 02H
MH_GOLD_INFORMATION, 02I
MH_GOLD_CONTACT, 02J
MH_GOLD_ACTION, 02K
MH_CUSTOMER_CRITERION, 02L
MH_GREENSHEET, 02M
MH_GREEN_GIVE_INFO, 02N
MH_GREEN_GET_INFO, 02O
MH_CONTACT_ROLE, 02P
MH_INFORMATION, 02Q
USER_PREFERENCE2, 02R
HTML_COMPONENT, 02S
CUSTOM_PAGE, 02T
CUSTOM_PAGE_ITEM, 02U
PAGE_COMPONENT, 02V
CUSTOM_PAGE_PROFILE, 02X
USER_COMPONENT_DATA, 02Y
ACCOUNT_CONTACT_ROLE, 02Z
CONTRACT_CONTACT_ROLE, 02a
COMPONENT_RESOURCE_LINK, 02b
DIVISION, 02d
DIVISION_WORKFLOW_RULE, 02e
DELEGATE_GROUP, 02f
DELEGATE_GROUP_MEMBER, 02g
DELEGATE_GROUP_GRANT, 02h
ASSET, 02i
PROFILE_ENTITY_PERMISSIONS, 02j
LIST_LAYOUT, 02k
OUTBOUND_QUEUE, 02l
CUSTOM_INDEX, 02m
CATEGORY_NODE, 02n
CATEGORY_DATA, 02o
DIV_TRANSFER_EVENT, 02p
LAYOUT_ITEM_COLUMN, 02q
OPPORTUNITY_ALERT, 02r
EMAIL_MESSAGE, 02s
EMAIL_ROUTING_ADDRESS, 02t
TAB_SET, 02u
TAB_SET_MEMBER, 02v
LOGIN_IP_RANGE, 02w
LOGIN_HOURS, 02x
REPORT_AGGREGATE, 02y
REPORT_COLOR_RANGE, 02z
PROFILE_TAB_SET, 030
USER_TAB_SET_MEMBER, 031
ACC_TERRITORY_RULE, 032
PROJECT, 033
PROJECT_MEMBER, 034
SELF_SERVICE_USER, 035
JOB_QUEUE, 036
REPORT_COLUMN, 037
REPORT_FILTER_ITEM, 038
REPORT_BREAK, 039
DEPENDENT_PICKLIST, 03a
PACKAGE_EXPORT, 03b
LAYOUT_RIGHT_PANEL, 03c
CUSTOM_SETUP_DEFINITION, 03e
CUSTOM_SETUP, 03f
REPORT_PARAM, 040
ACC_TERRITORY_ASSIGN, 041
ACC_TERR_ASSIGN_RULE_ITEM, 042
OUTBOUND_FIELD, 043
USER_TERRITORY, 04S
TERRITORY, 04T
DNB_ACCOUNT_MAPPING, 04U
DNB_FIELD, 04V
REVENUE_FORECAST_HISTORY, 04W
QUANTITY_FORECAST_HISTORY, 04X
CONTENTVERSION, 068
CONTENTDOCUMENT, 069
ENTITY_PERMISSION, 110
SFDC_PARTNER, 204
SFDC_DIVISION, 208
CASES, 500
SOLUTION, 501
BILLING_DIVISION, 600
BILLING_ORDER, 601
CURRENCY, 602
PLAN, 604
PRODUCT, 605
BILLING_ORDER_ITEM, 606
PLAN_PRODUCT, 607
CAMPAIGN, 701
FIELD_HISTORY, 737
UI_STYLE_DEFINITION, 766
UI_STYLE, 777
CONTRACT, 800
ORDERS, 801
ORDER_ITEM, 802
INVOICE, 803
INVOICE_ITEM, 804
PAYMENT, 805
APPROVAL, 806
URI_BLOCK_RULE, 807
CUSTOM_ENTITY_DATA, a00`
	return raw
}
