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

func newTAcuanKegiatanBckp(db *gorm.DB, opts ...gen.DOOption) tAcuanKegiatanBckp {
	_tAcuanKegiatanBckp := tAcuanKegiatanBckp{}

	_tAcuanKegiatanBckp.tAcuanKegiatanBckpDo.UseDB(db, opts...)
	_tAcuanKegiatanBckp.tAcuanKegiatanBckpDo.UseModel(&model.TAcuanKegiatanBckp{})

	tableName := _tAcuanKegiatanBckp.tAcuanKegiatanBckpDo.TableName()
	_tAcuanKegiatanBckp.ALL = field.NewAsterisk(tableName)
	_tAcuanKegiatanBckp.CKodeAcuan = field.NewString(tableName, "c_kode_acuan")
	_tAcuanKegiatanBckp.CKodeKelBarang = field.NewString(tableName, "c_kode_kel_barang")
	_tAcuanKegiatanBckp.CKodeKelAnggaran = field.NewString(tableName, "c_kode_kel_anggaran")
	_tAcuanKegiatanBckp.CKodeKepanitiaan = field.NewString(tableName, "c_kode_kepanitiaan")
	_tAcuanKegiatanBckp.CBidangPenanggungJawab = field.NewInt32(tableName, "c_bidang_penanggung_jawab")
	_tAcuanKegiatanBckp.CBidangTerkait = field.NewString(tableName, "c_bidang_terkait")
	_tAcuanKegiatanBckp.CNamaKegiatan = field.NewString(tableName, "c_nama_kegiatan")
	_tAcuanKegiatanBckp.CUpline = field.NewString(tableName, "c_upline")
	_tAcuanKegiatanBckp.CIsPromosi = field.NewInt16(tableName, "c_is_promosi")
	_tAcuanKegiatanBckp.CIsPusat = field.NewInt16(tableName, "c_is_pusat")
	_tAcuanKegiatanBckp.CStatus = field.NewString(tableName, "c_status")
	_tAcuanKegiatanBckp.CUpdater = field.NewString(tableName, "c_updater")
	_tAcuanKegiatanBckp.CLastUpdate = field.NewTime(tableName, "c_last_update")

	_tAcuanKegiatanBckp.fillFieldMap()

	return _tAcuanKegiatanBckp
}

type tAcuanKegiatanBckp struct {
	tAcuanKegiatanBckpDo

	ALL                    field.Asterisk
	CKodeAcuan             field.String
	CKodeKelBarang         field.String
	CKodeKelAnggaran       field.String
	CKodeKepanitiaan       field.String
	CBidangPenanggungJawab field.Int32
	CBidangTerkait         field.String
	CNamaKegiatan          field.String
	CUpline                field.String
	CIsPromosi             field.Int16
	CIsPusat               field.Int16
	CStatus                field.String
	CUpdater               field.String
	CLastUpdate            field.Time

	fieldMap map[string]field.Expr
}

func (t tAcuanKegiatanBckp) Table(newTableName string) *tAcuanKegiatanBckp {
	t.tAcuanKegiatanBckpDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tAcuanKegiatanBckp) As(alias string) *tAcuanKegiatanBckp {
	t.tAcuanKegiatanBckpDo.DO = *(t.tAcuanKegiatanBckpDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tAcuanKegiatanBckp) updateTableName(table string) *tAcuanKegiatanBckp {
	t.ALL = field.NewAsterisk(table)
	t.CKodeAcuan = field.NewString(table, "c_kode_acuan")
	t.CKodeKelBarang = field.NewString(table, "c_kode_kel_barang")
	t.CKodeKelAnggaran = field.NewString(table, "c_kode_kel_anggaran")
	t.CKodeKepanitiaan = field.NewString(table, "c_kode_kepanitiaan")
	t.CBidangPenanggungJawab = field.NewInt32(table, "c_bidang_penanggung_jawab")
	t.CBidangTerkait = field.NewString(table, "c_bidang_terkait")
	t.CNamaKegiatan = field.NewString(table, "c_nama_kegiatan")
	t.CUpline = field.NewString(table, "c_upline")
	t.CIsPromosi = field.NewInt16(table, "c_is_promosi")
	t.CIsPusat = field.NewInt16(table, "c_is_pusat")
	t.CStatus = field.NewString(table, "c_status")
	t.CUpdater = field.NewString(table, "c_updater")
	t.CLastUpdate = field.NewTime(table, "c_last_update")

	t.fillFieldMap()

	return t
}

func (t *tAcuanKegiatanBckp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tAcuanKegiatanBckp) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 13)
	t.fieldMap["c_kode_acuan"] = t.CKodeAcuan
	t.fieldMap["c_kode_kel_barang"] = t.CKodeKelBarang
	t.fieldMap["c_kode_kel_anggaran"] = t.CKodeKelAnggaran
	t.fieldMap["c_kode_kepanitiaan"] = t.CKodeKepanitiaan
	t.fieldMap["c_bidang_penanggung_jawab"] = t.CBidangPenanggungJawab
	t.fieldMap["c_bidang_terkait"] = t.CBidangTerkait
	t.fieldMap["c_nama_kegiatan"] = t.CNamaKegiatan
	t.fieldMap["c_upline"] = t.CUpline
	t.fieldMap["c_is_promosi"] = t.CIsPromosi
	t.fieldMap["c_is_pusat"] = t.CIsPusat
	t.fieldMap["c_status"] = t.CStatus
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_last_update"] = t.CLastUpdate
}

