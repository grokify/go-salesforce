package sobjects

import (
	"strings"
)

type IdSet struct {
	SObjectsInfo SObjectsInfo `json:"-"`
	IdMap        map[string]int
	IdMapByType  map[string]map[string]int
}

func NewIdSet() IdSet {
	set := IdSet{
		SObjectsInfo: NewSObjectsInfo(),
		IdMap:        map[string]int{},
		IdMapByType:  map[string]map[string]int{}}
	return set
}

func (set *IdSet) AddId(id string) {
	if len(id) < 1 {
		return
	}
	set.IdMap[id]++
	desc, err := set.SObjectsInfo.GetTypeForId(id)
	if err != nil {
		return
	}
	desc = strings.ToUpper(desc)
	if _, ok1 := set.IdMapByType[desc]; !ok1 {
		set.IdMapByType[desc] = map[string]int{}
	}
	if _, ok2 := set.IdMapByType[desc][id]; ok2 {
		set.IdMapByType[desc][id]++
	} else {
		set.IdMapByType[desc][id] = 1
	}
}

func (set *IdSet) GetIdsByType(sobjectType string) map[string]int {
	sobjectType = strings.ToUpper(sobjectType)
	if ids, ok := set.IdMapByType[sobjectType]; ok {
		return ids
	} else {
		return map[string]int{}
	}
}

func (set *IdSet) Merge(newSet IdSet) {
	for id := range newSet.IdMap {
		set.AddId(id)
	}
}

func (set *IdSet) MergeTypes(newSet IdSet, types []string) {
	for _, sObjectType := range types {
		ids := newSet.GetIdsByType(sObjectType)
		for id := range ids {
			set.AddId(id)
		}
	}
}

type IdSetMulti struct {
	Sets map[string]IdSet
}

func NewIdSetMulti() IdSetMulti {
	sets := IdSetMulti{Sets: map[string]IdSet{}}
	return sets
}

func (sets *IdSetMulti) MergeTypes(setName string, newSet IdSet, types []string) {
	if len(setName) < 1 {
		return
	}
	if _, ok := sets.Sets[setName]; !ok {
		sets.Sets[setName] = NewIdSet()
	}
	set := sets.Sets[setName]
	set.MergeTypes(newSet, types)
	sets.Sets[setName] = set
}
