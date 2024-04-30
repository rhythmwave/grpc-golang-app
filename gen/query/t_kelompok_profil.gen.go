// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"go-bponline/m/gen/model"
)

func newTKelompokProfil(db *gorm.DB, opts ...gen.DOOption) tKelompokProfil {
	_tKelompokProfil := tKelompokProfil{}

	_tKelompokProfil.tKelompokProfilDo.UseDB(db, opts...)
	_tKelompokProfil.tKelompokProfilDo.UseModel(&model.TKelompokProfil{})

	tableName := _tKelompokProfil.tKelompokProfilDo.TableName()
	_tKelompokProfil.ALL = field.NewAsterisk(tableName)
	_tKelompokProfil.CIDKelompok = field.NewInt32(tableName, "c_id_kelompok")
	_tKelompokProfil.CNamaKelompok = field.NewString(tableName, "c_nama_kelompok")
	_tKelompokProfil.CStatus = field.NewString(tableName, "c_status")
	_tKelompokProfil.CUpdater = field.NewString(tableName, "c_updater")
	_tKelompokProfil.CLastUpdate = field.NewTime(tableName, "c_last_update")

	_tKelompokProfil.fillFieldMap()

	return _tKelompokProfil
}

type tKelompokProfil struct {
	tKelompokProfilDo

	ALL           field.Asterisk
	CIDKelompok   field.Int32
	CNamaKelompok field.String
	CStatus       field.String
	CUpdater      field.String
	CLastUpdate   field.Time

	fieldMap map[string]field.Expr
}

func (t tKelompokProfil) Table(newTableName string) *tKelompokProfil {
	t.tKelompokProfilDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tKelompokProfil) As(alias string) *tKelompokProfil {
	t.tKelompokProfilDo.DO = *(t.tKelompokProfilDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tKelompokProfil) updateTableName(table string) *tKelompokProfil {
	t.ALL = field.NewAsterisk(table)
	t.CIDKelompok = field.NewInt32(table, "c_id_kelompok")
	t.CNamaKelompok = field.NewString(table, "c_nama_kelompok")
	t.CStatus = field.NewString(table, "c_status")
	t.CUpdater = field.NewString(table, "c_updater")
	t.CLastUpdate = field.NewTime(table, "c_last_update")

	t.fillFieldMap()

	return t
}

func (t *tKelompokProfil) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tKelompokProfil) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 5)
	t.fieldMap["c_id_kelompok"] = t.CIDKelompok
	t.fieldMap["c_nama_kelompok"] = t.CNamaKelompok
	t.fieldMap["c_status"] = t.CStatus
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_last_update"] = t.CLastUpdate
}

func (t tKelompokProfil) clone(db *gorm.DB) tKelompokProfil {
	t.tKelompokProfilDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tKelompokProfil) replaceDB(db *gorm.DB) tKelompokProfil {
	t.tKelompokProfilDo.ReplaceDB(db)
	return t
}

type tKelompokProfilDo struct{ gen.DO }

type ITKelompokProfilDo interface {
	gen.SubQuery
	Debug() ITKelompokProfilDo
	WithContext(ctx context.Context) ITKelompokProfilDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITKelompokProfilDo
	WriteDB() ITKelompokProfilDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITKelompokProfilDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITKelompokProfilDo
	Not(conds ...gen.Condition) ITKelompokProfilDo
	Or(conds ...gen.Condition) ITKelompokProfilDo
	Select(conds ...field.Expr) ITKelompokProfilDo
	Where(conds ...gen.Condition) ITKelompokProfilDo
	Order(conds ...field.Expr) ITKelompokProfilDo
	Distinct(cols ...field.Expr) ITKelompokProfilDo
	Omit(cols ...field.Expr) ITKelompokProfilDo
	Join(table schema.Tabler, on ...field.Expr) ITKelompokProfilDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITKelompokProfilDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITKelompokProfilDo
	Group(cols ...field.Expr) ITKelompokProfilDo
	Having(conds ...gen.Condition) ITKelompokProfilDo
	Limit(limit int) ITKelompokProfilDo
	Offset(offset int) ITKelompokProfilDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITKelompokProfilDo
	Unscoped() ITKelompokProfilDo
	Create(values ...*model.TKelompokProfil) error
	CreateInBatches(values []*model.TKelompokProfil, batchSize int) error
	Save(values ...*model.TKelompokProfil) error
	First() (*model.TKelompokProfil, error)
	Take() (*model.TKelompokProfil, error)
	Last() (*model.TKelompokProfil, error)
	Find() ([]*model.TKelompokProfil, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TKelompokProfil, err error)
	FindInBatches(result *[]*model.TKelompokProfil, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TKelompokProfil) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITKelompokProfilDo
	Assign(attrs ...field.AssignExpr) ITKelompokProfilDo
	Joins(fields ...field.RelationField) ITKelompokProfilDo
	Preload(fields ...field.RelationField) ITKelompokProfilDo
	FirstOrInit() (*model.TKelompokProfil, error)
	FirstOrCreate() (*model.TKelompokProfil, error)
	FindByPage(offset int, limit int) (result []*model.TKelompokProfil, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITKelompokProfilDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tKelompokProfilDo) Debug() ITKelompokProfilDo {
	return t.withDO(t.DO.Debug())
}

func (t tKelompokProfilDo) WithContext(ctx context.Context) ITKelompokProfilDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tKelompokProfilDo) ReadDB() ITKelompokProfilDo {
	return t.Clauses(dbresolver.Read)
}

