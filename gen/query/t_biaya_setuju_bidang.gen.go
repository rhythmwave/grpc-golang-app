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

func newTBiayaSetujuBidang(db *gorm.DB, opts ...gen.DOOption) tBiayaSetujuBidang {
	_tBiayaSetujuBidang := tBiayaSetujuBidang{}

	_tBiayaSetujuBidang.tBiayaSetujuBidangDo.UseDB(db, opts...)
	_tBiayaSetujuBidang.tBiayaSetujuBidangDo.UseModel(&model.TBiayaSetujuBidang{})

	tableName := _tBiayaSetujuBidang.tBiayaSetujuBidangDo.TableName()
	_tBiayaSetujuBidang.ALL = field.NewAsterisk(tableName)
	_tBiayaSetujuBidang.CIDKegiatan = field.NewInt32(tableName, "c_id_kegiatan")
	_tBiayaSetujuBidang.CIDJenis = field.NewInt32(tableName, "c_id_jenis")
	_tBiayaSetujuBidang.CBiayaSetuju = field.NewInt32(tableName, "c_biaya_setuju")
	_tBiayaSetujuBidang.CTanggalPersetujuanBiaya = field.NewTime(tableName, "c_tanggal_persetujuan_biaya")
	_tBiayaSetujuBidang.CKodePengajuanAnggaran = field.NewString(tableName, "c_kode_pengajuan_anggaran")
	_tBiayaSetujuBidang.CUpdater = field.NewString(tableName, "c_updater")
	_tBiayaSetujuBidang.CLastUpdate = field.NewTime(tableName, "c_last_update")

	_tBiayaSetujuBidang.fillFieldMap()

	return _tBiayaSetujuBidang
}

type tBiayaSetujuBidang struct {
	tBiayaSetujuBidangDo

	ALL                      field.Asterisk
	CIDKegiatan              field.Int32
	CIDJenis                 field.Int32
	CBiayaSetuju             field.Int32
	CTanggalPersetujuanBiaya field.Time
	CKodePengajuanAnggaran   field.String
	CUpdater                 field.String
	CLastUpdate              field.Time

	fieldMap map[string]field.Expr
}

func (t tBiayaSetujuBidang) Table(newTableName string) *tBiayaSetujuBidang {
	t.tBiayaSetujuBidangDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tBiayaSetujuBidang) As(alias string) *tBiayaSetujuBidang {
	t.tBiayaSetujuBidangDo.DO = *(t.tBiayaSetujuBidangDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tBiayaSetujuBidang) updateTableName(table string) *tBiayaSetujuBidang {
	t.ALL = field.NewAsterisk(table)
	t.CIDKegiatan = field.NewInt32(table, "c_id_kegiatan")
	t.CIDJenis = field.NewInt32(table, "c_id_jenis")
	t.CBiayaSetuju = field.NewInt32(table, "c_biaya_setuju")
	t.CTanggalPersetujuanBiaya = field.NewTime(table, "c_tanggal_persetujuan_biaya")
	t.CKodePengajuanAnggaran = field.NewString(table, "c_kode_pengajuan_anggaran")
	t.CUpdater = field.NewString(table, "c_updater")
	t.CLastUpdate = field.NewTime(table, "c_last_update")

	t.fillFieldMap()

	return t
}

func (t *tBiayaSetujuBidang) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tBiayaSetujuBidang) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 7)
	t.fieldMap["c_id_kegiatan"] = t.CIDKegiatan
	t.fieldMap["c_id_jenis"] = t.CIDJenis
	t.fieldMap["c_biaya_setuju"] = t.CBiayaSetuju
	t.fieldMap["c_tanggal_persetujuan_biaya"] = t.CTanggalPersetujuanBiaya
	t.fieldMap["c_kode_pengajuan_anggaran"] = t.CKodePengajuanAnggaran
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_last_update"] = t.CLastUpdate
}

