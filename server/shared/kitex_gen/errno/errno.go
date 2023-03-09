// Code generated by thriftgo (0.2.5). DO NOT EDIT.

package errno

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type Err int64

const (
	Err_Success            Err = 0
	Err_BadRequest         Err = 10000
	Err_ParamsErr          Err = 10001
	Err_AuthorizeFail      Err = 10002
	Err_ServiceErr         Err = 20000
	Err_RPCAuthSrvErr      Err = 30000
	Err_AuthSrvErr         Err = 30001
	Err_RPCBlobSrvErr      Err = 40000
	Err_BlobSrvErr         Err = 40001
	Err_RPCCarSrvErr       Err = 50000
	Err_CarSrvErr          Err = 50001
	Err_RPCProfileSrvErr   Err = 60000
	Err_ProfileSrvErr      Err = 60001
	Err_RPCTripSrvErr      Err = 70000
	Err_TripSrvErr         Err = 70001
	Err_RecordNotFound     Err = 80000
	Err_RecordAlreadyExist Err = 80001
	Err_DirtyData          Err = 80003
)

func (p Err) String() string {
	switch p {
	case Err_Success:
		return "Success"
	case Err_BadRequest:
		return "BadRequest"
	case Err_ParamsErr:
		return "ParamsErr"
	case Err_AuthorizeFail:
		return "AuthorizeFail"
	case Err_ServiceErr:
		return "ServiceErr"
	case Err_RPCAuthSrvErr:
		return "RPCAuthSrvErr"
	case Err_AuthSrvErr:
		return "AuthSrvErr"
	case Err_RPCBlobSrvErr:
		return "RPCBlobSrvErr"
	case Err_BlobSrvErr:
		return "BlobSrvErr"
	case Err_RPCCarSrvErr:
		return "RPCCarSrvErr"
	case Err_CarSrvErr:
		return "CarSrvErr"
	case Err_RPCProfileSrvErr:
		return "RPCProfileSrvErr"
	case Err_ProfileSrvErr:
		return "ProfileSrvErr"
	case Err_RPCTripSrvErr:
		return "RPCTripSrvErr"
	case Err_TripSrvErr:
		return "TripSrvErr"
	case Err_RecordNotFound:
		return "RecordNotFound"
	case Err_RecordAlreadyExist:
		return "RecordAlreadyExist"
	case Err_DirtyData:
		return "DirtyData"
	}
	return "<UNSET>"
}

func ErrFromString(s string) (Err, error) {
	switch s {
	case "Success":
		return Err_Success, nil
	case "BadRequest":
		return Err_BadRequest, nil
	case "ParamsErr":
		return Err_ParamsErr, nil
	case "AuthorizeFail":
		return Err_AuthorizeFail, nil
	case "ServiceErr":
		return Err_ServiceErr, nil
	case "RPCAuthSrvErr":
		return Err_RPCAuthSrvErr, nil
	case "AuthSrvErr":
		return Err_AuthSrvErr, nil
	case "RPCBlobSrvErr":
		return Err_RPCBlobSrvErr, nil
	case "BlobSrvErr":
		return Err_BlobSrvErr, nil
	case "RPCCarSrvErr":
		return Err_RPCCarSrvErr, nil
	case "CarSrvErr":
		return Err_CarSrvErr, nil
	case "RPCProfileSrvErr":
		return Err_RPCProfileSrvErr, nil
	case "ProfileSrvErr":
		return Err_ProfileSrvErr, nil
	case "RPCTripSrvErr":
		return Err_RPCTripSrvErr, nil
	case "TripSrvErr":
		return Err_TripSrvErr, nil
	case "RecordNotFound":
		return Err_RecordNotFound, nil
	case "RecordAlreadyExist":
		return Err_RecordAlreadyExist, nil
	case "DirtyData":
		return Err_DirtyData, nil
	}
	return Err(0), fmt.Errorf("not a valid Err string")
}

func ErrPtr(v Err) *Err { return &v }
func (p *Err) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = Err(result.Int64)
	return
}

func (p *Err) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}
