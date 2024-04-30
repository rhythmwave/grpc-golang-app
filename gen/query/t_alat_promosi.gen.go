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

func newTAlatPromosi(db *gorm.DB, opts ...gen.DOOption) tAlatPromosi {
	_tAlatPromosi := tAlatPromosi{}

	_tAlatPromosi.tAlatPromosiDo.UseDB(db, opts...)
	_tAlatPromosi.tAlatPromosiDo.UseModel(&model.TAlatPromosi{})

	tableName := _tAlatPromosi.tAlatPromosiDo.TableName()
	_tAlatPromosi.ALL = field.NewAsterisk(tableName)
	_tAlatPromosi.CIDAlat = field.NewInt32(tableName, "c_id_alat")
	_tAlatPromosi.CNamaAlat = field.NewString(tableName, "c_nama_alat")
	_tAlatPromosi.CTingkatKelas = field.NewString(tableName, "c_tingkat_kelas")
	_tAlatPromosi.CSemester = field.NewString(tableName, "c_semester")
	_tAlatPromosi.CUpdater = field.NewString(tableName, "c_updater")
	_tAlatPromosi.CLastUpdate = field.NewTime(tableName, "c_last_update")
	_tAlatPromosi.CIDBidang = field.NewInt16(tableName, "c_id_bidang")

	_tAlatPromosi.fillFieldMap()

	return _tAlatPromosi
}

type tAlatPromosi struct {
	tAlatPromosiDo

	ALL           field.Asterisk
	CIDAlat       field.Int32
	CNamaAlat     field.String
	CTingkatKelas field.String
	CSemester     field.String
	CUpdater      field.String
	CLastUpdate   field.Time
	CIDBidang     field.Int16

	fieldMap map[string]field.Expr
}

func (t tAlatPromosi) Table(newTableName string) *tAlatPromosi {
	t.tAlatPromosiDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tAlatPromosi) As(alias string) *tAlatPromosi {
	t.tAlatPromosiDo.DO = *(t.tAlatPromosiDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tAlatPromosi) updateTableName(table string) *tAlatPromosi {
	t.ALL = field.NewAsterisk(table)
	t.CIDAlat = field.NewInt32(table, "c_id_alat")
	t.CNamaAlat = field.NewString(table, "c_nama_alat")
	t.CTingkatKelas = field.NewString(table, "c_tingkat_kelas")
	t.CSemester = field.NewString(table, "c_semester")
	t.CUpdater = field.NewString(table, "c_updater")
	t.CLastUpdate = field.NewTime(table, "c_last_update")
	t.CIDBidang = field.NewInt16(table, "c_id_bidang")

	t.fillFieldMap()

	return t
}

func (t *tAlatPromosi) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tAlatPromosi) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 7)
	t.fieldMap["c_id_alat"] = t.CIDAlat
	t.fieldMap["c_nama_alat"] = t.CNamaAlat
	t.fieldMap["c_tingkat_kelas"] = t.CTingkatKelas
	t.fieldMap["c_semester"] = t.CSemester
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_last_update"] = t.CLastUpdate
	t.fieldMap["c_id_bidang"] = t.CIDBidang
}

func (t tAlatPromosi) clone(db *gorm.DB) tAlatPromosi {
	t.tAlatPromosiDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tAlatPromosi) replaceDB(db *gorm.DB) tAlatPromosi {
	t.tAlatPromosiDo.ReplaceDB(db)
	return t
}

type tAlatPromosiDo struct{ gen.DO }

type ITAlatPromosiDo interface {
	gen.SubQuery
	Debug() ITAlatPromosiDo
	WithContext(ctx context.Context) ITAlatPromosiDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITAlatPromosiDo
	WriteDB() ITAlatPromosiDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITAlatPromosiDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITAlatPromosiDo
	Not(conds ...gen.Condition) ITAlatPromosiDo
	Or(conds ...gen.Condition) ITAlatPromosiDo
	Select(conds ...field.Expr) ITAlatPromosiDo
	Where(conds ...gen.Condition) ITAlatPromosiDo
	Order(conds ...field.Expr) ITAlatPromosiDo
	Distinct(cols ...field.Expr) ITAlatPromosiDo
	Omit(cols ...field.Expr) ITAlatPromosiDo
	Join(table schema.Tabler, on ...field.Expr) ITAlatPromosiDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITAlatPromosiDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITAlatPromosiDo
	Group(cols ...field.Expr) ITAlatPromosiDo
	Having(conds ...gen.Condition) ITAlatPromosiDo
	Limit(limit int) ITAlatPromosiDo
	Offset(offset int) ITAlatPromosiDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITAlatPromosiDo
	Unscoped() ITAlatPromosiDo
	Create(values ...*model.TAlatPromosi) error
	CreateInBatches(values []*model.TAlatPromosi, batchSize int) error
	Save(values ...*model.TAlatPromosi) error
	First() (*model.TAlatPromosi, error)
	Take() (*model.TAlatPromosi, error)
	Last() (*model.TAlatPromosi, error)
	Find() ([]*model.TAlatPromosi, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TAlatPromosi, err error)
	FindInBatches(result *[]*model.TAlatPromosi, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TAlatPromosi) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITAlatPromosiDo
	Assign(attrs ...field.AssignExpr) ITAlatPromosiDo
	Joins(fields ...field.RelationField) ITAlatPromosiDo
	Preload(fields ...field.RelationField) ITAlatPromosiDo
	FirstOrInit() (*model.TAlatPromosi, error)
	FirstOrCreate() (*model.TAlatPromosi, error)
	FindByPage(offset int, limit int) (result []*model.TAlatPromosi, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITAlatPromosiDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tAlatPromosiDo) Debug() ITAlatPromosiDo {
	return t.withDO(t.DO.Debug())
}

func (t tAlatPromosiDo) WithContext(ctx context.Context) ITAlatPromosiDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tAlatPromosiDo) ReadDB() ITAlatPromosiDo {
	return t.Clauses(dbresolver.Read)
}