func (t tAcuanKegiatanBckp) clone(db *gorm.DB) tAcuanKegiatanBckp {
	t.tAcuanKegiatanBckpDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tAcuanKegiatanBckp) replaceDB(db *gorm.DB) tAcuanKegiatanBckp {
	t.tAcuanKegiatanBckpDo.ReplaceDB(db)
	return t
}

type tAcuanKegiatanBckpDo struct{ gen.DO }

type ITAcuanKegiatanBckpDo interface {
	gen.SubQuery
	Debug() ITAcuanKegiatanBckpDo
	WithContext(ctx context.Context) ITAcuanKegiatanBckpDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITAcuanKegiatanBckpDo
	WriteDB() ITAcuanKegiatanBckpDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITAcuanKegiatanBckpDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITAcuanKegiatanBckpDo
	Not(conds ...gen.Condition) ITAcuanKegiatanBckpDo
	Or(conds ...gen.Condition) ITAcuanKegiatanBckpDo
	Select(conds ...field.Expr) ITAcuanKegiatanBckpDo
	Where(conds ...gen.Condition) ITAcuanKegiatanBckpDo
	Order(conds ...field.Expr) ITAcuanKegiatanBckpDo
	Distinct(cols ...field.Expr) ITAcuanKegiatanBckpDo
	Omit(cols ...field.Expr) ITAcuanKegiatanBckpDo
	Join(table schema.Tabler, on ...field.Expr) ITAcuanKegiatanBckpDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITAcuanKegiatanBckpDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITAcuanKegiatanBckpDo
	Group(cols ...field.Expr) ITAcuanKegiatanBckpDo
	Having(conds ...gen.Condition) ITAcuanKegiatanBckpDo
	Limit(limit int) ITAcuanKegiatanBckpDo
	Offset(offset int) ITAcuanKegiatanBckpDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITAcuanKegiatanBckpDo
	Unscoped() ITAcuanKegiatanBckpDo
	Create(values ...*model.TAcuanKegiatanBckp) error
	CreateInBatches(values []*model.TAcuanKegiatanBckp, batchSize int) error
	Save(values ...*model.TAcuanKegiatanBckp) error
	First() (*model.TAcuanKegiatanBckp, error)
	Take() (*model.TAcuanKegiatanBckp, error)
	Last() (*model.TAcuanKegiatanBckp, error)
	Find() ([]*model.TAcuanKegiatanBckp, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TAcuanKegiatanBckp, err error)
	FindInBatches(result *[]*model.TAcuanKegiatanBckp, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TAcuanKegiatanBckp) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITAcuanKegiatanBckpDo
	Assign(attrs ...field.AssignExpr) ITAcuanKegiatanBckpDo
	Joins(fields ...field.RelationField) ITAcuanKegiatanBckpDo
	Preload(fields ...field.RelationField) ITAcuanKegiatanBckpDo
	FirstOrInit() (*model.TAcuanKegiatanBckp, error)
	FirstOrCreate() (*model.TAcuanKegiatanBckp, error)
	FindByPage(offset int, limit int) (result []*model.TAcuanKegiatanBckp, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITAcuanKegiatanBckpDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tAcuanKegiatanBckpDo) Debug() ITAcuanKegiatanBckpDo {
	return t.withDO(t.DO.Debug())
}

func (t tAcuanKegiatanBckpDo) WithContext(ctx context.Context) ITAcuanKegiatanBckpDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tAcuanKegiatanBckpDo) ReadDB() ITAcuanKegiatanBckpDo {
	return t.Clauses(dbresolver.Read)
}

func (t tAcuanKegiatanBckpDo) WriteDB() ITAcuanKegiatanBckpDo {
	return t.Clauses(dbresolver.Write)
}

func (t tAcuanKegiatanBckpDo) Session(config *gorm.Session) ITAcuanKegiatanBckpDo {
	return t.withDO(t.DO.Session(config))
}

