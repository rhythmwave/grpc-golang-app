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

func newTNamaBarang(db *gorm.DB, opts ...gen.DOOption) tNamaBarang {
	_tNamaBarang := tNamaBarang{}

	_tNamaBarang.tNamaBarangDo.UseDB(db, opts...)
	_tNamaBarang.tNamaBarangDo.UseModel(&model.TNamaBarang{})

	tableName := _tNamaBarang.tNamaBarangDo.TableName()
	_tNamaBarang.ALL = field.NewAsterisk(tableName)
	_tNamaBarang.CKodeBarang = field.NewString(tableName, "c_kode_barang")
	_tNamaBarang.CNamaBarang = field.NewString(tableName, "c_nama_barang")
	_tNamaBarang.CSumber = field.NewString(tableName, "c_sumber")
	_tNamaBarang.CKodeCOA = field.NewString(tableName, "c_kode_c_o_a")
	_tNamaBarang.CStatus = field.NewString(tableName, "c_status")
	_tNamaBarang.CUpdater = field.NewString(tableName, "c_updater")
	_tNamaBarang.CLastUpdate = field.NewTime(tableName, "c_last_update")

	_tNamaBarang.fillFieldMap()

	return _tNamaBarang
}

type tNamaBarang struct {
	tNamaBarangDo

	ALL         field.Asterisk
	CKodeBarang field.String
	CNamaBarang field.String
	CSumber     field.String
	CKodeCOA    field.String
	CStatus     field.String
	CUpdater    field.String
	CLastUpdate field.Time

	fieldMap map[string]field.Expr
}

func (t tNamaBarang) Table(newTableName string) *tNamaBarang {
	t.tNamaBarangDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tNamaBarang) As(alias string) *tNamaBarang {
	t.tNamaBarangDo.DO = *(t.tNamaBarangDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tNamaBarang) updateTableName(table string) *tNamaBarang {
	t.ALL = field.NewAsterisk(table)
	t.CKodeBarang = field.NewString(table, "c_kode_barang")
	t.CNamaBarang = field.NewString(table, "c_nama_barang")
	t.CSumber = field.NewString(table, "c_sumber")
	t.CKodeCOA = field.NewString(table, "c_kode_c_o_a")
	t.CStatus = field.NewString(table, "c_status")
	t.CUpdater = field.NewString(table, "c_updater")
	t.CLastUpdate = field.NewTime(table, "c_last_update")

	t.fillFieldMap()

	return t
}

func (t *tNamaBarang) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tNamaBarang) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 7)
	t.fieldMap["c_kode_barang"] = t.CKodeBarang
	t.fieldMap["c_nama_barang"] = t.CNamaBarang
	t.fieldMap["c_sumber"] = t.CSumber
	t.fieldMap["c_kode_c_o_a"] = t.CKodeCOA
	t.fieldMap["c_status"] = t.CStatus
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_last_update"] = t.CLastUpdate
}

func (t tNamaBarang) clone(db *gorm.DB) tNamaBarang {
	t.tNamaBarangDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tNamaBarang) replaceDB(db *gorm.DB) tNamaBarang {
	t.tNamaBarangDo.ReplaceDB(db)
	return t
}

type tNamaBarangDo struct{ gen.DO }

type ITNamaBarangDo interface {
	gen.SubQuery
	Debug() ITNamaBarangDo
	WithContext(ctx context.Context) ITNamaBarangDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITNamaBarangDo
	WriteDB() ITNamaBarangDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITNamaBarangDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITNamaBarangDo
	Not(conds ...gen.Condition) ITNamaBarangDo
	Or(conds ...gen.Condition) ITNamaBarangDo
	Select(conds ...field.Expr) ITNamaBarangDo
	Where(conds ...gen.Condition) ITNamaBarangDo
	Order(conds ...field.Expr) ITNamaBarangDo
	Distinct(cols ...field.Expr) ITNamaBarangDo
	Omit(cols ...field.Expr) ITNamaBarangDo
	Join(table schema.Tabler, on ...field.Expr) ITNamaBarangDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITNamaBarangDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITNamaBarangDo
	Group(cols ...field.Expr) ITNamaBarangDo
	Having(conds ...gen.Condition) ITNamaBarangDo
	Limit(limit int) ITNamaBarangDo
	Offset(offset int) ITNamaBarangDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITNamaBarangDo
	Unscoped() ITNamaBarangDo
	Create(values ...*model.TNamaBarang) error
	CreateInBatches(values []*model.TNamaBarang, batchSize int) error
	Save(values ...*model.TNamaBarang) error
	First() (*model.TNamaBarang, error)
	Take() (*model.TNamaBarang, error)
	Last() (*model.TNamaBarang, error)
	Find() ([]*model.TNamaBarang, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TNamaBarang, err error)
	FindInBatches(result *[]*model.TNamaBarang, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TNamaBarang) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITNamaBarangDo
	Assign(attrs ...field.AssignExpr) ITNamaBarangDo
	Joins(fields ...field.RelationField) ITNamaBarangDo
	Preload(fields ...field.RelationField) ITNamaBarangDo
	FirstOrInit() (*model.TNamaBarang, error)
	FirstOrCreate() (*model.TNamaBarang, error)
	FindByPage(offset int, limit int) (result []*model.TNamaBarang, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITNamaBarangDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tNamaBarangDo) Debug() ITNamaBarangDo {
	return t.withDO(t.DO.Debug())
}

func (t tNamaBarangDo) WithContext(ctx context.Context) ITNamaBarangDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tNamaBarangDo) ReadDB() ITNamaBarangDo {
	return t.Clauses(dbresolver.Read)
}

