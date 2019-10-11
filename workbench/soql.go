package workbench

import (
	"strings"
)

func SliceToSql(slice []string) string {
	newSlice := []string{}
	for _, el := range slice {
		newSlice = append(newSlice, "'"+el+"'")
	}
	return strings.Join(newSlice, ",")
}

func SliceToSqls(slice []string) []string {
	max := 18000
	strIdx := 0
	newSlicesWip := [][]string{}
	newSlicesWip = append(newSlicesWip, []string{})
	for _, el := range slice {
		newStr := "'" + el + "'"
		if (LenStringForSlice(newSlicesWip[strIdx], ",") + len(newStr)) > max {
			newSlicesWip = append(newSlicesWip, []string{})
			strIdx += 1
		}
		newSlicesWip[strIdx] = append(newSlicesWip[strIdx], newStr)
	}
	newSlices := []string{}
	for _, slice := range newSlicesWip {
		newSlices = append(newSlices, strings.Join(slice, ","))
	}
	return newSlices
}

func LenStringForSlice(slice []string, sep string) int {
	return len(strings.Join(slice, sep))
}
