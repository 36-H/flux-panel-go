package tunnel

import (
	"context"
	"flux-panel-go/internal/dao"
	"flux-panel-go/internal/model/entity"
	"flux-panel-go/internal/model/localErr"
	"time"

	"flux-panel-go/api/tunnel/v1"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	//验证隧道唯一性
	exist, err := dao.Tunnel.Ctx(ctx).Where(dao.Tunnel.Columns().Name, req.Name).Exist()
	if err != nil || exist {
		return nil, &localErr.CommonError{
			Code:   -1,
			ErrMsg: ErrorTunnelNameExists,
		}
	}
	//验证入口节点是否存在
	inNode := new(entity.Node)
	err = dao.Node.Ctx(ctx).Where(dao.Node.Columns().Id, req.InNodeId).Scan(inNode)
	if err != nil || inNode == nil {
		return nil, &localErr.CommonError{
			Code:   -1,
			ErrMsg: ErrorInNodeNotFound,
		}
	}
	if inNode.Status != NodeStatusOnline {
		return nil, &localErr.CommonError{
			Code:   -1,
			ErrMsg: ErrorInNodeOffline,
		}
	}
	//构建隧道实体
	tunnel := &entity.Tunnel{
		Name:          req.Name,
		InNodeId:      req.InNodeId,
		InIp:          inNode.Ip,
		Type:          req.Type,
		Flow:          req.Flow,
		TrafficRatio:  req.TrafficRatio,
		InterfaceName: req.InterfaceName,
		TcpListenAddr: req.TcpListenAddr,
		UdpListenAddr: req.UdpListenAddr,
		Status:        TunnelStatusActive,
		CreatedTime:   time.Now().UnixMilli(),
		UpdatedTime:   time.Now().UnixMilli(),
	}
	if req.Type == TunnelTypeTunnelForward {
		tunnel.Protocol = req.Protocol
	} else {
		tunnel.Protocol = ""
	}
	//根据隧道类型设置出口IP
	if req.Type == TunnelTypePortForward {
		tunnel.OutNodeId = tunnel.InNodeId
		tunnel.OutIp = inNode.ServerIp
	} else if req.Type == TunnelTypeTunnelForward {
		if req.OutNodeId == 0 {
			return nil, &localErr.CommonError{
				Code:   -1,
				ErrMsg: ErrorOutNodeRequired,
			}
		}
		//入口和出口不能是同一个节点
		if req.InNodeId == req.OutNodeId {
			return nil, &localErr.CommonError{
				Code:   -1,
				ErrMsg: ErrorSameNodeNotAllowed,
			}
		}
		//验证出口节点是否存在
		outNode := new(entity.Node)
		err = dao.Node.Ctx(ctx).Where(dao.Node.Columns().Id, req.OutNodeId).Scan(outNode)
		if err != nil || outNode == nil {
			return nil, &localErr.CommonError{
				Code:   -1,
				ErrMsg: ErrorOutNodeNotFound,
			}
		}
		if outNode.Status != NodeStatusOnline {
			return nil, &localErr.CommonError{
				Code:   -1,
				ErrMsg: ErrorOutNodeOffline,
			}
		}
		tunnel.OutNodeId = req.OutNodeId
		tunnel.OutIp = outNode.ServerIp
	}
	//创建隧道
	_, err = dao.Tunnel.Ctx(ctx).Save(tunnel)
	if err != nil {
		return nil, &localErr.CommonError{
			Code:   -1,
			ErrMsg: ErrorCreateMsg,
		}
	}
	r := v1.CreateRes(SuccessCreateMsg)
	return &r, nil
}
