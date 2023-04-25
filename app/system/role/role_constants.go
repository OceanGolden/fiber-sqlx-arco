package role

const (
	Table = "system_role"

	CreatedSuccess = "添加角色成功！"
	CreatedFail    = "添加角色失败！"
	UpdatedSuccess = "修改角色成功！"
	UpdatedFail    = "修改角色失败！"
	DeletedSuccess = "删除角色成功！"
	DeletedFail    = "删除角色失败！"

	ErrorNotExist   = "角色不存在！"
	ErrorNameRepeat = "角色名称重复，请重新输入！"
	ErrorCodeRepeat = "角色编码重复，请重新输入！"
	ErrorExistMenu  = "该角色下有菜单权限，无法删除! "
	ErrorExistStaff = "该角色下有员工，无法删除! "
)
