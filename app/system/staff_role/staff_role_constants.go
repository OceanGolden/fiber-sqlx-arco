package staff_role

const (
	Table = "system_staff_role"

	CreatedSuccess = "角色授权成功！"
	CreatedFail    = "角色-菜单授权失败！"
	UpdatedSuccess = "修改角色成功！"
	UpdatedFail    = "分配员工角色失败！"
	DeletedSuccess = "删除角色成功！"
	DeletedFail    = "删除员工角色失败！"

	ErrorNotExist   = "角色不存在！"
	ErrorNameRepeat = "角色名称重复，请重新输入！"
	ErrorCodeRepeat = "角色编码重复，请重新输入！"
	ErrorExistMenu  = "该角色下有菜单权限，无法删除! "
	ErrorExistStaff = "该角色下有员工，无法删除! "
)
