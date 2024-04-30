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

func newTKegiatan(db *gorm.DB, opts ...gen.DOOption) tKegiatan {
	_tKegiatan := tKegiatan{}

	_tKegiatan.tKegiatanDo.UseDB(db, opts...)
	_tKegiatan.tKegiatanDo.UseModel(&model.TKegiatan{})

	tableName := _tKegiatan.tKegiatanDo.TableName()
	_tKegiatan.ALL = field.NewAsterisk(tableName)
	_tKegiatan.CIDKegiatan = field.NewInt32(tableName, "c_id_kegiatan")
	_tKegiatan.CUpline = field.NewInt32(tableName, "c_upline")
	_tKegiatan.CIDAsal = field.NewInt32(tableName, "c_id_asal")
	_tKegiatan.CIDGedung = field.NewInt32(tableName, "c_id_gedung")
	_tKegiatan.CNamaKegiatan = field.NewString(tableName, "c_nama_kegiatan")
	_tKegiatan.CIDSekolah = field.NewInt32(tableName, "c_id_sekolah")
	_tKegiatan.CIDTingkat = field.NewInt32(tableName, "c_id_tingkat")
	_tKegiatan.CIDTujuanTingkat = field.NewInt32(tableName, "c_id_tujuan_tingkat")
	_tKegiatan.CPengelompokan = field.NewString(tableName, "c_pengelompokan")
	_tKegiatan.CTanggalAwal = field.NewTime(tableName, "c_tanggal_awal")
	_tKegiatan.CTanggalAkhir = field.NewTime(tableName, "c_tanggal_akhir")
	_tKegiatan.CPetugas = field.NewString(tableName, "c_petugas")
	_tKegiatan.CStatus = field.NewString(tableName, "c_status")
	_tKegiatan.CUpdater = field.NewString(tableName, "c_updater")
	_tKegiatan.CLastUpdate = field.NewTime(tableName, "c_last_update")
	_tKegiatan.CTahunAjaran = field.NewString(tableName, "c_tahun_ajaran")
	_tKegiatan.CIzinPrinsip = field.NewInt16(tableName, "c_izin_prinsip")

	_tKegiatan.fillFieldMap()

	return _tKegiatan
}

type tKegiatan struct {
	tKegiatanDo

	ALL              field.Asterisk
	CIDKegiatan      field.Int32
	CUpline          field.Int32
	CIDAsal          field.Int32
	CIDGedung        field.Int32
	CNamaKegiatan    field.String
	CIDSekolah       field.Int32
	CIDTingkat       field.Int32
	CIDTujuanTingkat field.Int32
	CPengelompokan   field.String
	CTanggalAwal     field.Time
	CTanggalAkhir    field.Time
	CPetugas         field.String
	CStatus          field.String
	CUpdater         field.String
	CLastUpdate      field.Time
	CTahunAjaran     field.String
	CIzinPrinsip     field.Int16

	fieldMap map[string]field.Expr
}

