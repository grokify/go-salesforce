package salesforcefsdb

import (
	"errors"
	"path"
	"strings"

	"github.com/grokify/go-salesforce/sobjects"
	"github.com/grokify/mogo/os/osutil"
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
	ok, err := osutil.IsFile(path1, true)
	if err != nil {
		return "", err
	}
	if ok {
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
		ok2, err := osutil.IsFile(path2, true)
		if err != nil {
			return "", err
		}
		if ok2 {
			return path2, nil
		}
	}
	return "", errors.New("cannot find valid path for Sfdc Id")
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
		return "", errors.New("no SFDC source provided")
	}
	if len(format) > 0 {
		parts = append(parts, strings.ToUpper(format))
	} else {
		return "", errors.New("no SObject format provided")
	}
	dir := path.Join(parts...)
	return dir, nil
}
