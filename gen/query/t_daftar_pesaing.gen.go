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

func newTDaftarPesaing(db *gorm.DB, opts ...gen.DOOption) tDaftarPesaing {
	_tDaftarPesaing := tDaftarPesaing{}

	_tDaftarPesaing.tDaftarPesaingDo.UseDB(db, opts...)
	_tDaftarPesaing.tDaftarPesaingDo.UseModel(&model.TDaftarPesaing{})

	tableName := _tDaftarPesaing.tDaftarPesaingDo.TableName()
	_tDaftarPesaing.ALL = field.NewAsterisk(tableName)
	_tDaftarPesaing.CIDPesaing = field.NewInt32(tableName, "c_id_pesaing")
	_tDaftarPesaing.CNamaPesaing = field.NewString(tableName, "c_nama_pesaing")
	_tDaftarPesaing.CAlamat = field.NewString(tableName, "c_alamat")
	_tDaftarPesaing.CUpdater = field.NewString(tableName, "c_updater")
	_tDaftarPesaing.CLastUpdate = field.NewTime(tableName, "c_last_update")

	_tDaftarPesaing.fillFieldMap()

	return _tDaftarPesaing
}

type tDaftarPesaing struct {
	tDaftarPesaingDo

	ALL          field.Asterisk
	CIDPesaing   field.Int32
	CNamaPesaing field.String
	CAlamat      field.String
	CUpdater     field.String
	CLastUpdate  field.Time

	fieldMap map[string]field.Expr
}

func (t tDaftarPesaing) Table(newTableName string) *tDaftarPesaing {
	t.tDaftarPesaingDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tDaftarPesaing) As(alias string) *tDaftarPesaing {
	t.tDaftarPesaingDo.DO = *(t.tDaftarPesaingDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tDaftarPesaing) updateTableName(table string) *tDaftarPesaing {
	t.ALL = field.NewAsterisk(table)
	t.CIDPesaing = field.NewInt32(table, "c_id_pesaing")
	t.CNamaPesaing = field.NewString(table, "c_nama_pesaing")
	t.CAlamat = field.NewString(table, "c_alamat")
	t.CUpdater = field.NewString(table, "c_updater")
	t.CLastUpdate = field.NewTime(table, "c_last_update")

	t.fillFieldMap()

	return t
}

func (t *tDaftarPesaing) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tDaftarPesaing) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 5)
	t.fieldMap["c_id_pesaing"] = t.CIDPesaing
	t.fieldMap["c_nama_pesaing"] = t.CNamaPesaing
	t.fieldMap["c_alamat"] = t.CAlamat
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_last_update"] = t.CLastUpdate
}

func (t tDaftarPesaing) clone(db *gorm.DB) tDaftarPesaing {
	t.tDaftarPesaingDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tDaftarPesaing) replaceDB(db *gorm.DB) tDaftarPesaing {
	t.tDaftarPesaingDo.ReplaceDB(db)
	return t
}

type tDaftarPesaingDo struct{ gen.DO }

type ITDaftarPesaingDo interface {
	gen.SubQuery
	Debug() ITDaftarPesaingDo
	WithContext(ctx context.Context) ITDaftarPesaingDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITDaftarPesaingDo
	WriteDB() ITDaftarPesaingDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITDaftarPesaingDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITDaftarPesaingDo
	Not(conds ...gen.Condition) ITDaftarPesaingDo
	Or(conds ...gen.Condition) ITDaftarPesaingDo
	Select(conds ...field.Expr) ITDaftarPesaingDo
	Where(conds ...gen.Condition) ITDaftarPesaingDo
	Order(conds ...field.Expr) ITDaftarPesaingDo
	Distinct(cols ...field.Expr) ITDaftarPesaingDo
	Omit(cols ...field.Expr) ITDaftarPesaingDo
	Join(table schema.Tabler, on ...field.Expr) ITDaftarPesaingDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITDaftarPesaingDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITDaftarPesaingDo
	Group(cols ...field.Expr) ITDaftarPesaingDo
	Having(conds ...gen.Condition) ITDaftarPesaingDo
	Limit(limit int) ITDaftarPesaingDo
	Offset(offset int) ITDaftarPesaingDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITDaftarPesaingDo
	Unscoped() ITDaftarPesaingDo
	Create(values ...*model.TDaftarPesaing) error
	CreateInBatches(values []*model.TDaftarPesaing, batchSize int) error
	Save(values ...*model.TDaftarPesaing) error
	First() (*model.TDaftarPesaing, error)
	Take() (*model.TDaftarPesaing, error)
	Last() (*model.TDaftarPesaing, error)
	Find() ([]*model.TDaftarPesaing, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TDaftarPesaing, err error)
	FindInBatches(result *[]*model.TDaftarPesaing, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TDaftarPesaing) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITDaftarPesaingDo
	Assign(attrs ...field.AssignExpr) ITDaftarPesaingDo
	Joins(fields ...field.RelationField) ITDaftarPesaingDo
	Preload(fields ...field.RelationField) ITDaftarPesaingDo
	FirstOrInit() (*model.TDaftarPesaing, error)
	FirstOrCreate() (*model.TDaftarPesaing, error)
	FindByPage(offset int, limit int) (result []*model.TDaftarPesaing, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITDaftarPesaingDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tDaftarPesaingDo) Debug() ITDaftarPesaingDo {
	return t.withDO(t.DO.Debug())
}

func (t tDaftarPesaingDo) WithContext(ctx context.Context) ITDaftarPesaingDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tDaftarPesaingDo) ReadDB() ITDaftarPesaingDo {
	return t.Clauses(dbresolver.Read)
}

