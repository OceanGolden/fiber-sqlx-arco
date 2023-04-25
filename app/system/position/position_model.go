package position

type Position struct {
	ID        string `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Code      string `db:"code" json:"code"`
	Status    string `db:"status" json:"status"`
	Sort      int32  `db:"sort" json:"sort"`
	Remark    string `db:"remark" json:"remark,omitempty"`
	CreatedAt int64  `db:"created_at" json:"-"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
	CreatedBy string `db:"created_by" json:"-"`
	UpdatedBy string `db:"updated_by" json:"updated_by"`
}

type CreateRequest struct {
	Name   string `zh:"职位名称" json:"name" validate:"required,min=2,max=32"`
	Code   string `zh:"职位编码" json:"code" validate:"required,min=2,max=64"`
	Status string `zh:"状态" json:"status" validate:"required,oneof=Disable Enable"`
	Sort   int32  `zh:"排序" json:"sort" validate:"number,gt=0"`
	Remark string `zh:"备注" json:"remark" validate:"omitempty,max=128"`
}

type UpdateRequest struct {
	ID     string `zh:"唯一标识符" json:"id" validate:"required"`
	Name   string `zh:"职位名称" json:"name" validate:"required,min=2,max=32"`
	Code   string `zh:"职位编码" json:"code" validate:"required,min=2,max=64"`
	Status string `zh:"状态" json:"status" validate:"required,oneof=Disable Enable"`
	Sort   int32  `zh:"排序" json:"sort" validate:"omitempty,number,gt=0"`
	Remark string `zh:"备注" json:"remark" validate:"omitempty,max=128"`
}

type DeleteRequest struct {
	ID string `zh:"唯一标识符" json:"id" validate:"required"`
}

type WhereParams struct {
	Name     string `zh:"职位名称" query:"name" json:"name" validate:"omitempty,max=32"`
	Code     string `zh:"职位编码" query:"code" json:"code" validate:"omitempty,max=64"`
	Status   string `zh:"状态" query:"status" json:"status" validate:"omitempty,oneof=Disable Enable"`
	Remark   string `zh:"备注" query:"remark" json:"remark" validate:"omitempty,max=128"`
	PageSize uint64 `zh:"分页数量" query:"pageSize" json:"pageSize" validate:"omitempty,number,gt=0,max=50"`
	Current  uint64 `zh:"页数" query:"current" json:"current" validate:"omitempty,number,gt=0"`
}
