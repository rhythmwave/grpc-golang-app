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

func newTPengirimanAlatPromosi(db *gorm.DB, opts ...gen.DOOption) tPengirimanAlatPromosi {
	_tPengirimanAlatPromosi := tPengirimanAlatPromosi{}

	_tPengirimanAlatPromosi.tPengirimanAlatPromosiDo.UseDB(db, opts...)
	_tPengirimanAlatPromosi.tPengirimanAlatPromosiDo.UseModel(&model.TPengirimanAlatPromosi{})

	tableName := _tPengirimanAlatPromosi.tPengirimanAlatPromosiDo.TableName()
	_tPengirimanAlatPromosi.ALL = field.NewAsterisk(tableName)
	_tPengirimanAlatPromosi.CIDAlatPromosi = field.NewInt32(tableName, "c_id_alat_promosi")
	_tPengirimanAlatPromosi.CIDPenanda = field.NewInt32(tableName, "c_id_penanda")
	_tPengirimanAlatPromosi.CJumlahKirim = field.NewInt32(tableName, "c_jumlah_kirim")
	_tPengirimanAlatPromosi.CTanggalKirim = field.NewTime(tableName, "c_tanggal_kirim")
	_tPengirimanAlatPromosi.CUpdater = field.NewString(tableName, "c_updater")
	_tPengirimanAlatPromosi.CLastUpdate = field.NewTime(tableName, "c_last_update")

	_tPengirimanAlatPromosi.fillFieldMap()

	return _tPengirimanAlatPromosi
}

type tPengirimanAlatPromosi struct {
	tPengirimanAlatPromosiDo

	ALL            field.Asterisk
	CIDAlatPromosi field.Int32
	CIDPenanda     field.Int32
	CJumlahKirim   field.Int32
	CTanggalKirim  field.Time
	CUpdater       field.String
	CLastUpdate    field.Time

	fieldMap map[string]field.Expr
}

func (t tPengirimanAlatPromosi) Table(newTableName string) *tPengirimanAlatPromosi {
	t.tPengirimanAlatPromosiDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tPengirimanAlatPromosi) As(alias string) *tPengirimanAlatPromosi {
	t.tPengirimanAlatPromosiDo.DO = *(t.tPengirimanAlatPromosiDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tPengirimanAlatPromosi) updateTableName(table string) *tPengirimanAlatPromosi {
	t.ALL = field.NewAsterisk(table)
	t.CIDAlatPromosi = field.NewInt32(table, "c_id_alat_promosi")
	t.CIDPenanda = field.NewInt32(table, "c_id_penanda")
	t.CJumlahKirim = field.NewInt32(table, "c_jumlah_kirim")
	t.CTanggalKirim = field.NewTime(table, "c_tanggal_kirim")
	t.CUpdater = field.NewString(table, "c_updater")
	t.CLastUpdate = field.NewTime(table, "c_last_update")

	t.fillFieldMap()

	return t
}

func (t *tPengirimanAlatPromosi) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tPengirimanAlatPromosi) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 6)
	t.fieldMap["c_id_alat_promosi"] = t.CIDAlatPromosi
	t.fieldMap["c_id_penanda"] = t.CIDPenanda
	t.fieldMap["c_jumlah_kirim"] = t.CJumlahKirim
	t.fieldMap["c_tanggal_kirim"] = t.CTanggalKirim
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_last_update"] = t.CLastUpdate
}

func (t tPengirimanAlatPromosi) clone(db *gorm.DB) tPengirimanAlatPromosi {
	t.tPengirimanAlatPromosiDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tPengirimanAlatPromosi) replaceDB(db *gorm.DB) tPengirimanAlatPromosi {
	t.tPengirimanAlatPromosiDo.ReplaceDB(db)
	return t
}

type tPengirimanAlatPromosiDo struct{ gen.DO }

