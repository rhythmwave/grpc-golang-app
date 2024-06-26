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

func newTPengajuanTransfer(db *gorm.DB, opts ...gen.DOOption) tPengajuanTransfer {
	_tPengajuanTransfer := tPengajuanTransfer{}

	_tPengajuanTransfer.tPengajuanTransferDo.UseDB(db, opts...)
	_tPengajuanTransfer.tPengajuanTransferDo.UseModel(&model.TPengajuanTransfer{})

	tableName := _tPengajuanTransfer.tPengajuanTransferDo.TableName()
	_tPengajuanTransfer.ALL = field.NewAsterisk(tableName)
	_tPengajuanTransfer.CIDKegiatan = field.NewInt32(tableName, "c_id_kegiatan")
	_tPengajuanTransfer.CIDJenis = field.NewInt32(tableName, "c_id_jenis")
	_tPengajuanTransfer.CBiayaSetujuAkuntansi = field.NewFloat64(tableName, "c_biaya_setuju_akuntansi")
	_tPengajuanTransfer.CUpdater = field.NewString(tableName, "c_updater")
	_tPengajuanTransfer.CLastUpdate = field.NewTime(tableName, "c_last_update")

	_tPengajuanTransfer.fillFieldMap()

	return _tPengajuanTransfer
}

type tPengajuanTransfer struct {
	tPengajuanTransferDo

	ALL                   field.Asterisk
	CIDKegiatan           field.Int32
	CIDJenis              field.Int32
	CBiayaSetujuAkuntansi field.Float64
	CUpdater              field.String
	CLastUpdate           field.Time

	fieldMap map[string]field.Expr
}

func (t tPengajuanTransfer) Table(newTableName string) *tPengajuanTransfer {
	t.tPengajuanTransferDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tPengajuanTransfer) As(alias string) *tPengajuanTransfer {
	t.tPengajuanTransferDo.DO = *(t.tPengajuanTransferDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tPengajuanTransfer) updateTableName(table string) *tPengajuanTransfer {
	t.ALL = field.NewAsterisk(table)
	t.CIDKegiatan = field.NewInt32(table, "c_id_kegiatan")
	t.CIDJenis = field.NewInt32(table, "c_id_jenis")
	t.CBiayaSetujuAkuntansi = field.NewFloat64(table, "c_biaya_setuju_akuntansi")
	t.CUpdater = field.NewString(table, "c_updater")
	t.CLastUpdate = field.NewTime(table, "c_last_update")

	t.fillFieldMap()

	return t
}

func (t *tPengajuanTransfer) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tPengajuanTransfer) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 5)
	t.fieldMap["c_id_kegiatan"] = t.CIDKegiatan
	t.fieldMap["c_id_jenis"] = t.CIDJenis
	t.fieldMap["c_biaya_setuju_akuntansi"] = t.CBiayaSetujuAkuntansi
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_last_update"] = t.CLastUpdate
}

func (t tPengajuanTransfer) clone(db *gorm.DB) tPengajuanTransfer {
	t.tPengajuanTransferDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tPengajuanTransfer) replaceDB(db *gorm.DB) tPengajuanTransfer {
	t.tPengajuanTransferDo.ReplaceDB(db)
	return t
}

type tPengajuanTransferDo struct{ gen.DO }

type ITPengajuanTransferDo interface {
	gen.SubQuery
	Debug() ITPengajuanTransferDo
	WithContext(ctx context.Context) ITPengajuanTransferDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITPengajuanTransferDo
	WriteDB() ITPengajuanTransferDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITPengajuanTransferDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITPengajuanTransferDo
	Not(conds ...gen.Condition) ITPengajuanTransferDo
	Or(conds ...gen.Condition) ITPengajuanTransferDo
	Select(conds ...field.Expr) ITPengajuanTransferDo
	Where(conds ...gen.Condition) ITPengajuanTransferDo
	Order(conds ...field.Expr) ITPengajuanTransferDo
	Distinct(cols ...field.Expr) ITPengajuanTransferDo
	Omit(cols ...field.Expr) ITPengajuanTransferDo
	Join(table schema.Tabler, on ...field.Expr) ITPengajuanTransferDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITPengajuanTransferDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITPengajuanTransferDo
	Group(cols ...field.Expr) ITPengajuanTransferDo
	Having(conds ...gen.Condition) ITPengajuanTransferDo
	Limit(limit int) ITPengajuanTransferDo
	Offset(offset int) ITPengajuanTransferDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITPengajuanTransferDo
	Unscoped() ITPengajuanTransferDo
	Create(values ...*model.TPengajuanTransfer) error
	CreateInBatches(values []*model.TPengajuanTransfer, batchSize int) error
	Save(values ...*model.TPengajuanTransfer) error
	First() (*model.TPengajuanTransfer, error)
	Take() (*model.TPengajuanTransfer, error)
	Last() (*model.TPengajuanTransfer, error)
	Find() ([]*model.TPengajuanTransfer, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TPengajuanTransfer, err error)
	FindInBatches(result *[]*model.TPengajuanTransfer, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TPengajuanTransfer) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITPengajuanTransferDo
	Assign(attrs ...field.AssignExpr) ITPengajuanTransferDo
	Joins(fields ...field.RelationField) ITPengajuanTransferDo
	Preload(fields ...field.RelationField) ITPengajuanTransferDo
	FirstOrInit() (*model.TPengajuanTransfer, error)
	FirstOrCreate() (*model.TPengajuanTransfer, error)
	FindByPage(offset int, limit int) (result []*model.TPengajuanTransfer, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITPengajuanTransferDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tPengajuanTransferDo) Debug() ITPengajuanTransferDo {
	return t.withDO(t.DO.Debug())
}

func (t tPengajuanTransferDo) WithContext(ctx context.Context) ITPengajuanTransferDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tPengajuanTransferDo) ReadDB() ITPengajuanTransferDo {
	return t.Clauses(dbresolver.Read)
}

