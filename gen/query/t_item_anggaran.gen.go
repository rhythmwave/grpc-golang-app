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

func newTItemAnggaran(db *gorm.DB, opts ...gen.DOOption) tItemAnggaran {
	_tItemAnggaran := tItemAnggaran{}

	_tItemAnggaran.tItemAnggaranDo.UseDB(db, opts...)
	_tItemAnggaran.tItemAnggaranDo.UseModel(&model.TItemAnggaran{})

	tableName := _tItemAnggaran.tItemAnggaranDo.TableName()
	_tItemAnggaran.ALL = field.NewAsterisk(tableName)
	_tItemAnggaran.CKodeItemAnggaran = field.NewString(tableName, "c_kode_item_anggaran")
	_tItemAnggaran.CNamaItemAnggaran = field.NewString(tableName, "c_nama_item_anggaran")
	_tItemAnggaran.CIsDokumen = field.NewInt16(tableName, "c_is_dokumen")
	_tItemAnggaran.CIDKelompokPajak = field.NewString(tableName, "c_id_kelompok_pajak")
	_tItemAnggaran.CKodeCOA = field.NewString(tableName, "c_kode_c_o_a")
	_tItemAnggaran.CKeterangan = field.NewString(tableName, "c_keterangan")
	_tItemAnggaran.CStatus = field.NewString(tableName, "c_status")
	_tItemAnggaran.CUpdater = field.NewString(tableName, "c_updater")
	_tItemAnggaran.CLastUpdate = field.NewTime(tableName, "c_last_update")

	_tItemAnggaran.fillFieldMap()

	return _tItemAnggaran
}

type tItemAnggaran struct {
	tItemAnggaranDo

	ALL               field.Asterisk
	CKodeItemAnggaran field.String
	CNamaItemAnggaran field.String
	CIsDokumen        field.Int16
	CIDKelompokPajak  field.String
	CKodeCOA          field.String
	CKeterangan       field.String
	CStatus           field.String
	CUpdater          field.String
	CLastUpdate       field.Time

	fieldMap map[string]field.Expr
}

func (t tItemAnggaran) Table(newTableName string) *tItemAnggaran {
	t.tItemAnggaranDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tItemAnggaran) As(alias string) *tItemAnggaran {
	t.tItemAnggaranDo.DO = *(t.tItemAnggaranDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tItemAnggaran) updateTableName(table string) *tItemAnggaran {
	t.ALL = field.NewAsterisk(table)
	t.CKodeItemAnggaran = field.NewString(table, "c_kode_item_anggaran")
	t.CNamaItemAnggaran = field.NewString(table, "c_nama_item_anggaran")
	t.CIsDokumen = field.NewInt16(table, "c_is_dokumen")
	t.CIDKelompokPajak = field.NewString(table, "c_id_kelompok_pajak")
	t.CKodeCOA = field.NewString(table, "c_kode_c_o_a")
	t.CKeterangan = field.NewString(table, "c_keterangan")
	t.CStatus = field.NewString(table, "c_status")
	t.CUpdater = field.NewString(table, "c_updater")
	t.CLastUpdate = field.NewTime(table, "c_last_update")

	t.fillFieldMap()

	return t
}

func (t *tItemAnggaran) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tItemAnggaran) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 9)
	t.fieldMap["c_kode_item_anggaran"] = t.CKodeItemAnggaran
	t.fieldMap["c_nama_item_anggaran"] = t.CNamaItemAnggaran
	t.fieldMap["c_is_dokumen"] = t.CIsDokumen
	t.fieldMap["c_id_kelompok_pajak"] = t.CIDKelompokPajak
	t.fieldMap["c_kode_c_o_a"] = t.CKodeCOA
	t.fieldMap["c_keterangan"] = t.CKeterangan
	t.fieldMap["c_status"] = t.CStatus
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_last_update"] = t.CLastUpdate
}

func (t tItemAnggaran) clone(db *gorm.DB) tItemAnggaran {
	t.tItemAnggaranDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tItemAnggaran) replaceDB(db *gorm.DB) tItemAnggaran {
	t.tItemAnggaranDo.ReplaceDB(db)
	return t
}

type tItemAnggaranDo struct{ gen.DO }

