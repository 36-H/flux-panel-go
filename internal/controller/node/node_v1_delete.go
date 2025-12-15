package node

import (
	"context"
	"flux-panel-go/internal/dao"
	"flux-panel-go/internal/model/localErr"

	"flux-panel-go/api/node/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	// 1. 验证节点是否存在
	exist, err := dao.Node.Ctx(ctx).Where(dao.Node.Columns().Id, req.ID).Exist()
	if err != nil || !exist {
		return nil, &localErr.CommonError{
			Code:   -1,
			ErrMsg: ERROR_NODE_NOT_FOUND,
		}
	}
	// TODO 2. 检查节点是否有正在运行的隧道

	// 3.删除节点
	_, err = dao.Node.Ctx(ctx).Where(dao.Node.Columns().Id, req.ID).Delete()
	if err != nil {
		return nil, &localErr.CommonError{
			Code:   -1,
			ErrMsg: ERROR_DELETE_MSG,
		}
	}
	r := v1.DeleteRes(SUCCESS_DELETE_MSG)
	return &r, nil
}