func (t tAlatPromosiDo) WriteDB() ITAlatPromosiDo {
	return t.Clauses(dbresolver.Write)
}

func (t tAlatPromosiDo) Session(config *gorm.Session) ITAlatPromosiDo {
	return t.withDO(t.DO.Session(config))
}

func (t tAlatPromosiDo) Clauses(conds ...clause.Expression) ITAlatPromosiDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tAlatPromosiDo) Returning(value interface{}, columns ...string) ITAlatPromosiDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tAlatPromosiDo) Not(conds ...gen.Condition) ITAlatPromosiDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tAlatPromosiDo) Or(conds ...gen.Condition) ITAlatPromosiDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tAlatPromosiDo) Select(conds ...field.Expr) ITAlatPromosiDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tAlatPromosiDo) Where(conds ...gen.Condition) ITAlatPromosiDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tAlatPromosiDo) Order(conds ...field.Expr) ITAlatPromosiDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tAlatPromosiDo) Distinct(cols ...field.Expr) ITAlatPromosiDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tAlatPromosiDo) Omit(cols ...field.Expr) ITAlatPromosiDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tAlatPromosiDo) Join(table schema.Tabler, on ...field.Expr) ITAlatPromosiDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tAlatPromosiDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITAlatPromosiDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tAlatPromosiDo) RightJoin(table schema.Tabler, on ...field.Expr) ITAlatPromosiDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tAlatPromosiDo) Group(cols ...field.Expr) ITAlatPromosiDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tAlatPromosiDo) Having(conds ...gen.Condition) ITAlatPromosiDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tAlatPromosiDo) Limit(limit int) ITAlatPromosiDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tAlatPromosiDo) Offset(offset int) ITAlatPromosiDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tAlatPromosiDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITAlatPromosiDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tAlatPromosiDo) Unscoped() ITAlatPromosiDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tAlatPromosiDo) Create(values ...*model.TAlatPromosi) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tAlatPromosiDo) CreateInBatches(values []*model.TAlatPromosi, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tAlatPromosiDo) Save(values ...*model.TAlatPromosi) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tAlatPromosiDo) First() (*model.TAlatPromosi, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAlatPromosi), nil
	}
}

func (t tAlatPromosiDo) Take() (*model.TAlatPromosi, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAlatPromosi), nil
	}
}

func (t tAlatPromosiDo) Last() (*model.TAlatPromosi, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAlatPromosi), nil
	}
}

func (t tAlatPromosiDo) Find() ([]*model.TAlatPromosi, error) {
	result, err := t.DO.Find()
	return result.([]*model.TAlatPromosi), err
}

func (t tAlatPromosiDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TAlatPromosi, err error) {
	buf := make([]*model.TAlatPromosi, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tAlatPromosiDo) FindInBatches(result *[]*model.TAlatPromosi, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tAlatPromosiDo) Attrs(attrs ...field.AssignExpr) ITAlatPromosiDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tAlatPromosiDo) Assign(attrs ...field.AssignExpr) ITAlatPromosiDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tAlatPromosiDo) Joins(fields ...field.RelationField) ITAlatPromosiDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tAlatPromosiDo) Preload(fields ...field.RelationField) ITAlatPromosiDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tAlatPromosiDo) FirstOrInit() (*model.TAlatPromosi, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAlatPromosi), nil
	}
}

func (t tAlatPromosiDo) FirstOrCreate() (*model.TAlatPromosi, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAlatPromosi), nil
	}
}

func (t tAlatPromosiDo) FindByPage(offset int, limit int) (result []*model.TAlatPromosi, count int64, err error) {
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

func (t tAlatPromosiDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tAlatPromosiDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tAlatPromosiDo) Delete(models ...*model.TAlatPromosi) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tAlatPromosiDo) withDO(do gen.Dao) *tAlatPromosiDo {
	t.DO = *do.(*gen.DO)
	return t
}