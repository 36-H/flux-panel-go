package node

import (
	"context"
	"fmt"
	"time"

	v1 "flux-panel-go/api/node/v1"
	"flux-panel-go/internal/dao"
	"flux-panel-go/internal/model/do"
	"flux-panel-go/internal/model/localErr"

	"github.com/google/uuid"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	// 验证端口范围
	if err = validatePortRange(req.PortStart, req.PortEnd); err != nil {
		return nil, err
	}
	secret := uuid.New()
	_, err = dao.Node.Ctx(ctx).Data(do.Node{
		Name:        req.Name,
		Ip:          req.IP,
		ServerIp:    req.ServerIp,
		PortSta:     req.PortStart,
		PortEnd:     req.PortEnd,
		Secret:      secret,
		Status:      NODE_STATUS_ACTIVE,
		CreatedTime: time.Now().UnixMilli(),
		UpdatedTime: time.Now().UnixMilli(),
	}).Save()
	if err == nil {
		r := v1.CreateRes(SUCCESS_CREATE_MSG)
		res = &r
	} else {
		r := v1.CreateRes(ERROR_CREATE_MSG)
		res = &r
	}
	return
}

func validatePortRange(portStart, portEnd int) error {
	if portStart < 1 || portStart > 65535 || portEnd < 1 || portEnd > 65535 {
		return &localErr.CommonError{
			Code:   -2,
			ErrMsg: fmt.Sprintf(ERROR_PORT_RANGE_INVALID, portStart, portEnd),
		}
	}
	if portEnd < portStart {
		return &localErr.CommonError{
			Code:   -2,
			ErrMsg: ERROR_PORT_ORDER_INVALID,
		}
	}
	return nil
}