func (t tAcuanKegiatanBckpDo) Clauses(conds ...clause.Expression) ITAcuanKegiatanBckpDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tAcuanKegiatanBckpDo) Returning(value interface{}, columns ...string) ITAcuanKegiatanBckpDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tAcuanKegiatanBckpDo) Not(conds ...gen.Condition) ITAcuanKegiatanBckpDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tAcuanKegiatanBckpDo) Or(conds ...gen.Condition) ITAcuanKegiatanBckpDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tAcuanKegiatanBckpDo) Select(conds ...field.Expr) ITAcuanKegiatanBckpDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tAcuanKegiatanBckpDo) Where(conds ...gen.Condition) ITAcuanKegiatanBckpDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tAcuanKegiatanBckpDo) Order(conds ...field.Expr) ITAcuanKegiatanBckpDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tAcuanKegiatanBckpDo) Distinct(cols ...field.Expr) ITAcuanKegiatanBckpDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tAcuanKegiatanBckpDo) Omit(cols ...field.Expr) ITAcuanKegiatanBckpDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tAcuanKegiatanBckpDo) Join(table schema.Tabler, on ...field.Expr) ITAcuanKegiatanBckpDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tAcuanKegiatanBckpDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITAcuanKegiatanBckpDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tAcuanKegiatanBckpDo) RightJoin(table schema.Tabler, on ...field.Expr) ITAcuanKegiatanBckpDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tAcuanKegiatanBckpDo) Group(cols ...field.Expr) ITAcuanKegiatanBckpDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tAcuanKegiatanBckpDo) Having(conds ...gen.Condition) ITAcuanKegiatanBckpDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tAcuanKegiatanBckpDo) Limit(limit int) ITAcuanKegiatanBckpDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tAcuanKegiatanBckpDo) Offset(offset int) ITAcuanKegiatanBckpDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tAcuanKegiatanBckpDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITAcuanKegiatanBckpDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tAcuanKegiatanBckpDo) Unscoped() ITAcuanKegiatanBckpDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tAcuanKegiatanBckpDo) Create(values ...*model.TAcuanKegiatanBckp) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tAcuanKegiatanBckpDo) CreateInBatches(values []*model.TAcuanKegiatanBckp, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tAcuanKegiatanBckpDo) Save(values ...*model.TAcuanKegiatanBckp) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tAcuanKegiatanBckpDo) First() (*model.TAcuanKegiatanBckp, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAcuanKegiatanBckp), nil
	}
}

func (t tAcuanKegiatanBckpDo) Take() (*model.TAcuanKegiatanBckp, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAcuanKegiatanBckp), nil
	}
}

func (t tAcuanKegiatanBckpDo) Last() (*model.TAcuanKegiatanBckp, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAcuanKegiatanBckp), nil
	}
}

func (t tAcuanKegiatanBckpDo) Find() ([]*model.TAcuanKegiatanBckp, error) {
	result, err := t.DO.Find()
	return result.([]*model.TAcuanKegiatanBckp), err
}

func (t tAcuanKegiatanBckpDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TAcuanKegiatanBckp, err error) {
	buf := make([]*model.TAcuanKegiatanBckp, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tAcuanKegiatanBckpDo) FindInBatches(result *[]*model.TAcuanKegiatanBckp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tAcuanKegiatanBckpDo) Attrs(attrs ...field.AssignExpr) ITAcuanKegiatanBckpDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tAcuanKegiatanBckpDo) Assign(attrs ...field.AssignExpr) ITAcuanKegiatanBckpDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tAcuanKegiatanBckpDo) Joins(fields ...field.RelationField) ITAcuanKegiatanBckpDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tAcuanKegiatanBckpDo) Preload(fields ...field.RelationField) ITAcuanKegiatanBckpDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tAcuanKegiatanBckpDo) FirstOrInit() (*model.TAcuanKegiatanBckp, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAcuanKegiatanBckp), nil
	}
}

func (t tAcuanKegiatanBckpDo) FirstOrCreate() (*model.TAcuanKegiatanBckp, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAcuanKegiatanBckp), nil
	}
}

func (t tAcuanKegiatanBckpDo) FindByPage(offset int, limit int) (result []*model.TAcuanKegiatanBckp, count int64, err error) {
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

func (t tAcuanKegiatanBckpDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tAcuanKegiatanBckpDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tAcuanKegiatanBckpDo) Delete(models ...*model.TAcuanKegiatanBckp) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tAcuanKegiatanBckpDo) withDO(do gen.Dao) *tAcuanKegiatanBckpDo {
	t.DO = *do.(*gen.DO)
	return t
}