func (t tKegiatan) Table(newTableName string) *tKegiatan {
	t.tKegiatanDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tKegiatan) As(alias string) *tKegiatan {
	t.tKegiatanDo.DO = *(t.tKegiatanDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tKegiatan) updateTableName(table string) *tKegiatan {
	t.ALL = field.NewAsterisk(table)
	t.CIDKegiatan = field.NewInt32(table, "c_id_kegiatan")
	t.CUpline = field.NewInt32(table, "c_upline")
	t.CIDAsal = field.NewInt32(table, "c_id_asal")
	t.CIDGedung = field.NewInt32(table, "c_id_gedung")
	t.CNamaKegiatan = field.NewString(table, "c_nama_kegiatan")
	t.CIDSekolah = field.NewInt32(table, "c_id_sekolah")
	t.CIDTingkat = field.NewInt32(table, "c_id_tingkat")
	t.CIDTujuanTingkat = field.NewInt32(table, "c_id_tujuan_tingkat")
	t.CPengelompokan = field.NewString(table, "c_pengelompokan")
	t.CTanggalAwal = field.NewTime(table, "c_tanggal_awal")
	t.CTanggalAkhir = field.NewTime(table, "c_tanggal_akhir")
	t.CPetugas = field.NewString(table, "c_petugas")
	t.CStatus = field.NewString(table, "c_status")
	t.CUpdater = field.NewString(table, "c_updater")
	t.CLastUpdate = field.NewTime(table, "c_last_update")
	t.CTahunAjaran = field.NewString(table, "c_tahun_ajaran")
	t.CIzinPrinsip = field.NewInt16(table, "c_izin_prinsip")

	t.fillFieldMap()

	return t
}

func (t *tKegiatan) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tKegiatan) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 17)
	t.fieldMap["c_id_kegiatan"] = t.CIDKegiatan
	t.fieldMap["c_upline"] = t.CUpline
	t.fieldMap["c_id_asal"] = t.CIDAsal
	t.fieldMap["c_id_gedung"] = t.CIDGedung
	t.fieldMap["c_nama_kegiatan"] = t.CNamaKegiatan
	t.fieldMap["c_id_sekolah"] = t.CIDSekolah
	t.fieldMap["c_id_tingkat"] = t.CIDTingkat
	t.fieldMap["c_id_tujuan_tingkat"] = t.CIDTujuanTingkat
	t.fieldMap["c_pengelompokan"] = t.CPengelompokan
	t.fieldMap["c_tanggal_awal"] = t.CTanggalAwal
	t.fieldMap["c_tanggal_akhir"] = t.CTanggalAkhir
	t.fieldMap["c_petugas"] = t.CPetugas
	t.fieldMap["c_status"] = t.CStatus
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_last_update"] = t.CLastUpdate
	t.fieldMap["c_tahun_ajaran"] = t.CTahunAjaran
	t.fieldMap["c_izin_prinsip"] = t.CIzinPrinsip
}

func (t tKegiatan) clone(db *gorm.DB) tKegiatan {
	t.tKegiatanDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tKegiatan) replaceDB(db *gorm.DB) tKegiatan {
	t.tKegiatanDo.ReplaceDB(db)
	return t
}

type tKegiatanDo struct{ gen.DO }

type ITKegiatanDo interface {
	gen.SubQuery
	Debug() ITKegiatanDo
	WithContext(ctx context.Context) ITKegiatanDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITKegiatanDo
	WriteDB() ITKegiatanDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITKegiatanDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITKegiatanDo
	Not(conds ...gen.Condition) ITKegiatanDo
	Or(conds ...gen.Condition) ITKegiatanDo
	Select(conds ...field.Expr) ITKegiatanDo
	Where(conds ...gen.Condition) ITKegiatanDo
	Order(conds ...field.Expr) ITKegiatanDo
	Distinct(cols ...field.Expr) ITKegiatanDo
	Omit(cols ...field.Expr) ITKegiatanDo
	Join(table schema.Tabler, on ...field.Expr) ITKegiatanDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITKegiatanDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITKegiatanDo
	Group(cols ...field.Expr) ITKegiatanDo
	Having(conds ...gen.Condition) ITKegiatanDo
	Limit(limit int) ITKegiatanDo
	Offset(offset int) ITKegiatanDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITKegiatanDo
	Unscoped() ITKegiatanDo
	Create(values ...*model.TKegiatan) error
	CreateInBatches(values []*model.TKegiatan, batchSize int) error
	Save(values ...*model.TKegiatan) error
	First() (*model.TKegiatan, error)
	Take() (*model.TKegiatan, error)
	Last() (*model.TKegiatan, error)
	Find() ([]*model.TKegiatan, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TKegiatan, err error)
	FindInBatches(result *[]*model.TKegiatan, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TKegiatan) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITKegiatanDo
	Assign(attrs ...field.AssignExpr) ITKegiatanDo
	Joins(fields ...field.RelationField) ITKegiatanDo
	Preload(fields ...field.RelationField) ITKegiatanDo
	FirstOrInit() (*model.TKegiatan, error)
	FirstOrCreate() (*model.TKegiatan, error)
	FindByPage(offset int, limit int) (result []*model.TKegiatan, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITKegiatanDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tKegiatanDo) Debug() ITKegiatanDo {
	return t.withDO(t.DO.Debug())
}

func (t tKegiatanDo) WithContext(ctx context.Context) ITKegiatanDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tKegiatanDo) ReadDB() ITKegiatanDo {
	return t.Clauses(dbresolver.Read)
}

