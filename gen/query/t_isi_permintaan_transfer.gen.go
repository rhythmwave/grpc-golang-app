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

func newTIsiPermintaanTransfer(db *gorm.DB, opts ...gen.DOOption) tIsiPermintaanTransfer {
	_tIsiPermintaanTransfer := tIsiPermintaanTransfer{}

	_tIsiPermintaanTransfer.tIsiPermintaanTransferDo.UseDB(db, opts...)
	_tIsiPermintaanTransfer.tIsiPermintaanTransferDo.UseModel(&model.TIsiPermintaanTransfer{})

	tableName := _tIsiPermintaanTransfer.tIsiPermintaanTransferDo.TableName()
	_tIsiPermintaanTransfer.ALL = field.NewAsterisk(tableName)
	_tIsiPermintaanTransfer.CIDPermintaanTransfer = field.NewString(tableName, "c_id_permintaan_transfer")
	_tIsiPermintaanTransfer.CKodeSplit = field.NewString(tableName, "c_kode_split")
	_tIsiPermintaanTransfer.CStatus = field.NewString(tableName, "c_status")
	_tIsiPermintaanTransfer.CUpdater = field.NewString(tableName, "c_updater")
	_tIsiPermintaanTransfer.CLastUpdate = field.NewTime(tableName, "c_last_update")

	_tIsiPermintaanTransfer.fillFieldMap()

	return _tIsiPermintaanTransfer
}

type tIsiPermintaanTransfer struct {
	tIsiPermintaanTransferDo

	ALL                   field.Asterisk
	CIDPermintaanTransfer field.String
	CKodeSplit            field.String
	CStatus               field.String
	CUpdater              field.String
	CLastUpdate           field.Time

	fieldMap map[string]field.Expr
}

func (t tIsiPermintaanTransfer) Table(newTableName string) *tIsiPermintaanTransfer {
	t.tIsiPermintaanTransferDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tIsiPermintaanTransfer) As(alias string) *tIsiPermintaanTransfer {
	t.tIsiPermintaanTransferDo.DO = *(t.tIsiPermintaanTransferDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tIsiPermintaanTransfer) updateTableName(table string) *tIsiPermintaanTransfer {
	t.ALL = field.NewAsterisk(table)
	t.CIDPermintaanTransfer = field.NewString(table, "c_id_permintaan_transfer")
	t.CKodeSplit = field.NewString(table, "c_kode_split")
	t.CStatus = field.NewString(table, "c_status")
	t.CUpdater = field.NewString(table, "c_updater")
	t.CLastUpdate = field.NewTime(table, "c_last_update")

	t.fillFieldMap()

	return t
}

func (t *tIsiPermintaanTransfer) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tIsiPermintaanTransfer) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 5)
	t.fieldMap["c_id_permintaan_transfer"] = t.CIDPermintaanTransfer
	t.fieldMap["c_kode_split"] = t.CKodeSplit
	t.fieldMap["c_status"] = t.CStatus
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_last_update"] = t.CLastUpdate
}

func (t tIsiPermintaanTransfer) clone(db *gorm.DB) tIsiPermintaanTransfer {
	t.tIsiPermintaanTransferDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tIsiPermintaanTransfer) replaceDB(db *gorm.DB) tIsiPermintaanTransfer {
	t.tIsiPermintaanTransferDo.ReplaceDB(db)
	return t
}

type tIsiPermintaanTransferDo struct{ gen.DO }

type ITIsiPermintaanTransferDo interface {
	gen.SubQuery
	Debug() ITIsiPermintaanTransferDo
	WithContext(ctx context.Context) ITIsiPermintaanTransferDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITIsiPermintaanTransferDo
	WriteDB() ITIsiPermintaanTransferDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITIsiPermintaanTransferDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITIsiPermintaanTransferDo
	Not(conds ...gen.Condition) ITIsiPermintaanTransferDo
	Or(conds ...gen.Condition) ITIsiPermintaanTransferDo
	Select(conds ...field.Expr) ITIsiPermintaanTransferDo
	Where(conds ...gen.Condition) ITIsiPermintaanTransferDo
	Order(conds ...field.Expr) ITIsiPermintaanTransferDo
	Distinct(cols ...field.Expr) ITIsiPermintaanTransferDo
	Omit(cols ...field.Expr) ITIsiPermintaanTransferDo
	Join(table schema.Tabler, on ...field.Expr) ITIsiPermintaanTransferDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITIsiPermintaanTransferDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITIsiPermintaanTransferDo
	Group(cols ...field.Expr) ITIsiPermintaanTransferDo
	Having(conds ...gen.Condition) ITIsiPermintaanTransferDo
	Limit(limit int) ITIsiPermintaanTransferDo
	Offset(offset int) ITIsiPermintaanTransferDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITIsiPermintaanTransferDo
	Unscoped() ITIsiPermintaanTransferDo
	Create(values ...*model.TIsiPermintaanTransfer) error
	CreateInBatches(values []*model.TIsiPermintaanTransfer, batchSize int) error
	Save(values ...*model.TIsiPermintaanTransfer) error
	First() (*model.TIsiPermintaanTransfer, error)
	Take() (*model.TIsiPermintaanTransfer, error)
	Last() (*model.TIsiPermintaanTransfer, error)
	Find() ([]*model.TIsiPermintaanTransfer, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TIsiPermintaanTransfer, err error)
	FindInBatches(result *[]*model.TIsiPermintaanTransfer, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TIsiPermintaanTransfer) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITIsiPermintaanTransferDo
	Assign(attrs ...field.AssignExpr) ITIsiPermintaanTransferDo
	Joins(fields ...field.RelationField) ITIsiPermintaanTransferDo
	Preload(fields ...field.RelationField) ITIsiPermintaanTransferDo
	FirstOrInit() (*model.TIsiPermintaanTransfer, error)
	FirstOrCreate() (*model.TIsiPermintaanTransfer, error)
	FindByPage(offset int, limit int) (result []*model.TIsiPermintaanTransfer, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITIsiPermintaanTransferDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tIsiPermintaanTransferDo) Debug() ITIsiPermintaanTransferDo {
	return t.withDO(t.DO.Debug())
}

func (t tIsiPermintaanTransferDo) WithContext(ctx context.Context) ITIsiPermintaanTransferDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tIsiPermintaanTransferDo) ReadDB() ITIsiPermintaanTransferDo {
	return t.Clauses(dbresolver.Read)
}

