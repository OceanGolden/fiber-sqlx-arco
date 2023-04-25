package casbin_adapter

//
//import (
//	"fiber-sqlx-arco/app/system/role"
//	"log"
//)
//
//type RoleManager interface {
//	Clear() error
//	// AddLink adds the inheritance link between two roles. role: name1 and role: name2.
//	// domain is a prefix to the roles (can be used for other purposes).
//	AddLink(name1 string, name2 string, domain ...string) error
//	// Deprecated: BuildRelationship is no longer required
//	BuildRelationship(name1 string, name2 string, domain ...string) error
//	// DeleteLink deletes the inheritance link between two roles. role: name1 and role: name2.
//	// domain is a prefix to the roles (can be used for other purposes).
//	DeleteLink(name1 string, name2 string, domain ...string) error
//	// HasLink determines whether a link exists between two roles. role: name1 inherits role: name2.
//	// domain is a prefix to the roles (can be used for other purposes).
//	HasLink(name1 string, name2 string, domain ...string) (bool, error)
//	// GetRoles gets the roles that a user inherits.
//	// domain is a prefix to the roles (can be used for other purposes).
//	GetRoles(name string, domain ...string) ([]string, error)
//	// GetUsers gets the users that inherits a role.
//	// domain is a prefix to the users (can be used for other purposes).
//	GetUsers(name string, domain ...string) ([]string, error)
//	// GetDomains gets domains that a user has
//	GetDomains(name string) ([]string, error)
//	// GetAllDomains gets all domains
//	GetAllDomains() ([]string, error)
//	// PrintRoles prints all the roles to log.
//	PrintRoles() error
//	// SetLogger sets role manager's logger.
//	SetLogger(logger log.Logger)
//	// Match matches the domain with the pattern
//	Match(str string, pattern string) bool
//	// AddMatchingFunc adds the matching function
//	AddMatchingFunc(name string, fn MatchingFunc)
//	// AddDomainMatchingFunc adds the domain matching function
//	AddDomainMatchingFunc(name string, fn MatchingFunc)
//}
//
//type roleManager struct {
//	allRoles          []*role.Role
//	maxHierarchyLevel int
//	roleService       role.Service
//	//staffRoleService  staff_role.Service
//}
//
//func NewRoleManager() RoleManager {
//	return &roleManager{}
//}
//
//func (r *roleManager) Clear() error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (r *roleManager) AddLink(name1 string, name2 string, domain ...string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (r *roleManager) BuildRelationship(name1 string, name2 string, domain ...string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (r *roleManager) DeleteLink(name1 string, name2 string, domain ...string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (r *roleManager) HasLink(name1 string, name2 string, domain ...string) (bool, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (r *roleManager) GetRoles(name string, domain ...string) ([]string, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (r *roleManager) GetUsers(name string, domain ...string) ([]string, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (r *roleManager) GetDomains(name string) ([]string, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (r *roleManager) GetAllDomains() ([]string, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (r *roleManager) PrintRoles() error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (r *roleManager) SetLogger(logger log.Logger) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (r *roleManager) Match(str string, pattern string) bool {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (r *roleManager) AddMatchingFunc(name string, fn interface{}) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (r *roleManager) AddDomainMatchingFunc(name string, fn interface{}) {
//	//TODO implement me
//	panic("implement me")
//}
