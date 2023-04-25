package organization

import (
	"errors"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetCount(params *WhereParams) (uint64, error)
	FindAll(params *WhereParams) ([]*Organization, error)
	FindById(id string) (*Organization, error)
	Create(req *Organization) error
	Update(req *Organization) error
	Delete(req *DeleteRequest) error
	FindByIdWithTx(id string, tx *sqlx.Tx) (*Organization, error)
	CreateWithTx(req *Organization, tx *sqlx.Tx) error
	UpdateWithTx(req *Organization, tx *sqlx.Tx) error
	DeleteWithTx(req *DeleteRequest, tx *sqlx.Tx) error
	CheckFields(checkEntity *Organization) (*Organization, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetCount(where *WhereParams) (uint64, error) {
	selectBuilder := sq.Select("COUNT('id')").From(Table)
	name := strings.TrimSpace(where.Name)
	code := strings.TrimSpace(where.Code)
	remark := strings.TrimSpace(where.Remark)
	status := strings.TrimSpace(where.Status)
	parentId := where.ParentID
	if len(name) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"name": fmt.Sprint("%", name, "%")})
	}
	if len(code) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"code": fmt.Sprint("%", code, "%")})
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

func (r *repository) FindAll(where *WhereParams) ([]*Organization, error) {
	selectBuilder := sq.Select(
		"id",
		"name",
		"code",
		"parent_id",
		"parent_ids",
		"status",
		"sort",
		"remark",
		"updated_at",
		"updated_by",
	).From(Table)
	name := strings.TrimSpace(where.Name)
	code := strings.TrimSpace(where.Code)
	remark := strings.TrimSpace(where.Remark)
	status := strings.TrimSpace(where.Status)
	if len(name) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"name": fmt.Sprint("%", name, "%")})
	}
	if len(code) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"code": fmt.Sprint("%", code, "%")})
	}
	if len(remark) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"remark": fmt.Sprint("%", remark, "%")})
	}
	if len(status) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"status": status})
	}
	selectBuilder = selectBuilder.OrderBy("sort")
	sql, args, _ := selectBuilder.ToSql()

	entities := []*Organization{}
	err := r.db.Select(&entities, sql, args...)
	return entities, err
}

func (r *repository) FindById(id string) (*Organization, error) {
	selectBuilder := sq.Select("*").From(Table)
	selectBuilder = selectBuilder.Where(sq.Eq{"id": id}).Limit(1)
	sql, args, _ := selectBuilder.ToSql()
	entity := &Organization{}
	err := r.db.Get(entity, sql, args...)
	return entity, err
}

