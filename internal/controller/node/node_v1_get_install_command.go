package node

import (
	"context"
	"flux-panel-go/internal/dao"
	"flux-panel-go/internal/model/entity"
	"flux-panel-go/internal/model/localErr"

	"flux-panel-go/api/node/v1"
)

func (c *ControllerV1) GetInstallCommand(ctx context.Context, req *v1.GetInstallCommandReq) (res *v1.GetInstallCommandRes, err error) {
	// 1. 验证节点是否存在
	exist, err := dao.Node.Ctx(ctx).Where(dao.Node.Columns().Id, req.ID).Exist()
	if err != nil || !exist {
		return nil, &localErr.CommonError{
			Code:   -1,
			ErrMsg: ErrorNodeNotFound,
		}
	}
	node := new(entity.Node)
	err = dao.Node.Ctx(ctx).Where(dao.Node.Columns().Id, req.ID).Scan(node)
	if err != nil {
		return nil, &localErr.CommonError{
			Code:   -1,
			ErrMsg: ErrorNodeNotFound,
		}
	}
	// 2. 构建安装命令
	command, err := buildInstallCommand(node, ctx)
	if err != nil {
		return nil, err
	}
	r := v1.GetInstallCommandRes(command)
	return &r, nil
}

func buildInstallCommand(node *entity.Node, ctx context.Context) (string, error) {
	viteConfig := new(entity.ViteConfig)
	// 从数据库中查询 节点的Vite配置
	err := dao.ViteConfig.Ctx(ctx).Where(dao.ViteConfig.Columns().Name, "ip").Scan(viteConfig)
	if err != nil {
		return "", &localErr.CommonError{
			Code:   -1,
			ErrMsg: ErrorViteIpConfig,
		}
	}
	// TODO 从配置中获取下载URL
	//downloadUrl := g.Cfg().MustGet(ctx, "agent.downloadUrl").String()
	//common := "curl -L " + downloadUrl + " -o install.sh && chmod +x install.sh && ./install.sh"
	return "TODO 未完成代码", nil
}
