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

func newTPesaing(db *gorm.DB, opts ...gen.DOOption) tPesaing {
	_tPesaing := tPesaing{}

	_tPesaing.tPesaingDo.UseDB(db, opts...)
	_tPesaing.tPesaingDo.UseModel(&model.TPesaing{})

	tableName := _tPesaing.tPesaingDo.TableName()
	_tPesaing.ALL = field.NewAsterisk(tableName)
	_tPesaing.CIDPesaing = field.NewInt32(tableName, "c_id_pesaing")
	_tPesaing.CNamaPesaing = field.NewString(tableName, "c_nama_pesaing")
	_tPesaing.CAlamatPesaing = field.NewString(tableName, "c_alamat_pesaing")
	_tPesaing.CIDGedung = field.NewInt32(tableName, "c_id_gedung")
	_tPesaing.CKeunikan = field.NewString(tableName, "c_keunikan")
	_tPesaing.CUpdater = field.NewString(tableName, "c_updater")
	_tPesaing.CLastUpdate = field.NewTime(tableName, "c_last_update")

	_tPesaing.fillFieldMap()

	return _tPesaing
}

type tPesaing struct {
	tPesaingDo

	ALL            field.Asterisk
	CIDPesaing     field.Int32
	CNamaPesaing   field.String
	CAlamatPesaing field.String
	CIDGedung      field.Int32
	CKeunikan      field.String
	CUpdater       field.String
	CLastUpdate    field.Time

	fieldMap map[string]field.Expr
}

func (t tPesaing) Table(newTableName string) *tPesaing {
	t.tPesaingDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tPesaing) As(alias string) *tPesaing {
	t.tPesaingDo.DO = *(t.tPesaingDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tPesaing) updateTableName(table string) *tPesaing {
	t.ALL = field.NewAsterisk(table)
	t.CIDPesaing = field.NewInt32(table, "c_id_pesaing")
	t.CNamaPesaing = field.NewString(table, "c_nama_pesaing")
	t.CAlamatPesaing = field.NewString(table, "c_alamat_pesaing")
	t.CIDGedung = field.NewInt32(table, "c_id_gedung")
	t.CKeunikan = field.NewString(table, "c_keunikan")
	t.CUpdater = field.NewString(table, "c_updater")
	t.CLastUpdate = field.NewTime(table, "c_last_update")

	t.fillFieldMap()

	return t
}

func (t *tPesaing) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tPesaing) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 7)
	t.fieldMap["c_id_pesaing"] = t.CIDPesaing
	t.fieldMap["c_nama_pesaing"] = t.CNamaPesaing
	t.fieldMap["c_alamat_pesaing"] = t.CAlamatPesaing
	t.fieldMap["c_id_gedung"] = t.CIDGedung
	t.fieldMap["c_keunikan"] = t.CKeunikan
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_last_update"] = t.CLastUpdate
}

func (t tPesaing) clone(db *gorm.DB) tPesaing {
	t.tPesaingDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tPesaing) replaceDB(db *gorm.DB) tPesaing {
	t.tPesaingDo.ReplaceDB(db)
	return t
}

type tPesaingDo struct{ gen.DO }

type ITPesaingDo interface {
	gen.SubQuery
	Debug() ITPesaingDo
	WithContext(ctx context.Context) ITPesaingDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITPesaingDo
	WriteDB() ITPesaingDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITPesaingDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITPesaingDo
	Not(conds ...gen.Condition) ITPesaingDo
	Or(conds ...gen.Condition) ITPesaingDo
	Select(conds ...field.Expr) ITPesaingDo
	Where(conds ...gen.Condition) ITPesaingDo
	Order(conds ...field.Expr) ITPesaingDo
	Distinct(cols ...field.Expr) ITPesaingDo
	Omit(cols ...field.Expr) ITPesaingDo
	Join(table schema.Tabler, on ...field.Expr) ITPesaingDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITPesaingDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITPesaingDo
	Group(cols ...field.Expr) ITPesaingDo
	Having(conds ...gen.Condition) ITPesaingDo
	Limit(limit int) ITPesaingDo
	Offset(offset int) ITPesaingDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITPesaingDo
	Unscoped() ITPesaingDo
	Create(values ...*model.TPesaing) error
	CreateInBatches(values []*model.TPesaing, batchSize int) error
	Save(values ...*model.TPesaing) error
	First() (*model.TPesaing, error)
	Take() (*model.TPesaing, error)
	Last() (*model.TPesaing, error)
	Find() ([]*model.TPesaing, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TPesaing, err error)
	FindInBatches(result *[]*model.TPesaing, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TPesaing) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITPesaingDo
	Assign(attrs ...field.AssignExpr) ITPesaingDo
	Joins(fields ...field.RelationField) ITPesaingDo
	Preload(fields ...field.RelationField) ITPesaingDo
	FirstOrInit() (*model.TPesaing, error)
	FirstOrCreate() (*model.TPesaing, error)
	FindByPage(offset int, limit int) (result []*model.TPesaing, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITPesaingDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tPesaingDo) Debug() ITPesaingDo {
	return t.withDO(t.DO.Debug())
}

func (t tPesaingDo) WithContext(ctx context.Context) ITPesaingDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tPesaingDo) ReadDB() ITPesaingDo {
	return t.Clauses(dbresolver.Read)
}

