package tunnel

import (
	"context"
	"flux-panel-go/internal/dao"
	"flux-panel-go/internal/model/entity"
	"flux-panel-go/internal/model/localErr"

	"flux-panel-go/api/tunnel/v1"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	//验证隧道是否存在
	tunnel := &entity.Tunnel{}
	err = dao.Tunnel.Ctx(ctx).Where(dao.Tunnel.Columns().Id, req.Id).Scan(tunnel)
	if err != nil || tunnel == nil {
		return nil, &localErr.CommonError{
			Code:   -1,
			ErrMsg: ErrorTunnelNotFound,
		}
	}
	//检测隧道名称是否存在
	exist, err := dao.Tunnel.Ctx(ctx).Where(dao.Tunnel.Columns().Name, req.Name).WhereNot(dao.Tunnel.Columns().Id, req.Id).Exist()
	if err != nil || exist {
		return nil, &localErr.CommonError{
			Code:   -1,
			ErrMsg: ErrorTunnelNameExists,
		}
	}
	IsForwardNeedUp := tunnel.TcpListenAddr != req.TcpListenAddr || tunnel.UdpListenAddr != req.UdpListenAddr || tunnel.Protocol != req.Protocol || tunnel.InterfaceName != req.InterfaceName
	//允许更新的字段
	tunnel.Name = req.Name
	tunnel.Flow = req.Flow
	tunnel.TcpListenAddr = req.TcpListenAddr
	tunnel.UdpListenAddr = req.UdpListenAddr
	tunnel.TrafficRatio = req.TrafficRatio
	tunnel.Protocol = req.Protocol
	tunnel.InterfaceName = req.InterfaceName
	_, err = dao.Tunnel.Ctx(ctx).Where(dao.Tunnel.Columns().Id, req.Id).Update(tunnel)
	if err != nil {
		return nil, &localErr.CommonError{
			Code:   -1,
			ErrMsg: ErrorTunnelUpdateFailed,
		}
	}
	if IsForwardNeedUp {
		//TODO 更新转发
	}
	r := v1.UpdateRes(SuccessUpdateMsg)
	return &r, nil
}