type ITItemAnggaranDo interface {
	gen.SubQuery
	Debug() ITItemAnggaranDo
	WithContext(ctx context.Context) ITItemAnggaranDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITItemAnggaranDo
	WriteDB() ITItemAnggaranDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITItemAnggaranDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITItemAnggaranDo
	Not(conds ...gen.Condition) ITItemAnggaranDo
	Or(conds ...gen.Condition) ITItemAnggaranDo
	Select(conds ...field.Expr) ITItemAnggaranDo
	Where(conds ...gen.Condition) ITItemAnggaranDo
	Order(conds ...field.Expr) ITItemAnggaranDo
	Distinct(cols ...field.Expr) ITItemAnggaranDo
	Omit(cols ...field.Expr) ITItemAnggaranDo
	Join(table schema.Tabler, on ...field.Expr) ITItemAnggaranDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITItemAnggaranDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITItemAnggaranDo
	Group(cols ...field.Expr) ITItemAnggaranDo
	Having(conds ...gen.Condition) ITItemAnggaranDo
	Limit(limit int) ITItemAnggaranDo
	Offset(offset int) ITItemAnggaranDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITItemAnggaranDo
	Unscoped() ITItemAnggaranDo
	Create(values ...*model.TItemAnggaran) error
	CreateInBatches(values []*model.TItemAnggaran, batchSize int) error
	Save(values ...*model.TItemAnggaran) error
	First() (*model.TItemAnggaran, error)
	Take() (*model.TItemAnggaran, error)
	Last() (*model.TItemAnggaran, error)
	Find() ([]*model.TItemAnggaran, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TItemAnggaran, err error)
	FindInBatches(result *[]*model.TItemAnggaran, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TItemAnggaran) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITItemAnggaranDo
	Assign(attrs ...field.AssignExpr) ITItemAnggaranDo
	Joins(fields ...field.RelationField) ITItemAnggaranDo
	Preload(fields ...field.RelationField) ITItemAnggaranDo
	FirstOrInit() (*model.TItemAnggaran, error)
	FirstOrCreate() (*model.TItemAnggaran, error)
	FindByPage(offset int, limit int) (result []*model.TItemAnggaran, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITItemAnggaranDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tItemAnggaranDo) Debug() ITItemAnggaranDo {
	return t.withDO(t.DO.Debug())
}

func (t tItemAnggaranDo) WithContext(ctx context.Context) ITItemAnggaranDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tItemAnggaranDo) ReadDB() ITItemAnggaranDo {
	return t.Clauses(dbresolver.Read)
}

func (t tItemAnggaranDo) WriteDB() ITItemAnggaranDo {
	return t.Clauses(dbresolver.Write)
}

func (t tItemAnggaranDo) Session(config *gorm.Session) ITItemAnggaranDo {
	return t.withDO(t.DO.Session(config))
}

func (t tItemAnggaranDo) Clauses(conds ...clause.Expression) ITItemAnggaranDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tItemAnggaranDo) Returning(value interface{}, columns ...string) ITItemAnggaranDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tItemAnggaranDo) Not(conds ...gen.Condition) ITItemAnggaranDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tItemAnggaranDo) Or(conds ...gen.Condition) ITItemAnggaranDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tItemAnggaranDo) Select(conds ...field.Expr) ITItemAnggaranDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tItemAnggaranDo) Where(conds ...gen.Condition) ITItemAnggaranDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tItemAnggaranDo) Order(conds ...field.Expr) ITItemAnggaranDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tItemAnggaranDo) Distinct(cols ...field.Expr) ITItemAnggaranDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tItemAnggaranDo) Omit(cols ...field.Expr) ITItemAnggaranDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tItemAnggaranDo) Join(table schema.Tabler, on ...field.Expr) ITItemAnggaranDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tItemAnggaranDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITItemAnggaranDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tItemAnggaranDo) RightJoin(table schema.Tabler, on ...field.Expr) ITItemAnggaranDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tItemAnggaranDo) Group(cols ...field.Expr) ITItemAnggaranDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tItemAnggaranDo) Having(conds ...gen.Condition) ITItemAnggaranDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tItemAnggaranDo) Limit(limit int) ITItemAnggaranDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tItemAnggaranDo) Offset(offset int) ITItemAnggaranDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tItemAnggaranDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITItemAnggaranDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tItemAnggaranDo) Unscoped() ITItemAnggaranDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tItemAnggaranDo) Create(values ...*model.TItemAnggaran) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tItemAnggaranDo) CreateInBatches(values []*model.TItemAnggaran, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tItemAnggaranDo) Save(values ...*model.TItemAnggaran) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tItemAnggaranDo) First() (*model.TItemAnggaran, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TItemAnggaran), nil
	}
}

func (t tItemAnggaranDo) Take() (*model.TItemAnggaran, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TItemAnggaran), nil
	}
}

func (t tItemAnggaranDo) Last() (*model.TItemAnggaran, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TItemAnggaran), nil
	}
}

func (t tItemAnggaranDo) Find() ([]*model.TItemAnggaran, error) {
	result, err := t.DO.Find()
	return result.([]*model.TItemAnggaran), err
}

func (t tItemAnggaranDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TItemAnggaran, err error) {
	buf := make([]*model.TItemAnggaran, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tItemAnggaranDo) FindInBatches(result *[]*model.TItemAnggaran, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tItemAnggaranDo) Attrs(attrs ...field.AssignExpr) ITItemAnggaranDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tItemAnggaranDo) Assign(attrs ...field.AssignExpr) ITItemAnggaranDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tItemAnggaranDo) Joins(fields ...field.RelationField) ITItemAnggaranDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tItemAnggaranDo) Preload(fields ...field.RelationField) ITItemAnggaranDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tItemAnggaranDo) FirstOrInit() (*model.TItemAnggaran, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TItemAnggaran), nil
	}
}

func (t tItemAnggaranDo) FirstOrCreate() (*model.TItemAnggaran, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TItemAnggaran), nil
	}
}

func (t tItemAnggaranDo) FindByPage(offset int, limit int) (result []*model.TItemAnggaran, count int64, err error) {
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

func (t tItemAnggaranDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tItemAnggaranDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tItemAnggaranDo) Delete(models ...*model.TItemAnggaran) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tItemAnggaranDo) withDO(do gen.Dao) *tItemAnggaranDo {
	t.DO = *do.(*gen.DO)
	return t
}