package po

type LoginInfo struct {
	UserId   uint64 `db:"id"`
	Password string `db:"password"`
	BranchId uint64 `db:"branch_id"`
}
