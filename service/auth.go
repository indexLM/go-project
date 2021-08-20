package service

import (
	"go-project/global"
	"go-project/model/po"
	"go-project/model/req"
	"go-project/utils"
	"golang.org/x/crypto/bcrypt"
)

func PasswordLogin(req *req.LoginRequest) (string, error) {
	var loginInfo po.LoginInfo
	err := global.MySqlx.Get(&loginInfo, "select id,branch_id,password from his_branch_admin hba where username=? and hba.is_delete=0;", req.Username)
	if err != nil {
		return "用户名不存在", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(loginInfo.Password), []byte(req.Password))
	if err != nil {
		return "密码错误", err
	}
	generate, err := utils.JwtGenerate(loginInfo.UserId, loginInfo.BranchId)
	if err != nil {
		return generate, err
	}
	return generate, nil
}
