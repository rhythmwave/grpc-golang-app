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

func newTVendorGedung(db *gorm.DB, opts ...gen.DOOption) tVendorGedung {
	_tVendorGedung := tVendorGedung{}

	_tVendorGedung.tVendorGedungDo.UseDB(db, opts...)
	_tVendorGedung.tVendorGedungDo.UseModel(&model.TVendorGedung{})

	tableName := _tVendorGedung.tVendorGedungDo.TableName()
	_tVendorGedung.ALL = field.NewAsterisk(tableName)
	_tVendorGedung.CIDVendor = field.NewInt32(tableName, "c_id_vendor")
	_tVendorGedung.CIDGedung = field.NewInt32(tableName, "c_id_gedung")
	_tVendorGedung.CUpdater = field.NewString(tableName, "c_updater")
	_tVendorGedung.CLastUpdate = field.NewTime(tableName, "c_last_update")

	_tVendorGedung.fillFieldMap()

	return _tVendorGedung
}

type tVendorGedung struct {
	tVendorGedungDo

	ALL         field.Asterisk
	CIDVendor   field.Int32
	CIDGedung   field.Int32
	CUpdater    field.String
	CLastUpdate field.Time

	fieldMap map[string]field.Expr
}

func (t tVendorGedung) Table(newTableName string) *tVendorGedung {
	t.tVendorGedungDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tVendorGedung) As(alias string) *tVendorGedung {
	t.tVendorGedungDo.DO = *(t.tVendorGedungDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tVendorGedung) updateTableName(table string) *tVendorGedung {
	t.ALL = field.NewAsterisk(table)
	t.CIDVendor = field.NewInt32(table, "c_id_vendor")
	t.CIDGedung = field.NewInt32(table, "c_id_gedung")
	t.CUpdater = field.NewString(table, "c_updater")
	t.CLastUpdate = field.NewTime(table, "c_last_update")

	t.fillFieldMap()

	return t
}

func (t *tVendorGedung) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tVendorGedung) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 4)
	t.fieldMap["c_id_vendor"] = t.CIDVendor
	t.fieldMap["c_id_gedung"] = t.CIDGedung
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_last_update"] = t.CLastUpdate
}

func (t tVendorGedung) clone(db *gorm.DB) tVendorGedung {
	t.tVendorGedungDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tVendorGedung) replaceDB(db *gorm.DB) tVendorGedung {
	t.tVendorGedungDo.ReplaceDB(db)
	return t
}

type tVendorGedungDo struct{ gen.DO }

type ITVendorGedungDo interface {
	gen.SubQuery
	Debug() ITVendorGedungDo
	WithContext(ctx context.Context) ITVendorGedungDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITVendorGedungDo
	WriteDB() ITVendorGedungDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITVendorGedungDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITVendorGedungDo
	Not(conds ...gen.Condition) ITVendorGedungDo
	Or(conds ...gen.Condition) ITVendorGedungDo
	Select(conds ...field.Expr) ITVendorGedungDo
	Where(conds ...gen.Condition) ITVendorGedungDo
	Order(conds ...field.Expr) ITVendorGedungDo
	Distinct(cols ...field.Expr) ITVendorGedungDo
	Omit(cols ...field.Expr) ITVendorGedungDo
	Join(table schema.Tabler, on ...field.Expr) ITVendorGedungDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITVendorGedungDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITVendorGedungDo
	Group(cols ...field.Expr) ITVendorGedungDo
	Having(conds ...gen.Condition) ITVendorGedungDo
	Limit(limit int) ITVendorGedungDo
	Offset(offset int) ITVendorGedungDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITVendorGedungDo
	Unscoped() ITVendorGedungDo
	Create(values ...*model.TVendorGedung) error
	CreateInBatches(values []*model.TVendorGedung, batchSize int) error
	Save(values ...*model.TVendorGedung) error
	First() (*model.TVendorGedung, error)
	Take() (*model.TVendorGedung, error)
	Last() (*model.TVendorGedung, error)
	Find() ([]*model.TVendorGedung, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TVendorGedung, err error)
	FindInBatches(result *[]*model.TVendorGedung, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TVendorGedung) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITVendorGedungDo
	Assign(attrs ...field.AssignExpr) ITVendorGedungDo
	Joins(fields ...field.RelationField) ITVendorGedungDo
	Preload(fields ...field.RelationField) ITVendorGedungDo
	FirstOrInit() (*model.TVendorGedung, error)
	FirstOrCreate() (*model.TVendorGedung, error)
	FindByPage(offset int, limit int) (result []*model.TVendorGedung, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITVendorGedungDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tVendorGedungDo) Debug() ITVendorGedungDo {
	return t.withDO(t.DO.Debug())
}

func (t tVendorGedungDo) WithContext(ctx context.Context) ITVendorGedungDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tVendorGedungDo) ReadDB() ITVendorGedungDo {
	return t.Clauses(dbresolver.Read)
}

