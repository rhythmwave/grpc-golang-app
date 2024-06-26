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

func newTAcuanHargaBarang(db *gorm.DB, opts ...gen.DOOption) tAcuanHargaBarang {
	_tAcuanHargaBarang := tAcuanHargaBarang{}

	_tAcuanHargaBarang.tAcuanHargaBarangDo.UseDB(db, opts...)
	_tAcuanHargaBarang.tAcuanHargaBarangDo.UseModel(&model.TAcuanHargaBarang{})

	tableName := _tAcuanHargaBarang.tAcuanHargaBarangDo.TableName()
	_tAcuanHargaBarang.ALL = field.NewAsterisk(tableName)
	_tAcuanHargaBarang.CKodeBarang = field.NewString(tableName, "c_kode_barang")
	_tAcuanHargaBarang.CIDSatuan = field.NewInt32(tableName, "c_id_satuan")
	_tAcuanHargaBarang.CIDKota = field.NewInt32(tableName, "c_id_kota")
	_tAcuanHargaBarang.CAcuanHarga = field.NewFloat64(tableName, "c_acuan_harga")
	_tAcuanHargaBarang.CUpdater = field.NewString(tableName, "c_updater")
	_tAcuanHargaBarang.CLastUpdate = field.NewTime(tableName, "c_last_update")

	_tAcuanHargaBarang.fillFieldMap()

	return _tAcuanHargaBarang
}

type tAcuanHargaBarang struct {
	tAcuanHargaBarangDo

	ALL         field.Asterisk
	CKodeBarang field.String
	CIDSatuan   field.Int32
	CIDKota     field.Int32
	CAcuanHarga field.Float64
	CUpdater    field.String
	CLastUpdate field.Time

	fieldMap map[string]field.Expr
}

func (t tAcuanHargaBarang) Table(newTableName string) *tAcuanHargaBarang {
	t.tAcuanHargaBarangDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tAcuanHargaBarang) As(alias string) *tAcuanHargaBarang {
	t.tAcuanHargaBarangDo.DO = *(t.tAcuanHargaBarangDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tAcuanHargaBarang) updateTableName(table string) *tAcuanHargaBarang {
	t.ALL = field.NewAsterisk(table)
	t.CKodeBarang = field.NewString(table, "c_kode_barang")
	t.CIDSatuan = field.NewInt32(table, "c_id_satuan")
	t.CIDKota = field.NewInt32(table, "c_id_kota")
	t.CAcuanHarga = field.NewFloat64(table, "c_acuan_harga")
	t.CUpdater = field.NewString(table, "c_updater")
	t.CLastUpdate = field.NewTime(table, "c_last_update")

	t.fillFieldMap()

	return t
}

func (t *tAcuanHargaBarang) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tAcuanHargaBarang) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 6)
	t.fieldMap["c_kode_barang"] = t.CKodeBarang
	t.fieldMap["c_id_satuan"] = t.CIDSatuan
	t.fieldMap["c_id_kota"] = t.CIDKota
	t.fieldMap["c_acuan_harga"] = t.CAcuanHarga
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_last_update"] = t.CLastUpdate
}

func (t tAcuanHargaBarang) clone(db *gorm.DB) tAcuanHargaBarang {
	t.tAcuanHargaBarangDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tAcuanHargaBarang) replaceDB(db *gorm.DB) tAcuanHargaBarang {
	t.tAcuanHargaBarangDo.ReplaceDB(db)
	return t
}

type tAcuanHargaBarangDo struct{ gen.DO }

type ITAcuanHargaBarangDo interface {
	gen.SubQuery
	Debug() ITAcuanHargaBarangDo
	WithContext(ctx context.Context) ITAcuanHargaBarangDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITAcuanHargaBarangDo
	WriteDB() ITAcuanHargaBarangDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITAcuanHargaBarangDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITAcuanHargaBarangDo
	Not(conds ...gen.Condition) ITAcuanHargaBarangDo
	Or(conds ...gen.Condition) ITAcuanHargaBarangDo
	Select(conds ...field.Expr) ITAcuanHargaBarangDo
	Where(conds ...gen.Condition) ITAcuanHargaBarangDo
	Order(conds ...field.Expr) ITAcuanHargaBarangDo
	Distinct(cols ...field.Expr) ITAcuanHargaBarangDo
	Omit(cols ...field.Expr) ITAcuanHargaBarangDo
	Join(table schema.Tabler, on ...field.Expr) ITAcuanHargaBarangDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITAcuanHargaBarangDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITAcuanHargaBarangDo
	Group(cols ...field.Expr) ITAcuanHargaBarangDo
	Having(conds ...gen.Condition) ITAcuanHargaBarangDo
	Limit(limit int) ITAcuanHargaBarangDo
	Offset(offset int) ITAcuanHargaBarangDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITAcuanHargaBarangDo
	Unscoped() ITAcuanHargaBarangDo
	Create(values ...*model.TAcuanHargaBarang) error
	CreateInBatches(values []*model.TAcuanHargaBarang, batchSize int) error
	Save(values ...*model.TAcuanHargaBarang) error
	First() (*model.TAcuanHargaBarang, error)
	Take() (*model.TAcuanHargaBarang, error)
	Last() (*model.TAcuanHargaBarang, error)
	Find() ([]*model.TAcuanHargaBarang, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TAcuanHargaBarang, err error)
	FindInBatches(result *[]*model.TAcuanHargaBarang, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TAcuanHargaBarang) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITAcuanHargaBarangDo
	Assign(attrs ...field.AssignExpr) ITAcuanHargaBarangDo
	Joins(fields ...field.RelationField) ITAcuanHargaBarangDo
	Preload(fields ...field.RelationField) ITAcuanHargaBarangDo
	FirstOrInit() (*model.TAcuanHargaBarang, error)
	FirstOrCreate() (*model.TAcuanHargaBarang, error)
	FindByPage(offset int, limit int) (result []*model.TAcuanHargaBarang, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITAcuanHargaBarangDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tAcuanHargaBarangDo) Debug() ITAcuanHargaBarangDo {
	return t.withDO(t.DO.Debug())
}

func (t tAcuanHargaBarangDo) WithContext(ctx context.Context) ITAcuanHargaBarangDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tAcuanHargaBarangDo) ReadDB() ITAcuanHargaBarangDo {
	return t.Clauses(dbresolver.Read)
}

