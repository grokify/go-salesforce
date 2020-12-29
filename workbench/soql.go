package workbench

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/grokify/simplego/type/stringsutil"
)

var maxInsertLength = 18000

var rxSplitLines = regexp.MustCompile(`(\r\n|\r|\n)`)

func SplitTextLines(text string) []string {
	return rxSplitLines.Split(text, -1)
}

func ReadFileCSVToSQLs(filename, sqlFormat string, skipHeader bool) ([]string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}
	lines := strings.Split(string(bytes), "\n")
	if len(lines) == 0 {
		return []string{}, nil
	}
	if skipHeader {
		lines = lines[1:]
	}
	if len(lines) == 0 {
		return []string{}, nil
	}
	values := stringsutil.SliceCondenseSpace(
		strings.Split(string(bytes), "\n"), true, true)
	sqls := BuildSQLsInStrings(sqlFormat, values)
	return sqls, nil
}

func BuildSQLsInStrings(sqlFormat string, values []string) []string {
	sqls := []string{}
	sqlIns := SliceToSQLs(values)
	for _, sqlIn := range sqlIns {
		sqls = append(sqls, fmt.Sprintf(sqlFormat, sqlIn))
	}
	return sqls
}

func SliceToSQL(slice []string) string {
	newSlice := []string{}
	for _, el := range slice {
		newSlice = append(newSlice, "'"+el+"'")
	}
	return strings.Join(newSlice, ",")
}

func SliceToSQLs(slice []string) []string {
	max := maxInsertLength
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
