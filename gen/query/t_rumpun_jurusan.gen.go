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

func newTRumpunJurusan(db *gorm.DB, opts ...gen.DOOption) tRumpunJurusan {
	_tRumpunJurusan := tRumpunJurusan{}

	_tRumpunJurusan.tRumpunJurusanDo.UseDB(db, opts...)
	_tRumpunJurusan.tRumpunJurusanDo.UseModel(&model.TRumpunJurusan{})

	tableName := _tRumpunJurusan.tRumpunJurusanDo.TableName()
	_tRumpunJurusan.ALL = field.NewAsterisk(tableName)
	_tRumpunJurusan.CKodeRumpunJurusan = field.NewString(tableName, "c_kode_rumpun_jurusan")
	_tRumpunJurusan.CDeskripsi = field.NewString(tableName, "c_deskripsi")
	_tRumpunJurusan.CUpline = field.NewString(tableName, "c_upline")
	_tRumpunJurusan.CKeterangan = field.NewString(tableName, "c_keterangan")
	_tRumpunJurusan.CStatus = field.NewString(tableName, "c_status")
	_tRumpunJurusan.CUpdater = field.NewString(tableName, "c_updater")
	_tRumpunJurusan.CCreatedAt = field.NewTime(tableName, "c_created_at")
	_tRumpunJurusan.CLastUpdate = field.NewTime(tableName, "c_last_update")

	_tRumpunJurusan.fillFieldMap()

	return _tRumpunJurusan
}

type tRumpunJurusan struct {
	tRumpunJurusanDo

	ALL                field.Asterisk
	CKodeRumpunJurusan field.String
	CDeskripsi         field.String
	CUpline            field.String
	CKeterangan        field.String
	CStatus            field.String
	CUpdater           field.String
	CCreatedAt         field.Time
	CLastUpdate        field.Time

	fieldMap map[string]field.Expr
}

func (t tRumpunJurusan) Table(newTableName string) *tRumpunJurusan {
	t.tRumpunJurusanDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tRumpunJurusan) As(alias string) *tRumpunJurusan {
	t.tRumpunJurusanDo.DO = *(t.tRumpunJurusanDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tRumpunJurusan) updateTableName(table string) *tRumpunJurusan {
	t.ALL = field.NewAsterisk(table)
	t.CKodeRumpunJurusan = field.NewString(table, "c_kode_rumpun_jurusan")
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

func (t *tRumpunJurusan) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tRumpunJurusan) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 8)
	t.fieldMap["c_kode_rumpun_jurusan"] = t.CKodeRumpunJurusan
	t.fieldMap["c_deskripsi"] = t.CDeskripsi
	t.fieldMap["c_upline"] = t.CUpline
	t.fieldMap["c_keterangan"] = t.CKeterangan
	t.fieldMap["c_status"] = t.CStatus
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_created_at"] = t.CCreatedAt
	t.fieldMap["c_last_update"] = t.CLastUpdate
}

func (t tRumpunJurusan) clone(db *gorm.DB) tRumpunJurusan {
	t.tRumpunJurusanDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tRumpunJurusan) replaceDB(db *gorm.DB) tRumpunJurusan {
	t.tRumpunJurusanDo.ReplaceDB(db)
	return t
}

type tRumpunJurusanDo struct{ gen.DO }

type ITRumpunJurusanDo interface {
	gen.SubQuery
	Debug() ITRumpunJurusanDo
	WithContext(ctx context.Context) ITRumpunJurusanDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITRumpunJurusanDo
	WriteDB() ITRumpunJurusanDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITRumpunJurusanDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITRumpunJurusanDo
	Not(conds ...gen.Condition) ITRumpunJurusanDo
	Or(conds ...gen.Condition) ITRumpunJurusanDo
	Select(conds ...field.Expr) ITRumpunJurusanDo
	Where(conds ...gen.Condition) ITRumpunJurusanDo
	Order(conds ...field.Expr) ITRumpunJurusanDo
	Distinct(cols ...field.Expr) ITRumpunJurusanDo
	Omit(cols ...field.Expr) ITRumpunJurusanDo
	Join(table schema.Tabler, on ...field.Expr) ITRumpunJurusanDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITRumpunJurusanDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITRumpunJurusanDo
	Group(cols ...field.Expr) ITRumpunJurusanDo
	Having(conds ...gen.Condition) ITRumpunJurusanDo
	Limit(limit int) ITRumpunJurusanDo
	Offset(offset int) ITRumpunJurusanDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITRumpunJurusanDo
	Unscoped() ITRumpunJurusanDo
	Create(values ...*model.TRumpunJurusan) error
	CreateInBatches(values []*model.TRumpunJurusan, batchSize int) error
	Save(values ...*model.TRumpunJurusan) error
	First() (*model.TRumpunJurusan, error)
	Take() (*model.TRumpunJurusan, error)
	Last() (*model.TRumpunJurusan, error)
	Find() ([]*model.TRumpunJurusan, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TRumpunJurusan, err error)
	FindInBatches(result *[]*model.TRumpunJurusan, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TRumpunJurusan) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITRumpunJurusanDo
	Assign(attrs ...field.AssignExpr) ITRumpunJurusanDo
	Joins(fields ...field.RelationField) ITRumpunJurusanDo
	Preload(fields ...field.RelationField) ITRumpunJurusanDo
	FirstOrInit() (*model.TRumpunJurusan, error)
	FirstOrCreate() (*model.TRumpunJurusan, error)
	FindByPage(offset int, limit int) (result []*model.TRumpunJurusan, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITRumpunJurusanDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tRumpunJurusanDo) Debug() ITRumpunJurusanDo {
	return t.withDO(t.DO.Debug())
}

func (t tRumpunJurusanDo) WithContext(ctx context.Context) ITRumpunJurusanDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tRumpunJurusanDo) ReadDB() ITRumpunJurusanDo {
	return t.Clauses(dbresolver.Read)
}