func (r *repository) Create(entity *Organization) error {
	insertBuilder := sq.Insert(Table).SetMap(sq.Eq{
		"id":         entity.ID,
		"name":       entity.Name,
		"code":       entity.Code,
		"parent_id":  entity.ParentID,
		"parent_ids": entity.ParentIDS,
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

func (r *repository) Update(entity *Organization) error {
	found, err := r.FindById(entity.ID)
	if err != nil {
		return errors.New(ErrorNotExist)
	}
	updatedBuilder := sq.Update(Table)
	if found.Name != entity.Name && len(entity.Name) > 0 {
		updatedBuilder = updatedBuilder.Set("name", entity.Name)
	}
	if found.Code != entity.Code && len(entity.Code) > 0 {
		updatedBuilder = updatedBuilder.Set("code", entity.Code)
	}
	if found.ParentID != entity.ParentID && len(entity.ParentID) > 0 {
		updatedBuilder = updatedBuilder.Set("parent_id", entity.ParentID)
	}
	if found.ParentIDS != entity.ParentIDS && len(entity.ParentIDS) > 0 {
		updatedBuilder = updatedBuilder.Set("parent_ids", entity.ParentIDS)
	}
	if found.Status != entity.Status && len(entity.Status) > 0 {
		updatedBuilder = updatedBuilder.Set("status", entity.Status)
	}
	if found.Sort != entity.Sort && entity.Sort > 0 {
		updatedBuilder = updatedBuilder.Set("sort", entity.Sort)
	}
	if found.Remark != entity.Remark {
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

func (r *repository) FindByIdWithTx(id string, tx *sqlx.Tx) (*Organization, error) {
	selectBuilder := sq.Select("*").From(Table)
	selectBuilder = selectBuilder.Where(sq.Eq{"id": id}).Limit(1)
	sql, args, _ := selectBuilder.ToSql()
	entity := &Organization{}
	err := tx.Get(entity, sql, args...)
	return entity, err
}

func (r *repository) CreateWithTx(entity *Organization, tx *sqlx.Tx) error {
	insertBuilder := sq.Insert(Table).SetMap(sq.Eq{
		"id":         entity.ID,
		"name":       entity.Name,
		"code":       entity.Code,
		"parent_id":  entity.ParentID,
		"parent_ids": entity.ParentIDS,
		"status":     entity.Status,
		"sort":       entity.Sort,
		"remark":     entity.Remark,
		"created_at": entity.CreatedAt,
		"updated_at": entity.UpdatedAt,
		"created_by": entity.CreatedBy,
		"updated_by": entity.UpdatedBy,
	})
	sql, args, _ := insertBuilder.ToSql()
	result, err := tx.Exec(sql, args...)
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(CreatedFail)
	}
	return err
}

func (r *repository) UpdateWithTx(entity *Organization, tx *sqlx.Tx) error {
	found, err := r.FindByIdWithTx(entity.ID, tx)
	if err != nil {
		return errors.New(ErrorNotExist)
	}
	updatedBuilder := sq.Update(Table)
	if found.Name != entity.Name && len(entity.Name) > 0 {
		updatedBuilder = updatedBuilder.Set("name", entity.Name)
	}
	if found.Code != entity.Code && len(entity.Code) > 0 {
		updatedBuilder = updatedBuilder.Set("code", entity.Code)
	}
	if found.ParentID != entity.ParentID && len(entity.ParentID) > 0 {
		updatedBuilder = updatedBuilder.Set("parent_id", entity.ParentID)
	}
	if found.ParentIDS != entity.ParentIDS && len(entity.ParentIDS) > 0 {
		updatedBuilder = updatedBuilder.Set("parent_ids", entity.ParentIDS)
	}
	if found.Status != entity.Status && len(entity.Status) > 0 {
		updatedBuilder = updatedBuilder.Set("status", entity.Status)
	}
	if found.Sort != entity.Sort && entity.Sort > 0 {
		updatedBuilder = updatedBuilder.Set("sort", entity.Sort)
	}
	if found.Remark != entity.Remark {
		updatedBuilder = updatedBuilder.Set("remark", entity.Remark)
	}
	updatedBuilder = updatedBuilder.SetMap(sq.Eq{
		"updated_at": entity.UpdatedAt,
		"updated_by": entity.UpdatedBy,
	}).Where(sq.Eq{"id": entity.ID}).Limit(1)
	sql, args, _ := updatedBuilder.ToSql()
	result, err := tx.Exec(sql, args...)
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(UpdatedFail)
	}
	return err
}

func (r *repository) DeleteWithTx(deleteReq *DeleteRequest, tx *sqlx.Tx) error {
	sql, args, _ := sq.Delete(Table).Where(sq.Eq{"id": deleteReq.ID}).ToSql()
	result, err := tx.Exec(sql, args...)
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(DeletedFail)
	}
	return err
}

func (r *repository) CheckFields(checkEntity *Organization) (*Organization, error) {
	name := checkEntity.Name
	code := checkEntity.Code
	id := checkEntity.ID
	// name, code 唯一
	selectBuilder := sq.Select(
		"id",
		"name",
		"code",
		"parent_id",
	).From(Table)
	var or []sq.Sqlizer
	if len(name) > 0 {
		or = append(or, sq.Eq{"name": name})
	}
	if len(code) > 0 {
		or = append(or, sq.Eq{"code": code})
	}
	selectBuilder = selectBuilder.Where(sq.Or(or))
	if len(id) > 0 {
		selectBuilder = selectBuilder.Where(sq.NotEq{"id": id})
	}
	selectBuilder = selectBuilder.Limit(1)
	sql, args, _ := selectBuilder.ToSql()
	entity := &Organization{}
	err := r.db.Get(entity, sql, args...)
	return entity, err
}
