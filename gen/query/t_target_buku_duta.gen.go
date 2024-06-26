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

func newTTargetBukuDutum(db *gorm.DB, opts ...gen.DOOption) tTargetBukuDutum {
	_tTargetBukuDutum := tTargetBukuDutum{}

	_tTargetBukuDutum.tTargetBukuDutumDo.UseDB(db, opts...)
	_tTargetBukuDutum.tTargetBukuDutumDo.UseModel(&model.TTargetBukuDutum{})

	tableName := _tTargetBukuDutum.tTargetBukuDutumDo.TableName()
	_tTargetBukuDutum.ALL = field.NewAsterisk(tableName)
	_tTargetBukuDutum.CIDGedung = field.NewInt32(tableName, "c_id_gedung")
	_tTargetBukuDutum.CTarget = field.NewInt32(tableName, "c_target")

	_tTargetBukuDutum.fillFieldMap()

	return _tTargetBukuDutum
}

type tTargetBukuDutum struct {
	tTargetBukuDutumDo

	ALL       field.Asterisk
	CIDGedung field.Int32
	CTarget   field.Int32

	fieldMap map[string]field.Expr
}

func (t tTargetBukuDutum) Table(newTableName string) *tTargetBukuDutum {
	t.tTargetBukuDutumDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tTargetBukuDutum) As(alias string) *tTargetBukuDutum {
	t.tTargetBukuDutumDo.DO = *(t.tTargetBukuDutumDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tTargetBukuDutum) updateTableName(table string) *tTargetBukuDutum {
	t.ALL = field.NewAsterisk(table)
	t.CIDGedung = field.NewInt32(table, "c_id_gedung")
	t.CTarget = field.NewInt32(table, "c_target")

	t.fillFieldMap()

	return t
}

func (t *tTargetBukuDutum) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tTargetBukuDutum) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 2)
	t.fieldMap["c_id_gedung"] = t.CIDGedung
	t.fieldMap["c_target"] = t.CTarget
}

func (t tTargetBukuDutum) clone(db *gorm.DB) tTargetBukuDutum {
	t.tTargetBukuDutumDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tTargetBukuDutum) replaceDB(db *gorm.DB) tTargetBukuDutum {
	t.tTargetBukuDutumDo.ReplaceDB(db)
	return t
}

type tTargetBukuDutumDo struct{ gen.DO }

type ITTargetBukuDutumDo interface {
	gen.SubQuery
	Debug() ITTargetBukuDutumDo
	WithContext(ctx context.Context) ITTargetBukuDutumDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITTargetBukuDutumDo
	WriteDB() ITTargetBukuDutumDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITTargetBukuDutumDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITTargetBukuDutumDo
	Not(conds ...gen.Condition) ITTargetBukuDutumDo
	Or(conds ...gen.Condition) ITTargetBukuDutumDo
	Select(conds ...field.Expr) ITTargetBukuDutumDo
	Where(conds ...gen.Condition) ITTargetBukuDutumDo
	Order(conds ...field.Expr) ITTargetBukuDutumDo
	Distinct(cols ...field.Expr) ITTargetBukuDutumDo
	Omit(cols ...field.Expr) ITTargetBukuDutumDo
	Join(table schema.Tabler, on ...field.Expr) ITTargetBukuDutumDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITTargetBukuDutumDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITTargetBukuDutumDo
	Group(cols ...field.Expr) ITTargetBukuDutumDo
	Having(conds ...gen.Condition) ITTargetBukuDutumDo
	Limit(limit int) ITTargetBukuDutumDo
	Offset(offset int) ITTargetBukuDutumDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITTargetBukuDutumDo
	Unscoped() ITTargetBukuDutumDo
	Create(values ...*model.TTargetBukuDutum) error
	CreateInBatches(values []*model.TTargetBukuDutum, batchSize int) error
	Save(values ...*model.TTargetBukuDutum) error
	First() (*model.TTargetBukuDutum, error)
	Take() (*model.TTargetBukuDutum, error)
	Last() (*model.TTargetBukuDutum, error)
	Find() ([]*model.TTargetBukuDutum, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TTargetBukuDutum, err error)
	FindInBatches(result *[]*model.TTargetBukuDutum, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TTargetBukuDutum) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITTargetBukuDutumDo
	Assign(attrs ...field.AssignExpr) ITTargetBukuDutumDo
	Joins(fields ...field.RelationField) ITTargetBukuDutumDo
	Preload(fields ...field.RelationField) ITTargetBukuDutumDo
	FirstOrInit() (*model.TTargetBukuDutum, error)
	FirstOrCreate() (*model.TTargetBukuDutum, error)
	FindByPage(offset int, limit int) (result []*model.TTargetBukuDutum, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITTargetBukuDutumDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tTargetBukuDutumDo) Debug() ITTargetBukuDutumDo {
	return t.withDO(t.DO.Debug())
}

func (t tTargetBukuDutumDo) WithContext(ctx context.Context) ITTargetBukuDutumDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tTargetBukuDutumDo) ReadDB() ITTargetBukuDutumDo {
	return t.Clauses(dbresolver.Read)
}

