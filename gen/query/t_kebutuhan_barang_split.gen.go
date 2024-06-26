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

func newTKebutuhanBarangSplit(db *gorm.DB, opts ...gen.DOOption) tKebutuhanBarangSplit {
	_tKebutuhanBarangSplit := tKebutuhanBarangSplit{}

	_tKebutuhanBarangSplit.tKebutuhanBarangSplitDo.UseDB(db, opts...)
	_tKebutuhanBarangSplit.tKebutuhanBarangSplitDo.UseModel(&model.TKebutuhanBarangSplit{})

	tableName := _tKebutuhanBarangSplit.tKebutuhanBarangSplitDo.TableName()
	_tKebutuhanBarangSplit.ALL = field.NewAsterisk(tableName)
	_tKebutuhanBarangSplit.CKodeSplitBarang = field.NewString(tableName, "c_kode_split_barang")
	_tKebutuhanBarangSplit.CKodeKegiatan = field.NewString(tableName, "c_kode_kegiatan")
	_tKebutuhanBarangSplit.CKodeBarang = field.NewString(tableName, "c_kode_barang")
	_tKebutuhanBarangSplit.CTanggalPakai = field.NewTime(tableName, "c_tanggal_pakai")
	_tKebutuhanBarangSplit.CTanggalPakaiAcc = field.NewTime(tableName, "c_tanggal_pakai_acc")
	_tKebutuhanBarangSplit.CPengajuan = field.NewFloat64(tableName, "c_pengajuan")
	_tKebutuhanBarangSplit.CSetujuKacab = field.NewFloat64(tableName, "c_setuju_kacab")
	_tKebutuhanBarangSplit.CSetujuBidang = field.NewFloat64(tableName, "c_setuju_bidang")
	_tKebutuhanBarangSplit.CSetujuLog = field.NewFloat64(tableName, "c_setuju_log")
	_tKebutuhanBarangSplit.CSetujuAkuntansi = field.NewFloat64(tableName, "c_setuju_akuntansi")
	_tKebutuhanBarangSplit.CStatus = field.NewString(tableName, "c_status")
	_tKebutuhanBarangSplit.CDokumen = field.NewString(tableName, "c_dokumen")
	_tKebutuhanBarangSplit.CIDVendor = field.NewInt32(tableName, "c_id_vendor")
	_tKebutuhanBarangSplit.CIsPusat = field.NewInt16(tableName, "c_is_pusat")
	_tKebutuhanBarangSplit.CFlagSpta = field.NewInt16(tableName, "c_flag_spta")
	_tKebutuhanBarangSplit.CFlagDrt = field.NewInt16(tableName, "c_flag_drt")
	_tKebutuhanBarangSplit.CUpdater = field.NewString(tableName, "c_updater")
	_tKebutuhanBarangSplit.CLastUpdate = field.NewTime(tableName, "c_last_update")

	_tKebutuhanBarangSplit.fillFieldMap()

	return _tKebutuhanBarangSplit
}

type tKebutuhanBarangSplit struct {
	tKebutuhanBarangSplitDo

	ALL              field.Asterisk
	CKodeSplitBarang field.String
	CKodeKegiatan    field.String
	CKodeBarang      field.String
	CTanggalPakai    field.Time
	CTanggalPakaiAcc field.Time
	CPengajuan       field.Float64
	CSetujuKacab     field.Float64
	CSetujuBidang    field.Float64
	CSetujuLog       field.Float64
	CSetujuAkuntansi field.Float64
	CStatus          field.String
	CDokumen         field.String
	CIDVendor        field.Int32
	CIsPusat         field.Int16
	CFlagSpta        field.Int16
	CFlagDrt         field.Int16
	CUpdater         field.String
	CLastUpdate      field.Time

	fieldMap map[string]field.Expr
}

