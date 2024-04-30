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

func newTKebutuhanKerta(db *gorm.DB, opts ...gen.DOOption) tKebutuhanKerta {
	_tKebutuhanKerta := tKebutuhanKerta{}

	_tKebutuhanKerta.tKebutuhanKertaDo.UseDB(db, opts...)
	_tKebutuhanKerta.tKebutuhanKertaDo.UseModel(&model.TKebutuhanKerta{})

	tableName := _tKebutuhanKerta.tKebutuhanKertaDo.TableName()
	_tKebutuhanKerta.ALL = field.NewAsterisk(tableName)
	_tKebutuhanKerta.CIDKebutuhanKertas = field.NewInt32(tableName, "c_id_kebutuhan_kertas")
	_tKebutuhanKerta.CIDPusatRisso = field.NewInt32(tableName, "c_id_pusat_risso")
	_tKebutuhanKerta.CIDBidang = field.NewInt32(tableName, "c_id_bidang")
	_tKebutuhanKerta.CIDKegiatanAcuan = field.NewInt32(tableName, "c_id_kegiatan_acuan")
	_tKebutuhanKerta.CIDKebutuhan = field.NewInt32(tableName, "c_id_kebutuhan")
	_tKebutuhanKerta.CIDJenisKertas = field.NewInt32(tableName, "c_id_jenis_kertas")
	_tKebutuhanKerta.CPembagiKertas = field.NewInt32(tableName, "c_pembagi_kertas")
	_tKebutuhanKerta.CPengaliKertas = field.NewInt32(tableName, "c_pengali_kertas")
	_tKebutuhanKerta.CTargetKeikutsertaan = field.NewInt32(tableName, "c_target_keikutsertaan")
	_tKebutuhanKerta.CIsEvent = field.NewString(tableName, "c_is_event")
	_tKebutuhanKerta.CTanggalAwal = field.NewTime(tableName, "c_tanggal_awal")
	_tKebutuhanKerta.CTanggalAkhir = field.NewTime(tableName, "c_tanggal_akhir")
	_tKebutuhanKerta.CJumlah = field.NewInt32(tableName, "c_jumlah")
	_tKebutuhanKerta.CJumlahAkhir = field.NewInt32(tableName, "c_jumlah_akhir")
	_tKebutuhanKerta.CUpdater = field.NewString(tableName, "c_updater")
	_tKebutuhanKerta.CLastUpdate = field.NewTime(tableName, "c_last_update")

	_tKebutuhanKerta.fillFieldMap()

	return _tKebutuhanKerta
}

type tKebutuhanKerta struct {
	tKebutuhanKertaDo

	ALL                  field.Asterisk
	CIDKebutuhanKertas   field.Int32
	CIDPusatRisso        field.Int32
	CIDBidang            field.Int32
	CIDKegiatanAcuan     field.Int32
	CIDKebutuhan         field.Int32
	CIDJenisKertas       field.Int32
	CPembagiKertas       field.Int32
	CPengaliKertas       field.Int32
	CTargetKeikutsertaan field.Int32
	CIsEvent             field.String
	CTanggalAwal         field.Time
	CTanggalAkhir        field.Time
	CJumlah              field.Int32
	CJumlahAkhir         field.Int32
	CUpdater             field.String
	CLastUpdate          field.Time

	fieldMap map[string]field.Expr
}