func (t tTargetBukuDutumDo) WriteDB() ITTargetBukuDutumDo {
	return t.Clauses(dbresolver.Write)
}

func (t tTargetBukuDutumDo) Session(config *gorm.Session) ITTargetBukuDutumDo {
	return t.withDO(t.DO.Session(config))
}

func (t tTargetBukuDutumDo) Clauses(conds ...clause.Expression) ITTargetBukuDutumDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tTargetBukuDutumDo) Returning(value interface{}, columns ...string) ITTargetBukuDutumDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tTargetBukuDutumDo) Not(conds ...gen.Condition) ITTargetBukuDutumDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tTargetBukuDutumDo) Or(conds ...gen.Condition) ITTargetBukuDutumDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tTargetBukuDutumDo) Select(conds ...field.Expr) ITTargetBukuDutumDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tTargetBukuDutumDo) Where(conds ...gen.Condition) ITTargetBukuDutumDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tTargetBukuDutumDo) Order(conds ...field.Expr) ITTargetBukuDutumDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tTargetBukuDutumDo) Distinct(cols ...field.Expr) ITTargetBukuDutumDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tTargetBukuDutumDo) Omit(cols ...field.Expr) ITTargetBukuDutumDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tTargetBukuDutumDo) Join(table schema.Tabler, on ...field.Expr) ITTargetBukuDutumDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tTargetBukuDutumDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITTargetBukuDutumDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tTargetBukuDutumDo) RightJoin(table schema.Tabler, on ...field.Expr) ITTargetBukuDutumDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tTargetBukuDutumDo) Group(cols ...field.Expr) ITTargetBukuDutumDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tTargetBukuDutumDo) Having(conds ...gen.Condition) ITTargetBukuDutumDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tTargetBukuDutumDo) Limit(limit int) ITTargetBukuDutumDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tTargetBukuDutumDo) Offset(offset int) ITTargetBukuDutumDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tTargetBukuDutumDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITTargetBukuDutumDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tTargetBukuDutumDo) Unscoped() ITTargetBukuDutumDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tTargetBukuDutumDo) Create(values ...*model.TTargetBukuDutum) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tTargetBukuDutumDo) CreateInBatches(values []*model.TTargetBukuDutum, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tTargetBukuDutumDo) Save(values ...*model.TTargetBukuDutum) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tTargetBukuDutumDo) First() (*model.TTargetBukuDutum, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TTargetBukuDutum), nil
	}
}

func (t tTargetBukuDutumDo) Take() (*model.TTargetBukuDutum, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TTargetBukuDutum), nil
	}
}

func (t tTargetBukuDutumDo) Last() (*model.TTargetBukuDutum, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TTargetBukuDutum), nil
	}
}

func (t tTargetBukuDutumDo) Find() ([]*model.TTargetBukuDutum, error) {
	result, err := t.DO.Find()
	return result.([]*model.TTargetBukuDutum), err
}

func (t tTargetBukuDutumDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TTargetBukuDutum, err error) {
	buf := make([]*model.TTargetBukuDutum, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tTargetBukuDutumDo) FindInBatches(result *[]*model.TTargetBukuDutum, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tTargetBukuDutumDo) Attrs(attrs ...field.AssignExpr) ITTargetBukuDutumDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tTargetBukuDutumDo) Assign(attrs ...field.AssignExpr) ITTargetBukuDutumDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tTargetBukuDutumDo) Joins(fields ...field.RelationField) ITTargetBukuDutumDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tTargetBukuDutumDo) Preload(fields ...field.RelationField) ITTargetBukuDutumDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tTargetBukuDutumDo) FirstOrInit() (*model.TTargetBukuDutum, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TTargetBukuDutum), nil
	}
}

func (t tTargetBukuDutumDo) FirstOrCreate() (*model.TTargetBukuDutum, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TTargetBukuDutum), nil
	}
}

func (t tTargetBukuDutumDo) FindByPage(offset int, limit int) (result []*model.TTargetBukuDutum, count int64, err error) {
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

func (t tTargetBukuDutumDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tTargetBukuDutumDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tTargetBukuDutumDo) Delete(models ...*model.TTargetBukuDutum) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tTargetBukuDutumDo) withDO(do gen.Dao) *tTargetBukuDutumDo {
	t.DO = *do.(*gen.DO)
	return t
}
