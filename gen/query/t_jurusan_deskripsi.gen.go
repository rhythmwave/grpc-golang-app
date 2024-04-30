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

func newTJurusanDeskripsi(db *gorm.DB, opts ...gen.DOOption) tJurusanDeskripsi {
	_tJurusanDeskripsi := tJurusanDeskripsi{}

	_tJurusanDeskripsi.tJurusanDeskripsiDo.UseDB(db, opts...)
	_tJurusanDeskripsi.tJurusanDeskripsiDo.UseModel(&model.TJurusanDeskripsi{})

	tableName := _tJurusanDeskripsi.tJurusanDeskripsiDo.TableName()
	_tJurusanDeskripsi.ALL = field.NewAsterisk(tableName)
	_tJurusanDeskripsi.CIDJurusan = field.NewInt32(tableName, "c_id_jurusan")
	_tJurusanDeskripsi.CDeskripsi = field.NewString(tableName, "c_deskripsi")
	_tJurusanDeskripsi.CLapanganKerja = field.NewString(tableName, "c_lapangan_kerja")

	_tJurusanDeskripsi.fillFieldMap()

	return _tJurusanDeskripsi
}

type tJurusanDeskripsi struct {
	tJurusanDeskripsiDo

	ALL            field.Asterisk
	CIDJurusan     field.Int32
	CDeskripsi     field.String
	CLapanganKerja field.String

	fieldMap map[string]field.Expr
}

func (t tJurusanDeskripsi) Table(newTableName string) *tJurusanDeskripsi {
	t.tJurusanDeskripsiDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tJurusanDeskripsi) As(alias string) *tJurusanDeskripsi {
	t.tJurusanDeskripsiDo.DO = *(t.tJurusanDeskripsiDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tJurusanDeskripsi) updateTableName(table string) *tJurusanDeskripsi {
	t.ALL = field.NewAsterisk(table)
	t.CIDJurusan = field.NewInt32(table, "c_id_jurusan")
	t.CDeskripsi = field.NewString(table, "c_deskripsi")
	t.CLapanganKerja = field.NewString(table, "c_lapangan_kerja")

	t.fillFieldMap()

	return t
}

func (t *tJurusanDeskripsi) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tJurusanDeskripsi) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 3)
	t.fieldMap["c_id_jurusan"] = t.CIDJurusan
	t.fieldMap["c_deskripsi"] = t.CDeskripsi
	t.fieldMap["c_lapangan_kerja"] = t.CLapanganKerja
}

func (t tJurusanDeskripsi) clone(db *gorm.DB) tJurusanDeskripsi {
	t.tJurusanDeskripsiDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tJurusanDeskripsi) replaceDB(db *gorm.DB) tJurusanDeskripsi {
	t.tJurusanDeskripsiDo.ReplaceDB(db)
	return t
}

type tJurusanDeskripsiDo struct{ gen.DO }

type ITJurusanDeskripsiDo interface {
	gen.SubQuery
	Debug() ITJurusanDeskripsiDo
	WithContext(ctx context.Context) ITJurusanDeskripsiDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITJurusanDeskripsiDo
	WriteDB() ITJurusanDeskripsiDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITJurusanDeskripsiDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITJurusanDeskripsiDo
	Not(conds ...gen.Condition) ITJurusanDeskripsiDo
	Or(conds ...gen.Condition) ITJurusanDeskripsiDo
	Select(conds ...field.Expr) ITJurusanDeskripsiDo
	Where(conds ...gen.Condition) ITJurusanDeskripsiDo
	Order(conds ...field.Expr) ITJurusanDeskripsiDo
	Distinct(cols ...field.Expr) ITJurusanDeskripsiDo
	Omit(cols ...field.Expr) ITJurusanDeskripsiDo
	Join(table schema.Tabler, on ...field.Expr) ITJurusanDeskripsiDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITJurusanDeskripsiDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITJurusanDeskripsiDo
	Group(cols ...field.Expr) ITJurusanDeskripsiDo
	Having(conds ...gen.Condition) ITJurusanDeskripsiDo
	Limit(limit int) ITJurusanDeskripsiDo
	Offset(offset int) ITJurusanDeskripsiDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITJurusanDeskripsiDo
	Unscoped() ITJurusanDeskripsiDo
	Create(values ...*model.TJurusanDeskripsi) error
	CreateInBatches(values []*model.TJurusanDeskripsi, batchSize int) error
	Save(values ...*model.TJurusanDeskripsi) error
	First() (*model.TJurusanDeskripsi, error)
	Take() (*model.TJurusanDeskripsi, error)
	Last() (*model.TJurusanDeskripsi, error)
	Find() ([]*model.TJurusanDeskripsi, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TJurusanDeskripsi, err error)
	FindInBatches(result *[]*model.TJurusanDeskripsi, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TJurusanDeskripsi) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITJurusanDeskripsiDo
	Assign(attrs ...field.AssignExpr) ITJurusanDeskripsiDo
	Joins(fields ...field.RelationField) ITJurusanDeskripsiDo
	Preload(fields ...field.RelationField) ITJurusanDeskripsiDo
	FirstOrInit() (*model.TJurusanDeskripsi, error)
	FirstOrCreate() (*model.TJurusanDeskripsi, error)
	FindByPage(offset int, limit int) (result []*model.TJurusanDeskripsi, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITJurusanDeskripsiDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tJurusanDeskripsiDo) Debug() ITJurusanDeskripsiDo {
	return t.withDO(t.DO.Debug())
}

func (t tJurusanDeskripsiDo) WithContext(ctx context.Context) ITJurusanDeskripsiDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tJurusanDeskripsiDo) ReadDB() ITJurusanDeskripsiDo {
	return t.Clauses(dbresolver.Read)
}