func (t tKelompokProfilDo) WriteDB() ITKelompokProfilDo {
	return t.Clauses(dbresolver.Write)
}

func (t tKelompokProfilDo) Session(config *gorm.Session) ITKelompokProfilDo {
	return t.withDO(t.DO.Session(config))
}

func (t tKelompokProfilDo) Clauses(conds ...clause.Expression) ITKelompokProfilDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tKelompokProfilDo) Returning(value interface{}, columns ...string) ITKelompokProfilDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tKelompokProfilDo) Not(conds ...gen.Condition) ITKelompokProfilDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tKelompokProfilDo) Or(conds ...gen.Condition) ITKelompokProfilDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tKelompokProfilDo) Select(conds ...field.Expr) ITKelompokProfilDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tKelompokProfilDo) Where(conds ...gen.Condition) ITKelompokProfilDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tKelompokProfilDo) Order(conds ...field.Expr) ITKelompokProfilDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tKelompokProfilDo) Distinct(cols ...field.Expr) ITKelompokProfilDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tKelompokProfilDo) Omit(cols ...field.Expr) ITKelompokProfilDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tKelompokProfilDo) Join(table schema.Tabler, on ...field.Expr) ITKelompokProfilDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tKelompokProfilDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITKelompokProfilDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tKelompokProfilDo) RightJoin(table schema.Tabler, on ...field.Expr) ITKelompokProfilDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tKelompokProfilDo) Group(cols ...field.Expr) ITKelompokProfilDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tKelompokProfilDo) Having(conds ...gen.Condition) ITKelompokProfilDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tKelompokProfilDo) Limit(limit int) ITKelompokProfilDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tKelompokProfilDo) Offset(offset int) ITKelompokProfilDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tKelompokProfilDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITKelompokProfilDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tKelompokProfilDo) Unscoped() ITKelompokProfilDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tKelompokProfilDo) Create(values ...*model.TKelompokProfil) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tKelompokProfilDo) CreateInBatches(values []*model.TKelompokProfil, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tKelompokProfilDo) Save(values ...*model.TKelompokProfil) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tKelompokProfilDo) First() (*model.TKelompokProfil, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKelompokProfil), nil
	}
}

func (t tKelompokProfilDo) Take() (*model.TKelompokProfil, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKelompokProfil), nil
	}
}

func (t tKelompokProfilDo) Last() (*model.TKelompokProfil, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKelompokProfil), nil
	}
}

func (t tKelompokProfilDo) Find() ([]*model.TKelompokProfil, error) {
	result, err := t.DO.Find()
	return result.([]*model.TKelompokProfil), err
}

func (t tKelompokProfilDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TKelompokProfil, err error) {
	buf := make([]*model.TKelompokProfil, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tKelompokProfilDo) FindInBatches(result *[]*model.TKelompokProfil, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tKelompokProfilDo) Attrs(attrs ...field.AssignExpr) ITKelompokProfilDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tKelompokProfilDo) Assign(attrs ...field.AssignExpr) ITKelompokProfilDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tKelompokProfilDo) Joins(fields ...field.RelationField) ITKelompokProfilDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tKelompokProfilDo) Preload(fields ...field.RelationField) ITKelompokProfilDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tKelompokProfilDo) FirstOrInit() (*model.TKelompokProfil, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKelompokProfil), nil
	}
}

func (t tKelompokProfilDo) FirstOrCreate() (*model.TKelompokProfil, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKelompokProfil), nil
	}
}

func (t tKelompokProfilDo) FindByPage(offset int, limit int) (result []*model.TKelompokProfil, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tKelompokProfilDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tKelompokProfilDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tKelompokProfilDo) Delete(models ...*model.TKelompokProfil) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tKelompokProfilDo) withDO(do gen.Dao) *tKelompokProfilDo {
	t.DO = *do.(*gen.DO)
	return t
}
