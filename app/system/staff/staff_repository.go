package staff

import (
	"errors"
	"fiber-sqlx-arco/pkg/common/constants"
	"fmt"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetCount(params *WhereParams) (uint64, error)
	FindPage(params *WhereParams) ([]*Staff, error)
	FindAll(params *WhereParams) ([]*Staff, error)
	FindById(id string) (*Staff, error)
	Create(req *Staff) error
	Update(req *Staff) error
	Delete(req *DeleteRequest) error
	CreateWithTx(req *Staff, tx *sqlx.Tx) error
	UpdateWithTx(req *Staff, tx *sqlx.Tx) error
	DeleteWithTx(req *DeleteRequest, tx *sqlx.Tx) error
	CheckFields(checkEntity *Staff) (*Staff, error)
	FindOneByUsername(username string) (*Staff, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetCount(where *WhereParams) (uint64, error) {
	selectBuilder := sq.Select("COUNT('id')").From(Table)
	username := strings.TrimSpace(where.Username)
	name := strings.TrimSpace(where.Name)
	email := strings.TrimSpace(where.Email)
	mobile := strings.TrimSpace(where.Mobile)
	gender := strings.TrimSpace(where.Gender)
	organizationId := where.OrganizationID
	positionId := where.PositionID
	workStatus := strings.TrimSpace(where.WorkStatus)
	status := where.Status
	remark := strings.TrimSpace(where.Remark)
	if len(username) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"username": fmt.Sprint("%", username, "%")})
	}
	if len(name) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"name": fmt.Sprint("%", name, "%")})
	}
	if len(email) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"email": fmt.Sprint("%", email, "%")})
	}
	if len(mobile) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"mobile": fmt.Sprint("%", mobile, "%")})
	}
	if len(gender) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"gender": gender})
	}
	if len(organizationId) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"organization_id": organizationId})
	}
	if len(positionId) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"position_id": positionId})
	}
	if len(workStatus) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"work_status": workStatus})
	}
	if len(status) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"status": status})
	}
	if len(remark) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"remark": fmt.Sprint("%", remark, "%")})
	}
	selectBuilder = selectBuilder.Where(sq.Eq{constants.DeletedFlag: constants.DeletedValue})
	sql, args, _ := selectBuilder.ToSql()
	var count uint64
	err := r.db.Get(&count, sql, args...)
	return count, err
}

func (r *repository) FindPage(where *WhereParams) ([]*Staff, error) {
	selectBuilder := sq.Select(
		"id",
		"username",
		"name",
		"email",
		"mobile",
		"avatar",
		"gender",
		"organization_id",
		"position_id",
		"work_status",
		"status",
		"sort",
		"remark",
		"updated_at",
		"updated_by",
	).From(Table)
	username := strings.TrimSpace(where.Username)
	name := strings.TrimSpace(where.Name)
	email := strings.TrimSpace(where.Email)
	mobile := strings.TrimSpace(where.Mobile)
	gender := strings.TrimSpace(where.Gender)
	organizationId := where.OrganizationID
	positionId := where.PositionID
	workStatus := strings.TrimSpace(where.WorkStatus)
	status := strings.TrimSpace(where.Status)
	remark := strings.TrimSpace(where.Remark)
	if len(username) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"username": fmt.Sprint("%", username, "%")})
	}
	if len(name) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"name": fmt.Sprint("%", name, "%")})
	}
	if len(email) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"email": fmt.Sprint("%", email, "%")})
	}
	if len(mobile) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"mobile": fmt.Sprint("%", mobile, "%")})
	}
	if len(gender) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"gender": gender})
	}
	if len(organizationId) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"organization_id": organizationId})
	}
	if len(positionId) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"position_id": positionId})
	}
	if len(workStatus) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"work_status": workStatus})
	}
	if len(status) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"status": status})
	}
	if len(remark) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"remark": fmt.Sprint("%", remark, "%")})
	}
	selectBuilder = selectBuilder.Where(sq.Eq{constants.DeletedFlag: constants.DeletedValue})
	selectBuilder = selectBuilder.OrderBy("sort").Limit(where.PageSize).Offset((where.Current - 1) * where.PageSize)
	sql, args, _ := selectBuilder.ToSql()
	entities := []*Staff{}
	err := r.db.Select(&entities, sql, args...)
	return entities, err
}

