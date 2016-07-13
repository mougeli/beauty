package beauty

import (
	"encoding/json"
	"errors"
)

import (
	"gopkg.in/macaron.v1"
	"runtime"
)

type Render interface {
	OK(data interface{})
	Error(err interface{})
	Interface(data interface{})
}

type BeautyRender struct {
	Ctx *macaron.Context
}

func StackTrace(all bool) string {
	// Reserve 10K buffer at first
	buf := make([]byte, 10240)
	size := runtime.Stack(buf, all)
	return string(buf[0:size])
}

// 支持string,error,ErrorResponse三种参数
func (this BeautyRender) Error(err interface{}) {

	if logger.Level >= DebugLevel {
		logger.Debug(StackTrace(false))
	}

	switch err.(type) {
	case ErrResp:
		this.Ctx.JSON(200, err)
	case error:
		this.Ctx.JSON(200, NewUnknownErrResp(err.(error)))
	case string:
		this.Ctx.JSON(200, NewUnknownErrResp(errors.New(err.(string))))
	default:
		this.Ctx.JSON(200, UnknownErrResp)
	}
}

// 支持[]byte或者interface{}
func (this BeautyRender) OK(data interface{}) {
	switch data.(type) {
	case []byte:
		result := make(map[string]interface{})
		if err := json.Unmarshal(data.([]byte), &result); err != nil {
			this.Ctx.RawData(200, data.([]byte))
			return
		}
		this.Ctx.JSON(200, DataResp{
			Resp: OK,
			Data: result,
		})
	default:
		this.Ctx.JSON(200, DataResp{
			Resp: OK,
			Data: data,
		})
	}
}

// 支持Error和OK两种
func (this BeautyRender) Interface(data interface{}) {
	var bytes []byte
	switch data.(type) {
	case ErrResp:
		this.Error(data)
		return
	case error:
		this.Error(data)
		return
	case string:
		bytes = []byte(data.(string))
	case []byte:
		bytes = data.([]byte)
	}

	result := make(map[string]interface{})
	err := json.Unmarshal(bytes, &result)
	if err != nil {
		// 不是JSON
		this.Error(string(bytes))
		return
	}
	// 返回JSON
	this.OK(result)
}

// 注册用
func Renderer() macaron.Handler {
	return func(ctx *macaron.Context) {
		ctx.MapTo(&BeautyRender{ctx}, (*Render)(nil))
	}
}
