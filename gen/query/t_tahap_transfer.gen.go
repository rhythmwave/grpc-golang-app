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

func newTTahapTransfer(db *gorm.DB, opts ...gen.DOOption) tTahapTransfer {
	_tTahapTransfer := tTahapTransfer{}

	_tTahapTransfer.tTahapTransferDo.UseDB(db, opts...)
	_tTahapTransfer.tTahapTransferDo.UseModel(&model.TTahapTransfer{})

	tableName := _tTahapTransfer.tTahapTransferDo.TableName()
	_tTahapTransfer.ALL = field.NewAsterisk(tableName)
	_tTahapTransfer.CKodePengajuanAnggaran = field.NewString(tableName, "c_kode_pengajuan_anggaran")
	_tTahapTransfer.CTahapKe = field.NewInt16(tableName, "c_tahap_ke")
	_tTahapTransfer.CJumlahTransfer = field.NewInt32(tableName, "c_jumlah_transfer")
	_tTahapTransfer.CTanggalTransfer = field.NewTime(tableName, "c_tanggal_transfer")
	_tTahapTransfer.CIsTransfer = field.NewString(tableName, "c_is_transfer")
	_tTahapTransfer.CUpdater = field.NewString(tableName, "c_updater")
	_tTahapTransfer.CLastUpdate = field.NewTime(tableName, "c_last_update")

	_tTahapTransfer.fillFieldMap()

	return _tTahapTransfer
}

type tTahapTransfer struct {
	tTahapTransferDo

	ALL                    field.Asterisk
	CKodePengajuanAnggaran field.String
	CTahapKe               field.Int16
	CJumlahTransfer        field.Int32
	CTanggalTransfer       field.Time
	CIsTransfer            field.String
	CUpdater               field.String
	CLastUpdate            field.Time

	fieldMap map[string]field.Expr
}

func (t tTahapTransfer) Table(newTableName string) *tTahapTransfer {
	t.tTahapTransferDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tTahapTransfer) As(alias string) *tTahapTransfer {
	t.tTahapTransferDo.DO = *(t.tTahapTransferDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tTahapTransfer) updateTableName(table string) *tTahapTransfer {
	t.ALL = field.NewAsterisk(table)
	t.CKodePengajuanAnggaran = field.NewString(table, "c_kode_pengajuan_anggaran")
	t.CTahapKe = field.NewInt16(table, "c_tahap_ke")
	t.CJumlahTransfer = field.NewInt32(table, "c_jumlah_transfer")
	t.CTanggalTransfer = field.NewTime(table, "c_tanggal_transfer")
	t.CIsTransfer = field.NewString(table, "c_is_transfer")
	t.CUpdater = field.NewString(table, "c_updater")
	t.CLastUpdate = field.NewTime(table, "c_last_update")

	t.fillFieldMap()

	return t
}

func (t *tTahapTransfer) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tTahapTransfer) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 7)
	t.fieldMap["c_kode_pengajuan_anggaran"] = t.CKodePengajuanAnggaran
	t.fieldMap["c_tahap_ke"] = t.CTahapKe
	t.fieldMap["c_jumlah_transfer"] = t.CJumlahTransfer
	t.fieldMap["c_tanggal_transfer"] = t.CTanggalTransfer
	t.fieldMap["c_is_transfer"] = t.CIsTransfer
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_last_update"] = t.CLastUpdate
}

func (t tTahapTransfer) clone(db *gorm.DB) tTahapTransfer {
	t.tTahapTransferDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tTahapTransfer) replaceDB(db *gorm.DB) tTahapTransfer {
	t.tTahapTransferDo.ReplaceDB(db)
	return t
}

type tTahapTransferDo struct{ gen.DO }

type ITTahapTransferDo interface {
	gen.SubQuery
	Debug() ITTahapTransferDo
	WithContext(ctx context.Context) ITTahapTransferDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITTahapTransferDo
	WriteDB() ITTahapTransferDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITTahapTransferDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITTahapTransferDo
	Not(conds ...gen.Condition) ITTahapTransferDo
	Or(conds ...gen.Condition) ITTahapTransferDo
	Select(conds ...field.Expr) ITTahapTransferDo
	Where(conds ...gen.Condition) ITTahapTransferDo
	Order(conds ...field.Expr) ITTahapTransferDo
	Distinct(cols ...field.Expr) ITTahapTransferDo
	Omit(cols ...field.Expr) ITTahapTransferDo
	Join(table schema.Tabler, on ...field.Expr) ITTahapTransferDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITTahapTransferDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITTahapTransferDo
	Group(cols ...field.Expr) ITTahapTransferDo
	Having(conds ...gen.Condition) ITTahapTransferDo
	Limit(limit int) ITTahapTransferDo
	Offset(offset int) ITTahapTransferDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITTahapTransferDo
	Unscoped() ITTahapTransferDo
	Create(values ...*model.TTahapTransfer) error
	CreateInBatches(values []*model.TTahapTransfer, batchSize int) error
	Save(values ...*model.TTahapTransfer) error
	First() (*model.TTahapTransfer, error)
	Take() (*model.TTahapTransfer, error)
	Last() (*model.TTahapTransfer, error)
	Find() ([]*model.TTahapTransfer, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TTahapTransfer, err error)
	FindInBatches(result *[]*model.TTahapTransfer, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TTahapTransfer) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITTahapTransferDo
	Assign(attrs ...field.AssignExpr) ITTahapTransferDo
	Joins(fields ...field.RelationField) ITTahapTransferDo
	Preload(fields ...field.RelationField) ITTahapTransferDo
	FirstOrInit() (*model.TTahapTransfer, error)
	FirstOrCreate() (*model.TTahapTransfer, error)
	FindByPage(offset int, limit int) (result []*model.TTahapTransfer, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITTahapTransferDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tTahapTransferDo) Debug() ITTahapTransferDo {
	return t.withDO(t.DO.Debug())
}

func (t tTahapTransferDo) WithContext(ctx context.Context) ITTahapTransferDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tTahapTransferDo) ReadDB() ITTahapTransferDo {
	return t.Clauses(dbresolver.Read)
}