func (t tPengajuanTransferDo) WriteDB() ITPengajuanTransferDo {
	return t.Clauses(dbresolver.Write)
}

func (t tPengajuanTransferDo) Session(config *gorm.Session) ITPengajuanTransferDo {
	return t.withDO(t.DO.Session(config))
}

func (t tPengajuanTransferDo) Clauses(conds ...clause.Expression) ITPengajuanTransferDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tPengajuanTransferDo) Returning(value interface{}, columns ...string) ITPengajuanTransferDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tPengajuanTransferDo) Not(conds ...gen.Condition) ITPengajuanTransferDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tPengajuanTransferDo) Or(conds ...gen.Condition) ITPengajuanTransferDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tPengajuanTransferDo) Select(conds ...field.Expr) ITPengajuanTransferDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tPengajuanTransferDo) Where(conds ...gen.Condition) ITPengajuanTransferDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tPengajuanTransferDo) Order(conds ...field.Expr) ITPengajuanTransferDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tPengajuanTransferDo) Distinct(cols ...field.Expr) ITPengajuanTransferDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tPengajuanTransferDo) Omit(cols ...field.Expr) ITPengajuanTransferDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tPengajuanTransferDo) Join(table schema.Tabler, on ...field.Expr) ITPengajuanTransferDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tPengajuanTransferDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITPengajuanTransferDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tPengajuanTransferDo) RightJoin(table schema.Tabler, on ...field.Expr) ITPengajuanTransferDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tPengajuanTransferDo) Group(cols ...field.Expr) ITPengajuanTransferDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tPengajuanTransferDo) Having(conds ...gen.Condition) ITPengajuanTransferDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tPengajuanTransferDo) Limit(limit int) ITPengajuanTransferDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tPengajuanTransferDo) Offset(offset int) ITPengajuanTransferDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tPengajuanTransferDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITPengajuanTransferDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tPengajuanTransferDo) Unscoped() ITPengajuanTransferDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tPengajuanTransferDo) Create(values ...*model.TPengajuanTransfer) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tPengajuanTransferDo) CreateInBatches(values []*model.TPengajuanTransfer, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tPengajuanTransferDo) Save(values ...*model.TPengajuanTransfer) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tPengajuanTransferDo) First() (*model.TPengajuanTransfer, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPengajuanTransfer), nil
	}
}

func (t tPengajuanTransferDo) Take() (*model.TPengajuanTransfer, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPengajuanTransfer), nil
	}
}

func (t tPengajuanTransferDo) Last() (*model.TPengajuanTransfer, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPengajuanTransfer), nil
	}
}

func (t tPengajuanTransferDo) Find() ([]*model.TPengajuanTransfer, error) {
	result, err := t.DO.Find()
	return result.([]*model.TPengajuanTransfer), err
}

func (t tPengajuanTransferDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TPengajuanTransfer, err error) {
	buf := make([]*model.TPengajuanTransfer, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tPengajuanTransferDo) FindInBatches(result *[]*model.TPengajuanTransfer, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tPengajuanTransferDo) Attrs(attrs ...field.AssignExpr) ITPengajuanTransferDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tPengajuanTransferDo) Assign(attrs ...field.AssignExpr) ITPengajuanTransferDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tPengajuanTransferDo) Joins(fields ...field.RelationField) ITPengajuanTransferDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tPengajuanTransferDo) Preload(fields ...field.RelationField) ITPengajuanTransferDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tPengajuanTransferDo) FirstOrInit() (*model.TPengajuanTransfer, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPengajuanTransfer), nil
	}
}

func (t tPengajuanTransferDo) FirstOrCreate() (*model.TPengajuanTransfer, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPengajuanTransfer), nil
	}
}

func (t tPengajuanTransferDo) FindByPage(offset int, limit int) (result []*model.TPengajuanTransfer, count int64, err error) {
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

func (t tPengajuanTransferDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tPengajuanTransferDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tPengajuanTransferDo) Delete(models ...*model.TPengajuanTransfer) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tPengajuanTransferDo) withDO(do gen.Dao) *tPengajuanTransferDo {
	t.DO = *do.(*gen.DO)
	return t
}
