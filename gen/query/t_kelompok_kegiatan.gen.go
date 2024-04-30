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

func newTKelompokKegiatan(db *gorm.DB, opts ...gen.DOOption) tKelompokKegiatan {
	_tKelompokKegiatan := tKelompokKegiatan{}

	_tKelompokKegiatan.tKelompokKegiatanDo.UseDB(db, opts...)
	_tKelompokKegiatan.tKelompokKegiatanDo.UseModel(&model.TKelompokKegiatan{})

	tableName := _tKelompokKegiatan.tKelompokKegiatanDo.TableName()
	_tKelompokKegiatan.ALL = field.NewAsterisk(tableName)
	_tKelompokKegiatan.CIDKelompokKegiatan = field.NewInt32(tableName, "c_id_kelompok_kegiatan")
	_tKelompokKegiatan.CNamaKelompokKegiatan = field.NewString(tableName, "c_nama_kelompok_kegiatan")
	_tKelompokKegiatan.CStatus = field.NewString(tableName, "c_status")
	_tKelompokKegiatan.CUpdater = field.NewString(tableName, "c_updater")
	_tKelompokKegiatan.CLastUpdate = field.NewTime(tableName, "c_last_update")

	_tKelompokKegiatan.fillFieldMap()

	return _tKelompokKegiatan
}

type tKelompokKegiatan struct {
	tKelompokKegiatanDo

	ALL                   field.Asterisk
	CIDKelompokKegiatan   field.Int32
	CNamaKelompokKegiatan field.String
	CStatus               field.String
	CUpdater              field.String
	CLastUpdate           field.Time

	fieldMap map[string]field.Expr
}

func (t tKelompokKegiatan) Table(newTableName string) *tKelompokKegiatan {
	t.tKelompokKegiatanDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tKelompokKegiatan) As(alias string) *tKelompokKegiatan {
	t.tKelompokKegiatanDo.DO = *(t.tKelompokKegiatanDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tKelompokKegiatan) updateTableName(table string) *tKelompokKegiatan {
	t.ALL = field.NewAsterisk(table)
	t.CIDKelompokKegiatan = field.NewInt32(table, "c_id_kelompok_kegiatan")
	t.CNamaKelompokKegiatan = field.NewString(table, "c_nama_kelompok_kegiatan")
	t.CStatus = field.NewString(table, "c_status")
	t.CUpdater = field.NewString(table, "c_updater")
	t.CLastUpdate = field.NewTime(table, "c_last_update")

	t.fillFieldMap()

	return t
}

func (t *tKelompokKegiatan) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tKelompokKegiatan) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 5)
	t.fieldMap["c_id_kelompok_kegiatan"] = t.CIDKelompokKegiatan
	t.fieldMap["c_nama_kelompok_kegiatan"] = t.CNamaKelompokKegiatan
	t.fieldMap["c_status"] = t.CStatus
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_last_update"] = t.CLastUpdate
}

func (t tKelompokKegiatan) clone(db *gorm.DB) tKelompokKegiatan {
	t.tKelompokKegiatanDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tKelompokKegiatan) replaceDB(db *gorm.DB) tKelompokKegiatan {
	t.tKelompokKegiatanDo.ReplaceDB(db)
	return t
}

type tKelompokKegiatanDo struct{ gen.DO }

type ITKelompokKegiatanDo interface {
	gen.SubQuery
	Debug() ITKelompokKegiatanDo
	WithContext(ctx context.Context) ITKelompokKegiatanDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITKelompokKegiatanDo
	WriteDB() ITKelompokKegiatanDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITKelompokKegiatanDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITKelompokKegiatanDo
	Not(conds ...gen.Condition) ITKelompokKegiatanDo
	Or(conds ...gen.Condition) ITKelompokKegiatanDo
	Select(conds ...field.Expr) ITKelompokKegiatanDo
	Where(conds ...gen.Condition) ITKelompokKegiatanDo
	Order(conds ...field.Expr) ITKelompokKegiatanDo
	Distinct(cols ...field.Expr) ITKelompokKegiatanDo
	Omit(cols ...field.Expr) ITKelompokKegiatanDo
	Join(table schema.Tabler, on ...field.Expr) ITKelompokKegiatanDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITKelompokKegiatanDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITKelompokKegiatanDo
	Group(cols ...field.Expr) ITKelompokKegiatanDo
	Having(conds ...gen.Condition) ITKelompokKegiatanDo
	Limit(limit int) ITKelompokKegiatanDo
	Offset(offset int) ITKelompokKegiatanDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITKelompokKegiatanDo
	Unscoped() ITKelompokKegiatanDo
	Create(values ...*model.TKelompokKegiatan) error
	CreateInBatches(values []*model.TKelompokKegiatan, batchSize int) error
	Save(values ...*model.TKelompokKegiatan) error
	First() (*model.TKelompokKegiatan, error)
	Take() (*model.TKelompokKegiatan, error)
	Last() (*model.TKelompokKegiatan, error)
	Find() ([]*model.TKelompokKegiatan, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TKelompokKegiatan, err error)
	FindInBatches(result *[]*model.TKelompokKegiatan, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TKelompokKegiatan) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITKelompokKegiatanDo
	Assign(attrs ...field.AssignExpr) ITKelompokKegiatanDo
	Joins(fields ...field.RelationField) ITKelompokKegiatanDo
	Preload(fields ...field.RelationField) ITKelompokKegiatanDo
	FirstOrInit() (*model.TKelompokKegiatan, error)
	FirstOrCreate() (*model.TKelompokKegiatan, error)
	FindByPage(offset int, limit int) (result []*model.TKelompokKegiatan, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITKelompokKegiatanDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tKelompokKegiatanDo) Debug() ITKelompokKegiatanDo {
	return t.withDO(t.DO.Debug())
}

func (t tKelompokKegiatanDo) WithContext(ctx context.Context) ITKelompokKegiatanDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tKelompokKegiatanDo) ReadDB() ITKelompokKegiatanDo {
	return t.Clauses(dbresolver.Read)
}

