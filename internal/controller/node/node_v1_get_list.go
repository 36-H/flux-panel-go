package node

import (
	"context"

	v1 "flux-panel-go/api/node/v1"
	"flux-panel-go/internal/dao"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	res = &v1.GetListRes{}
	err = dao.Node.Ctx(ctx).Scan(&res.Nodes)
	for _, node := range res.Nodes {
		node.Secret = ""
	}
	return
}
