package dictionary_item

import (
	"errors"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetCount(params *WhereParams) (uint64, error)
	FindPage(params *WhereParams) ([]*DictionaryItem, error)
	FindById(id string) (*DictionaryItem, error)
	Create(req *DictionaryItem) error
	Update(req *DictionaryItem) error
	Delete(req *DeleteRequest) error
	CreateWithTx(req *DictionaryItem, tx *sqlx.Tx) error
	UpdateWithTx(req *DictionaryItem, tx *sqlx.Tx) error
	DeleteWithTx(req *DeleteRequest, tx *sqlx.Tx) error
	CheckFields(checkEntity *DictionaryItem) (*DictionaryItem, error)
	FindAllByDictionaryID(dictionaryId string) ([]*DictionaryItem, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetCount(where *WhereParams) (uint64, error) {
	selectBuilder := sq.Select("count('id')").From(Table)
	label := strings.TrimSpace(where.Label)
	value := strings.TrimSpace(where.Value)
	remark := strings.TrimSpace(where.Remark)
	status := strings.TrimSpace(where.Status)
	if len(label) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"label": fmt.Sprint("%", label, "%")})
	}
	if len(value) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"value": fmt.Sprint("%", value, "%")})
	}
	if len(remark) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"remark": fmt.Sprint("%", remark, "%")})
	}
	if len(status) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"status": status})
	}
	selectBuilder = selectBuilder.Where(sq.Eq{"dictionary_id": where.DictionaryID})
	sql, args, _ := selectBuilder.ToSql()
	var count uint64
	err := r.db.Get(&count, sql, args...)
	return count, err
}

func (r *repository) FindPage(where *WhereParams) ([]*DictionaryItem, error) {
	selectBuilder := sq.Select(
		"id",
		"dictionary_id",
		"label",
		"value",
		"color",
		"status",
		"sort",
		"remark",
		"updated_at",
		"updated_by",
	).From(Table)
	label := strings.TrimSpace(where.Label)
	value := strings.TrimSpace(where.Value)
	remark := strings.TrimSpace(where.Remark)
	status := strings.TrimSpace(where.Status)
	if len(where.DictionaryID) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"dictionary_id": where.DictionaryID})
	}
	if len(label) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"label": fmt.Sprint("%", label, "%")})
	}
	if len(value) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"value": fmt.Sprint("%", value, "%")})
	}
	if len(remark) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"remark": fmt.Sprint("%", remark, "%")})
	}
	if len(status) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"status": status})
	}
	selectBuilder = selectBuilder.OrderBy("sort").Limit(where.PageSize).Offset((where.Current - 1) * where.PageSize)
	sql, args, _ := selectBuilder.ToSql()
	entities := []*DictionaryItem{}
	err := r.db.Select(&entities, sql, args...)
	return entities, err
}

func (r *repository) FindById(id string) (*DictionaryItem, error) {
	selectBuilder := sq.Select(
		"id",
		"dictionary_id",
		"label",
		"value",
		"color",
		"status",
		"sort",
		"remark",
		"updated_at",
		"updated_by",
	).From(Table)
	selectBuilder = selectBuilder.Where(sq.Eq{"id": id}).Limit(1)
	sql, args, _ := selectBuilder.ToSql()
	entity := &DictionaryItem{}
	err := r.db.Get(entity, sql, args...)
	return entity, err
}

func (r *repository) Create(entity *DictionaryItem) error {
	insertBuilder := sq.Insert(Table).SetMap(sq.Eq{
		"id":            entity.ID,
		"label":         entity.Label,
		"value":         entity.Value,
		"color":         entity.Color,
		"dictionary_id": entity.DictionaryID,
		"status":        entity.Status,
		"sort":          entity.Sort,
		"remark":        entity.Remark,
		"created_at":    entity.CreatedAt,
		"updated_at":    entity.UpdatedAt,
		"created_by":    entity.CreatedBy,
		"updated_by":    entity.UpdatedBy,
	})
	sql, args, _ := insertBuilder.ToSql()
	result, err := r.db.Exec(sql, args...)
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(CreatedFail)
	}
	return err
}