func (t tAcuanHargaBarangDo) WriteDB() ITAcuanHargaBarangDo {
	return t.Clauses(dbresolver.Write)
}

func (t tAcuanHargaBarangDo) Session(config *gorm.Session) ITAcuanHargaBarangDo {
	return t.withDO(t.DO.Session(config))
}

func (t tAcuanHargaBarangDo) Clauses(conds ...clause.Expression) ITAcuanHargaBarangDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tAcuanHargaBarangDo) Returning(value interface{}, columns ...string) ITAcuanHargaBarangDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tAcuanHargaBarangDo) Not(conds ...gen.Condition) ITAcuanHargaBarangDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tAcuanHargaBarangDo) Or(conds ...gen.Condition) ITAcuanHargaBarangDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tAcuanHargaBarangDo) Select(conds ...field.Expr) ITAcuanHargaBarangDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tAcuanHargaBarangDo) Where(conds ...gen.Condition) ITAcuanHargaBarangDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tAcuanHargaBarangDo) Order(conds ...field.Expr) ITAcuanHargaBarangDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tAcuanHargaBarangDo) Distinct(cols ...field.Expr) ITAcuanHargaBarangDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tAcuanHargaBarangDo) Omit(cols ...field.Expr) ITAcuanHargaBarangDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tAcuanHargaBarangDo) Join(table schema.Tabler, on ...field.Expr) ITAcuanHargaBarangDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tAcuanHargaBarangDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITAcuanHargaBarangDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tAcuanHargaBarangDo) RightJoin(table schema.Tabler, on ...field.Expr) ITAcuanHargaBarangDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tAcuanHargaBarangDo) Group(cols ...field.Expr) ITAcuanHargaBarangDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tAcuanHargaBarangDo) Having(conds ...gen.Condition) ITAcuanHargaBarangDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tAcuanHargaBarangDo) Limit(limit int) ITAcuanHargaBarangDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tAcuanHargaBarangDo) Offset(offset int) ITAcuanHargaBarangDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tAcuanHargaBarangDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITAcuanHargaBarangDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tAcuanHargaBarangDo) Unscoped() ITAcuanHargaBarangDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tAcuanHargaBarangDo) Create(values ...*model.TAcuanHargaBarang) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tAcuanHargaBarangDo) CreateInBatches(values []*model.TAcuanHargaBarang, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tAcuanHargaBarangDo) Save(values ...*model.TAcuanHargaBarang) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tAcuanHargaBarangDo) First() (*model.TAcuanHargaBarang, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAcuanHargaBarang), nil
	}
}

func (t tAcuanHargaBarangDo) Take() (*model.TAcuanHargaBarang, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAcuanHargaBarang), nil
	}
}

func (t tAcuanHargaBarangDo) Last() (*model.TAcuanHargaBarang, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAcuanHargaBarang), nil
	}
}

func (t tAcuanHargaBarangDo) Find() ([]*model.TAcuanHargaBarang, error) {
	result, err := t.DO.Find()
	return result.([]*model.TAcuanHargaBarang), err
}

func (t tAcuanHargaBarangDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TAcuanHargaBarang, err error) {
	buf := make([]*model.TAcuanHargaBarang, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tAcuanHargaBarangDo) FindInBatches(result *[]*model.TAcuanHargaBarang, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tAcuanHargaBarangDo) Attrs(attrs ...field.AssignExpr) ITAcuanHargaBarangDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tAcuanHargaBarangDo) Assign(attrs ...field.AssignExpr) ITAcuanHargaBarangDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tAcuanHargaBarangDo) Joins(fields ...field.RelationField) ITAcuanHargaBarangDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tAcuanHargaBarangDo) Preload(fields ...field.RelationField) ITAcuanHargaBarangDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tAcuanHargaBarangDo) FirstOrInit() (*model.TAcuanHargaBarang, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAcuanHargaBarang), nil
	}
}

func (t tAcuanHargaBarangDo) FirstOrCreate() (*model.TAcuanHargaBarang, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAcuanHargaBarang), nil
	}
}

func (t tAcuanHargaBarangDo) FindByPage(offset int, limit int) (result []*model.TAcuanHargaBarang, count int64, err error) {
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

func (t tAcuanHargaBarangDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tAcuanHargaBarangDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tAcuanHargaBarangDo) Delete(models ...*model.TAcuanHargaBarang) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tAcuanHargaBarangDo) withDO(do gen.Dao) *tAcuanHargaBarangDo {
	t.DO = *do.(*gen.DO)
	return t
}
