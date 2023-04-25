package role_menu

type RoleMenu struct {
	RoleID    string `db:"role_id" json:"role_id"`
	MenuID    string `db:"menu_id" json:"menu_id"`
	CreatedAt int64  `db:"created_at" json:"-"`
	CreatedBy string `db:"created_by" json:"-"`
}

type Request struct {
	RoleID  string   `zh:"角色唯一标识符" json:"role_id" validate:"required"`
	MenuIDs []string `zh:"菜单ID列表" json:"menu_ids" validate:"required"`
}

type DeleteRequest struct {
	RoleID string `zh:"角色唯一标识符" json:"role_id" validate:"required"`
}

type WhereParams struct {
	RoleID  string   `zh:"角色ID" query:"role_id" json:"role_id" validate:"omitempty"`
	RoleIDs []string `zh:"角色ID列表" query:"role_ids" json:"role_ids" validate:"omitempty"`
}
