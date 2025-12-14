package user

import (
	"context"

	v1 "flux-panel-go/api/user/v1"
	"flux-panel-go/internal/dao"
	"flux-panel-go/internal/model/do"
	"flux-panel-go/internal/model/entity"
	"flux-panel-go/internal/model/localErr"
	"flux-panel-go/internal/utils"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	//1.验证账号与密码
	loginValidationResult := loginValidation(ctx, req)
	if loginValidationResult.HasError {
		return nil, &localErr.CommonError{
			Code:   -1,
			ErrMsg: loginValidationResult.ErrorMessage,
		}
	}
	user := loginValidationResult.User
	//2.生成token
	token, err := utils.GenerateToken(user)
	if err != nil {
		return nil, err
	}
	//3.判断是否需要修改密码
	requirePasswordChange := isDefaultCredentials(req.Username, req.Password)
	res = &v1.LoginRes{
		Token:                 token,
		Name:                  user.User,
		RoleId:                user.RoleId,
		RequirePasswordChange: requirePasswordChange,
	}
	return res, nil
}

func isDefaultCredentials(username, password string) bool {
	return username == DefaultUsername && password == DefaultPassword
}

func loginValidation(ctx context.Context, req *v1.LoginReq) *LoginValidationResult {
	user := new(entity.User)
	err := dao.User.Ctx(ctx).Where(do.User{
		User: req.Username,
	}).Scan(user)
	if err != nil || !(user.Pwd == utils.MD5(req.Password)) {
		return errorLoginValidationResult(ErrorLoginCredentials)
	}
	if user.Status == UserStatusDisabled {
		return errorLoginValidationResult(ErrorAccountDisabled)
	}
	return successLoginValidationResult(user)
}

func successLoginValidationResult(user *entity.User) *LoginValidationResult {
	return newLoginValidationResult(false, "", user)
}

func errorLoginValidationResult(errorMessage string) *LoginValidationResult {
	return newLoginValidationResult(true, errorMessage, nil)
}

func newLoginValidationResult(hasError bool, errorMessage string, user *entity.User) *LoginValidationResult {
	return &LoginValidationResult{
		HasError:     hasError,
		ErrorMessage: errorMessage,
		User:         user,
	}
}
