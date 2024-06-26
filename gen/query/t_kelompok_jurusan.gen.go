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

func newTKelompokJurusan(db *gorm.DB, opts ...gen.DOOption) tKelompokJurusan {
	_tKelompokJurusan := tKelompokJurusan{}

	_tKelompokJurusan.tKelompokJurusanDo.UseDB(db, opts...)
	_tKelompokJurusan.tKelompokJurusanDo.UseModel(&model.TKelompokJurusan{})

	tableName := _tKelompokJurusan.tKelompokJurusanDo.TableName()
	_tKelompokJurusan.ALL = field.NewAsterisk(tableName)
	_tKelompokJurusan.CKodeKelompokJurusan = field.NewString(tableName, "c_kode_kelompok_jurusan")
	_tKelompokJurusan.CDeskripsi = field.NewString(tableName, "c_deskripsi")
	_tKelompokJurusan.CUpline = field.NewString(tableName, "c_upline")
	_tKelompokJurusan.CKeterangan = field.NewString(tableName, "c_keterangan")
	_tKelompokJurusan.CStatus = field.NewString(tableName, "c_status")
	_tKelompokJurusan.CUpdater = field.NewString(tableName, "c_updater")
	_tKelompokJurusan.CCreatedAt = field.NewTime(tableName, "c_created_at")
	_tKelompokJurusan.CLastUpdate = field.NewTime(tableName, "c_last_update")

	_tKelompokJurusan.fillFieldMap()

	return _tKelompokJurusan
}

type tKelompokJurusan struct {
	tKelompokJurusanDo

	ALL                  field.Asterisk
	CKodeKelompokJurusan field.String
	CDeskripsi           field.String
	CUpline              field.String
	CKeterangan          field.String
	CStatus              field.String
	CUpdater             field.String
	CCreatedAt           field.Time
	CLastUpdate          field.Time

	fieldMap map[string]field.Expr
}

func (t tKelompokJurusan) Table(newTableName string) *tKelompokJurusan {
	t.tKelompokJurusanDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tKelompokJurusan) As(alias string) *tKelompokJurusan {
	t.tKelompokJurusanDo.DO = *(t.tKelompokJurusanDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tKelompokJurusan) updateTableName(table string) *tKelompokJurusan {
	t.ALL = field.NewAsterisk(table)
	t.CKodeKelompokJurusan = field.NewString(table, "c_kode_kelompok_jurusan")
	t.CDeskripsi = field.NewString(table, "c_deskripsi")
	t.CUpline = field.NewString(table, "c_upline")
	t.CKeterangan = field.NewString(table, "c_keterangan")
	t.CStatus = field.NewString(table, "c_status")
	t.CUpdater = field.NewString(table, "c_updater")
	t.CCreatedAt = field.NewTime(table, "c_created_at")
	t.CLastUpdate = field.NewTime(table, "c_last_update")

	t.fillFieldMap()

	return t
}

func (t *tKelompokJurusan) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tKelompokJurusan) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 8)
	t.fieldMap["c_kode_kelompok_jurusan"] = t.CKodeKelompokJurusan
	t.fieldMap["c_deskripsi"] = t.CDeskripsi
	t.fieldMap["c_upline"] = t.CUpline
	t.fieldMap["c_keterangan"] = t.CKeterangan
	t.fieldMap["c_status"] = t.CStatus
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_created_at"] = t.CCreatedAt
	t.fieldMap["c_last_update"] = t.CLastUpdate
}

func (t tKelompokJurusan) clone(db *gorm.DB) tKelompokJurusan {
	t.tKelompokJurusanDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tKelompokJurusan) replaceDB(db *gorm.DB) tKelompokJurusan {
	t.tKelompokJurusanDo.ReplaceDB(db)
	return t
}

type tKelompokJurusanDo struct{ gen.DO }

type ITKelompokJurusanDo interface {
	gen.SubQuery
	Debug() ITKelompokJurusanDo
	WithContext(ctx context.Context) ITKelompokJurusanDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITKelompokJurusanDo
	WriteDB() ITKelompokJurusanDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITKelompokJurusanDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITKelompokJurusanDo
	Not(conds ...gen.Condition) ITKelompokJurusanDo
	Or(conds ...gen.Condition) ITKelompokJurusanDo
	Select(conds ...field.Expr) ITKelompokJurusanDo
	Where(conds ...gen.Condition) ITKelompokJurusanDo
	Order(conds ...field.Expr) ITKelompokJurusanDo
	Distinct(cols ...field.Expr) ITKelompokJurusanDo
	Omit(cols ...field.Expr) ITKelompokJurusanDo
	Join(table schema.Tabler, on ...field.Expr) ITKelompokJurusanDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITKelompokJurusanDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITKelompokJurusanDo
	Group(cols ...field.Expr) ITKelompokJurusanDo
	Having(conds ...gen.Condition) ITKelompokJurusanDo
	Limit(limit int) ITKelompokJurusanDo
	Offset(offset int) ITKelompokJurusanDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITKelompokJurusanDo
	Unscoped() ITKelompokJurusanDo
	Create(values ...*model.TKelompokJurusan) error
	CreateInBatches(values []*model.TKelompokJurusan, batchSize int) error
	Save(values ...*model.TKelompokJurusan) error
	First() (*model.TKelompokJurusan, error)
	Take() (*model.TKelompokJurusan, error)
	Last() (*model.TKelompokJurusan, error)
	Find() ([]*model.TKelompokJurusan, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TKelompokJurusan, err error)
	FindInBatches(result *[]*model.TKelompokJurusan, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TKelompokJurusan) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITKelompokJurusanDo
	Assign(attrs ...field.AssignExpr) ITKelompokJurusanDo
	Joins(fields ...field.RelationField) ITKelompokJurusanDo
	Preload(fields ...field.RelationField) ITKelompokJurusanDo
	FirstOrInit() (*model.TKelompokJurusan, error)
	FirstOrCreate() (*model.TKelompokJurusan, error)
	FindByPage(offset int, limit int) (result []*model.TKelompokJurusan, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITKelompokJurusanDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tKelompokJurusanDo) Debug() ITKelompokJurusanDo {
	return t.withDO(t.DO.Debug())
}

func (t tKelompokJurusanDo) WithContext(ctx context.Context) ITKelompokJurusanDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tKelompokJurusanDo) ReadDB() ITKelompokJurusanDo {
	return t.Clauses(dbresolver.Read)
}

