package sobjects

import (
	"strings"
)

type IDSet struct {
	SObjectsInfo SObjectsInfo `json:"-"`
	IDMap        map[string]int
	IDMapByType  map[string]map[string]int
}

func NewIDSet() IDSet {
	set := IDSet{
		SObjectsInfo: NewSObjectsInfo(),
		IDMap:        map[string]int{},
		IDMapByType:  map[string]map[string]int{}}
	return set
}

func (set *IDSet) AddID(id string) {
	if len(id) < 1 {
		return
	}
	set.IDMap[id]++
	desc, err := set.SObjectsInfo.GetTypeForID(id)
	if err != nil {
		return
	}
	desc = strings.ToUpper(desc)
	if _, ok1 := set.IDMapByType[desc]; !ok1 {
		set.IDMapByType[desc] = map[string]int{}
	}
	set.IDMapByType[desc][id]++
}

func (set *IDSet) GetIDsByType(sobjectType string) map[string]int {
	sobjectType = strings.ToUpper(sobjectType)
	if ids, ok := set.IDMapByType[sobjectType]; ok {
		return ids
	} else {
		return map[string]int{}
	}
}

func (set *IDSet) Merge(newSet IDSet) {
	for id := range newSet.IDMap {
		set.AddID(id)
	}
}

func (set *IDSet) MergeTypes(newSet IDSet, types []string) {
	for _, sObjectType := range types {
		ids := newSet.GetIDsByType(sObjectType)
		for id := range ids {
			set.AddID(id)
		}
	}
}

type IDSetMulti struct {
	Sets map[string]IDSet
}

func NewIDSetMulti() IDSetMulti {
	sets := IDSetMulti{Sets: map[string]IDSet{}}
	return sets
}

func (sets *IDSetMulti) MergeTypes(setName string, newSet IDSet, types []string) {
	if len(setName) < 1 {
		return
	}
	if _, ok := sets.Sets[setName]; !ok {
		sets.Sets[setName] = NewIDSet()
	}
	set := sets.Sets[setName]
	set.MergeTypes(newSet, types)
	sets.Sets[setName] = set
}
