package salesforcefsdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"

	httputil "github.com/grokify/gotilla/net/httputilmore"
)

type FsdbClient struct {
	Config     SalesforceClientConfig
	RestClient RestClient
}

func NewFsdbClient(cfg SalesforceClientConfig) FsdbClient {
	cl := FsdbClient{}
	cl.Config = cfg
	cl.RestClient = NewRestClient(cfg)
	return cl
}

func (fc *FsdbClient) GetPathForSfidAndType(sSfid string, sType string) (string, error) {
	sFilename := fc.GetFilenameForSfidAndType(sSfid, sType)
	sDir, err := fc.GetDirForType(sType)
	sPath := path.Join(sDir, sFilename)
	return sPath, err
}

func (fc *FsdbClient) GetFilenameForSfidAndType(sSfid string, sType string) string {
	filename := "sf_" + sType + "_" + sSfid + ".json"
	return filename
}

func (fc *FsdbClient) GetDirForType(sType string) (string, error) {
	sDir := path.Join(fc.Config.ConfigGeneral.DataDir, sType)
	err := os.MkdirAll(sDir, 0755)
	return sDir, err
}

func (fc *FsdbClient) GetSobjectForSfidAndType(sSfidTry string, sType string) (SobjectFsdb, error) {
	sobTry, err := fc.GetSobjectForSfidAndTypeFromLocal(sSfidTry, sType)
	if err == nil {
		if sobTry.Meta.HttpStatusCodeI32 == int32(200) && sobTry.Meta.EpochRetrievedSourceI64 > 0 {
			now := time.Now()
			iEpochNow := now.Unix()
			diff := iEpochNow - sobTry.Meta.EpochRetrievedSourceI64
			if diff < fc.Config.ConfigGeneral.MaxAgeSec {
				return sobTry, nil
			}
		}
	}
	sobTry, err = fc.GetSobjectForSfidAndTypeFromRemote(sSfidTry, sType)
	return sobTry, err
}

func (fc *FsdbClient) GetSobjectForSfidAndTypeFromLocal(sSfidTry string, sType string) (SobjectFsdb, error) {
	sobTry := NewSobjectFsdb()
	sPath, err := fc.GetPathForSfidAndType(sSfidTry, sType)
	if err != nil {
		return sobTry, err
	}
	if _, err := os.Stat(sPath); os.IsNotExist(err) {
		return sobTry, err
	}
	abData, err := ioutil.ReadFile(sPath)
	if err != nil {
		return sobTry, err
	}
	err = json.Unmarshal(abData, &sobTry)
	if err == nil && sobTry.Meta.HttpStatusCodeI32 == int32(301) && len(sobTry.Meta.RedirectSfidS) > 0 {
		sobTry2, err := fc.GetSobjectForSfidAndTypeFromLocal(sobTry.Meta.RedirectSfidS, sType)
		if err == nil {
			return sobTry2, nil
		}
	}
	return sobTry, err
}

func (fc *FsdbClient) GetSobjectForSfidAndTypeFromRemote(sSfidTry string, sType string) (SobjectFsdb, error) {
	if fc.Config.ConfigGeneral.FlagDisableRemote == true {
		err := errors.New("404 File Not Found")
		return NewSobjectFsdb(), err
	}
	resTry, err := fc.RestClient.GetSobjectResponseForSfidAndType(sSfidTry, sType)
	if err != nil {
		return NewSobjectFsdb(), err
	}
	sobTry := NewSobjectFsdbForResponse(resTry)

	if resTry.StatusCode == 404 && sType == "Account" {
		resOpp, err := fc.RestClient.GetSobjectResponseForSfidAndType(sSfidTry, "Opportunity")
		if err == nil {
			sobOpp := NewSobjectFsdbForResponse(resOpp)
			fc.WriteSobjectFsdb(sSfidTry, "Opportunity", sobOpp)
			sSfidAct := fmt.Sprintf("%s", sobOpp.Data["AccountId"])
			if len(sSfidAct) > 0 {
				resAct, err := fc.RestClient.GetSobjectResponseForSfidAndType(sSfidAct, "Account")
				if err != nil {
					return sobTry, err
				}
				sobAct := NewSobjectFsdbForResponse(resAct)
				fc.WriteSobjectFsdb(sSfidAct, "Account", sobAct)
				sobjAct301 := NewSobjectFsdb()
				sobjAct301.SetEpochRetrievedSource()
				sobjAct301.Meta.HttpStatusCodeI32 = int32(301)
				sobjAct301.Meta.RedirectSfidS = sSfidAct
				fc.WriteSobjectFsdb(sSfidTry, "Account", sobjAct301)
				return sobAct, nil
			}
		}
	}
	if resTry.StatusCode >= 400 {
		err := errors.New(resTry.Status)
		return sobTry, err
	}
	err = fc.WriteSobjectFsdb(sSfidTry, sType, sobTry)
	return sobTry, err
}

func (fc *FsdbClient) WriteSobjectFsdb(sSfid string, sType string, sobjectFsdb SobjectFsdb) error {
	if fc.Config.ConfigGeneral.FlagSaveFs == false {
		return nil
	}
	j, err := json.MarshalIndent(sobjectFsdb, "", "  ")
	if err != nil {
		return err
	}
	sPath, err := fc.GetPathForSfidAndType(sSfid, sType)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(sPath, j, 0755)
	if err != nil {
		return err
	}
	return nil
}

type SobjectFsdb struct {
	Meta SobjectFsdbMeta
	Data map[string]interface{}
}

func NewSobjectFsdb() SobjectFsdb {
	sobjectFsdb := SobjectFsdb{Data: map[string]interface{}{}, Meta: SobjectFsdbMeta{}}
	sobjectFsdb.Meta.EpochRetrievedSourceI64 = int64(0)
	sobjectFsdb.Meta.HttpStatusCodeI32 = int32(0)
	return sobjectFsdb
}

func NewSobjectFsdbForResponse(res *http.Response) SobjectFsdb {
	sobjectFsdb := NewSobjectFsdb()
	sobjectFsdb.LoadResponse(res)
	return sobjectFsdb
}

type SobjectFsdbMeta struct {
	EpochRetrievedSourceI64 int64
	HttpStatusCodeI32       int32
	RedirectSfidS           string
}

func (so *SobjectFsdb) SetEpochRetrievedSource() {
	now := time.Now()
	so.Meta.EpochRetrievedSourceI64 = now.Unix()
}

func (so *SobjectFsdb) LoadResponse(res *http.Response) error {
	body, err := httputil.ResponseBody(res)
	if err != nil {
		return err
	}
	msi := map[string]interface{}{}
	json.Unmarshal(body, &msi)
	so.Data = msi
	now := time.Now()
	so.Meta.EpochRetrievedSourceI64 = now.Unix()
	so.Meta.HttpStatusCodeI32 = int32(res.StatusCode)
	return nil
}