func (t tKelompokJurusanDo) WriteDB() ITKelompokJurusanDo {
	return t.Clauses(dbresolver.Write)
}

func (t tKelompokJurusanDo) Session(config *gorm.Session) ITKelompokJurusanDo {
	return t.withDO(t.DO.Session(config))
}

func (t tKelompokJurusanDo) Clauses(conds ...clause.Expression) ITKelompokJurusanDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tKelompokJurusanDo) Returning(value interface{}, columns ...string) ITKelompokJurusanDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tKelompokJurusanDo) Not(conds ...gen.Condition) ITKelompokJurusanDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tKelompokJurusanDo) Or(conds ...gen.Condition) ITKelompokJurusanDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tKelompokJurusanDo) Select(conds ...field.Expr) ITKelompokJurusanDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tKelompokJurusanDo) Where(conds ...gen.Condition) ITKelompokJurusanDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tKelompokJurusanDo) Order(conds ...field.Expr) ITKelompokJurusanDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tKelompokJurusanDo) Distinct(cols ...field.Expr) ITKelompokJurusanDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tKelompokJurusanDo) Omit(cols ...field.Expr) ITKelompokJurusanDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tKelompokJurusanDo) Join(table schema.Tabler, on ...field.Expr) ITKelompokJurusanDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tKelompokJurusanDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITKelompokJurusanDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tKelompokJurusanDo) RightJoin(table schema.Tabler, on ...field.Expr) ITKelompokJurusanDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tKelompokJurusanDo) Group(cols ...field.Expr) ITKelompokJurusanDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tKelompokJurusanDo) Having(conds ...gen.Condition) ITKelompokJurusanDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tKelompokJurusanDo) Limit(limit int) ITKelompokJurusanDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tKelompokJurusanDo) Offset(offset int) ITKelompokJurusanDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tKelompokJurusanDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITKelompokJurusanDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tKelompokJurusanDo) Unscoped() ITKelompokJurusanDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tKelompokJurusanDo) Create(values ...*model.TKelompokJurusan) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tKelompokJurusanDo) CreateInBatches(values []*model.TKelompokJurusan, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tKelompokJurusanDo) Save(values ...*model.TKelompokJurusan) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tKelompokJurusanDo) First() (*model.TKelompokJurusan, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKelompokJurusan), nil
	}
}

func (t tKelompokJurusanDo) Take() (*model.TKelompokJurusan, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKelompokJurusan), nil
	}
}

func (t tKelompokJurusanDo) Last() (*model.TKelompokJurusan, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKelompokJurusan), nil
	}
}

func (t tKelompokJurusanDo) Find() ([]*model.TKelompokJurusan, error) {
	result, err := t.DO.Find()
	return result.([]*model.TKelompokJurusan), err
}

func (t tKelompokJurusanDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TKelompokJurusan, err error) {
	buf := make([]*model.TKelompokJurusan, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tKelompokJurusanDo) FindInBatches(result *[]*model.TKelompokJurusan, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tKelompokJurusanDo) Attrs(attrs ...field.AssignExpr) ITKelompokJurusanDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tKelompokJurusanDo) Assign(attrs ...field.AssignExpr) ITKelompokJurusanDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tKelompokJurusanDo) Joins(fields ...field.RelationField) ITKelompokJurusanDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tKelompokJurusanDo) Preload(fields ...field.RelationField) ITKelompokJurusanDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tKelompokJurusanDo) FirstOrInit() (*model.TKelompokJurusan, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKelompokJurusan), nil
	}
}

func (t tKelompokJurusanDo) FirstOrCreate() (*model.TKelompokJurusan, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TKelompokJurusan), nil
	}
}

func (t tKelompokJurusanDo) FindByPage(offset int, limit int) (result []*model.TKelompokJurusan, count int64, err error) {
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

func (t tKelompokJurusanDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tKelompokJurusanDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tKelompokJurusanDo) Delete(models ...*model.TKelompokJurusan) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tKelompokJurusanDo) withDO(do gen.Dao) *tKelompokJurusanDo {
	t.DO = *do.(*gen.DO)
	return t
}
