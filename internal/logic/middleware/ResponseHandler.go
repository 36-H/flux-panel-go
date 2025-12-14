package middleware

import (
	"errors"
	"flux-panel-go/internal/model"
	"flux-panel-go/internal/model/localErr"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// ResponseHandler 统一响应处理中间件
func ResponseHandler(r *ghttp.Request) {
	// 执行后续的逻辑
	r.Middleware.Next()

	// 如果已经有返回内容，则不做处理
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		code = gcode.CodeOK.Code()
		res  = r.GetHandlerResponse()
		err  = r.GetError()
	)

	if err != nil {
		var commonErr *localErr.CommonError
		if errors.As(err, &commonErr) {
			code = commonErr.Code
			msg = commonErr.ErrMsg
		} else {
			// 获取错误码
			code = gerror.Code(err).Code()

			// 如果是系统级错误（无特定错误码），记录详细日志
			if code == gcode.CodeNil.Code() {
				code = gcode.CodeInternalError.Code()
				// 记录堆栈信息到日志，方便排查
				g.Log().Error(r.Context(), err)
				// 生产环境可能不想把具体报错暴露给前端，可修改 msg 为 "服务器内部错误"
				// msg = "服务器内部错误"
				msg = err.Error()
			} else {
				// 业务错误，直接返回错误信息
				msg = err.Error()
			}
		}
	} else {
		msg = "操作成功"
	}

	r.Response.WriteJson(model.BaseRes{
		Code:      code,
		Msg:       msg,
		TimeStamp: time.Now().UnixMilli(),
		Data:      res,
	})
}