func (t tBiayaSetujuBidang) clone(db *gorm.DB) tBiayaSetujuBidang {
	t.tBiayaSetujuBidangDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tBiayaSetujuBidang) replaceDB(db *gorm.DB) tBiayaSetujuBidang {
	t.tBiayaSetujuBidangDo.ReplaceDB(db)
	return t
}

type tBiayaSetujuBidangDo struct{ gen.DO }

type ITBiayaSetujuBidangDo interface {
	gen.SubQuery
	Debug() ITBiayaSetujuBidangDo
	WithContext(ctx context.Context) ITBiayaSetujuBidangDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITBiayaSetujuBidangDo
	WriteDB() ITBiayaSetujuBidangDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITBiayaSetujuBidangDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITBiayaSetujuBidangDo
	Not(conds ...gen.Condition) ITBiayaSetujuBidangDo
	Or(conds ...gen.Condition) ITBiayaSetujuBidangDo
	Select(conds ...field.Expr) ITBiayaSetujuBidangDo
	Where(conds ...gen.Condition) ITBiayaSetujuBidangDo
	Order(conds ...field.Expr) ITBiayaSetujuBidangDo
	Distinct(cols ...field.Expr) ITBiayaSetujuBidangDo
	Omit(cols ...field.Expr) ITBiayaSetujuBidangDo
	Join(table schema.Tabler, on ...field.Expr) ITBiayaSetujuBidangDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITBiayaSetujuBidangDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITBiayaSetujuBidangDo
	Group(cols ...field.Expr) ITBiayaSetujuBidangDo
	Having(conds ...gen.Condition) ITBiayaSetujuBidangDo
	Limit(limit int) ITBiayaSetujuBidangDo
	Offset(offset int) ITBiayaSetujuBidangDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITBiayaSetujuBidangDo
	Unscoped() ITBiayaSetujuBidangDo
	Create(values ...*model.TBiayaSetujuBidang) error
	CreateInBatches(values []*model.TBiayaSetujuBidang, batchSize int) error
	Save(values ...*model.TBiayaSetujuBidang) error
	First() (*model.TBiayaSetujuBidang, error)
	Take() (*model.TBiayaSetujuBidang, error)
	Last() (*model.TBiayaSetujuBidang, error)
	Find() ([]*model.TBiayaSetujuBidang, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TBiayaSetujuBidang, err error)
	FindInBatches(result *[]*model.TBiayaSetujuBidang, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TBiayaSetujuBidang) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITBiayaSetujuBidangDo
	Assign(attrs ...field.AssignExpr) ITBiayaSetujuBidangDo
	Joins(fields ...field.RelationField) ITBiayaSetujuBidangDo
	Preload(fields ...field.RelationField) ITBiayaSetujuBidangDo
	FirstOrInit() (*model.TBiayaSetujuBidang, error)
	FirstOrCreate() (*model.TBiayaSetujuBidang, error)
	FindByPage(offset int, limit int) (result []*model.TBiayaSetujuBidang, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITBiayaSetujuBidangDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tBiayaSetujuBidangDo) Debug() ITBiayaSetujuBidangDo {
	return t.withDO(t.DO.Debug())
}

func (t tBiayaSetujuBidangDo) WithContext(ctx context.Context) ITBiayaSetujuBidangDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tBiayaSetujuBidangDo) ReadDB() ITBiayaSetujuBidangDo {
	return t.Clauses(dbresolver.Read)
}

func (t tBiayaSetujuBidangDo) WriteDB() ITBiayaSetujuBidangDo {
	return t.Clauses(dbresolver.Write)
}

func (t tBiayaSetujuBidangDo) Session(config *gorm.Session) ITBiayaSetujuBidangDo {
	return t.withDO(t.DO.Session(config))
}

