package staff_role

type StaffRole struct {
	ID        string `db:"id" json:"id"`
	StaffID   string `db:"staff_id" json:"staff_id"`
	RoleID    string `db:"role_id" json:"role_id"`
	CreatedAt int64  `db:"created_at" json:"-"`
	CreatedBy string `db:"created_by" json:"-"`
}

type Request struct {
	StaffID string   `zh:"员工ID" json:"staff_id" validate:"required"`
	RoleIDs []string `zh:"角色ID"  json:"role_ids" validate:"omitempty"`
}

type DeleteRequest struct {
	StaffID string `zh:"员工ID" json:"staff_id" validate:"required"`
}

type WhereParams struct {
	StaffID  string   `zh:"用户ID" query:"staff_id" json:"staff_id" validate:"omitempty"`
	StaffIDs []string `zh:"用户ID列表" query:"staff_ids" json:"staff_ids" validate:"omitempty"`
	RoleID   string   `zh:"角色ID" query:"role_id" json:"role_id" validate:"omitempty"`
	RoleIDs  []string `zh:"角色ID列表" query:"role_ids" json:"role_ids" validate:"omitempty"`
}
