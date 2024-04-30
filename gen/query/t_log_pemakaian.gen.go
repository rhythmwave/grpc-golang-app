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

func newTLogPemakaian(db *gorm.DB, opts ...gen.DOOption) tLogPemakaian {
	_tLogPemakaian := tLogPemakaian{}

	_tLogPemakaian.tLogPemakaianDo.UseDB(db, opts...)
	_tLogPemakaian.tLogPemakaianDo.UseModel(&model.TLogPemakaian{})

	tableName := _tLogPemakaian.tLogPemakaianDo.TableName()
	_tLogPemakaian.ALL = field.NewAsterisk(tableName)
	_tLogPemakaian.CIDLog = field.NewInt32(tableName, "c_id_log")
	_tLogPemakaian.CIDPenanda = field.NewInt32(tableName, "c_id_penanda")
	_tLogPemakaian.CPenanda = field.NewString(tableName, "c_penanda")
	_tLogPemakaian.CIDAlat = field.NewInt32(tableName, "c_id_alat")
	_tLogPemakaian.CNamaAlat = field.NewString(tableName, "c_nama_alat")
	_tLogPemakaian.CJumlah = field.NewInt32(tableName, "c_jumlah")
	_tLogPemakaian.CIDKegiatan = field.NewInt32(tableName, "c_id_kegiatan")
	_tLogPemakaian.CNamaKegiatan = field.NewString(tableName, "c_nama_kegiatan")
	_tLogPemakaian.CIDSekolah = field.NewInt32(tableName, "c_id_sekolah")
	_tLogPemakaian.CSekolah = field.NewString(tableName, "c_sekolah")
	_tLogPemakaian.CUpdater = field.NewString(tableName, "c_updater")
	_tLogPemakaian.CLastUpdate = field.NewTime(tableName, "c_last_update")

	_tLogPemakaian.fillFieldMap()

	return _tLogPemakaian
}

type tLogPemakaian struct {
	tLogPemakaianDo

	ALL           field.Asterisk
	CIDLog        field.Int32
	CIDPenanda    field.Int32
	CPenanda      field.String
	CIDAlat       field.Int32
	CNamaAlat     field.String
	CJumlah       field.Int32
	CIDKegiatan   field.Int32
	CNamaKegiatan field.String
	CIDSekolah    field.Int32
	CSekolah      field.String
	CUpdater      field.String
	CLastUpdate   field.Time

	fieldMap map[string]field.Expr
}

func (t tLogPemakaian) Table(newTableName string) *tLogPemakaian {
	t.tLogPemakaianDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tLogPemakaian) As(alias string) *tLogPemakaian {
	t.tLogPemakaianDo.DO = *(t.tLogPemakaianDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tLogPemakaian) updateTableName(table string) *tLogPemakaian {
	t.ALL = field.NewAsterisk(table)
	t.CIDLog = field.NewInt32(table, "c_id_log")
	t.CIDPenanda = field.NewInt32(table, "c_id_penanda")
	t.CPenanda = field.NewString(table, "c_penanda")
	t.CIDAlat = field.NewInt32(table, "c_id_alat")
	t.CNamaAlat = field.NewString(table, "c_nama_alat")
	t.CJumlah = field.NewInt32(table, "c_jumlah")
	t.CIDKegiatan = field.NewInt32(table, "c_id_kegiatan")
	t.CNamaKegiatan = field.NewString(table, "c_nama_kegiatan")
	t.CIDSekolah = field.NewInt32(table, "c_id_sekolah")
	t.CSekolah = field.NewString(table, "c_sekolah")
	t.CUpdater = field.NewString(table, "c_updater")
	t.CLastUpdate = field.NewTime(table, "c_last_update")

	t.fillFieldMap()

	return t
}

func (t *tLogPemakaian) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tLogPemakaian) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 12)
	t.fieldMap["c_id_log"] = t.CIDLog
	t.fieldMap["c_id_penanda"] = t.CIDPenanda
	t.fieldMap["c_penanda"] = t.CPenanda
	t.fieldMap["c_id_alat"] = t.CIDAlat
	t.fieldMap["c_nama_alat"] = t.CNamaAlat
	t.fieldMap["c_jumlah"] = t.CJumlah
	t.fieldMap["c_id_kegiatan"] = t.CIDKegiatan
	t.fieldMap["c_nama_kegiatan"] = t.CNamaKegiatan
	t.fieldMap["c_id_sekolah"] = t.CIDSekolah
	t.fieldMap["c_sekolah"] = t.CSekolah
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_last_update"] = t.CLastUpdate
}

func (t tLogPemakaian) clone(db *gorm.DB) tLogPemakaian {
	t.tLogPemakaianDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tLogPemakaian) replaceDB(db *gorm.DB) tLogPemakaian {
	t.tLogPemakaianDo.ReplaceDB(db)
	return t
}

type tLogPemakaianDo struct{ gen.DO }