func (t tKebutuhanBarangSplit) Table(newTableName string) *tKebutuhanBarangSplit {
	t.tKebutuhanBarangSplitDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tKebutuhanBarangSplit) As(alias string) *tKebutuhanBarangSplit {
	t.tKebutuhanBarangSplitDo.DO = *(t.tKebutuhanBarangSplitDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tKebutuhanBarangSplit) updateTableName(table string) *tKebutuhanBarangSplit {
	t.ALL = field.NewAsterisk(table)
	t.CKodeSplitBarang = field.NewString(table, "c_kode_split_barang")
	t.CKodeKegiatan = field.NewString(table, "c_kode_kegiatan")
	t.CKodeBarang = field.NewString(table, "c_kode_barang")
	t.CTanggalPakai = field.NewTime(table, "c_tanggal_pakai")
	t.CTanggalPakaiAcc = field.NewTime(table, "c_tanggal_pakai_acc")
	t.CPengajuan = field.NewFloat64(table, "c_pengajuan")
	t.CSetujuKacab = field.NewFloat64(table, "c_setuju_kacab")
	t.CSetujuBidang = field.NewFloat64(table, "c_setuju_bidang")
	t.CSetujuLog = field.NewFloat64(table, "c_setuju_log")
	t.CSetujuAkuntansi = field.NewFloat64(table, "c_setuju_akuntansi")
	t.CStatus = field.NewString(table, "c_status")
	t.CDokumen = field.NewString(table, "c_dokumen")
	t.CIDVendor = field.NewInt32(table, "c_id_vendor")
	t.CIsPusat = field.NewInt16(table, "c_is_pusat")
	t.CFlagSpta = field.NewInt16(table, "c_flag_spta")
	t.CFlagDrt = field.NewInt16(table, "c_flag_drt")
	t.CUpdater = field.NewString(table, "c_updater")
	t.CLastUpdate = field.NewTime(table, "c_last_update")

	t.fillFieldMap()

	return t
}

func (t *tKebutuhanBarangSplit) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tKebutuhanBarangSplit) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 18)
	t.fieldMap["c_kode_split_barang"] = t.CKodeSplitBarang
	t.fieldMap["c_kode_kegiatan"] = t.CKodeKegiatan
	t.fieldMap["c_kode_barang"] = t.CKodeBarang
	t.fieldMap["c_tanggal_pakai"] = t.CTanggalPakai
	t.fieldMap["c_tanggal_pakai_acc"] = t.CTanggalPakaiAcc
	t.fieldMap["c_pengajuan"] = t.CPengajuan
	t.fieldMap["c_setuju_kacab"] = t.CSetujuKacab
	t.fieldMap["c_setuju_bidang"] = t.CSetujuBidang
	t.fieldMap["c_setuju_log"] = t.CSetujuLog
	t.fieldMap["c_setuju_akuntansi"] = t.CSetujuAkuntansi
	t.fieldMap["c_status"] = t.CStatus
	t.fieldMap["c_dokumen"] = t.CDokumen
	t.fieldMap["c_id_vendor"] = t.CIDVendor
	t.fieldMap["c_is_pusat"] = t.CIsPusat
	t.fieldMap["c_flag_spta"] = t.CFlagSpta
	t.fieldMap["c_flag_drt"] = t.CFlagDrt
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_last_update"] = t.CLastUpdate
}

func (t tKebutuhanBarangSplit) clone(db *gorm.DB) tKebutuhanBarangSplit {
	t.tKebutuhanBarangSplitDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tKebutuhanBarangSplit) replaceDB(db *gorm.DB) tKebutuhanBarangSplit {
	t.tKebutuhanBarangSplitDo.ReplaceDB(db)
	return t
}

type tKebutuhanBarangSplitDo struct{ gen.DO }

