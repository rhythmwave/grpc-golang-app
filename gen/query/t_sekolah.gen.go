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

func newTSekolah(db *gorm.DB, opts ...gen.DOOption) tSekolah {
	_tSekolah := tSekolah{}

	_tSekolah.tSekolahDo.UseDB(db, opts...)
	_tSekolah.tSekolahDo.UseModel(&model.TSekolah{})

	tableName := _tSekolah.tSekolahDo.TableName()
	_tSekolah.ALL = field.NewAsterisk(tableName)
	_tSekolah.CIDSekolah = field.NewInt32(tableName, "c_id_sekolah")
	_tSekolah.CNamaSekolah = field.NewString(tableName, "c_nama_sekolah")
	_tSekolah.CIDDistrict = field.NewInt32(tableName, "c_id_district")
	_tSekolah.CStatus = field.NewString(tableName, "c_status")
	_tSekolah.CUpdater = field.NewString(tableName, "c_updater")
	_tSekolah.CLastUpdate = field.NewTime(tableName, "c_last_update")
	_tSekolah.CJenjangPendidikan = field.NewString(tableName, "c_jenjang_pendidikan")
	_tSekolah.CAlamatSekolah = field.NewString(tableName, "c_alamat_sekolah")

	_tSekolah.fillFieldMap()

	return _tSekolah
}

type tSekolah struct {
	tSekolahDo

	ALL                field.Asterisk
	CIDSekolah         field.Int32
	CNamaSekolah       field.String // Ref : t_Idn_District
	CIDDistrict        field.Int32
	CStatus            field.String
	CUpdater           field.String
	CLastUpdate        field.Time
	CJenjangPendidikan field.String
	CAlamatSekolah     field.String

	fieldMap map[string]field.Expr
}

func (t tSekolah) Table(newTableName string) *tSekolah {
	t.tSekolahDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tSekolah) As(alias string) *tSekolah {
	t.tSekolahDo.DO = *(t.tSekolahDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tSekolah) updateTableName(table string) *tSekolah {
	t.ALL = field.NewAsterisk(table)
	t.CIDSekolah = field.NewInt32(table, "c_id_sekolah")
	t.CNamaSekolah = field.NewString(table, "c_nama_sekolah")
	t.CIDDistrict = field.NewInt32(table, "c_id_district")
	t.CStatus = field.NewString(table, "c_status")
	t.CUpdater = field.NewString(table, "c_updater")
	t.CLastUpdate = field.NewTime(table, "c_last_update")
	t.CJenjangPendidikan = field.NewString(table, "c_jenjang_pendidikan")
	t.CAlamatSekolah = field.NewString(table, "c_alamat_sekolah")

	t.fillFieldMap()

	return t
}

func (t *tSekolah) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tSekolah) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 8)
	t.fieldMap["c_id_sekolah"] = t.CIDSekolah
	t.fieldMap["c_nama_sekolah"] = t.CNamaSekolah
	t.fieldMap["c_id_district"] = t.CIDDistrict
	t.fieldMap["c_status"] = t.CStatus
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_last_update"] = t.CLastUpdate
	t.fieldMap["c_jenjang_pendidikan"] = t.CJenjangPendidikan
	t.fieldMap["c_alamat_sekolah"] = t.CAlamatSekolah
}

func (t tSekolah) clone(db *gorm.DB) tSekolah {
	t.tSekolahDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tSekolah) replaceDB(db *gorm.DB) tSekolah {
	t.tSekolahDo.ReplaceDB(db)
	return t
}

type tSekolahDo struct{ gen.DO }