func (t tKelompokKegiatanDo) WriteDB() ITKelompokKegiatanDo {
	return t.Clauses(dbresolver.Write)
}

func (t tKelompokKegiatanDo) Session(config *gorm.Session) ITKelompokKegiatanDo {
	return t.withDO(t.DO.Session(config))
}

func (t tKelompokKegiatanDo) Clauses(conds ...clause.Expression) ITKelompokKegiatanDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tKelompokKegiatanDo) Returning(value interface{}, columns ...string) ITKelompokKegiatanDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tKelompokKegiatanDo) Not(conds ...gen.Condition) ITKelompokKegiatanDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tKelompokKegiatanDo) Or(conds ...gen.Condition) ITKelompokKegiatanDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tKelompokKegiatanDo) Select(conds ...field.Expr) ITKelompokKegiatanDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tKelompokKegiatanDo) Where(conds ...gen.Condition) ITKelompokKegiatanDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tKelompokKegiatanDo) Order(conds ...field.Expr) ITKelompokKegiatanDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tKelompokKegiatanDo) Distinct(cols ...field.Expr) ITKelompokKegiatanDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tKelompokKegiatanDo) Omit(cols ...field.Expr) ITKelompokKegiatanDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tKelompokKegiatanDo) Join(table schema.Tabler, on ...field.Expr) ITKelompokKegiatanDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tKelompokKegiatanDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITKelompokKegiatanDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tKelompokKegiatanDo) RightJoin(table schema.Tabler, on ...field.Expr) ITKelompokKegiatanDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tKelompokKegiatanDo) Group(cols ...field.Expr) ITKelompokKegiatanDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tKelompokKegiatanDo) Having(conds ...gen.Condition) ITKelompokKegiatanDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tKelompokKegiatanDo) Limit(limit int) ITKelompokKegiatanDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tKelompokKegiatanDo) Offset(offset int) ITKelompokKegiatanDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tKelompokKegiatanDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITKelompokKegiatanDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tKelompokKegiatanDo) Unscoped() ITKelompokKegiatanDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tKelompokKegiatanDo) Create(values ...*model.TKelompokKegiatan) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tKelompokKegiatanDo) CreateInBatches(values []*model.TKelompokKegiatan, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tKelompokKegiatanDo) Save(values ...*model.TKelompokKegiatan) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tKelompokKegiatanDo) First() (*model.TKelompokKegiatan, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKelompokKegiatan), nil
	}
}

func (t tKelompokKegiatanDo) Take() (*model.TKelompokKegiatan, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKelompokKegiatan), nil
	}
}

func (t tKelompokKegiatanDo) Last() (*model.TKelompokKegiatan, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKelompokKegiatan), nil
	}
}

func (t tKelompokKegiatanDo) Find() ([]*model.TKelompokKegiatan, error) {
	result, err := t.DO.Find()
	return result.([]*model.TKelompokKegiatan), err
}

func (t tKelompokKegiatanDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TKelompokKegiatan, err error) {
	buf := make([]*model.TKelompokKegiatan, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tKelompokKegiatanDo) FindInBatches(result *[]*model.TKelompokKegiatan, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tKelompokKegiatanDo) Attrs(attrs ...field.AssignExpr) ITKelompokKegiatanDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tKelompokKegiatanDo) Assign(attrs ...field.AssignExpr) ITKelompokKegiatanDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tKelompokKegiatanDo) Joins(fields ...field.RelationField) ITKelompokKegiatanDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tKelompokKegiatanDo) Preload(fields ...field.RelationField) ITKelompokKegiatanDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tKelompokKegiatanDo) FirstOrInit() (*model.TKelompokKegiatan, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKelompokKegiatan), nil
	}
}

func (t tKelompokKegiatanDo) FirstOrCreate() (*model.TKelompokKegiatan, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKelompokKegiatan), nil
	}
}

func (t tKelompokKegiatanDo) FindByPage(offset int, limit int) (result []*model.TKelompokKegiatan, count int64, err error) {
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

func (t tKelompokKegiatanDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tKelompokKegiatanDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tKelompokKegiatanDo) Delete(models ...*model.TKelompokKegiatan) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tKelompokKegiatanDo) withDO(do gen.Dao) *tKelompokKegiatanDo {
	t.DO = *do.(*gen.DO)
	return t
}