func (t tRumpunJurusanDo) WriteDB() ITRumpunJurusanDo {
	return t.Clauses(dbresolver.Write)
}

func (t tRumpunJurusanDo) Session(config *gorm.Session) ITRumpunJurusanDo {
	return t.withDO(t.DO.Session(config))
}

func (t tRumpunJurusanDo) Clauses(conds ...clause.Expression) ITRumpunJurusanDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tRumpunJurusanDo) Returning(value interface{}, columns ...string) ITRumpunJurusanDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tRumpunJurusanDo) Not(conds ...gen.Condition) ITRumpunJurusanDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tRumpunJurusanDo) Or(conds ...gen.Condition) ITRumpunJurusanDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tRumpunJurusanDo) Select(conds ...field.Expr) ITRumpunJurusanDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tRumpunJurusanDo) Where(conds ...gen.Condition) ITRumpunJurusanDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tRumpunJurusanDo) Order(conds ...field.Expr) ITRumpunJurusanDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tRumpunJurusanDo) Distinct(cols ...field.Expr) ITRumpunJurusanDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tRumpunJurusanDo) Omit(cols ...field.Expr) ITRumpunJurusanDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tRumpunJurusanDo) Join(table schema.Tabler, on ...field.Expr) ITRumpunJurusanDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tRumpunJurusanDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITRumpunJurusanDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tRumpunJurusanDo) RightJoin(table schema.Tabler, on ...field.Expr) ITRumpunJurusanDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tRumpunJurusanDo) Group(cols ...field.Expr) ITRumpunJurusanDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tRumpunJurusanDo) Having(conds ...gen.Condition) ITRumpunJurusanDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tRumpunJurusanDo) Limit(limit int) ITRumpunJurusanDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tRumpunJurusanDo) Offset(offset int) ITRumpunJurusanDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tRumpunJurusanDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITRumpunJurusanDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tRumpunJurusanDo) Unscoped() ITRumpunJurusanDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tRumpunJurusanDo) Create(values ...*model.TRumpunJurusan) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tRumpunJurusanDo) CreateInBatches(values []*model.TRumpunJurusan, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tRumpunJurusanDo) Save(values ...*model.TRumpunJurusan) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tRumpunJurusanDo) First() (*model.TRumpunJurusan, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TRumpunJurusan), nil
	}
}

func (t tRumpunJurusanDo) Take() (*model.TRumpunJurusan, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TRumpunJurusan), nil
	}
}

func (t tRumpunJurusanDo) Last() (*model.TRumpunJurusan, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TRumpunJurusan), nil
	}
}

func (t tRumpunJurusanDo) Find() ([]*model.TRumpunJurusan, error) {
	result, err := t.DO.Find()
	return result.([]*model.TRumpunJurusan), err
}

func (t tRumpunJurusanDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TRumpunJurusan, err error) {
	buf := make([]*model.TRumpunJurusan, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tRumpunJurusanDo) FindInBatches(result *[]*model.TRumpunJurusan, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tRumpunJurusanDo) Attrs(attrs ...field.AssignExpr) ITRumpunJurusanDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tRumpunJurusanDo) Assign(attrs ...field.AssignExpr) ITRumpunJurusanDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tRumpunJurusanDo) Joins(fields ...field.RelationField) ITRumpunJurusanDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tRumpunJurusanDo) Preload(fields ...field.RelationField) ITRumpunJurusanDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tRumpunJurusanDo) FirstOrInit() (*model.TRumpunJurusan, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TRumpunJurusan), nil
	}
}

func (t tRumpunJurusanDo) FirstOrCreate() (*model.TRumpunJurusan, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TRumpunJurusan), nil
	}
}

func (t tRumpunJurusanDo) FindByPage(offset int, limit int) (result []*model.TRumpunJurusan, count int64, err error) {
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

func (t tRumpunJurusanDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tRumpunJurusanDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tRumpunJurusanDo) Delete(models ...*model.TRumpunJurusan) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tRumpunJurusanDo) withDO(do gen.Dao) *tRumpunJurusanDo {
	t.DO = *do.(*gen.DO)
	return t
}
