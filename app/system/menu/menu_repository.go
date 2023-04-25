package menu

import (
	"errors"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetTotal(params *WhereParams) (uint64, error)
	FindAll(params *WhereParams) ([]*Menu, error)
	FindById(id string) (*Menu, error)
	Create(req *Menu) error
	Update(req *Menu) error
	Delete(req *DeleteRequest) error
	CreateWithTx(req *Menu, tx *sqlx.Tx) error
	UpdateWithTx(req *Menu, tx *sqlx.Tx) error
	DeleteWithTx(req *DeleteRequest, tx *sqlx.Tx) error
	CheckFields(checkEntity *Menu) (*Menu, error)
	FindAllByIDs(ids []string) ([]*Menu, error)
	//FindAllByStaffID(staffID string) ([]*Menu, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetTotal(where *WhereParams) (uint64, error) {
	selectBuilder := sq.Select("count('id')").From(Table)
	name := strings.TrimSpace(where.Name)
	remark := strings.TrimSpace(where.Remark)
	status := strings.TrimSpace(where.Status)
	parentId := where.ParentID
	if len(name) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"name": fmt.Sprint("%", name, "%")})
	}
	if len(remark) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"remark": fmt.Sprint("%", remark, "%")})
	}
	if len(status) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"status": status})
	}
	if len(parentId) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"parent_id": parentId})
	}
	sql, args, _ := selectBuilder.ToSql()
	var count uint64
	err := r.db.Get(&count, sql, args...)
	return count, err
}

func (r *repository) FindAll(where *WhereParams) ([]*Menu, error) {
	selectBuilder := sq.Select(
		"id",
		"name",
		"parent_id",
		"icon",
		"path",
		"permission",
		"component",
		"type",
		"method",
		"status",
		"sort",
		"remark",
		"updated_at",
		"updated_by",
	).From(Table)
	name := strings.TrimSpace(where.Name)
	remark := strings.TrimSpace(where.Remark)
	status := strings.TrimSpace(where.Status)
	if len(name) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"name": fmt.Sprint("%", name, "%")})
	}
	if len(remark) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"remark": fmt.Sprint("%", remark, "%")})
	}
	if len(status) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"status": status})
	}
	selectBuilder = selectBuilder.OrderBy("sort")
	sql, args, _ := selectBuilder.ToSql()

	entities := []*Menu{}
	err := r.db.Select(&entities, sql, args...)
	return entities, err
}

func (r *repository) FindById(id string) (*Menu, error) {
	selectBuilder := sq.Select("*").From(Table)
	selectBuilder = selectBuilder.Where(sq.Eq{"id": id}).Limit(1)
	sql, args, _ := selectBuilder.ToSql()
	entity := &Menu{}
	err := r.db.Get(entity, sql, args...)
	return entity, err
}

func (r *repository) Create(entity *Menu) error {
	insertBuilder := sq.Insert(Table).SetMap(sq.Eq{
		"id":         entity.ID,
		"name":       entity.Name,
		"parent_id":  entity.ParentID,
		"icon":       entity.Icon,
		"path":       entity.Path,
		"permission": entity.Permission,
		"component":  entity.Component,
		"type":       entity.Type,
		"method":     entity.Method,
		"status":     entity.Status,
		"sort":       entity.Sort,
		"remark":     entity.Remark,
		"created_at": entity.CreatedAt,
		"updated_at": entity.UpdatedAt,
		"created_by": entity.CreatedBy,
		"updated_by": entity.UpdatedBy,
	})
	sql, args, _ := insertBuilder.ToSql()
	result, err := r.db.Exec(sql, args...)
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(CreatedFail)
	}
	return err
}

func (r *repository) Update(entity *Menu) error {
	found, err := r.FindById(entity.ID)
	if err != nil {
		return errors.New(ErrorNotExist)
	}
	updatedBuilder := sq.Update(Table)
	if found.Name != entity.Name && len(entity.Name) > 0 {
		updatedBuilder = updatedBuilder.Set("name", entity.Name)
	}
	if found.ParentID != entity.ParentID && len(entity.ParentID) > 0 {
		updatedBuilder = updatedBuilder.Set("parent_id", entity.ParentID)
	}
	if found.Icon != entity.Icon && len(entity.Icon) > 0 {
		updatedBuilder = updatedBuilder.Set("icon", entity.Icon)
	}
	if found.Path != entity.Path && len(entity.Path) > 0 {
		updatedBuilder = updatedBuilder.Set("path", entity.Path)
	}
	if found.Permission != entity.Permission && len(entity.Permission) > 0 {
		updatedBuilder = updatedBuilder.Set("permission", entity.Permission)
	}
	if found.Component != entity.Component && len(entity.Component) > 0 {
		updatedBuilder = updatedBuilder.Set("component", entity.Component)
	}
	if found.Type != entity.Type && len(entity.Type) > 0 {
		updatedBuilder = updatedBuilder.Set("type", entity.Type)
	}
	if found.Method != entity.Method && len(entity.Method) > 0 {
		updatedBuilder = updatedBuilder.Set("method", entity.Method)
	}
	if found.Status != entity.Status && len(entity.Status) > 0 {
		updatedBuilder = updatedBuilder.Set("status", entity.Status)
	}
	if found.Sort != entity.Sort && entity.Sort > 0 {
		updatedBuilder = updatedBuilder.Set("sort", entity.Sort)
	}
	if found.Remark != entity.Remark && len(entity.Remark) > 0 {
		updatedBuilder = updatedBuilder.Set("remark", entity.Remark)
	}
	updatedBuilder = updatedBuilder.SetMap(sq.Eq{
		"updated_at": entity.UpdatedAt,
		"updated_by": entity.UpdatedBy,
	}).Where(sq.Eq{"id": entity.ID}).Limit(1)
	sql, args, _ := updatedBuilder.ToSql()
	result, err := r.db.Exec(sql, args...)
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(UpdatedFail)
	}
	return err
}

