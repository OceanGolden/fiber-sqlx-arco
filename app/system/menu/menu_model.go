package menu

type Menu struct {
	ID         string  `db:"id" json:"id"`
	Name       string  `db:"name" json:"name"`
	ParentID   string  `db:"parent_id" json:"parent_id"`
	ParentIDS  string  `db:"parent_ids" json:"parent_ids"`
	Icon       string  `db:"icon" json:"icon"`
	Path       string  `db:"path" json:"path"`
	Permission string  `db:"permission" json:"permission"`
	Type       string  `db:"type" json:"type"`
	Method     string  `db:"method" json:"method"`
	Component  string  `db:"component" json:"component"`
	Link       string  `db:"link" json:"link"`
	Visible    string  `db:"visible" json:"visible"`
	Redirect   string  `db:"redirect" json:"redirect"`
	Status     string  `db:"status" json:"status"`
	Sort       uint32  `db:"sort" json:"sort"`
	Remark     string  `db:"remark" json:"remark,omitempty"`
	CreatedAt  int64   `db:"created_at" json:"-"`
	UpdatedAt  int64   `db:"updated_at" json:"updated_at"`
	CreatedBy  string  `db:"created_by" json:"-"`
	UpdatedBy  string  `db:"updated_by" json:"-"`
	Children   []*Menu `json:"children,omitempty"`
}

type CreateRequest struct {
	Name       string `zh:"菜单名称" json:"name" validate:"required,min=2,max=32"`
	ParentID   string `zh:"父级菜单" json:"parent_id" validate:"required"`
	Icon       string `zh:"菜单图标" json:"icon"`
	Path       string `zh:"路径" json:"path"`
	Permission string `zh:"权限" json:"permission"`
	Component  string `zh:"组件" json:"component"`
	Type       string `zh:"类型" json:"type" validate:"required,oneof=Catalog Menu Button"`
	Method     string `zh:"方法" json:"method"`
	Status     string `zh:"状态" json:"status" validate:"required,oneof=Disable Enable"`
	Sort       uint32 `zh:"排序" json:"sort" validate:"number,gt=0"`
	Remark     string `zh:"备注" json:"remark" validate:"omitempty,max=128"`
}

type UpdateRequest struct {
	ID         string `zh:"唯一标识符" json:"id" validate:"required"`
	Name       string `zh:"菜单名称" json:"name" validate:"required,min=2,max=32"`
	ParentID   string `zh:"父级菜单" json:"parent_id" validate:"required"`
	Icon       string `zh:"菜单图标" json:"icon"`
	Path       string `zh:"路径" json:"path"`
	Permission string `zh:"权限" json:"permission"`
	Component  string `zh:"组件" json:"component"`
	Type       string `zh:"类型" json:"type" validate:"required,oneof=Catalog Menu Button"`
	Method     string `zh:"方法" json:"method"`
	Status     string `zh:"状态" json:"status" validate:"required,oneof=Disable Enable"`
	Sort       uint32 `zh:"排序" json:"sort" validate:"number,gt=0"`
	Remark     string `zh:"备注" json:"remark" validate:"omitempty,max=128"`
}

type DeleteRequest struct {
	ID string `zh:"唯一标识符" json:"id" validate:"required"`
}

type WhereParams struct {
	Name     string `zh:"菜单名称" query:"name" json:"name" validate:"omitempty,max=32"`
	ParentID string `zh:"父级菜单" query:"parent_id" json:"parent_id,string" validate:"omitempty"`
	Status   string `zh:"状态" query:"status" json:"status" validate:"omitempty,oneof=Disable Enable"`
	Remark   string `zh:"备注" query:"remark" json:"remark" validate:"omitempty,max=128"`
	PageSize uint64 `zh:"分页数量" query:"pageSize" json:"pageSize" validate:"omitempty,number,gt=0,max=50"`
	Current  uint64 `zh:"页数" query:"current" json:"current" validate:"omitempty,number,gt=0"`
}
