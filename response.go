package beauty

import (
	"encoding/json"
	"errors"
)

var (
	UnknownCode int64 = -1
)

var (
	UnknownErr error = errors.New("unknown err")
)

var (
	OK             Resp    = Resp{Code: 0, Msg: "OK"}
	UnknownErrResp ErrResp = NewUnknownErrResp(UnknownErr)
)

type Resp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type DataResp struct {
	Resp
	Data interface{} `json:"data"`
}

type ErrResp struct {
	Resp
}

func (e ErrResp) Error() string {
	bytes, err := json.Marshal(e)
	if err != nil {
		return err.Error()

	}
	return string(bytes)
}

func (e *ErrResp) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

type Response interface {
	Error() string
	Marshal() ([]byte, error)
}

func NewUnknownErrResp(err error) (resp ErrResp) {
	return NewErrResp(UnknownCode, err)
}

func NewErrResp(code int64, err error) (resp ErrResp) {
	return ErrResp{Resp{code, err.Error()}}
}