func (r *repository) Delete(deleteReq *DeleteRequest) error {
	sql, args, _ := sq.Delete(Table).Where(sq.Eq{"id": deleteReq.ID}).ToSql()
	result, err := r.db.Exec(sql, args...)
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(DeletedFail)
	}
	return err
}

func (r *repository) CreateWithTx(entity *Menu, db *sqlx.Tx) error {
	insertBuilder := sq.Insert(Table).SetMap(sq.Eq{
		"id":         entity.ID,
		"name":       entity.Name,
		"parent_id":  entity.ParentID,
		"icon":       entity.Icon,
		"path":       entity.Path,
		"permission": entity.Permission,
		"component":  entity.Component,
		"type":       entity.Type,
		"method":     entity.Method,
		"status":     entity.Status,
		"sort":       entity.Sort,
		"remark":     entity.Remark,
		"created_at": entity.CreatedAt,
		"updated_at": entity.UpdatedAt,
		"created_by": entity.CreatedBy,
		"updated_by": entity.UpdatedBy,
	})
	sql, args, _ := insertBuilder.ToSql()
	result, err := db.Exec(sql, args...)
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(CreatedFail)
	}
	return err
}

func (r *repository) UpdateWithTx(entity *Menu, db *sqlx.Tx) error {
	found, err := r.FindById(entity.ID)
	if err != nil {
		return errors.New(ErrorNotExist)
	}
	updatedBuilder := sq.Update(Table)
	if found.Name != entity.Name && len(entity.Name) > 0 {
		updatedBuilder = updatedBuilder.Set("name", entity.Name)
	}
	if found.ParentID != entity.ParentID && len(entity.ParentID) > 0 {
		updatedBuilder = updatedBuilder.Set("parent_id", entity.ParentID)
	}
	if found.Icon != entity.Icon && len(entity.Icon) > 0 {
		updatedBuilder = updatedBuilder.Set("icon", entity.Icon)
	}
	if found.Path != entity.Path && len(entity.Path) > 0 {
		updatedBuilder = updatedBuilder.Set("path", entity.Path)
	}
	if found.Permission != entity.Permission && len(entity.Permission) > 0 {
		updatedBuilder = updatedBuilder.Set("permission", entity.Permission)
	}
	if found.Component != entity.Component && len(entity.Component) > 0 {
		updatedBuilder = updatedBuilder.Set("component", entity.Component)
	}
	if found.Type != entity.Type && len(entity.Type) > 0 {
		updatedBuilder = updatedBuilder.Set("type", entity.Type)
	}
	if found.Method != entity.Method && len(entity.Method) > 0 {
		updatedBuilder = updatedBuilder.Set("method", entity.Method)
	}
	if found.Status != entity.Status && len(entity.Status) > 0 {
		updatedBuilder = updatedBuilder.Set("status", entity.Status)
	}
	if found.Sort != entity.Sort && entity.Sort > 0 {
		updatedBuilder = updatedBuilder.Set("sort", entity.Sort)
	}
	if found.Remark != entity.Remark && len(entity.Remark) > 0 {
		updatedBuilder = updatedBuilder.Set("remark", entity.Remark)
	}
	updatedBuilder = updatedBuilder.SetMap(sq.Eq{
		"updated_at": entity.UpdatedAt,
		"updated_by": entity.UpdatedBy,
	}).Where(sq.Eq{"id": entity.ID}).Limit(1)
	sql, args, _ := updatedBuilder.ToSql()
	result, err := db.Exec(sql, args...)
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(UpdatedFail)
	}
	return err
}

func (r *repository) DeleteWithTx(deleteReq *DeleteRequest, db *sqlx.Tx) error {
	sql, args, _ := sq.Delete(Table).Where(sq.Eq{"id": deleteReq.ID}).ToSql()
	result, err := db.Exec(sql, args...)
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(DeletedFail)
	}
	return err
}

func (r *repository) CheckFields(checkEntity *Menu) (*Menu, error) {
	name := checkEntity.Name
	id := checkEntity.ID
	// name 唯一
	selectBuilder := sq.Select(
		"id",
		"name",
	).From(Table)
	var or []sq.Sqlizer
	if len(name) > 0 {
		or = append(or, sq.Eq{"name": name})
	}
	selectBuilder = selectBuilder.Where(sq.Or(or))
	if len(id) > 0 {
		selectBuilder = selectBuilder.Where(sq.NotEq{"id": id})
	}
	selectBuilder = selectBuilder.Limit(1)
	sql, args, _ := selectBuilder.ToSql()
	entity := &Menu{}
	err := r.db.Get(entity, sql, args...)
	return entity, err
}

func (r *repository) FindAllByIDs(ids []string) ([]*Menu, error) {
	selectBuilder := sq.Select(
		"id",
		"name",
		"parent_id",
		"icon",
		"path",
		"permission",
		"type",
		"method",
		"component",
		"link",
		"visible",
		"redirect",
		"status",
		"sort",
		"remark",
		"updated_at",
		"updated_by",
	).From(Table)
	selectBuilder = selectBuilder.Where(sq.Eq{"id": ids})
	selectBuilder = selectBuilder.OrderBy("sort")
	sql, args, _ := selectBuilder.ToSql()
	entities := []*Menu{}
	err := r.db.Select(&entities, sql, args...)

	return entities, err
}
