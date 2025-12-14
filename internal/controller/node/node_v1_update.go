package node

import (
	"context"
	"time"

	v1 "flux-panel-go/api/node/v1"
	"flux-panel-go/internal/dao"
	"flux-panel-go/internal/model/do"
	"flux-panel-go/internal/model/localErr"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	// 1. 验证节点是否存在
	exist, err := dao.Node.Ctx(ctx).Where(dao.Node.Columns().Id, req.ID).Exist()
	if err != nil || !exist {
		return nil, &localErr.CommonError{
			Code:   -2,
			ErrMsg: ERROR_NODE_NOT_FOUND,
		}
	}
	//1.1 如果节点在线 且传入更新的 http/tls/socks 任意一项与数据库不一致，则通过 WS 通知节点更新设置

	// 2. 构建更新对象并执行更新
	_, err = dao.Node.Ctx(ctx).Where(dao.Node.Columns().Id, req.ID).Data(do.Node{
		Name:        req.Name,
		Ip:          req.IP,
		ServerIp:    req.ServerIp,
		PortSta:     req.PortSta,
		PortEnd:     req.PortEnd,
		Http:        req.Http,
		Tls:         req.Tls,
		Socks:       req.Socks,
		UpdatedTime: time.Now().UnixMilli(),
	}).Update()
	//TODO 更新隧道入口ip
	//TODO 更新服务器出口ip
	if err != nil {
		return nil, &localErr.CommonError{
			Code:   -1,
			ErrMsg: ERROR_UPDATE_MSG,
		}
	}
	r := v1.UpdateRes(SUCCESS_UPDATE_MSG)
	return &r, nil
}