type ITKebutuhanBarangSplitDo interface {
	gen.SubQuery
	Debug() ITKebutuhanBarangSplitDo
	WithContext(ctx context.Context) ITKebutuhanBarangSplitDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITKebutuhanBarangSplitDo
	WriteDB() ITKebutuhanBarangSplitDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITKebutuhanBarangSplitDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITKebutuhanBarangSplitDo
	Not(conds ...gen.Condition) ITKebutuhanBarangSplitDo
	Or(conds ...gen.Condition) ITKebutuhanBarangSplitDo
	Select(conds ...field.Expr) ITKebutuhanBarangSplitDo
	Where(conds ...gen.Condition) ITKebutuhanBarangSplitDo
	Order(conds ...field.Expr) ITKebutuhanBarangSplitDo
	Distinct(cols ...field.Expr) ITKebutuhanBarangSplitDo
	Omit(cols ...field.Expr) ITKebutuhanBarangSplitDo
	Join(table schema.Tabler, on ...field.Expr) ITKebutuhanBarangSplitDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITKebutuhanBarangSplitDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITKebutuhanBarangSplitDo
	Group(cols ...field.Expr) ITKebutuhanBarangSplitDo
	Having(conds ...gen.Condition) ITKebutuhanBarangSplitDo
	Limit(limit int) ITKebutuhanBarangSplitDo
	Offset(offset int) ITKebutuhanBarangSplitDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITKebutuhanBarangSplitDo
	Unscoped() ITKebutuhanBarangSplitDo
	Create(values ...*model.TKebutuhanBarangSplit) error
	CreateInBatches(values []*model.TKebutuhanBarangSplit, batchSize int) error
	Save(values ...*model.TKebutuhanBarangSplit) error
	First() (*model.TKebutuhanBarangSplit, error)
	Take() (*model.TKebutuhanBarangSplit, error)
	Last() (*model.TKebutuhanBarangSplit, error)
	Find() ([]*model.TKebutuhanBarangSplit, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TKebutuhanBarangSplit, err error)
	FindInBatches(result *[]*model.TKebutuhanBarangSplit, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TKebutuhanBarangSplit) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITKebutuhanBarangSplitDo
	Assign(attrs ...field.AssignExpr) ITKebutuhanBarangSplitDo
	Joins(fields ...field.RelationField) ITKebutuhanBarangSplitDo
	Preload(fields ...field.RelationField) ITKebutuhanBarangSplitDo
	FirstOrInit() (*model.TKebutuhanBarangSplit, error)
	FirstOrCreate() (*model.TKebutuhanBarangSplit, error)
	FindByPage(offset int, limit int) (result []*model.TKebutuhanBarangSplit, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITKebutuhanBarangSplitDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tKebutuhanBarangSplitDo) Debug() ITKebutuhanBarangSplitDo {
	return t.withDO(t.DO.Debug())
}

func (t tKebutuhanBarangSplitDo) WithContext(ctx context.Context) ITKebutuhanBarangSplitDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tKebutuhanBarangSplitDo) ReadDB() ITKebutuhanBarangSplitDo {
	return t.Clauses(dbresolver.Read)
}

func (t tKebutuhanBarangSplitDo) WriteDB() ITKebutuhanBarangSplitDo {
	return t.Clauses(dbresolver.Write)
}

func (t tKebutuhanBarangSplitDo) Session(config *gorm.Session) ITKebutuhanBarangSplitDo {
	return t.withDO(t.DO.Session(config))
}

