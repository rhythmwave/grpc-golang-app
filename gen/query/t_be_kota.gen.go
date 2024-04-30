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

func newTBeKotum(db *gorm.DB, opts ...gen.DOOption) tBeKotum {
	_tBeKotum := tBeKotum{}

	_tBeKotum.tBeKotumDo.UseDB(db, opts...)
	_tBeKotum.tBeKotumDo.UseModel(&model.TBeKotum{})

	tableName := _tBeKotum.tBeKotumDo.TableName()
	_tBeKotum.ALL = field.NewAsterisk(tableName)
	_tBeKotum.CIDBeKota = field.NewInt32(tableName, "c_id_be_kota")
	_tBeKotum.CIDPenanda = field.NewInt32(tableName, "c_id_penanda")
	_tBeKotum.CIDIsiBe = field.NewInt32(tableName, "c_id_isi_be")
	_tBeKotum.CKondisiBisnis = field.NewString(tableName, "c_kondisi_bisnis")
	_tBeKotum.CIsFollowUp = field.NewString(tableName, "c_is_follow_up")
	_tBeKotum.CTahunAjaran = field.NewString(tableName, "c_tahun_ajaran")
	_tBeKotum.CStatus = field.NewString(tableName, "c_status")
	_tBeKotum.CUpdater = field.NewString(tableName, "c_updater")
	_tBeKotum.CLastUpdate = field.NewTime(tableName, "c_last_update")

	_tBeKotum.fillFieldMap()

	return _tBeKotum
}

type tBeKotum struct {
	tBeKotumDo

	ALL            field.Asterisk
	CIDBeKota      field.Int32
	CIDPenanda     field.Int32
	CIDIsiBe       field.Int32
	CKondisiBisnis field.String
	CIsFollowUp    field.String
	CTahunAjaran   field.String
	CStatus        field.String
	CUpdater       field.String
	CLastUpdate    field.Time

	fieldMap map[string]field.Expr
}

func (t tBeKotum) Table(newTableName string) *tBeKotum {
	t.tBeKotumDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tBeKotum) As(alias string) *tBeKotum {
	t.tBeKotumDo.DO = *(t.tBeKotumDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tBeKotum) updateTableName(table string) *tBeKotum {
	t.ALL = field.NewAsterisk(table)
	t.CIDBeKota = field.NewInt32(table, "c_id_be_kota")
	t.CIDPenanda = field.NewInt32(table, "c_id_penanda")
	t.CIDIsiBe = field.NewInt32(table, "c_id_isi_be")
	t.CKondisiBisnis = field.NewString(table, "c_kondisi_bisnis")
	t.CIsFollowUp = field.NewString(table, "c_is_follow_up")
	t.CTahunAjaran = field.NewString(table, "c_tahun_ajaran")
	t.CStatus = field.NewString(table, "c_status")
	t.CUpdater = field.NewString(table, "c_updater")
	t.CLastUpdate = field.NewTime(table, "c_last_update")

	t.fillFieldMap()

	return t
}

func (t *tBeKotum) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tBeKotum) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 9)
	t.fieldMap["c_id_be_kota"] = t.CIDBeKota
	t.fieldMap["c_id_penanda"] = t.CIDPenanda
	t.fieldMap["c_id_isi_be"] = t.CIDIsiBe
	t.fieldMap["c_kondisi_bisnis"] = t.CKondisiBisnis
	t.fieldMap["c_is_follow_up"] = t.CIsFollowUp
	t.fieldMap["c_tahun_ajaran"] = t.CTahunAjaran
	t.fieldMap["c_status"] = t.CStatus
	t.fieldMap["c_updater"] = t.CUpdater
	t.fieldMap["c_last_update"] = t.CLastUpdate
}

func (t tBeKotum) clone(db *gorm.DB) tBeKotum {
	t.tBeKotumDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tBeKotum) replaceDB(db *gorm.DB) tBeKotum {
	t.tBeKotumDo.ReplaceDB(db)
	return t
}

type tBeKotumDo struct{ gen.DO }

