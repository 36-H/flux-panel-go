package node

import (
	"context"
	"time"

	v1 "flux-panel-go/api/node/v1"
	"flux-panel-go/internal/dao"
	"flux-panel-go/internal/model/do"

	"github.com/google/uuid"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	// 验证端口范围
	if err = validatePortRange(req.PortSta, req.PortEnd); err != nil {
		return nil, err
	}
	secret := uuid.New()
	_, err = dao.Node.Ctx(ctx).Data(do.Node{
		Name:        req.Name,
		Ip:          req.IP,
		ServerIp:    req.ServerIp,
		PortSta:     req.PortSta,
		PortEnd:     req.PortEnd,
		Secret:      secret,
		Status:      NodeStatusActive,
		CreatedTime: time.Now().UnixMilli(),
		UpdatedTime: time.Now().UnixMilli(),
	}).Save()
	if err == nil {
		r := v1.CreateRes(SuccessCreateMsg)
		res = &r
	} else {
		r := v1.CreateRes(ErrorCreateMsg)
		res = &r
	}
	return
}