func (r *repository) FindAll(where *WhereParams) ([]*Staff, error) {
	selectBuilder := sq.Select(
		"id",
		"username",
		"name",
		"email",
		"mobile",
		"avatar",
		"gender",
		"organization_id",
		"position_id",
		"work_status",
		"status",
		"sort",
		"remark",
		"updated_at",
		"updated_by",
	).From(Table)
	username := strings.TrimSpace(where.Username)
	name := strings.TrimSpace(where.Name)
	email := strings.TrimSpace(where.Email)
	mobile := strings.TrimSpace(where.Mobile)
	gender := strings.TrimSpace(where.Gender)
	organizationId := where.OrganizationID
	positionId := where.PositionID
	workStatus := strings.TrimSpace(where.WorkStatus)
	status := strings.TrimSpace(where.Status)
	remark := strings.TrimSpace(where.Remark)
	if len(username) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"username": fmt.Sprint("%", username, "%")})
	}
	if len(name) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"name": fmt.Sprint("%", name, "%")})
	}
	if len(email) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"email": fmt.Sprint("%", email, "%")})
	}
	if len(mobile) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"mobile": fmt.Sprint("%", mobile, "%")})
	}
	if len(gender) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"gender": gender})
	}
	if len(organizationId) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"organization_id": organizationId})
	}
	if len(positionId) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"position_id": positionId})
	}
	if len(workStatus) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"work_status": workStatus})
	}
	if len(status) > 0 {
		selectBuilder = selectBuilder.Where(sq.Eq{"status": status})
	}
	if len(remark) > 0 {
		selectBuilder = selectBuilder.Where(sq.Like{"remark": fmt.Sprint("%", remark, "%")})
	}
	selectBuilder = selectBuilder.Where(sq.Eq{constants.DeletedFlag: constants.DeletedValue})
	selectBuilder = selectBuilder.OrderBy("sort")
	sql, args, _ := selectBuilder.ToSql()
	entities := []*Staff{}
	err := r.db.Select(&entities, sql, args...)
	return entities, err
}

func (r *repository) FindById(id string) (*Staff, error) {
	selectBuilder := sq.Select("*").From(Table)
	selectBuilder = selectBuilder.Where(sq.Eq{"id": id}).Where(sq.Eq{constants.DeletedFlag: constants.DeletedValue}).Limit(1)
	sql, args, _ := selectBuilder.ToSql()
	entity := &Staff{}
	err := r.db.Get(entity, sql, args...)
	return entity, err
}

func (r *repository) Create(entity *Staff) error {
	insertBuilder := sq.Insert(Table).SetMap(sq.Eq{
		"id":                  entity.ID,
		"username":            entity.Username,
		"password":            entity.Password,
		"name":                entity.Name,
		"email":               entity.Email,
		"mobile":              entity.Mobile,
		"avatar":              entity.Avatar,
		"gender":              entity.Gender,
		"organization_id":     entity.OrganizationID,
		"position_id":         entity.PositionID,
		"work_status":         entity.WorkStatus,
		"status":              entity.Status,
		"sort":                entity.Sort,
		"remark":              entity.Remark,
		"created_at":          entity.CreatedAt,
		"updated_at":          entity.UpdatedAt,
		"created_by":          entity.CreatedBy,
		"updated_by":          entity.UpdatedBy,
		constants.DeletedFlag: constants.DeletedValue,
	})
	sql, args, _ := insertBuilder.ToSql()
	result, err := r.db.Exec(sql, args...)
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(CreatedFail)
	}
	return err
}

func (r *repository) Update(entity *Staff) error {
	found, err := r.FindById(entity.ID)
	if err != nil {
		return errors.New(ErrorNotExist)
	}
	updatedBuilder := sq.Update(Table)
	if found.Username != entity.Username && len(entity.Username) > 0 {
		updatedBuilder = updatedBuilder.Set("username", entity.Username)
	}
	if found.Name != entity.Name && len(entity.Name) > 0 {
		updatedBuilder = updatedBuilder.Set("name", entity.Name)
	}
	if found.Email != entity.Email && len(entity.Email) > 0 {
		updatedBuilder = updatedBuilder.Set("email", entity.Email)
	}
	if found.Mobile != entity.Mobile && len(entity.Mobile) > 0 {
		updatedBuilder = updatedBuilder.Set("mobile", entity.Mobile)
	}
	if found.Avatar != entity.Avatar && len(entity.Avatar) > 0 {
		updatedBuilder = updatedBuilder.Set("avatar", entity.Avatar)
	}
	if found.Gender != entity.Gender && len(entity.Gender) > 0 {
		updatedBuilder = updatedBuilder.Set("gender", entity.Gender)
	}
	if found.OrganizationID != entity.OrganizationID && len(entity.OrganizationID) > 0 {
		updatedBuilder = updatedBuilder.Set("organization_id", entity.OrganizationID)
	}
	if found.PositionID != entity.PositionID && len(entity.PositionID) > 0 {
		updatedBuilder = updatedBuilder.Set("position_id", entity.PositionID)
	}
	if found.WorkStatus != entity.WorkStatus && len(entity.WorkStatus) > 0 {
		updatedBuilder = updatedBuilder.Set("work_status", entity.WorkStatus)
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
	sql, args, _ := sq.Update(Table).Set(constants.DeletedFlag, time.Now().Unix()).Where(sq.Eq{"id": deleteReq.ID}).Where(sq.Eq{constants.DeletedFlag: constants.DeletedValue}).ToSql()
	result, err := r.db.Exec(sql, args...)
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(DeletedFail)
	}
	return err
}

