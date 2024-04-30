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

func newTIsiItemKelAnggaran(db *gorm.DB, opts ...gen.DOOption) tIsiItemKelAnggaran {
	_tIsiItemKelAnggaran := tIsiItemKelAnggaran{}

	_tIsiItemKelAnggaran.tIsiItemKelAnggaranDo.UseDB(db, opts...)
	_tIsiItemKelAnggaran.tIsiItemKelAnggaranDo.UseModel(&model.TIsiItemKelAnggaran{})

	tableName := _tIsiItemKelAnggaran.tIsiItemKelAnggaranDo.TableName()
	_tIsiItemKelAnggaran.ALL = field.NewAsterisk(tableName)
	_tIsiItemKelAnggaran.CKodeKelAnggaran = field.NewString(tableName, "c_kode_kel_anggaran")
	_tIsiItemKelAnggaran.CKodeItemAnggaran = field.NewString(tableName, "c_kode_item_anggaran")
	_tIsiItemKelAnggaran.CUpdater = field.NewString(tableName, "c_updater")
	_tIsiItemKelAnggaran.CLastUpdate = field.NewTime(tableName, "c_last_update")

	_tIsiItemKelAnggaran.fillFieldMap()

	return _tIsiItemKelAnggaran
}

type tIsiItemKelAnggaran struct {
	tIsiItemKelAnggaranDo

	ALL               field.Asterisk
	CKodeKelAnggaran  field.String
	CKodeItemAnggaran field.String
	CUpdater          field.String
	CLastUpdate       field.Time

	fieldMap map[string]field.Expr
}

func (t tIsiItemKelAnggaran) Table(newTableName string) *tIsiItemKelAnggaran {
	t.tIsiItemKelAnggaranDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tIsiItemKelAnggaran) As(alias string) *tIsiItemKelAnggaran {
	t.tIsiItemKelAnggaranDo.DO = *(t.tIsiItemKelAnggaranDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tIsiItemKelAnggaran) updateTableName(table string) *tIsiItemKelAnggaran {
	t.ALL = field.NewAsterisk(table)
	t.CKodeKelAnggaran = field.NewString(table, "c_kode_kel_anggaran")
	t.CKodeItemAnggaran = field.NewString(table, "c_kode_item_anggaran")
	t.CUpdater = field.NewString(table, "c_updater")
	t.CLastUpdate = field.NewTime(table, "c_last_update")

	t.fillFieldMap()

	return t
}

func (t *tIsiItemKelAnggaran) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tIsiItemKelAnggaran) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 4)
	t.fieldMap["c_kode_kel_anggaran"] = t.CKodeKelAnggaran
	t.fieldMap["c_kode_item_anggaran"] = t.CKodeItemAnggaran
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_last_update"] = t.CLastUpdate
}

func (t tIsiItemKelAnggaran) clone(db *gorm.DB) tIsiItemKelAnggaran {
	t.tIsiItemKelAnggaranDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tIsiItemKelAnggaran) replaceDB(db *gorm.DB) tIsiItemKelAnggaran {
	t.tIsiItemKelAnggaranDo.ReplaceDB(db)
	return t
}

type tIsiItemKelAnggaranDo struct{ gen.DO }

type ITIsiItemKelAnggaranDo interface {
	gen.SubQuery
	Debug() ITIsiItemKelAnggaranDo
	WithContext(ctx context.Context) ITIsiItemKelAnggaranDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITIsiItemKelAnggaranDo
	WriteDB() ITIsiItemKelAnggaranDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITIsiItemKelAnggaranDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITIsiItemKelAnggaranDo
	Not(conds ...gen.Condition) ITIsiItemKelAnggaranDo
	Or(conds ...gen.Condition) ITIsiItemKelAnggaranDo
	Select(conds ...field.Expr) ITIsiItemKelAnggaranDo
	Where(conds ...gen.Condition) ITIsiItemKelAnggaranDo
	Order(conds ...field.Expr) ITIsiItemKelAnggaranDo
	Distinct(cols ...field.Expr) ITIsiItemKelAnggaranDo
	Omit(cols ...field.Expr) ITIsiItemKelAnggaranDo
	Join(table schema.Tabler, on ...field.Expr) ITIsiItemKelAnggaranDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITIsiItemKelAnggaranDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITIsiItemKelAnggaranDo
	Group(cols ...field.Expr) ITIsiItemKelAnggaranDo
	Having(conds ...gen.Condition) ITIsiItemKelAnggaranDo
	Limit(limit int) ITIsiItemKelAnggaranDo
	Offset(offset int) ITIsiItemKelAnggaranDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITIsiItemKelAnggaranDo
	Unscoped() ITIsiItemKelAnggaranDo
	Create(values ...*model.TIsiItemKelAnggaran) error
	CreateInBatches(values []*model.TIsiItemKelAnggaran, batchSize int) error
	Save(values ...*model.TIsiItemKelAnggaran) error
	First() (*model.TIsiItemKelAnggaran, error)
	Take() (*model.TIsiItemKelAnggaran, error)
	Last() (*model.TIsiItemKelAnggaran, error)
	Find() ([]*model.TIsiItemKelAnggaran, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TIsiItemKelAnggaran, err error)
	FindInBatches(result *[]*model.TIsiItemKelAnggaran, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TIsiItemKelAnggaran) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITIsiItemKelAnggaranDo
	Assign(attrs ...field.AssignExpr) ITIsiItemKelAnggaranDo
	Joins(fields ...field.RelationField) ITIsiItemKelAnggaranDo
	Preload(fields ...field.RelationField) ITIsiItemKelAnggaranDo
	FirstOrInit() (*model.TIsiItemKelAnggaran, error)
	FirstOrCreate() (*model.TIsiItemKelAnggaran, error)
	FindByPage(offset int, limit int) (result []*model.TIsiItemKelAnggaran, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITIsiItemKelAnggaranDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tIsiItemKelAnggaranDo) Debug() ITIsiItemKelAnggaranDo {
	return t.withDO(t.DO.Debug())
}

func (t tIsiItemKelAnggaranDo) WithContext(ctx context.Context) ITIsiItemKelAnggaranDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tIsiItemKelAnggaranDo) ReadDB() ITIsiItemKelAnggaranDo {
	return t.Clauses(dbresolver.Read)
}

