package salesforcefsdb

import (
	"errors"
	"path"
	"strings"

	"github.com/grokify/go-salesforce/sobjects"
	"github.com/grokify/simplego/io/ioutilmore"
)

type FsdbPathUtil struct {
	BaseDir      string
	SObjectsInfo sobjects.SObjectsInfo
}

func NewFsdbPathUtil(baseDir string) FsdbPathUtil {
	fsdb := FsdbPathUtil{BaseDir: baseDir}
	fsdb.SObjectsInfo = sobjects.NewSObjectsInfo()
	return fsdb
}

func (fsdb *FsdbPathUtil) GetPathForId(id string, source string, format string) (string, error) {
	dir, err := fsdb.GetDirForId(id, source, format)
	if err != nil {
		return "", err
	}
	filename, err := fsdb.GetFileForId(id, format)
	if err != nil {
		return "", err
	}
	filepath := path.Join(dir, filename)
	return filepath, nil
}

func (fsdb *FsdbPathUtil) GetValidPathForId(id string, source string, format string) (string, error) {
	path1, err := fsdb.GetPathForId(id, source, format)
	if err != nil {
		return "", err
	}
	ok, err := ioutilmore.IsFileWithSizeGtZero(path1)
	if err != nil {
		return "", err
	}
	if ok == true {
		return path1, nil
	}
	if len(id) == 18 {
		id15, err := fsdb.SObjectsInfo.GetId15ForId(id)
		if err != nil {
			return "", err
		}
		path2, err := fsdb.GetPathForId(id15, source, format)
		if err != nil {
			return "", err
		}
		ok2, err := ioutilmore.IsFileWithSizeGtZero(path2)
		if err != nil {
			return "", err
		}
		if ok2 == true {
			return path2, nil
		}
	}
	return "", errors.New("Cannot find valid path for Sfdc Id")
}

func (fsdb *FsdbPathUtil) GetFileForId(id string, format string) (string, error) {
	filename := "sfid_" + id + "." + strings.ToLower(format)
	return filename, nil
}

func (fsdb *FsdbPathUtil) GetDirForId(id string, source string, format string) (string, error) {
	parts := []string{fsdb.BaseDir}
	// Add SObject Desc
	sobjectDesc, err := fsdb.SObjectsInfo.GetTypeForId(id)
	if err != nil {
		return "", err
	} else {
		parts = append(parts, strings.ToUpper(sobjectDesc))
	}
	// Add SObject Source
	if len(source) > 0 {
		parts = append(parts, strings.ToUpper(source))
	} else {
		return "", errors.New("No SFDC Source Provided")
	}
	if len(format) > 0 {
		parts = append(parts, strings.ToUpper(format))
	} else {
		return "", errors.New("No SObject Format Provided")
	}
	dir := path.Join(parts...)
	return dir, nil
}
