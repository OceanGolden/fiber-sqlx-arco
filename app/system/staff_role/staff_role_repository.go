package staff_role

import (
	"errors"
	"fiber-sqlx-arco/pkg/global"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"strings"
)

type Repository interface {
	GetCount(where *WhereParams) (uint64, error)
	FindAll(where *WhereParams) ([]*StaffRole, error)
	Delete(req *DeleteRequest) error
	GetCountWithTx(where *WhereParams, db *sqlx.Tx) (uint64, error)
	CreateBatchWithTx(req []*StaffRole, db *sqlx.Tx) error
	DeleteWithTx(req *DeleteRequest, db *sqlx.Tx) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetCount(where *WhereParams) (uint64, error) {
	selectBuilder := sq.Select("COUNT('*')").From(Table)
	roleID := strings.TrimSpace(where.RoleID)
	if len(roleID) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"role_id": roleID})
	}
	sql, args, _ := selectBuilder.ToSql()
	var count uint64
	err := r.db.Get(&count, sql, args...)
	return count, err
}

func (r *repository) FindAll(where *WhereParams) ([]*StaffRole, error) {
	selectBuilder := sq.Select(
		"staff_id",
		"role_id",
	).From(Table)
	staffID := where.StaffID
	staffIDs := where.StaffIDs
	roleID := where.RoleID
	roleIDs := where.RoleIDs
	if len(staffID) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"staff_id": staffID})
	}
	if len(staffIDs) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"staff_id": staffIDs})
	}
	if len(roleID) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"role_id": roleID})
	}
	if len(roleIDs) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"role_ids": roleIDs})
	}
	sql, args, _ := selectBuilder.ToSql()
	entities := []*StaffRole{}
	err := global.DB.Select(&entities, sql, args...)
	return entities, err
}

func (r *repository) Delete(req *DeleteRequest) error {
	sql, args, _ := sq.Delete(Table).Where(sq.Eq{"staff_id": req.StaffID}).ToSql()
	result, err := r.db.Exec(sql, args...)
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(DeletedFail)
	}
	return err
}

func (r *repository) CreateBatchWithTx(entities []*StaffRole, db *sqlx.Tx) error {
	insertBuilder := sq.Insert(Table).Columns("staff_id", "role_id", "created_at", "created_by")
	for _, entity := range entities {
		insertBuilder = insertBuilder.Values(entity.StaffID, entity.RoleID, entity.CreatedAt, entity.CreatedBy)
	}
	sql, args, _ := insertBuilder.ToSql()
	result, err := db.Exec(sql, args...)
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(UpdatedFail)
	}
	return err
}

func (r *repository) GetCountWithTx(where *WhereParams, db *sqlx.Tx) (uint64, error) {
	selectBuilder := sq.Select("Count(*)").From(Table)
	staffID := where.StaffID
	if len(staffID) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"staff_id": staffID})
	}
	sql, args, _ := selectBuilder.ToSql()
	var count uint64
	err := db.Get(&count, sql, args...)
	return count, err
}

func (r *repository) DeleteWithTx(req *DeleteRequest, db *sqlx.Tx) error {
	sql, args, _ := sq.Delete(Table).Where(sq.Eq{"staff_id": req.StaffID}).ToSql()
	result, err := db.Exec(sql, args...)
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(DeletedFail)
	}
	return err
}