func (t tTahapTransferDo) WriteDB() ITTahapTransferDo {
	return t.Clauses(dbresolver.Write)
}

func (t tTahapTransferDo) Session(config *gorm.Session) ITTahapTransferDo {
	return t.withDO(t.DO.Session(config))
}

func (t tTahapTransferDo) Clauses(conds ...clause.Expression) ITTahapTransferDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tTahapTransferDo) Returning(value interface{}, columns ...string) ITTahapTransferDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tTahapTransferDo) Not(conds ...gen.Condition) ITTahapTransferDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tTahapTransferDo) Or(conds ...gen.Condition) ITTahapTransferDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tTahapTransferDo) Select(conds ...field.Expr) ITTahapTransferDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tTahapTransferDo) Where(conds ...gen.Condition) ITTahapTransferDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tTahapTransferDo) Order(conds ...field.Expr) ITTahapTransferDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tTahapTransferDo) Distinct(cols ...field.Expr) ITTahapTransferDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tTahapTransferDo) Omit(cols ...field.Expr) ITTahapTransferDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tTahapTransferDo) Join(table schema.Tabler, on ...field.Expr) ITTahapTransferDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tTahapTransferDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITTahapTransferDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tTahapTransferDo) RightJoin(table schema.Tabler, on ...field.Expr) ITTahapTransferDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tTahapTransferDo) Group(cols ...field.Expr) ITTahapTransferDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tTahapTransferDo) Having(conds ...gen.Condition) ITTahapTransferDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tTahapTransferDo) Limit(limit int) ITTahapTransferDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tTahapTransferDo) Offset(offset int) ITTahapTransferDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tTahapTransferDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITTahapTransferDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tTahapTransferDo) Unscoped() ITTahapTransferDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tTahapTransferDo) Create(values ...*model.TTahapTransfer) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tTahapTransferDo) CreateInBatches(values []*model.TTahapTransfer, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tTahapTransferDo) Save(values ...*model.TTahapTransfer) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tTahapTransferDo) First() (*model.TTahapTransfer, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TTahapTransfer), nil
	}
}

func (t tTahapTransferDo) Take() (*model.TTahapTransfer, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TTahapTransfer), nil
	}
}

func (t tTahapTransferDo) Last() (*model.TTahapTransfer, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TTahapTransfer), nil
	}
}

func (t tTahapTransferDo) Find() ([]*model.TTahapTransfer, error) {
	result, err := t.DO.Find()
	return result.([]*model.TTahapTransfer), err
}

func (t tTahapTransferDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TTahapTransfer, err error) {
	buf := make([]*model.TTahapTransfer, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tTahapTransferDo) FindInBatches(result *[]*model.TTahapTransfer, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tTahapTransferDo) Attrs(attrs ...field.AssignExpr) ITTahapTransferDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tTahapTransferDo) Assign(attrs ...field.AssignExpr) ITTahapTransferDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tTahapTransferDo) Joins(fields ...field.RelationField) ITTahapTransferDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tTahapTransferDo) Preload(fields ...field.RelationField) ITTahapTransferDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tTahapTransferDo) FirstOrInit() (*model.TTahapTransfer, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TTahapTransfer), nil
	}
}

func (t tTahapTransferDo) FirstOrCreate() (*model.TTahapTransfer, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TTahapTransfer), nil
	}
}

func (t tTahapTransferDo) FindByPage(offset int, limit int) (result []*model.TTahapTransfer, count int64, err error) {
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

func (t tTahapTransferDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tTahapTransferDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tTahapTransferDo) Delete(models ...*model.TTahapTransfer) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tTahapTransferDo) withDO(do gen.Dao) *tTahapTransferDo {
	t.DO = *do.(*gen.DO)
	return t
}