type ITPengirimanAlatPromosiDo interface {
	gen.SubQuery
	Debug() ITPengirimanAlatPromosiDo
	WithContext(ctx context.Context) ITPengirimanAlatPromosiDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITPengirimanAlatPromosiDo
	WriteDB() ITPengirimanAlatPromosiDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITPengirimanAlatPromosiDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITPengirimanAlatPromosiDo
	Not(conds ...gen.Condition) ITPengirimanAlatPromosiDo
	Or(conds ...gen.Condition) ITPengirimanAlatPromosiDo
	Select(conds ...field.Expr) ITPengirimanAlatPromosiDo
	Where(conds ...gen.Condition) ITPengirimanAlatPromosiDo
	Order(conds ...field.Expr) ITPengirimanAlatPromosiDo
	Distinct(cols ...field.Expr) ITPengirimanAlatPromosiDo
	Omit(cols ...field.Expr) ITPengirimanAlatPromosiDo
	Join(table schema.Tabler, on ...field.Expr) ITPengirimanAlatPromosiDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITPengirimanAlatPromosiDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITPengirimanAlatPromosiDo
	Group(cols ...field.Expr) ITPengirimanAlatPromosiDo
	Having(conds ...gen.Condition) ITPengirimanAlatPromosiDo
	Limit(limit int) ITPengirimanAlatPromosiDo
	Offset(offset int) ITPengirimanAlatPromosiDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITPengirimanAlatPromosiDo
	Unscoped() ITPengirimanAlatPromosiDo
	Create(values ...*model.TPengirimanAlatPromosi) error
	CreateInBatches(values []*model.TPengirimanAlatPromosi, batchSize int) error
	Save(values ...*model.TPengirimanAlatPromosi) error
	First() (*model.TPengirimanAlatPromosi, error)
	Take() (*model.TPengirimanAlatPromosi, error)
	Last() (*model.TPengirimanAlatPromosi, error)
	Find() ([]*model.TPengirimanAlatPromosi, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TPengirimanAlatPromosi, err error)
	FindInBatches(result *[]*model.TPengirimanAlatPromosi, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TPengirimanAlatPromosi) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITPengirimanAlatPromosiDo
	Assign(attrs ...field.AssignExpr) ITPengirimanAlatPromosiDo
	Joins(fields ...field.RelationField) ITPengirimanAlatPromosiDo
	Preload(fields ...field.RelationField) ITPengirimanAlatPromosiDo
	FirstOrInit() (*model.TPengirimanAlatPromosi, error)
	FirstOrCreate() (*model.TPengirimanAlatPromosi, error)
	FindByPage(offset int, limit int) (result []*model.TPengirimanAlatPromosi, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITPengirimanAlatPromosiDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tPengirimanAlatPromosiDo) Debug() ITPengirimanAlatPromosiDo {
	return t.withDO(t.DO.Debug())
}

func (t tPengirimanAlatPromosiDo) WithContext(ctx context.Context) ITPengirimanAlatPromosiDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tPengirimanAlatPromosiDo) ReadDB() ITPengirimanAlatPromosiDo {
	return t.Clauses(dbresolver.Read)
}

func (t tPengirimanAlatPromosiDo) WriteDB() ITPengirimanAlatPromosiDo {
	return t.Clauses(dbresolver.Write)
}

func (t tPengirimanAlatPromosiDo) Session(config *gorm.Session) ITPengirimanAlatPromosiDo {
	return t.withDO(t.DO.Session(config))
}