func (r *repository) CreateWithTx(entity *Staff, tx *sqlx.Tx) error {
	insertBuilder := sq.Insert(Table).SetMap(sq.Eq{
		"id":              entity.ID,
		"username":        entity.Username,
		"password":        entity.Password,
		"name":            entity.Name,
		"email":           entity.Email,
		"mobile":          entity.Mobile,
		"avatar":          entity.Avatar,
		"gender":          entity.Gender,
		"organization_id": entity.OrganizationID,
		"position_id":     entity.PositionID,
		"work_status":     entity.WorkStatus,
		"status":          entity.Status,
		"sort":            entity.Sort,
		"remark":          entity.Remark,
		"created_at":      entity.CreatedAt,
		"updated_at":      entity.UpdatedAt,
		"created_by":      entity.CreatedBy,
		"updated_by":      entity.UpdatedBy,
	})
	insertBuilder = insertBuilder.SetMap(sq.Eq{constants.DeletedFlag: constants.DeletedValue})
	sql, args, _ := insertBuilder.ToSql()
	result, err := tx.Exec(sql, args...)
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(CreatedFail)
	}
	return err
}

func (r *repository) UpdateWithTx(entity *Staff, tx *sqlx.Tx) error {
	found, err := r.FindById(entity.ID)
	if err != nil {
		return errors.New(ErrorNotExist)
	}
	updatedBuilder := sq.Update(Table)
	if found.Username != entity.Username && len(entity.Username) > 0 {
		updatedBuilder = updatedBuilder.Set("username", entity.Username)
	}
	if found.Name != entity.Name && len(entity.Name) > 0 {
		updatedBuilder = updatedBuilder.Set("name", entity.Name)
	}
	if found.Email != entity.Email && len(entity.Email) > 0 {
		updatedBuilder = updatedBuilder.Set("email", entity.Email)
	}
	if found.Mobile != entity.Mobile && len(entity.Mobile) > 0 {
		updatedBuilder = updatedBuilder.Set("mobile", entity.Mobile)
	}
	if found.Avatar != entity.Avatar && len(entity.Avatar) > 0 {
		updatedBuilder = updatedBuilder.Set("avatar", entity.Avatar)
	}
	if found.Gender != entity.Gender && len(entity.Gender) > 0 {
		updatedBuilder = updatedBuilder.Set("gender", entity.Gender)
	}
	if found.OrganizationID != entity.OrganizationID && len(entity.OrganizationID) > 0 {
		updatedBuilder = updatedBuilder.Set("organization_id", entity.OrganizationID)
	}
	if found.PositionID != entity.PositionID && len(entity.PositionID) > 0 {
		updatedBuilder = updatedBuilder.Set("position_id", entity.PositionID)
	}
	if found.WorkStatus != entity.WorkStatus && len(entity.WorkStatus) > 0 {
		updatedBuilder = updatedBuilder.Set("work_status", entity.WorkStatus)
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
	sql, args, _ := sq.Update(Table).Set(constants.DeletedFlag, time.Now().Unix()).Where(sq.Eq{"id": deleteReq.ID}).Where(sq.Eq{constants.DeletedFlag: constants.DeletedValue}).ToSql()
	result, err := tx.Exec(sql, args...)
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(DeletedFail)
	}
	return err
}

func (r *repository) CheckFields(checkEntity *Staff) (*Staff, error) {
	username := checkEntity.Username
	email := checkEntity.Email
	mobile := checkEntity.Mobile
	id := checkEntity.ID
	// username, email, mobile 唯一
	selectBuilder := sq.Select(
		"id",
		"username",
		"email",
		"mobile",
	).From(Table)
	var or []sq.Sqlizer
	if len(username) > 0 {
		or = append(or, sq.Eq{"username": username})
	}
	if len(email) > 0 {
		or = append(or, sq.Eq{"email": email})
	}
	if len(mobile) > 0 {
		or = append(or, sq.Eq{"mobile": mobile})
	}
	selectBuilder = selectBuilder.Where(sq.Or(or))
	if len(id) > 0 {
		selectBuilder = selectBuilder.Where(sq.NotEq{"id": id})
	}
	selectBuilder = selectBuilder.Limit(1)
	sql, args, _ := selectBuilder.ToSql()
	entity := &Staff{}
	err := r.db.Get(entity, sql, args...)
	return entity, err
}

func (r *repository) FindOneByUsername(username string) (*Staff, error) {
	selectBuilder := sq.Select("*").From(Table)
	selectBuilder = selectBuilder.Where(sq.Eq{"username": username}).Where(sq.Eq{constants.DeletedFlag: constants.DeletedValue}).Limit(1)
	sql, args, _ := selectBuilder.ToSql()
	entity := &Staff{}
	err := r.db.Get(entity, sql, args...)
	return entity, err
}
