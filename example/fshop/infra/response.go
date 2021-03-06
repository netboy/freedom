// Code generated by 'freedom new-project fshop'
package infra

import (
	"encoding/json"
	"strconv"

	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/hero"
)

type JSONResponse struct {
	Code        int
	Error       error
	contentType string
	content     []byte
	Object      interface{}
}

func (jrep JSONResponse) Dispatch(ctx context.Context) {
	jrep.contentType = "application/json"
	var repData struct {
		Code  int         `json:"code"`
		Error string      `json:"error"`
		Data  interface{} `json:"data,omitempty"`
	}

	repData.Data = jrep.Object
	repData.Code = jrep.Code
	if jrep.Error != nil {
		repData.Error = jrep.Error.Error()
	}
	if repData.Error != "" && repData.Code == 0 {
		repData.Code = 501
	}
	ctx.Values().Set("code", strconv.Itoa(repData.Code))

	jrep.content, _ = json.Marshal(repData)
	ctx.Values().Set("response", string(jrep.content))
	hero.DispatchCommon(ctx, 0, jrep.contentType, jrep.content, nil, nil, true)
}