func (t tJurusanDeskripsiDo) WriteDB() ITJurusanDeskripsiDo {
	return t.Clauses(dbresolver.Write)
}

func (t tJurusanDeskripsiDo) Session(config *gorm.Session) ITJurusanDeskripsiDo {
	return t.withDO(t.DO.Session(config))
}

func (t tJurusanDeskripsiDo) Clauses(conds ...clause.Expression) ITJurusanDeskripsiDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tJurusanDeskripsiDo) Returning(value interface{}, columns ...string) ITJurusanDeskripsiDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tJurusanDeskripsiDo) Not(conds ...gen.Condition) ITJurusanDeskripsiDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tJurusanDeskripsiDo) Or(conds ...gen.Condition) ITJurusanDeskripsiDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tJurusanDeskripsiDo) Select(conds ...field.Expr) ITJurusanDeskripsiDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tJurusanDeskripsiDo) Where(conds ...gen.Condition) ITJurusanDeskripsiDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tJurusanDeskripsiDo) Order(conds ...field.Expr) ITJurusanDeskripsiDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tJurusanDeskripsiDo) Distinct(cols ...field.Expr) ITJurusanDeskripsiDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tJurusanDeskripsiDo) Omit(cols ...field.Expr) ITJurusanDeskripsiDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tJurusanDeskripsiDo) Join(table schema.Tabler, on ...field.Expr) ITJurusanDeskripsiDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tJurusanDeskripsiDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITJurusanDeskripsiDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tJurusanDeskripsiDo) RightJoin(table schema.Tabler, on ...field.Expr) ITJurusanDeskripsiDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tJurusanDeskripsiDo) Group(cols ...field.Expr) ITJurusanDeskripsiDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tJurusanDeskripsiDo) Having(conds ...gen.Condition) ITJurusanDeskripsiDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tJurusanDeskripsiDo) Limit(limit int) ITJurusanDeskripsiDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tJurusanDeskripsiDo) Offset(offset int) ITJurusanDeskripsiDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tJurusanDeskripsiDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITJurusanDeskripsiDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tJurusanDeskripsiDo) Unscoped() ITJurusanDeskripsiDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tJurusanDeskripsiDo) Create(values ...*model.TJurusanDeskripsi) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tJurusanDeskripsiDo) CreateInBatches(values []*model.TJurusanDeskripsi, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tJurusanDeskripsiDo) Save(values ...*model.TJurusanDeskripsi) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tJurusanDeskripsiDo) First() (*model.TJurusanDeskripsi, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TJurusanDeskripsi), nil
	}
}

func (t tJurusanDeskripsiDo) Take() (*model.TJurusanDeskripsi, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TJurusanDeskripsi), nil
	}
}

func (t tJurusanDeskripsiDo) Last() (*model.TJurusanDeskripsi, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TJurusanDeskripsi), nil
	}
}

func (t tJurusanDeskripsiDo) Find() ([]*model.TJurusanDeskripsi, error) {
	result, err := t.DO.Find()
	return result.([]*model.TJurusanDeskripsi), err
}

func (t tJurusanDeskripsiDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TJurusanDeskripsi, err error) {
	buf := make([]*model.TJurusanDeskripsi, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tJurusanDeskripsiDo) FindInBatches(result *[]*model.TJurusanDeskripsi, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tJurusanDeskripsiDo) Attrs(attrs ...field.AssignExpr) ITJurusanDeskripsiDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tJurusanDeskripsiDo) Assign(attrs ...field.AssignExpr) ITJurusanDeskripsiDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tJurusanDeskripsiDo) Joins(fields ...field.RelationField) ITJurusanDeskripsiDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tJurusanDeskripsiDo) Preload(fields ...field.RelationField) ITJurusanDeskripsiDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tJurusanDeskripsiDo) FirstOrInit() (*model.TJurusanDeskripsi, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TJurusanDeskripsi), nil
	}
}

func (t tJurusanDeskripsiDo) FirstOrCreate() (*model.TJurusanDeskripsi, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TJurusanDeskripsi), nil
	}
}

func (t tJurusanDeskripsiDo) FindByPage(offset int, limit int) (result []*model.TJurusanDeskripsi, count int64, err error) {
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

func (t tJurusanDeskripsiDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tJurusanDeskripsiDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tJurusanDeskripsiDo) Delete(models ...*model.TJurusanDeskripsi) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tJurusanDeskripsiDo) withDO(do gen.Dao) *tJurusanDeskripsiDo {
	t.DO = *do.(*gen.DO)
	return t
}
