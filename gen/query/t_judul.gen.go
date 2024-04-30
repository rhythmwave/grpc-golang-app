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

func newTJudul(db *gorm.DB, opts ...gen.DOOption) tJudul {
	_tJudul := tJudul{}

	_tJudul.tJudulDo.UseDB(db, opts...)
	_tJudul.tJudulDo.UseModel(&model.TJudul{})

	tableName := _tJudul.tJudulDo.TableName()
	_tJudul.ALL = field.NewAsterisk(tableName)
	_tJudul.CIDJudul = field.NewInt32(tableName, "c_id_judul")
	_tJudul.CJudul = field.NewString(tableName, "c_judul")
	_tJudul.CUpdater = field.NewString(tableName, "c_updater")
	_tJudul.CLastUpdate = field.NewTime(tableName, "c_last_update")

	_tJudul.fillFieldMap()

	return _tJudul
}

type tJudul struct {
	tJudulDo

	ALL         field.Asterisk
	CIDJudul    field.Int32
	CJudul      field.String
	CUpdater    field.String
	CLastUpdate field.Time

	fieldMap map[string]field.Expr
}

func (t tJudul) Table(newTableName string) *tJudul {
	t.tJudulDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tJudul) As(alias string) *tJudul {
	t.tJudulDo.DO = *(t.tJudulDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tJudul) updateTableName(table string) *tJudul {
	t.ALL = field.NewAsterisk(table)
	t.CIDJudul = field.NewInt32(table, "c_id_judul")
	t.CJudul = field.NewString(table, "c_judul")
	t.CUpdater = field.NewString(table, "c_updater")
	t.CLastUpdate = field.NewTime(table, "c_last_update")

	t.fillFieldMap()

	return t
}

func (t *tJudul) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tJudul) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 4)
	t.fieldMap["c_id_judul"] = t.CIDJudul
	t.fieldMap["c_judul"] = t.CJudul
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_last_update"] = t.CLastUpdate
}

func (t tJudul) clone(db *gorm.DB) tJudul {
	t.tJudulDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tJudul) replaceDB(db *gorm.DB) tJudul {
	t.tJudulDo.ReplaceDB(db)
	return t
}

type tJudulDo struct{ gen.DO }

type ITJudulDo interface {
	gen.SubQuery
	Debug() ITJudulDo
	WithContext(ctx context.Context) ITJudulDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITJudulDo
	WriteDB() ITJudulDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITJudulDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITJudulDo
	Not(conds ...gen.Condition) ITJudulDo
	Or(conds ...gen.Condition) ITJudulDo
	Select(conds ...field.Expr) ITJudulDo
	Where(conds ...gen.Condition) ITJudulDo
	Order(conds ...field.Expr) ITJudulDo
	Distinct(cols ...field.Expr) ITJudulDo
	Omit(cols ...field.Expr) ITJudulDo
	Join(table schema.Tabler, on ...field.Expr) ITJudulDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITJudulDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITJudulDo
	Group(cols ...field.Expr) ITJudulDo
	Having(conds ...gen.Condition) ITJudulDo
	Limit(limit int) ITJudulDo
	Offset(offset int) ITJudulDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITJudulDo
	Unscoped() ITJudulDo
	Create(values ...*model.TJudul) error
	CreateInBatches(values []*model.TJudul, batchSize int) error
	Save(values ...*model.TJudul) error
	First() (*model.TJudul, error)
	Take() (*model.TJudul, error)
	Last() (*model.TJudul, error)
	Find() ([]*model.TJudul, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TJudul, err error)
	FindInBatches(result *[]*model.TJudul, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TJudul) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITJudulDo
	Assign(attrs ...field.AssignExpr) ITJudulDo
	Joins(fields ...field.RelationField) ITJudulDo
	Preload(fields ...field.RelationField) ITJudulDo
	FirstOrInit() (*model.TJudul, error)
	FirstOrCreate() (*model.TJudul, error)
	FindByPage(offset int, limit int) (result []*model.TJudul, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITJudulDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tJudulDo) Debug() ITJudulDo {
	return t.withDO(t.DO.Debug())
}

func (t tJudulDo) WithContext(ctx context.Context) ITJudulDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tJudulDo) ReadDB() ITJudulDo {
	return t.Clauses(dbresolver.Read)
}

func (t tJudulDo) WriteDB() ITJudulDo {
	return t.Clauses(dbresolver.Write)
}

func (t tJudulDo) Session(config *gorm.Session) ITJudulDo {
	return t.withDO(t.DO.Session(config))
}

func (t tJudulDo) Clauses(conds ...clause.Expression) ITJudulDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tJudulDo) Returning(value interface{}, columns ...string) ITJudulDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tJudulDo) Not(conds ...gen.Condition) ITJudulDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tJudulDo) Or(conds ...gen.Condition) ITJudulDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tJudulDo) Select(conds ...field.Expr) ITJudulDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tJudulDo) Where(conds ...gen.Condition) ITJudulDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tJudulDo) Order(conds ...field.Expr) ITJudulDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tJudulDo) Distinct(cols ...field.Expr) ITJudulDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tJudulDo) Omit(cols ...field.Expr) ITJudulDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tJudulDo) Join(table schema.Tabler, on ...field.Expr) ITJudulDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tJudulDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITJudulDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tJudulDo) RightJoin(table schema.Tabler, on ...field.Expr) ITJudulDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tJudulDo) Group(cols ...field.Expr) ITJudulDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tJudulDo) Having(conds ...gen.Condition) ITJudulDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tJudulDo) Limit(limit int) ITJudulDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tJudulDo) Offset(offset int) ITJudulDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tJudulDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITJudulDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tJudulDo) Unscoped() ITJudulDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tJudulDo) Create(values ...*model.TJudul) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tJudulDo) CreateInBatches(values []*model.TJudul, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tJudulDo) Save(values ...*model.TJudul) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tJudulDo) First() (*model.TJudul, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TJudul), nil
	}
}

func (t tJudulDo) Take() (*model.TJudul, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TJudul), nil
	}
}

func (t tJudulDo) Last() (*model.TJudul, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TJudul), nil
	}
}

func (t tJudulDo) Find() ([]*model.TJudul, error) {
	result, err := t.DO.Find()
	return result.([]*model.TJudul), err
}

func (t tJudulDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TJudul, err error) {
	buf := make([]*model.TJudul, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tJudulDo) FindInBatches(result *[]*model.TJudul, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tJudulDo) Attrs(attrs ...field.AssignExpr) ITJudulDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tJudulDo) Assign(attrs ...field.AssignExpr) ITJudulDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tJudulDo) Joins(fields ...field.RelationField) ITJudulDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tJudulDo) Preload(fields ...field.RelationField) ITJudulDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tJudulDo) FirstOrInit() (*model.TJudul, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TJudul), nil
	}
}

func (t tJudulDo) FirstOrCreate() (*model.TJudul, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TJudul), nil
	}
}

func (t tJudulDo) FindByPage(offset int, limit int) (result []*model.TJudul, count int64, err error) {
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

func (t tJudulDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tJudulDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tJudulDo) Delete(models ...*model.TJudul) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tJudulDo) withDO(do gen.Dao) *tJudulDo {
	t.DO = *do.(*gen.DO)
	return t
}