func (t tPesaingDo) WriteDB() ITPesaingDo {
	return t.Clauses(dbresolver.Write)
}

func (t tPesaingDo) Session(config *gorm.Session) ITPesaingDo {
	return t.withDO(t.DO.Session(config))
}

func (t tPesaingDo) Clauses(conds ...clause.Expression) ITPesaingDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tPesaingDo) Returning(value interface{}, columns ...string) ITPesaingDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tPesaingDo) Not(conds ...gen.Condition) ITPesaingDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tPesaingDo) Or(conds ...gen.Condition) ITPesaingDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tPesaingDo) Select(conds ...field.Expr) ITPesaingDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tPesaingDo) Where(conds ...gen.Condition) ITPesaingDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tPesaingDo) Order(conds ...field.Expr) ITPesaingDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tPesaingDo) Distinct(cols ...field.Expr) ITPesaingDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tPesaingDo) Omit(cols ...field.Expr) ITPesaingDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tPesaingDo) Join(table schema.Tabler, on ...field.Expr) ITPesaingDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tPesaingDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITPesaingDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tPesaingDo) RightJoin(table schema.Tabler, on ...field.Expr) ITPesaingDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tPesaingDo) Group(cols ...field.Expr) ITPesaingDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tPesaingDo) Having(conds ...gen.Condition) ITPesaingDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tPesaingDo) Limit(limit int) ITPesaingDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tPesaingDo) Offset(offset int) ITPesaingDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tPesaingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITPesaingDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tPesaingDo) Unscoped() ITPesaingDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tPesaingDo) Create(values ...*model.TPesaing) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tPesaingDo) CreateInBatches(values []*model.TPesaing, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tPesaingDo) Save(values ...*model.TPesaing) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tPesaingDo) First() (*model.TPesaing, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPesaing), nil
	}
}

func (t tPesaingDo) Take() (*model.TPesaing, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPesaing), nil
	}
}

func (t tPesaingDo) Last() (*model.TPesaing, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPesaing), nil
	}
}

func (t tPesaingDo) Find() ([]*model.TPesaing, error) {
	result, err := t.DO.Find()
	return result.([]*model.TPesaing), err
}

func (t tPesaingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TPesaing, err error) {
	buf := make([]*model.TPesaing, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tPesaingDo) FindInBatches(result *[]*model.TPesaing, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tPesaingDo) Attrs(attrs ...field.AssignExpr) ITPesaingDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tPesaingDo) Assign(attrs ...field.AssignExpr) ITPesaingDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tPesaingDo) Joins(fields ...field.RelationField) ITPesaingDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tPesaingDo) Preload(fields ...field.RelationField) ITPesaingDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tPesaingDo) FirstOrInit() (*model.TPesaing, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPesaing), nil
	}
}

func (t tPesaingDo) FirstOrCreate() (*model.TPesaing, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPesaing), nil
	}
}

func (t tPesaingDo) FindByPage(offset int, limit int) (result []*model.TPesaing, count int64, err error) {
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

func (t tPesaingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tPesaingDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tPesaingDo) Delete(models ...*model.TPesaing) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tPesaingDo) withDO(do gen.Dao) *tPesaingDo {
	t.DO = *do.(*gen.DO)
	return t
}
