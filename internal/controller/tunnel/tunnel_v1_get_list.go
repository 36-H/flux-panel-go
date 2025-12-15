package tunnel

import (
	"context"
	"flux-panel-go/internal/dao"

	"flux-panel-go/api/tunnel/v1"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	res = &v1.GetListRes{}
	err = dao.Tunnel.Ctx(ctx).Scan(&res.Tunnels)
	return
}