func (t tIsiItemKelAnggaranDo) WriteDB() ITIsiItemKelAnggaranDo {
	return t.Clauses(dbresolver.Write)
}

func (t tIsiItemKelAnggaranDo) Session(config *gorm.Session) ITIsiItemKelAnggaranDo {
	return t.withDO(t.DO.Session(config))
}

func (t tIsiItemKelAnggaranDo) Clauses(conds ...clause.Expression) ITIsiItemKelAnggaranDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tIsiItemKelAnggaranDo) Returning(value interface{}, columns ...string) ITIsiItemKelAnggaranDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tIsiItemKelAnggaranDo) Not(conds ...gen.Condition) ITIsiItemKelAnggaranDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tIsiItemKelAnggaranDo) Or(conds ...gen.Condition) ITIsiItemKelAnggaranDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tIsiItemKelAnggaranDo) Select(conds ...field.Expr) ITIsiItemKelAnggaranDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tIsiItemKelAnggaranDo) Where(conds ...gen.Condition) ITIsiItemKelAnggaranDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tIsiItemKelAnggaranDo) Order(conds ...field.Expr) ITIsiItemKelAnggaranDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tIsiItemKelAnggaranDo) Distinct(cols ...field.Expr) ITIsiItemKelAnggaranDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tIsiItemKelAnggaranDo) Omit(cols ...field.Expr) ITIsiItemKelAnggaranDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tIsiItemKelAnggaranDo) Join(table schema.Tabler, on ...field.Expr) ITIsiItemKelAnggaranDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tIsiItemKelAnggaranDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITIsiItemKelAnggaranDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tIsiItemKelAnggaranDo) RightJoin(table schema.Tabler, on ...field.Expr) ITIsiItemKelAnggaranDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tIsiItemKelAnggaranDo) Group(cols ...field.Expr) ITIsiItemKelAnggaranDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tIsiItemKelAnggaranDo) Having(conds ...gen.Condition) ITIsiItemKelAnggaranDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tIsiItemKelAnggaranDo) Limit(limit int) ITIsiItemKelAnggaranDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tIsiItemKelAnggaranDo) Offset(offset int) ITIsiItemKelAnggaranDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tIsiItemKelAnggaranDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITIsiItemKelAnggaranDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tIsiItemKelAnggaranDo) Unscoped() ITIsiItemKelAnggaranDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tIsiItemKelAnggaranDo) Create(values ...*model.TIsiItemKelAnggaran) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tIsiItemKelAnggaranDo) CreateInBatches(values []*model.TIsiItemKelAnggaran, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tIsiItemKelAnggaranDo) Save(values ...*model.TIsiItemKelAnggaran) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tIsiItemKelAnggaranDo) First() (*model.TIsiItemKelAnggaran, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TIsiItemKelAnggaran), nil
	}
}

func (t tIsiItemKelAnggaranDo) Take() (*model.TIsiItemKelAnggaran, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TIsiItemKelAnggaran), nil
	}
}

func (t tIsiItemKelAnggaranDo) Last() (*model.TIsiItemKelAnggaran, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TIsiItemKelAnggaran), nil
	}
}

func (t tIsiItemKelAnggaranDo) Find() ([]*model.TIsiItemKelAnggaran, error) {
	result, err := t.DO.Find()
	return result.([]*model.TIsiItemKelAnggaran), err
}

func (t tIsiItemKelAnggaranDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TIsiItemKelAnggaran, err error) {
	buf := make([]*model.TIsiItemKelAnggaran, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tIsiItemKelAnggaranDo) FindInBatches(result *[]*model.TIsiItemKelAnggaran, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tIsiItemKelAnggaranDo) Attrs(attrs ...field.AssignExpr) ITIsiItemKelAnggaranDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tIsiItemKelAnggaranDo) Assign(attrs ...field.AssignExpr) ITIsiItemKelAnggaranDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tIsiItemKelAnggaranDo) Joins(fields ...field.RelationField) ITIsiItemKelAnggaranDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tIsiItemKelAnggaranDo) Preload(fields ...field.RelationField) ITIsiItemKelAnggaranDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tIsiItemKelAnggaranDo) FirstOrInit() (*model.TIsiItemKelAnggaran, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TIsiItemKelAnggaran), nil
	}
}

func (t tIsiItemKelAnggaranDo) FirstOrCreate() (*model.TIsiItemKelAnggaran, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TIsiItemKelAnggaran), nil
	}
}

func (t tIsiItemKelAnggaranDo) FindByPage(offset int, limit int) (result []*model.TIsiItemKelAnggaran, count int64, err error) {
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

func (t tIsiItemKelAnggaranDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tIsiItemKelAnggaranDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tIsiItemKelAnggaranDo) Delete(models ...*model.TIsiItemKelAnggaran) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tIsiItemKelAnggaranDo) withDO(do gen.Dao) *tIsiItemKelAnggaranDo {
	t.DO = *do.(*gen.DO)
	return t
}