func (t tKebutuhanKerta) Table(newTableName string) *tKebutuhanKerta {
	t.tKebutuhanKertaDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tKebutuhanKerta) As(alias string) *tKebutuhanKerta {
	t.tKebutuhanKertaDo.DO = *(t.tKebutuhanKertaDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tKebutuhanKerta) updateTableName(table string) *tKebutuhanKerta {
	t.ALL = field.NewAsterisk(table)
	t.CIDKebutuhanKertas = field.NewInt32(table, "c_id_kebutuhan_kertas")
	t.CIDPusatRisso = field.NewInt32(table, "c_id_pusat_risso")
	t.CIDBidang = field.NewInt32(table, "c_id_bidang")
	t.CIDKegiatanAcuan = field.NewInt32(table, "c_id_kegiatan_acuan")
	t.CIDKebutuhan = field.NewInt32(table, "c_id_kebutuhan")
	t.CIDJenisKertas = field.NewInt32(table, "c_id_jenis_kertas")
	t.CPembagiKertas = field.NewInt32(table, "c_pembagi_kertas")
	t.CPengaliKertas = field.NewInt32(table, "c_pengali_kertas")
	t.CTargetKeikutsertaan = field.NewInt32(table, "c_target_keikutsertaan")
	t.CIsEvent = field.NewString(table, "c_is_event")
	t.CTanggalAwal = field.NewTime(table, "c_tanggal_awal")
	t.CTanggalAkhir = field.NewTime(table, "c_tanggal_akhir")
	t.CJumlah = field.NewInt32(table, "c_jumlah")
	t.CJumlahAkhir = field.NewInt32(table, "c_jumlah_akhir")
	t.CUpdater = field.NewString(table, "c_updater")
	t.CLastUpdate = field.NewTime(table, "c_last_update")

	t.fillFieldMap()

	return t
}

func (t *tKebutuhanKerta) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tKebutuhanKerta) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 16)
	t.fieldMap["c_id_kebutuhan_kertas"] = t.CIDKebutuhanKertas
	t.fieldMap["c_id_pusat_risso"] = t.CIDPusatRisso
	t.fieldMap["c_id_bidang"] = t.CIDBidang
	t.fieldMap["c_id_kegiatan_acuan"] = t.CIDKegiatanAcuan
	t.fieldMap["c_id_kebutuhan"] = t.CIDKebutuhan
	t.fieldMap["c_id_jenis_kertas"] = t.CIDJenisKertas
	t.fieldMap["c_pembagi_kertas"] = t.CPembagiKertas
	t.fieldMap["c_pengali_kertas"] = t.CPengaliKertas
	t.fieldMap["c_target_keikutsertaan"] = t.CTargetKeikutsertaan
	t.fieldMap["c_is_event"] = t.CIsEvent
	t.fieldMap["c_tanggal_awal"] = t.CTanggalAwal
	t.fieldMap["c_tanggal_akhir"] = t.CTanggalAkhir
	t.fieldMap["c_jumlah"] = t.CJumlah
	t.fieldMap["c_jumlah_akhir"] = t.CJumlahAkhir
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_last_update"] = t.CLastUpdate
}

func (t tKebutuhanKerta) clone(db *gorm.DB) tKebutuhanKerta {
	t.tKebutuhanKertaDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tKebutuhanKerta) replaceDB(db *gorm.DB) tKebutuhanKerta {
	t.tKebutuhanKertaDo.ReplaceDB(db)
	return t
}

type tKebutuhanKertaDo struct{ gen.DO }

type ITKebutuhanKertaDo interface {
	gen.SubQuery
	Debug() ITKebutuhanKertaDo
	WithContext(ctx context.Context) ITKebutuhanKertaDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITKebutuhanKertaDo
	WriteDB() ITKebutuhanKertaDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITKebutuhanKertaDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITKebutuhanKertaDo
	Not(conds ...gen.Condition) ITKebutuhanKertaDo
	Or(conds ...gen.Condition) ITKebutuhanKertaDo
	Select(conds ...field.Expr) ITKebutuhanKertaDo
	Where(conds ...gen.Condition) ITKebutuhanKertaDo
	Order(conds ...field.Expr) ITKebutuhanKertaDo
	Distinct(cols ...field.Expr) ITKebutuhanKertaDo
	Omit(cols ...field.Expr) ITKebutuhanKertaDo
	Join(table schema.Tabler, on ...field.Expr) ITKebutuhanKertaDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITKebutuhanKertaDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITKebutuhanKertaDo
	Group(cols ...field.Expr) ITKebutuhanKertaDo
	Having(conds ...gen.Condition) ITKebutuhanKertaDo
	Limit(limit int) ITKebutuhanKertaDo
	Offset(offset int) ITKebutuhanKertaDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITKebutuhanKertaDo
	Unscoped() ITKebutuhanKertaDo
	Create(values ...*model.TKebutuhanKerta) error
	CreateInBatches(values []*model.TKebutuhanKerta, batchSize int) error
	Save(values ...*model.TKebutuhanKerta) error
	First() (*model.TKebutuhanKerta, error)
	Take() (*model.TKebutuhanKerta, error)
	Last() (*model.TKebutuhanKerta, error)
	Find() ([]*model.TKebutuhanKerta, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TKebutuhanKerta, err error)
	FindInBatches(result *[]*model.TKebutuhanKerta, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TKebutuhanKerta) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITKebutuhanKertaDo
	Assign(attrs ...field.AssignExpr) ITKebutuhanKertaDo
	Joins(fields ...field.RelationField) ITKebutuhanKertaDo
	Preload(fields ...field.RelationField) ITKebutuhanKertaDo
	FirstOrInit() (*model.TKebutuhanKerta, error)
	FirstOrCreate() (*model.TKebutuhanKerta, error)
	FindByPage(offset int, limit int) (result []*model.TKebutuhanKerta, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITKebutuhanKertaDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tKebutuhanKertaDo) Debug() ITKebutuhanKertaDo {
	return t.withDO(t.DO.Debug())
}

func (t tKebutuhanKertaDo) WithContext(ctx context.Context) ITKebutuhanKertaDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tKebutuhanKertaDo) ReadDB() ITKebutuhanKertaDo {
	return t.Clauses(dbresolver.Read)
}

