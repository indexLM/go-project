package po

type LoginInfo struct {
	UserId   uint64 `db:"user_id"`
	Password string `db:"password"`
	NickName string
}