func (t tIsiPermintaanTransferDo) WriteDB() ITIsiPermintaanTransferDo {
	return t.Clauses(dbresolver.Write)
}

func (t tIsiPermintaanTransferDo) Session(config *gorm.Session) ITIsiPermintaanTransferDo {
	return t.withDO(t.DO.Session(config))
}

func (t tIsiPermintaanTransferDo) Clauses(conds ...clause.Expression) ITIsiPermintaanTransferDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tIsiPermintaanTransferDo) Returning(value interface{}, columns ...string) ITIsiPermintaanTransferDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tIsiPermintaanTransferDo) Not(conds ...gen.Condition) ITIsiPermintaanTransferDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tIsiPermintaanTransferDo) Or(conds ...gen.Condition) ITIsiPermintaanTransferDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tIsiPermintaanTransferDo) Select(conds ...field.Expr) ITIsiPermintaanTransferDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tIsiPermintaanTransferDo) Where(conds ...gen.Condition) ITIsiPermintaanTransferDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tIsiPermintaanTransferDo) Order(conds ...field.Expr) ITIsiPermintaanTransferDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tIsiPermintaanTransferDo) Distinct(cols ...field.Expr) ITIsiPermintaanTransferDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tIsiPermintaanTransferDo) Omit(cols ...field.Expr) ITIsiPermintaanTransferDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tIsiPermintaanTransferDo) Join(table schema.Tabler, on ...field.Expr) ITIsiPermintaanTransferDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tIsiPermintaanTransferDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITIsiPermintaanTransferDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tIsiPermintaanTransferDo) RightJoin(table schema.Tabler, on ...field.Expr) ITIsiPermintaanTransferDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tIsiPermintaanTransferDo) Group(cols ...field.Expr) ITIsiPermintaanTransferDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tIsiPermintaanTransferDo) Having(conds ...gen.Condition) ITIsiPermintaanTransferDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tIsiPermintaanTransferDo) Limit(limit int) ITIsiPermintaanTransferDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tIsiPermintaanTransferDo) Offset(offset int) ITIsiPermintaanTransferDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tIsiPermintaanTransferDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITIsiPermintaanTransferDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tIsiPermintaanTransferDo) Unscoped() ITIsiPermintaanTransferDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tIsiPermintaanTransferDo) Create(values ...*model.TIsiPermintaanTransfer) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tIsiPermintaanTransferDo) CreateInBatches(values []*model.TIsiPermintaanTransfer, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tIsiPermintaanTransferDo) Save(values ...*model.TIsiPermintaanTransfer) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tIsiPermintaanTransferDo) First() (*model.TIsiPermintaanTransfer, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TIsiPermintaanTransfer), nil
	}
}

func (t tIsiPermintaanTransferDo) Take() (*model.TIsiPermintaanTransfer, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TIsiPermintaanTransfer), nil
	}
}

func (t tIsiPermintaanTransferDo) Last() (*model.TIsiPermintaanTransfer, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TIsiPermintaanTransfer), nil
	}
}

func (t tIsiPermintaanTransferDo) Find() ([]*model.TIsiPermintaanTransfer, error) {
	result, err := t.DO.Find()
	return result.([]*model.TIsiPermintaanTransfer), err
}

func (t tIsiPermintaanTransferDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TIsiPermintaanTransfer, err error) {
	buf := make([]*model.TIsiPermintaanTransfer, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tIsiPermintaanTransferDo) FindInBatches(result *[]*model.TIsiPermintaanTransfer, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tIsiPermintaanTransferDo) Attrs(attrs ...field.AssignExpr) ITIsiPermintaanTransferDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tIsiPermintaanTransferDo) Assign(attrs ...field.AssignExpr) ITIsiPermintaanTransferDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tIsiPermintaanTransferDo) Joins(fields ...field.RelationField) ITIsiPermintaanTransferDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tIsiPermintaanTransferDo) Preload(fields ...field.RelationField) ITIsiPermintaanTransferDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tIsiPermintaanTransferDo) FirstOrInit() (*model.TIsiPermintaanTransfer, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TIsiPermintaanTransfer), nil
	}
}

func (t tIsiPermintaanTransferDo) FirstOrCreate() (*model.TIsiPermintaanTransfer, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TIsiPermintaanTransfer), nil
	}
}

func (t tIsiPermintaanTransferDo) FindByPage(offset int, limit int) (result []*model.TIsiPermintaanTransfer, count int64, err error) {
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

func (t tIsiPermintaanTransferDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tIsiPermintaanTransferDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tIsiPermintaanTransferDo) Delete(models ...*model.TIsiPermintaanTransfer) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tIsiPermintaanTransferDo) withDO(do gen.Dao) *tIsiPermintaanTransferDo {
	t.DO = *do.(*gen.DO)
	return t
}