func (t tDaftarPesaingDo) WriteDB() ITDaftarPesaingDo {
	return t.Clauses(dbresolver.Write)
}

func (t tDaftarPesaingDo) Session(config *gorm.Session) ITDaftarPesaingDo {
	return t.withDO(t.DO.Session(config))
}

func (t tDaftarPesaingDo) Clauses(conds ...clause.Expression) ITDaftarPesaingDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tDaftarPesaingDo) Returning(value interface{}, columns ...string) ITDaftarPesaingDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tDaftarPesaingDo) Not(conds ...gen.Condition) ITDaftarPesaingDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tDaftarPesaingDo) Or(conds ...gen.Condition) ITDaftarPesaingDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tDaftarPesaingDo) Select(conds ...field.Expr) ITDaftarPesaingDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tDaftarPesaingDo) Where(conds ...gen.Condition) ITDaftarPesaingDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tDaftarPesaingDo) Order(conds ...field.Expr) ITDaftarPesaingDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tDaftarPesaingDo) Distinct(cols ...field.Expr) ITDaftarPesaingDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tDaftarPesaingDo) Omit(cols ...field.Expr) ITDaftarPesaingDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tDaftarPesaingDo) Join(table schema.Tabler, on ...field.Expr) ITDaftarPesaingDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tDaftarPesaingDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITDaftarPesaingDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tDaftarPesaingDo) RightJoin(table schema.Tabler, on ...field.Expr) ITDaftarPesaingDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tDaftarPesaingDo) Group(cols ...field.Expr) ITDaftarPesaingDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tDaftarPesaingDo) Having(conds ...gen.Condition) ITDaftarPesaingDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tDaftarPesaingDo) Limit(limit int) ITDaftarPesaingDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tDaftarPesaingDo) Offset(offset int) ITDaftarPesaingDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tDaftarPesaingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITDaftarPesaingDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tDaftarPesaingDo) Unscoped() ITDaftarPesaingDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tDaftarPesaingDo) Create(values ...*model.TDaftarPesaing) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tDaftarPesaingDo) CreateInBatches(values []*model.TDaftarPesaing, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tDaftarPesaingDo) Save(values ...*model.TDaftarPesaing) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tDaftarPesaingDo) First() (*model.TDaftarPesaing, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TDaftarPesaing), nil
	}
}

func (t tDaftarPesaingDo) Take() (*model.TDaftarPesaing, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TDaftarPesaing), nil
	}
}

func (t tDaftarPesaingDo) Last() (*model.TDaftarPesaing, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TDaftarPesaing), nil
	}
}

func (t tDaftarPesaingDo) Find() ([]*model.TDaftarPesaing, error) {
	result, err := t.DO.Find()
	return result.([]*model.TDaftarPesaing), err
}

func (t tDaftarPesaingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TDaftarPesaing, err error) {
	buf := make([]*model.TDaftarPesaing, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tDaftarPesaingDo) FindInBatches(result *[]*model.TDaftarPesaing, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tDaftarPesaingDo) Attrs(attrs ...field.AssignExpr) ITDaftarPesaingDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tDaftarPesaingDo) Assign(attrs ...field.AssignExpr) ITDaftarPesaingDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tDaftarPesaingDo) Joins(fields ...field.RelationField) ITDaftarPesaingDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tDaftarPesaingDo) Preload(fields ...field.RelationField) ITDaftarPesaingDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tDaftarPesaingDo) FirstOrInit() (*model.TDaftarPesaing, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TDaftarPesaing), nil
	}
}

func (t tDaftarPesaingDo) FirstOrCreate() (*model.TDaftarPesaing, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TDaftarPesaing), nil
	}
}

func (t tDaftarPesaingDo) FindByPage(offset int, limit int) (result []*model.TDaftarPesaing, count int64, err error) {
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

func (t tDaftarPesaingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tDaftarPesaingDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tDaftarPesaingDo) Delete(models ...*model.TDaftarPesaing) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tDaftarPesaingDo) withDO(do gen.Dao) *tDaftarPesaingDo {
	t.DO = *do.(*gen.DO)
	return t
}