func (r *repository) Update(entity *DictionaryItem) error {
	found, err := r.FindById(entity.ID)
	if err != nil {
		return errors.New(ErrorNotExist)
	}
	updatedBuilder := sq.Update(Table)
	if found.Label != entity.Label && len(entity.Label) > 0 {
		updatedBuilder = updatedBuilder.Set("label", entity.Label)
	}
	if found.Value != entity.Value && len(entity.Value) > 0 {
		updatedBuilder = updatedBuilder.Set("value", entity.Value)
	}
	if found.Color != entity.Color && len(entity.Color) > 0 {
		updatedBuilder = updatedBuilder.Set("color", entity.Color)
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

func (r *repository) CreateWithTx(entity *DictionaryItem, tx *sqlx.Tx) error {
	insertBuilder := sq.Insert(Table).SetMap(sq.Eq{
		"id":            entity.ID,
		"dictionary_id": entity.DictionaryID,
		"label":         entity.Label,
		"value":         entity.Value,
		"color":         entity.Color,
		"status":        entity.Status,
		"sort":          entity.Sort,
		"remark":        entity.Remark,
		"created_at":    entity.CreatedAt,
		"updated_at":    entity.UpdatedAt,
		"created_by":    entity.CreatedBy,
		"updated_by":    entity.UpdatedBy,
	})
	sql, args, _ := insertBuilder.ToSql()
	result, err := tx.Exec(sql, args...)
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(CreatedFail)
	}
	return err
}

func (r *repository) UpdateWithTx(entity *DictionaryItem, tx *sqlx.Tx) error {
	found, err := r.FindById(entity.ID)
	if err != nil {
		return errors.New(ErrorNotExist)
	}
	updatedBuilder := sq.Update(Table)
	if found.Label != entity.Label && len(entity.Label) > 0 {
		updatedBuilder = updatedBuilder.Set("label", entity.Label)
	}
	if found.Value != entity.Value && len(entity.Value) > 0 {
		updatedBuilder = updatedBuilder.Set("value", entity.Value)
	}
	if found.Color != entity.Color && len(entity.Color) > 0 {
		updatedBuilder = updatedBuilder.Set("color", entity.Color)
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

func (r *repository) CheckFields(checkEntity *DictionaryItem) (*DictionaryItem, error) {
	label := checkEntity.Label
	value := checkEntity.Value
	dictionaryId := checkEntity.DictionaryID
	id := checkEntity.ID
	// 当 dictionary_id 相等的时候 name 和 code 相同 则返回结果
	// dictionary_id name 联合唯一
	// dictionary_id code 联合唯一
	selectBuilder := sq.Select(
		"id",
		"label",
		"value",
		"dictionary_id",
	).From(Table)
	var or []sq.Sqlizer
	if len(label) > 0 {
		or = append(or, sq.Eq{"label": label})
	}
	if len(value) > 0 {
		or = append(or, sq.Eq{"value": value})
	}
	selectBuilder = selectBuilder.Where(sq.Or(or))
	selectBuilder = selectBuilder.Where(sq.Eq{"dictionary_id": dictionaryId})
	if len(id) > 0 {
		selectBuilder = selectBuilder.Where(sq.NotEq{"id": id})
	}
	selectBuilder = selectBuilder.Limit(1)
	sql, args, _ := selectBuilder.ToSql()

	entity := &DictionaryItem{}
	err := r.db.Get(entity, sql, args...)
	return entity, err
}

func (r *repository) FindAllByDictionaryID(dictionaryId string) ([]*DictionaryItem, error) {
	selectBuilder := sq.Select(
		"id",
		"label",
		"value",
		"color",
		"dictionary_id",
		"status",
		"sort",
		"remark",
		"updated_at",
		"updated_by",
	).From(Table)
	selectBuilder = selectBuilder.Where(sq.Eq{"dictionary_id": dictionaryId}).OrderBy("sort")
	sql, args, _ := selectBuilder.ToSql()
	entities := []*DictionaryItem{}
	err := r.db.Select(&entities, sql, args...)
	return entities, err
}