func (t tVendorGedungDo) WriteDB() ITVendorGedungDo {
	return t.Clauses(dbresolver.Write)
}

func (t tVendorGedungDo) Session(config *gorm.Session) ITVendorGedungDo {
	return t.withDO(t.DO.Session(config))
}

func (t tVendorGedungDo) Clauses(conds ...clause.Expression) ITVendorGedungDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tVendorGedungDo) Returning(value interface{}, columns ...string) ITVendorGedungDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tVendorGedungDo) Not(conds ...gen.Condition) ITVendorGedungDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tVendorGedungDo) Or(conds ...gen.Condition) ITVendorGedungDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tVendorGedungDo) Select(conds ...field.Expr) ITVendorGedungDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tVendorGedungDo) Where(conds ...gen.Condition) ITVendorGedungDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tVendorGedungDo) Order(conds ...field.Expr) ITVendorGedungDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tVendorGedungDo) Distinct(cols ...field.Expr) ITVendorGedungDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tVendorGedungDo) Omit(cols ...field.Expr) ITVendorGedungDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tVendorGedungDo) Join(table schema.Tabler, on ...field.Expr) ITVendorGedungDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tVendorGedungDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITVendorGedungDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tVendorGedungDo) RightJoin(table schema.Tabler, on ...field.Expr) ITVendorGedungDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tVendorGedungDo) Group(cols ...field.Expr) ITVendorGedungDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tVendorGedungDo) Having(conds ...gen.Condition) ITVendorGedungDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tVendorGedungDo) Limit(limit int) ITVendorGedungDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tVendorGedungDo) Offset(offset int) ITVendorGedungDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tVendorGedungDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITVendorGedungDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tVendorGedungDo) Unscoped() ITVendorGedungDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tVendorGedungDo) Create(values ...*model.TVendorGedung) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tVendorGedungDo) CreateInBatches(values []*model.TVendorGedung, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tVendorGedungDo) Save(values ...*model.TVendorGedung) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tVendorGedungDo) First() (*model.TVendorGedung, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TVendorGedung), nil
	}
}

func (t tVendorGedungDo) Take() (*model.TVendorGedung, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TVendorGedung), nil
	}
}

func (t tVendorGedungDo) Last() (*model.TVendorGedung, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TVendorGedung), nil
	}
}

func (t tVendorGedungDo) Find() ([]*model.TVendorGedung, error) {
	result, err := t.DO.Find()
	return result.([]*model.TVendorGedung), err
}

func (t tVendorGedungDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TVendorGedung, err error) {
	buf := make([]*model.TVendorGedung, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tVendorGedungDo) FindInBatches(result *[]*model.TVendorGedung, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tVendorGedungDo) Attrs(attrs ...field.AssignExpr) ITVendorGedungDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tVendorGedungDo) Assign(attrs ...field.AssignExpr) ITVendorGedungDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tVendorGedungDo) Joins(fields ...field.RelationField) ITVendorGedungDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tVendorGedungDo) Preload(fields ...field.RelationField) ITVendorGedungDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tVendorGedungDo) FirstOrInit() (*model.TVendorGedung, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TVendorGedung), nil
	}
}

func (t tVendorGedungDo) FirstOrCreate() (*model.TVendorGedung, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TVendorGedung), nil
	}
}

func (t tVendorGedungDo) FindByPage(offset int, limit int) (result []*model.TVendorGedung, count int64, err error) {
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

func (t tVendorGedungDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tVendorGedungDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tVendorGedungDo) Delete(models ...*model.TVendorGedung) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tVendorGedungDo) withDO(do gen.Dao) *tVendorGedungDo {
	t.DO = *do.(*gen.DO)
	return t
}