type ITLogPemakaianDo interface {
	gen.SubQuery
	Debug() ITLogPemakaianDo
	WithContext(ctx context.Context) ITLogPemakaianDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITLogPemakaianDo
	WriteDB() ITLogPemakaianDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITLogPemakaianDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITLogPemakaianDo
	Not(conds ...gen.Condition) ITLogPemakaianDo
	Or(conds ...gen.Condition) ITLogPemakaianDo
	Select(conds ...field.Expr) ITLogPemakaianDo
	Where(conds ...gen.Condition) ITLogPemakaianDo
	Order(conds ...field.Expr) ITLogPemakaianDo
	Distinct(cols ...field.Expr) ITLogPemakaianDo
	Omit(cols ...field.Expr) ITLogPemakaianDo
	Join(table schema.Tabler, on ...field.Expr) ITLogPemakaianDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITLogPemakaianDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITLogPemakaianDo
	Group(cols ...field.Expr) ITLogPemakaianDo
	Having(conds ...gen.Condition) ITLogPemakaianDo
	Limit(limit int) ITLogPemakaianDo
	Offset(offset int) ITLogPemakaianDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITLogPemakaianDo
	Unscoped() ITLogPemakaianDo
	Create(values ...*model.TLogPemakaian) error
	CreateInBatches(values []*model.TLogPemakaian, batchSize int) error
	Save(values ...*model.TLogPemakaian) error
	First() (*model.TLogPemakaian, error)
	Take() (*model.TLogPemakaian, error)
	Last() (*model.TLogPemakaian, error)
	Find() ([]*model.TLogPemakaian, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TLogPemakaian, err error)
	FindInBatches(result *[]*model.TLogPemakaian, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TLogPemakaian) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITLogPemakaianDo
	Assign(attrs ...field.AssignExpr) ITLogPemakaianDo
	Joins(fields ...field.RelationField) ITLogPemakaianDo
	Preload(fields ...field.RelationField) ITLogPemakaianDo
	FirstOrInit() (*model.TLogPemakaian, error)
	FirstOrCreate() (*model.TLogPemakaian, error)
	FindByPage(offset int, limit int) (result []*model.TLogPemakaian, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITLogPemakaianDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tLogPemakaianDo) Debug() ITLogPemakaianDo {
	return t.withDO(t.DO.Debug())
}

func (t tLogPemakaianDo) WithContext(ctx context.Context) ITLogPemakaianDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tLogPemakaianDo) ReadDB() ITLogPemakaianDo {
	return t.Clauses(dbresolver.Read)
}

func (t tLogPemakaianDo) WriteDB() ITLogPemakaianDo {
	return t.Clauses(dbresolver.Write)
}

func (t tLogPemakaianDo) Session(config *gorm.Session) ITLogPemakaianDo {
	return t.withDO(t.DO.Session(config))
}

func (t tLogPemakaianDo) Clauses(conds ...clause.Expression) ITLogPemakaianDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tLogPemakaianDo) Returning(value interface{}, columns ...string) ITLogPemakaianDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tLogPemakaianDo) Not(conds ...gen.Condition) ITLogPemakaianDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tLogPemakaianDo) Or(conds ...gen.Condition) ITLogPemakaianDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tLogPemakaianDo) Select(conds ...field.Expr) ITLogPemakaianDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tLogPemakaianDo) Where(conds ...gen.Condition) ITLogPemakaianDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tLogPemakaianDo) Order(conds ...field.Expr) ITLogPemakaianDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tLogPemakaianDo) Distinct(cols ...field.Expr) ITLogPemakaianDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tLogPemakaianDo) Omit(cols ...field.Expr) ITLogPemakaianDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tLogPemakaianDo) Join(table schema.Tabler, on ...field.Expr) ITLogPemakaianDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tLogPemakaianDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITLogPemakaianDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tLogPemakaianDo) RightJoin(table schema.Tabler, on ...field.Expr) ITLogPemakaianDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tLogPemakaianDo) Group(cols ...field.Expr) ITLogPemakaianDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tLogPemakaianDo) Having(conds ...gen.Condition) ITLogPemakaianDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tLogPemakaianDo) Limit(limit int) ITLogPemakaianDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tLogPemakaianDo) Offset(offset int) ITLogPemakaianDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tLogPemakaianDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITLogPemakaianDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tLogPemakaianDo) Unscoped() ITLogPemakaianDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tLogPemakaianDo) Create(values ...*model.TLogPemakaian) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tLogPemakaianDo) CreateInBatches(values []*model.TLogPemakaian, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tLogPemakaianDo) Save(values ...*model.TLogPemakaian) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tLogPemakaianDo) First() (*model.TLogPemakaian, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TLogPemakaian), nil
	}
}

func (t tLogPemakaianDo) Take() (*model.TLogPemakaian, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TLogPemakaian), nil
	}
}

func (t tLogPemakaianDo) Last() (*model.TLogPemakaian, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TLogPemakaian), nil
	}
}

func (t tLogPemakaianDo) Find() ([]*model.TLogPemakaian, error) {
	result, err := t.DO.Find()
	return result.([]*model.TLogPemakaian), err
}

func (t tLogPemakaianDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TLogPemakaian, err error) {
	buf := make([]*model.TLogPemakaian, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tLogPemakaianDo) FindInBatches(result *[]*model.TLogPemakaian, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tLogPemakaianDo) Attrs(attrs ...field.AssignExpr) ITLogPemakaianDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tLogPemakaianDo) Assign(attrs ...field.AssignExpr) ITLogPemakaianDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tLogPemakaianDo) Joins(fields ...field.RelationField) ITLogPemakaianDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tLogPemakaianDo) Preload(fields ...field.RelationField) ITLogPemakaianDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tLogPemakaianDo) FirstOrInit() (*model.TLogPemakaian, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TLogPemakaian), nil
	}
}

func (t tLogPemakaianDo) FirstOrCreate() (*model.TLogPemakaian, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TLogPemakaian), nil
	}
}

func (t tLogPemakaianDo) FindByPage(offset int, limit int) (result []*model.TLogPemakaian, count int64, err error) {
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

func (t tLogPemakaianDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tLogPemakaianDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tLogPemakaianDo) Delete(models ...*model.TLogPemakaian) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tLogPemakaianDo) withDO(do gen.Dao) *tLogPemakaianDo {
	t.DO = *do.(*gen.DO)
	return t
}