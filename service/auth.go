package service

import (
	"go-project/global"
	"go-project/model/po"
	"go-project/model/request"
	"golang.org/x/crypto/bcrypt"
)

func PasswordLogin(req *request.LoginRequest) (string, error) {
	var loginInfo po.LoginInfo
	err := global.MySqlx.Get(&loginInfo, "select su.user_id,su.password from system_users_account sua inner join system_users su on sua.user_id=su.user_id where sua.account=? and sua.status=1 and  su.status =1  limit 1;", req.Username)
	if err != nil {
		return "用户名不存在", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(loginInfo.Password), []byte(req.Password))
	if err != nil {
		return "密码错误", err
	}
	return "登录成功", nil
}