func (t tKebutuhanBarangSplitDo) Clauses(conds ...clause.Expression) ITKebutuhanBarangSplitDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tKebutuhanBarangSplitDo) Returning(value interface{}, columns ...string) ITKebutuhanBarangSplitDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tKebutuhanBarangSplitDo) Not(conds ...gen.Condition) ITKebutuhanBarangSplitDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tKebutuhanBarangSplitDo) Or(conds ...gen.Condition) ITKebutuhanBarangSplitDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tKebutuhanBarangSplitDo) Select(conds ...field.Expr) ITKebutuhanBarangSplitDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tKebutuhanBarangSplitDo) Where(conds ...gen.Condition) ITKebutuhanBarangSplitDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tKebutuhanBarangSplitDo) Order(conds ...field.Expr) ITKebutuhanBarangSplitDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tKebutuhanBarangSplitDo) Distinct(cols ...field.Expr) ITKebutuhanBarangSplitDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tKebutuhanBarangSplitDo) Omit(cols ...field.Expr) ITKebutuhanBarangSplitDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tKebutuhanBarangSplitDo) Join(table schema.Tabler, on ...field.Expr) ITKebutuhanBarangSplitDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tKebutuhanBarangSplitDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITKebutuhanBarangSplitDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tKebutuhanBarangSplitDo) RightJoin(table schema.Tabler, on ...field.Expr) ITKebutuhanBarangSplitDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tKebutuhanBarangSplitDo) Group(cols ...field.Expr) ITKebutuhanBarangSplitDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tKebutuhanBarangSplitDo) Having(conds ...gen.Condition) ITKebutuhanBarangSplitDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tKebutuhanBarangSplitDo) Limit(limit int) ITKebutuhanBarangSplitDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tKebutuhanBarangSplitDo) Offset(offset int) ITKebutuhanBarangSplitDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tKebutuhanBarangSplitDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITKebutuhanBarangSplitDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tKebutuhanBarangSplitDo) Unscoped() ITKebutuhanBarangSplitDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tKebutuhanBarangSplitDo) Create(values ...*model.TKebutuhanBarangSplit) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tKebutuhanBarangSplitDo) CreateInBatches(values []*model.TKebutuhanBarangSplit, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tKebutuhanBarangSplitDo) Save(values ...*model.TKebutuhanBarangSplit) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tKebutuhanBarangSplitDo) First() (*model.TKebutuhanBarangSplit, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKebutuhanBarangSplit), nil
	}
}

func (t tKebutuhanBarangSplitDo) Take() (*model.TKebutuhanBarangSplit, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKebutuhanBarangSplit), nil
	}
}

func (t tKebutuhanBarangSplitDo) Last() (*model.TKebutuhanBarangSplit, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKebutuhanBarangSplit), nil
	}
}

func (t tKebutuhanBarangSplitDo) Find() ([]*model.TKebutuhanBarangSplit, error) {
	result, err := t.DO.Find()
	return result.([]*model.TKebutuhanBarangSplit), err
}

func (t tKebutuhanBarangSplitDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TKebutuhanBarangSplit, err error) {
	buf := make([]*model.TKebutuhanBarangSplit, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tKebutuhanBarangSplitDo) FindInBatches(result *[]*model.TKebutuhanBarangSplit, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tKebutuhanBarangSplitDo) Attrs(attrs ...field.AssignExpr) ITKebutuhanBarangSplitDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tKebutuhanBarangSplitDo) Assign(attrs ...field.AssignExpr) ITKebutuhanBarangSplitDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tKebutuhanBarangSplitDo) Joins(fields ...field.RelationField) ITKebutuhanBarangSplitDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tKebutuhanBarangSplitDo) Preload(fields ...field.RelationField) ITKebutuhanBarangSplitDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tKebutuhanBarangSplitDo) FirstOrInit() (*model.TKebutuhanBarangSplit, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKebutuhanBarangSplit), nil
	}
}

func (t tKebutuhanBarangSplitDo) FirstOrCreate() (*model.TKebutuhanBarangSplit, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKebutuhanBarangSplit), nil
	}
}

func (t tKebutuhanBarangSplitDo) FindByPage(offset int, limit int) (result []*model.TKebutuhanBarangSplit, count int64, err error) {
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

func (t tKebutuhanBarangSplitDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tKebutuhanBarangSplitDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tKebutuhanBarangSplitDo) Delete(models ...*model.TKebutuhanBarangSplit) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tKebutuhanBarangSplitDo) withDO(do gen.Dao) *tKebutuhanBarangSplitDo {
	t.DO = *do.(*gen.DO)
	return t
}