type ITBeKotumDo interface {
	gen.SubQuery
	Debug() ITBeKotumDo
	WithContext(ctx context.Context) ITBeKotumDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITBeKotumDo
	WriteDB() ITBeKotumDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITBeKotumDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITBeKotumDo
	Not(conds ...gen.Condition) ITBeKotumDo
	Or(conds ...gen.Condition) ITBeKotumDo
	Select(conds ...field.Expr) ITBeKotumDo
	Where(conds ...gen.Condition) ITBeKotumDo
	Order(conds ...field.Expr) ITBeKotumDo
	Distinct(cols ...field.Expr) ITBeKotumDo
	Omit(cols ...field.Expr) ITBeKotumDo
	Join(table schema.Tabler, on ...field.Expr) ITBeKotumDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITBeKotumDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITBeKotumDo
	Group(cols ...field.Expr) ITBeKotumDo
	Having(conds ...gen.Condition) ITBeKotumDo
	Limit(limit int) ITBeKotumDo
	Offset(offset int) ITBeKotumDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITBeKotumDo
	Unscoped() ITBeKotumDo
	Create(values ...*model.TBeKotum) error
	CreateInBatches(values []*model.TBeKotum, batchSize int) error
	Save(values ...*model.TBeKotum) error
	First() (*model.TBeKotum, error)
	Take() (*model.TBeKotum, error)
	Last() (*model.TBeKotum, error)
	Find() ([]*model.TBeKotum, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TBeKotum, err error)
	FindInBatches(result *[]*model.TBeKotum, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TBeKotum) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITBeKotumDo
	Assign(attrs ...field.AssignExpr) ITBeKotumDo
	Joins(fields ...field.RelationField) ITBeKotumDo
	Preload(fields ...field.RelationField) ITBeKotumDo
	FirstOrInit() (*model.TBeKotum, error)
	FirstOrCreate() (*model.TBeKotum, error)
	FindByPage(offset int, limit int) (result []*model.TBeKotum, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITBeKotumDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tBeKotumDo) Debug() ITBeKotumDo {
	return t.withDO(t.DO.Debug())
}

func (t tBeKotumDo) WithContext(ctx context.Context) ITBeKotumDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tBeKotumDo) ReadDB() ITBeKotumDo {
	return t.Clauses(dbresolver.Read)
}

func (t tBeKotumDo) WriteDB() ITBeKotumDo {
	return t.Clauses(dbresolver.Write)
}

func (t tBeKotumDo) Session(config *gorm.Session) ITBeKotumDo {
	return t.withDO(t.DO.Session(config))
}

func (t tBeKotumDo) Clauses(conds ...clause.Expression) ITBeKotumDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tBeKotumDo) Returning(value interface{}, columns ...string) ITBeKotumDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tBeKotumDo) Not(conds ...gen.Condition) ITBeKotumDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tBeKotumDo) Or(conds ...gen.Condition) ITBeKotumDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tBeKotumDo) Select(conds ...field.Expr) ITBeKotumDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tBeKotumDo) Where(conds ...gen.Condition) ITBeKotumDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tBeKotumDo) Order(conds ...field.Expr) ITBeKotumDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tBeKotumDo) Distinct(cols ...field.Expr) ITBeKotumDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tBeKotumDo) Omit(cols ...field.Expr) ITBeKotumDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tBeKotumDo) Join(table schema.Tabler, on ...field.Expr) ITBeKotumDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tBeKotumDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITBeKotumDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tBeKotumDo) RightJoin(table schema.Tabler, on ...field.Expr) ITBeKotumDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tBeKotumDo) Group(cols ...field.Expr) ITBeKotumDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tBeKotumDo) Having(conds ...gen.Condition) ITBeKotumDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tBeKotumDo) Limit(limit int) ITBeKotumDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tBeKotumDo) Offset(offset int) ITBeKotumDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tBeKotumDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITBeKotumDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tBeKotumDo) Unscoped() ITBeKotumDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tBeKotumDo) Create(values ...*model.TBeKotum) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tBeKotumDo) CreateInBatches(values []*model.TBeKotum, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tBeKotumDo) Save(values ...*model.TBeKotum) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tBeKotumDo) First() (*model.TBeKotum, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TBeKotum), nil
	}
}

func (t tBeKotumDo) Take() (*model.TBeKotum, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TBeKotum), nil
	}
}

func (t tBeKotumDo) Last() (*model.TBeKotum, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TBeKotum), nil
	}
}

func (t tBeKotumDo) Find() ([]*model.TBeKotum, error) {
	result, err := t.DO.Find()
	return result.([]*model.TBeKotum), err
}

func (t tBeKotumDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TBeKotum, err error) {
	buf := make([]*model.TBeKotum, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tBeKotumDo) FindInBatches(result *[]*model.TBeKotum, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tBeKotumDo) Attrs(attrs ...field.AssignExpr) ITBeKotumDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tBeKotumDo) Assign(attrs ...field.AssignExpr) ITBeKotumDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tBeKotumDo) Joins(fields ...field.RelationField) ITBeKotumDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tBeKotumDo) Preload(fields ...field.RelationField) ITBeKotumDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tBeKotumDo) FirstOrInit() (*model.TBeKotum, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TBeKotum), nil
	}
}

func (t tBeKotumDo) FirstOrCreate() (*model.TBeKotum, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TBeKotum), nil
	}
}

func (t tBeKotumDo) FindByPage(offset int, limit int) (result []*model.TBeKotum, count int64, err error) {
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

func (t tBeKotumDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tBeKotumDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tBeKotumDo) Delete(models ...*model.TBeKotum) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tBeKotumDo) withDO(do gen.Dao) *tBeKotumDo {
	t.DO = *do.(*gen.DO)
	return t
}