func (t tPengirimanAlatPromosiDo) Clauses(conds ...clause.Expression) ITPengirimanAlatPromosiDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tPengirimanAlatPromosiDo) Returning(value interface{}, columns ...string) ITPengirimanAlatPromosiDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tPengirimanAlatPromosiDo) Not(conds ...gen.Condition) ITPengirimanAlatPromosiDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tPengirimanAlatPromosiDo) Or(conds ...gen.Condition) ITPengirimanAlatPromosiDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tPengirimanAlatPromosiDo) Select(conds ...field.Expr) ITPengirimanAlatPromosiDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tPengirimanAlatPromosiDo) Where(conds ...gen.Condition) ITPengirimanAlatPromosiDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tPengirimanAlatPromosiDo) Order(conds ...field.Expr) ITPengirimanAlatPromosiDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tPengirimanAlatPromosiDo) Distinct(cols ...field.Expr) ITPengirimanAlatPromosiDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tPengirimanAlatPromosiDo) Omit(cols ...field.Expr) ITPengirimanAlatPromosiDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tPengirimanAlatPromosiDo) Join(table schema.Tabler, on ...field.Expr) ITPengirimanAlatPromosiDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tPengirimanAlatPromosiDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITPengirimanAlatPromosiDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tPengirimanAlatPromosiDo) RightJoin(table schema.Tabler, on ...field.Expr) ITPengirimanAlatPromosiDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tPengirimanAlatPromosiDo) Group(cols ...field.Expr) ITPengirimanAlatPromosiDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tPengirimanAlatPromosiDo) Having(conds ...gen.Condition) ITPengirimanAlatPromosiDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tPengirimanAlatPromosiDo) Limit(limit int) ITPengirimanAlatPromosiDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tPengirimanAlatPromosiDo) Offset(offset int) ITPengirimanAlatPromosiDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tPengirimanAlatPromosiDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITPengirimanAlatPromosiDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tPengirimanAlatPromosiDo) Unscoped() ITPengirimanAlatPromosiDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tPengirimanAlatPromosiDo) Create(values ...*model.TPengirimanAlatPromosi) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tPengirimanAlatPromosiDo) CreateInBatches(values []*model.TPengirimanAlatPromosi, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tPengirimanAlatPromosiDo) Save(values ...*model.TPengirimanAlatPromosi) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tPengirimanAlatPromosiDo) First() (*model.TPengirimanAlatPromosi, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPengirimanAlatPromosi), nil
	}
}

func (t tPengirimanAlatPromosiDo) Take() (*model.TPengirimanAlatPromosi, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPengirimanAlatPromosi), nil
	}
}

func (t tPengirimanAlatPromosiDo) Last() (*model.TPengirimanAlatPromosi, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPengirimanAlatPromosi), nil
	}
}

func (t tPengirimanAlatPromosiDo) Find() ([]*model.TPengirimanAlatPromosi, error) {
	result, err := t.DO.Find()
	return result.([]*model.TPengirimanAlatPromosi), err
}

func (t tPengirimanAlatPromosiDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TPengirimanAlatPromosi, err error) {
	buf := make([]*model.TPengirimanAlatPromosi, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tPengirimanAlatPromosiDo) FindInBatches(result *[]*model.TPengirimanAlatPromosi, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tPengirimanAlatPromosiDo) Attrs(attrs ...field.AssignExpr) ITPengirimanAlatPromosiDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tPengirimanAlatPromosiDo) Assign(attrs ...field.AssignExpr) ITPengirimanAlatPromosiDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tPengirimanAlatPromosiDo) Joins(fields ...field.RelationField) ITPengirimanAlatPromosiDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tPengirimanAlatPromosiDo) Preload(fields ...field.RelationField) ITPengirimanAlatPromosiDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tPengirimanAlatPromosiDo) FirstOrInit() (*model.TPengirimanAlatPromosi, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPengirimanAlatPromosi), nil
	}
}

func (t tPengirimanAlatPromosiDo) FirstOrCreate() (*model.TPengirimanAlatPromosi, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPengirimanAlatPromosi), nil
	}
}

func (t tPengirimanAlatPromosiDo) FindByPage(offset int, limit int) (result []*model.TPengirimanAlatPromosi, count int64, err error) {
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

func (t tPengirimanAlatPromosiDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tPengirimanAlatPromosiDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tPengirimanAlatPromosiDo) Delete(models ...*model.TPengirimanAlatPromosi) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tPengirimanAlatPromosiDo) withDO(do gen.Dao) *tPengirimanAlatPromosiDo {
	t.DO = *do.(*gen.DO)
	return t
}