type ITSekolahDo interface {
	gen.SubQuery
	Debug() ITSekolahDo
	WithContext(ctx context.Context) ITSekolahDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITSekolahDo
	WriteDB() ITSekolahDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITSekolahDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITSekolahDo
	Not(conds ...gen.Condition) ITSekolahDo
	Or(conds ...gen.Condition) ITSekolahDo
	Select(conds ...field.Expr) ITSekolahDo
	Where(conds ...gen.Condition) ITSekolahDo
	Order(conds ...field.Expr) ITSekolahDo
	Distinct(cols ...field.Expr) ITSekolahDo
	Omit(cols ...field.Expr) ITSekolahDo
	Join(table schema.Tabler, on ...field.Expr) ITSekolahDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITSekolahDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITSekolahDo
	Group(cols ...field.Expr) ITSekolahDo
	Having(conds ...gen.Condition) ITSekolahDo
	Limit(limit int) ITSekolahDo
	Offset(offset int) ITSekolahDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITSekolahDo
	Unscoped() ITSekolahDo
	Create(values ...*model.TSekolah) error
	CreateInBatches(values []*model.TSekolah, batchSize int) error
	Save(values ...*model.TSekolah) error
	First() (*model.TSekolah, error)
	Take() (*model.TSekolah, error)
	Last() (*model.TSekolah, error)
	Find() ([]*model.TSekolah, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TSekolah, err error)
	FindInBatches(result *[]*model.TSekolah, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TSekolah) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITSekolahDo
	Assign(attrs ...field.AssignExpr) ITSekolahDo
	Joins(fields ...field.RelationField) ITSekolahDo
	Preload(fields ...field.RelationField) ITSekolahDo
	FirstOrInit() (*model.TSekolah, error)
	FirstOrCreate() (*model.TSekolah, error)
	FindByPage(offset int, limit int) (result []*model.TSekolah, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITSekolahDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tSekolahDo) Debug() ITSekolahDo {
	return t.withDO(t.DO.Debug())
}

func (t tSekolahDo) WithContext(ctx context.Context) ITSekolahDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tSekolahDo) ReadDB() ITSekolahDo {
	return t.Clauses(dbresolver.Read)
}

func (t tSekolahDo) WriteDB() ITSekolahDo {
	return t.Clauses(dbresolver.Write)
}

func (t tSekolahDo) Session(config *gorm.Session) ITSekolahDo {
	return t.withDO(t.DO.Session(config))
}

func (t tSekolahDo) Clauses(conds ...clause.Expression) ITSekolahDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tSekolahDo) Returning(value interface{}, columns ...string) ITSekolahDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tSekolahDo) Not(conds ...gen.Condition) ITSekolahDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tSekolahDo) Or(conds ...gen.Condition) ITSekolahDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tSekolahDo) Select(conds ...field.Expr) ITSekolahDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tSekolahDo) Where(conds ...gen.Condition) ITSekolahDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tSekolahDo) Order(conds ...field.Expr) ITSekolahDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tSekolahDo) Distinct(cols ...field.Expr) ITSekolahDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tSekolahDo) Omit(cols ...field.Expr) ITSekolahDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tSekolahDo) Join(table schema.Tabler, on ...field.Expr) ITSekolahDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tSekolahDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITSekolahDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tSekolahDo) RightJoin(table schema.Tabler, on ...field.Expr) ITSekolahDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tSekolahDo) Group(cols ...field.Expr) ITSekolahDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tSekolahDo) Having(conds ...gen.Condition) ITSekolahDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tSekolahDo) Limit(limit int) ITSekolahDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tSekolahDo) Offset(offset int) ITSekolahDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tSekolahDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITSekolahDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tSekolahDo) Unscoped() ITSekolahDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tSekolahDo) Create(values ...*model.TSekolah) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tSekolahDo) CreateInBatches(values []*model.TSekolah, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tSekolahDo) Save(values ...*model.TSekolah) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tSekolahDo) First() (*model.TSekolah, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSekolah), nil
	}
}

func (t tSekolahDo) Take() (*model.TSekolah, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSekolah), nil
	}
}

func (t tSekolahDo) Last() (*model.TSekolah, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSekolah), nil
	}
}

func (t tSekolahDo) Find() ([]*model.TSekolah, error) {
	result, err := t.DO.Find()
	return result.([]*model.TSekolah), err
}

func (t tSekolahDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TSekolah, err error) {
	buf := make([]*model.TSekolah, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tSekolahDo) FindInBatches(result *[]*model.TSekolah, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tSekolahDo) Attrs(attrs ...field.AssignExpr) ITSekolahDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tSekolahDo) Assign(attrs ...field.AssignExpr) ITSekolahDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tSekolahDo) Joins(fields ...field.RelationField) ITSekolahDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tSekolahDo) Preload(fields ...field.RelationField) ITSekolahDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tSekolahDo) FirstOrInit() (*model.TSekolah, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSekolah), nil
	}
}

func (t tSekolahDo) FirstOrCreate() (*model.TSekolah, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSekolah), nil
	}
}

func (t tSekolahDo) FindByPage(offset int, limit int) (result []*model.TSekolah, count int64, err error) {
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

func (t tSekolahDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tSekolahDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tSekolahDo) Delete(models ...*model.TSekolah) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tSekolahDo) withDO(do gen.Dao) *tSekolahDo {
	t.DO = *do.(*gen.DO)
	return t
}