func (t tBiayaSetujuBidangDo) Clauses(conds ...clause.Expression) ITBiayaSetujuBidangDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tBiayaSetujuBidangDo) Returning(value interface{}, columns ...string) ITBiayaSetujuBidangDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tBiayaSetujuBidangDo) Not(conds ...gen.Condition) ITBiayaSetujuBidangDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tBiayaSetujuBidangDo) Or(conds ...gen.Condition) ITBiayaSetujuBidangDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tBiayaSetujuBidangDo) Select(conds ...field.Expr) ITBiayaSetujuBidangDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tBiayaSetujuBidangDo) Where(conds ...gen.Condition) ITBiayaSetujuBidangDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tBiayaSetujuBidangDo) Order(conds ...field.Expr) ITBiayaSetujuBidangDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tBiayaSetujuBidangDo) Distinct(cols ...field.Expr) ITBiayaSetujuBidangDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tBiayaSetujuBidangDo) Omit(cols ...field.Expr) ITBiayaSetujuBidangDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tBiayaSetujuBidangDo) Join(table schema.Tabler, on ...field.Expr) ITBiayaSetujuBidangDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tBiayaSetujuBidangDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITBiayaSetujuBidangDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tBiayaSetujuBidangDo) RightJoin(table schema.Tabler, on ...field.Expr) ITBiayaSetujuBidangDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tBiayaSetujuBidangDo) Group(cols ...field.Expr) ITBiayaSetujuBidangDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tBiayaSetujuBidangDo) Having(conds ...gen.Condition) ITBiayaSetujuBidangDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tBiayaSetujuBidangDo) Limit(limit int) ITBiayaSetujuBidangDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tBiayaSetujuBidangDo) Offset(offset int) ITBiayaSetujuBidangDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tBiayaSetujuBidangDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITBiayaSetujuBidangDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tBiayaSetujuBidangDo) Unscoped() ITBiayaSetujuBidangDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tBiayaSetujuBidangDo) Create(values ...*model.TBiayaSetujuBidang) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tBiayaSetujuBidangDo) CreateInBatches(values []*model.TBiayaSetujuBidang, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tBiayaSetujuBidangDo) Save(values ...*model.TBiayaSetujuBidang) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tBiayaSetujuBidangDo) First() (*model.TBiayaSetujuBidang, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TBiayaSetujuBidang), nil
	}
}

func (t tBiayaSetujuBidangDo) Take() (*model.TBiayaSetujuBidang, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TBiayaSetujuBidang), nil
	}
}

func (t tBiayaSetujuBidangDo) Last() (*model.TBiayaSetujuBidang, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TBiayaSetujuBidang), nil
	}
}

func (t tBiayaSetujuBidangDo) Find() ([]*model.TBiayaSetujuBidang, error) {
	result, err := t.DO.Find()
	return result.([]*model.TBiayaSetujuBidang), err
}

func (t tBiayaSetujuBidangDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TBiayaSetujuBidang, err error) {
	buf := make([]*model.TBiayaSetujuBidang, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tBiayaSetujuBidangDo) FindInBatches(result *[]*model.TBiayaSetujuBidang, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tBiayaSetujuBidangDo) Attrs(attrs ...field.AssignExpr) ITBiayaSetujuBidangDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tBiayaSetujuBidangDo) Assign(attrs ...field.AssignExpr) ITBiayaSetujuBidangDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tBiayaSetujuBidangDo) Joins(fields ...field.RelationField) ITBiayaSetujuBidangDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tBiayaSetujuBidangDo) Preload(fields ...field.RelationField) ITBiayaSetujuBidangDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tBiayaSetujuBidangDo) FirstOrInit() (*model.TBiayaSetujuBidang, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TBiayaSetujuBidang), nil
	}
}

func (t tBiayaSetujuBidangDo) FirstOrCreate() (*model.TBiayaSetujuBidang, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TBiayaSetujuBidang), nil
	}
}

func (t tBiayaSetujuBidangDo) FindByPage(offset int, limit int) (result []*model.TBiayaSetujuBidang, count int64, err error) {
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

func (t tBiayaSetujuBidangDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tBiayaSetujuBidangDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tBiayaSetujuBidangDo) Delete(models ...*model.TBiayaSetujuBidang) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tBiayaSetujuBidangDo) withDO(do gen.Dao) *tBiayaSetujuBidangDo {
	t.DO = *do.(*gen.DO)
	return t
}