func (t tKegiatanDo) WriteDB() ITKegiatanDo {
	return t.Clauses(dbresolver.Write)
}

func (t tKegiatanDo) Session(config *gorm.Session) ITKegiatanDo {
	return t.withDO(t.DO.Session(config))
}

func (t tKegiatanDo) Clauses(conds ...clause.Expression) ITKegiatanDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tKegiatanDo) Returning(value interface{}, columns ...string) ITKegiatanDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tKegiatanDo) Not(conds ...gen.Condition) ITKegiatanDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tKegiatanDo) Or(conds ...gen.Condition) ITKegiatanDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tKegiatanDo) Select(conds ...field.Expr) ITKegiatanDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tKegiatanDo) Where(conds ...gen.Condition) ITKegiatanDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tKegiatanDo) Order(conds ...field.Expr) ITKegiatanDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tKegiatanDo) Distinct(cols ...field.Expr) ITKegiatanDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tKegiatanDo) Omit(cols ...field.Expr) ITKegiatanDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tKegiatanDo) Join(table schema.Tabler, on ...field.Expr) ITKegiatanDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tKegiatanDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITKegiatanDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tKegiatanDo) RightJoin(table schema.Tabler, on ...field.Expr) ITKegiatanDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tKegiatanDo) Group(cols ...field.Expr) ITKegiatanDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tKegiatanDo) Having(conds ...gen.Condition) ITKegiatanDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tKegiatanDo) Limit(limit int) ITKegiatanDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tKegiatanDo) Offset(offset int) ITKegiatanDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tKegiatanDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITKegiatanDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tKegiatanDo) Unscoped() ITKegiatanDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tKegiatanDo) Create(values ...*model.TKegiatan) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tKegiatanDo) CreateInBatches(values []*model.TKegiatan, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tKegiatanDo) Save(values ...*model.TKegiatan) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tKegiatanDo) First() (*model.TKegiatan, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKegiatan), nil
	}
}

func (t tKegiatanDo) Take() (*model.TKegiatan, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKegiatan), nil
	}
}

func (t tKegiatanDo) Last() (*model.TKegiatan, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKegiatan), nil
	}
}

func (t tKegiatanDo) Find() ([]*model.TKegiatan, error) {
	result, err := t.DO.Find()
	return result.([]*model.TKegiatan), err
}

func (t tKegiatanDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TKegiatan, err error) {
	buf := make([]*model.TKegiatan, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tKegiatanDo) FindInBatches(result *[]*model.TKegiatan, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tKegiatanDo) Attrs(attrs ...field.AssignExpr) ITKegiatanDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tKegiatanDo) Assign(attrs ...field.AssignExpr) ITKegiatanDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tKegiatanDo) Joins(fields ...field.RelationField) ITKegiatanDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tKegiatanDo) Preload(fields ...field.RelationField) ITKegiatanDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tKegiatanDo) FirstOrInit() (*model.TKegiatan, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKegiatan), nil
	}
}

func (t tKegiatanDo) FirstOrCreate() (*model.TKegiatan, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKegiatan), nil
	}
}

func (t tKegiatanDo) FindByPage(offset int, limit int) (result []*model.TKegiatan, count int64, err error) {
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

func (t tKegiatanDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tKegiatanDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tKegiatanDo) Delete(models ...*model.TKegiatan) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tKegiatanDo) withDO(do gen.Dao) *tKegiatanDo {
	t.DO = *do.(*gen.DO)
	return t
}