func (t tNamaBarangDo) WriteDB() ITNamaBarangDo {
	return t.Clauses(dbresolver.Write)
}

func (t tNamaBarangDo) Session(config *gorm.Session) ITNamaBarangDo {
	return t.withDO(t.DO.Session(config))
}

func (t tNamaBarangDo) Clauses(conds ...clause.Expression) ITNamaBarangDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tNamaBarangDo) Returning(value interface{}, columns ...string) ITNamaBarangDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tNamaBarangDo) Not(conds ...gen.Condition) ITNamaBarangDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tNamaBarangDo) Or(conds ...gen.Condition) ITNamaBarangDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tNamaBarangDo) Select(conds ...field.Expr) ITNamaBarangDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tNamaBarangDo) Where(conds ...gen.Condition) ITNamaBarangDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tNamaBarangDo) Order(conds ...field.Expr) ITNamaBarangDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tNamaBarangDo) Distinct(cols ...field.Expr) ITNamaBarangDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tNamaBarangDo) Omit(cols ...field.Expr) ITNamaBarangDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tNamaBarangDo) Join(table schema.Tabler, on ...field.Expr) ITNamaBarangDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tNamaBarangDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITNamaBarangDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tNamaBarangDo) RightJoin(table schema.Tabler, on ...field.Expr) ITNamaBarangDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tNamaBarangDo) Group(cols ...field.Expr) ITNamaBarangDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tNamaBarangDo) Having(conds ...gen.Condition) ITNamaBarangDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tNamaBarangDo) Limit(limit int) ITNamaBarangDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tNamaBarangDo) Offset(offset int) ITNamaBarangDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tNamaBarangDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITNamaBarangDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tNamaBarangDo) Unscoped() ITNamaBarangDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tNamaBarangDo) Create(values ...*model.TNamaBarang) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tNamaBarangDo) CreateInBatches(values []*model.TNamaBarang, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tNamaBarangDo) Save(values ...*model.TNamaBarang) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tNamaBarangDo) First() (*model.TNamaBarang, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TNamaBarang), nil
	}
}

func (t tNamaBarangDo) Take() (*model.TNamaBarang, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TNamaBarang), nil
	}
}

func (t tNamaBarangDo) Last() (*model.TNamaBarang, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TNamaBarang), nil
	}
}

func (t tNamaBarangDo) Find() ([]*model.TNamaBarang, error) {
	result, err := t.DO.Find()
	return result.([]*model.TNamaBarang), err
}

func (t tNamaBarangDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TNamaBarang, err error) {
	buf := make([]*model.TNamaBarang, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tNamaBarangDo) FindInBatches(result *[]*model.TNamaBarang, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tNamaBarangDo) Attrs(attrs ...field.AssignExpr) ITNamaBarangDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tNamaBarangDo) Assign(attrs ...field.AssignExpr) ITNamaBarangDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tNamaBarangDo) Joins(fields ...field.RelationField) ITNamaBarangDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tNamaBarangDo) Preload(fields ...field.RelationField) ITNamaBarangDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tNamaBarangDo) FirstOrInit() (*model.TNamaBarang, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TNamaBarang), nil
	}
}

func (t tNamaBarangDo) FirstOrCreate() (*model.TNamaBarang, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TNamaBarang), nil
	}
}

func (t tNamaBarangDo) FindByPage(offset int, limit int) (result []*model.TNamaBarang, count int64, err error) {
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

func (t tNamaBarangDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tNamaBarangDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tNamaBarangDo) Delete(models ...*model.TNamaBarang) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tNamaBarangDo) withDO(do gen.Dao) *tNamaBarangDo {
	t.DO = *do.(*gen.DO)
	return t
}