func (t tKebutuhanKertaDo) WriteDB() ITKebutuhanKertaDo {
	return t.Clauses(dbresolver.Write)
}

func (t tKebutuhanKertaDo) Session(config *gorm.Session) ITKebutuhanKertaDo {
	return t.withDO(t.DO.Session(config))
}

func (t tKebutuhanKertaDo) Clauses(conds ...clause.Expression) ITKebutuhanKertaDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tKebutuhanKertaDo) Returning(value interface{}, columns ...string) ITKebutuhanKertaDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tKebutuhanKertaDo) Not(conds ...gen.Condition) ITKebutuhanKertaDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tKebutuhanKertaDo) Or(conds ...gen.Condition) ITKebutuhanKertaDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tKebutuhanKertaDo) Select(conds ...field.Expr) ITKebutuhanKertaDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tKebutuhanKertaDo) Where(conds ...gen.Condition) ITKebutuhanKertaDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tKebutuhanKertaDo) Order(conds ...field.Expr) ITKebutuhanKertaDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tKebutuhanKertaDo) Distinct(cols ...field.Expr) ITKebutuhanKertaDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tKebutuhanKertaDo) Omit(cols ...field.Expr) ITKebutuhanKertaDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tKebutuhanKertaDo) Join(table schema.Tabler, on ...field.Expr) ITKebutuhanKertaDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tKebutuhanKertaDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITKebutuhanKertaDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tKebutuhanKertaDo) RightJoin(table schema.Tabler, on ...field.Expr) ITKebutuhanKertaDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tKebutuhanKertaDo) Group(cols ...field.Expr) ITKebutuhanKertaDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tKebutuhanKertaDo) Having(conds ...gen.Condition) ITKebutuhanKertaDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tKebutuhanKertaDo) Limit(limit int) ITKebutuhanKertaDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tKebutuhanKertaDo) Offset(offset int) ITKebutuhanKertaDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tKebutuhanKertaDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITKebutuhanKertaDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tKebutuhanKertaDo) Unscoped() ITKebutuhanKertaDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tKebutuhanKertaDo) Create(values ...*model.TKebutuhanKerta) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tKebutuhanKertaDo) CreateInBatches(values []*model.TKebutuhanKerta, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tKebutuhanKertaDo) Save(values ...*model.TKebutuhanKerta) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tKebutuhanKertaDo) First() (*model.TKebutuhanKerta, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKebutuhanKerta), nil
	}
}

func (t tKebutuhanKertaDo) Take() (*model.TKebutuhanKerta, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKebutuhanKerta), nil
	}
}

func (t tKebutuhanKertaDo) Last() (*model.TKebutuhanKerta, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKebutuhanKerta), nil
	}
}

func (t tKebutuhanKertaDo) Find() ([]*model.TKebutuhanKerta, error) {
	result, err := t.DO.Find()
	return result.([]*model.TKebutuhanKerta), err
}

func (t tKebutuhanKertaDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TKebutuhanKerta, err error) {
	buf := make([]*model.TKebutuhanKerta, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tKebutuhanKertaDo) FindInBatches(result *[]*model.TKebutuhanKerta, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tKebutuhanKertaDo) Attrs(attrs ...field.AssignExpr) ITKebutuhanKertaDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tKebutuhanKertaDo) Assign(attrs ...field.AssignExpr) ITKebutuhanKertaDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tKebutuhanKertaDo) Joins(fields ...field.RelationField) ITKebutuhanKertaDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tKebutuhanKertaDo) Preload(fields ...field.RelationField) ITKebutuhanKertaDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tKebutuhanKertaDo) FirstOrInit() (*model.TKebutuhanKerta, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKebutuhanKerta), nil
	}
}

func (t tKebutuhanKertaDo) FirstOrCreate() (*model.TKebutuhanKerta, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKebutuhanKerta), nil
	}
}

func (t tKebutuhanKertaDo) FindByPage(offset int, limit int) (result []*model.TKebutuhanKerta, count int64, err error) {
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

func (t tKebutuhanKertaDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tKebutuhanKertaDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tKebutuhanKertaDo) Delete(models ...*model.TKebutuhanKerta) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tKebutuhanKertaDo) withDO(do gen.Dao) *tKebutuhanKertaDo {
	t.DO = *do.(*gen.DO)
	return t
}