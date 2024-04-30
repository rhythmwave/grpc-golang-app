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

func newTInformasiPesertaStrategicMeeting(db *gorm.DB, opts ...gen.DOOption) tInformasiPesertaStrategicMeeting {
	_tInformasiPesertaStrategicMeeting := tInformasiPesertaStrategicMeeting{}

	_tInformasiPesertaStrategicMeeting.tInformasiPesertaStrategicMeetingDo.UseDB(db, opts...)
	_tInformasiPesertaStrategicMeeting.tInformasiPesertaStrategicMeetingDo.UseModel(&model.TInformasiPesertaStrategicMeeting{})

	tableName := _tInformasiPesertaStrategicMeeting.tInformasiPesertaStrategicMeetingDo.TableName()
	_tInformasiPesertaStrategicMeeting.ALL = field.NewAsterisk(tableName)
	_tInformasiPesertaStrategicMeeting.CNik = field.NewString(tableName, "c_nik")
	_tInformasiPesertaStrategicMeeting.CPhoto = field.NewString(tableName, "c_photo")
	_tInformasiPesertaStrategicMeeting.CBernyanyi = field.NewString(tableName, "c_bernyanyi")
	_tInformasiPesertaStrategicMeeting.CRoundTable = field.NewString(tableName, "c_round_table")

	_tInformasiPesertaStrategicMeeting.fillFieldMap()

	return _tInformasiPesertaStrategicMeeting
}

type tInformasiPesertaStrategicMeeting struct {
	tInformasiPesertaStrategicMeetingDo

	ALL         field.Asterisk
	CNik        field.String
	CPhoto      field.String
	CBernyanyi  field.String
	CRoundTable field.String

	fieldMap map[string]field.Expr
}

func (t tInformasiPesertaStrategicMeeting) Table(newTableName string) *tInformasiPesertaStrategicMeeting {
	t.tInformasiPesertaStrategicMeetingDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tInformasiPesertaStrategicMeeting) As(alias string) *tInformasiPesertaStrategicMeeting {
	t.tInformasiPesertaStrategicMeetingDo.DO = *(t.tInformasiPesertaStrategicMeetingDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tInformasiPesertaStrategicMeeting) updateTableName(table string) *tInformasiPesertaStrategicMeeting {
	t.ALL = field.NewAsterisk(table)
	t.CNik = field.NewString(table, "c_nik")
	t.CPhoto = field.NewString(table, "c_photo")
	t.CBernyanyi = field.NewString(table, "c_bernyanyi")
	t.CRoundTable = field.NewString(table, "c_round_table")

	t.fillFieldMap()

	return t
}

func (t *tInformasiPesertaStrategicMeeting) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tInformasiPesertaStrategicMeeting) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 4)
	t.fieldMap["c_nik"] = t.CNik
	t.fieldMap["c_photo"] = t.CPhoto
	t.fieldMap["c_bernyanyi"] = t.CBernyanyi
	t.fieldMap["c_round_table"] = t.CRoundTable
}

func (t tInformasiPesertaStrategicMeeting) clone(db *gorm.DB) tInformasiPesertaStrategicMeeting {
	t.tInformasiPesertaStrategicMeetingDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tInformasiPesertaStrategicMeeting) replaceDB(db *gorm.DB) tInformasiPesertaStrategicMeeting {
	t.tInformasiPesertaStrategicMeetingDo.ReplaceDB(db)
	return t
}

type tInformasiPesertaStrategicMeetingDo struct{ gen.DO }

type ITInformasiPesertaStrategicMeetingDo interface {
	gen.SubQuery
	Debug() ITInformasiPesertaStrategicMeetingDo
	WithContext(ctx context.Context) ITInformasiPesertaStrategicMeetingDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITInformasiPesertaStrategicMeetingDo
	WriteDB() ITInformasiPesertaStrategicMeetingDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITInformasiPesertaStrategicMeetingDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITInformasiPesertaStrategicMeetingDo
	Not(conds ...gen.Condition) ITInformasiPesertaStrategicMeetingDo
	Or(conds ...gen.Condition) ITInformasiPesertaStrategicMeetingDo
	Select(conds ...field.Expr) ITInformasiPesertaStrategicMeetingDo
	Where(conds ...gen.Condition) ITInformasiPesertaStrategicMeetingDo
	Order(conds ...field.Expr) ITInformasiPesertaStrategicMeetingDo
	Distinct(cols ...field.Expr) ITInformasiPesertaStrategicMeetingDo
	Omit(cols ...field.Expr) ITInformasiPesertaStrategicMeetingDo
	Join(table schema.Tabler, on ...field.Expr) ITInformasiPesertaStrategicMeetingDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITInformasiPesertaStrategicMeetingDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITInformasiPesertaStrategicMeetingDo
	Group(cols ...field.Expr) ITInformasiPesertaStrategicMeetingDo
	Having(conds ...gen.Condition) ITInformasiPesertaStrategicMeetingDo
	Limit(limit int) ITInformasiPesertaStrategicMeetingDo
	Offset(offset int) ITInformasiPesertaStrategicMeetingDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITInformasiPesertaStrategicMeetingDo
	Unscoped() ITInformasiPesertaStrategicMeetingDo
	Create(values ...*model.TInformasiPesertaStrategicMeeting) error
	CreateInBatches(values []*model.TInformasiPesertaStrategicMeeting, batchSize int) error
	Save(values ...*model.TInformasiPesertaStrategicMeeting) error
	First() (*model.TInformasiPesertaStrategicMeeting, error)
	Take() (*model.TInformasiPesertaStrategicMeeting, error)
	Last() (*model.TInformasiPesertaStrategicMeeting, error)
	Find() ([]*model.TInformasiPesertaStrategicMeeting, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TInformasiPesertaStrategicMeeting, err error)
	FindInBatches(result *[]*model.TInformasiPesertaStrategicMeeting, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TInformasiPesertaStrategicMeeting) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITInformasiPesertaStrategicMeetingDo
	Assign(attrs ...field.AssignExpr) ITInformasiPesertaStrategicMeetingDo
	Joins(fields ...field.RelationField) ITInformasiPesertaStrategicMeetingDo
	Preload(fields ...field.RelationField) ITInformasiPesertaStrategicMeetingDo
	FirstOrInit() (*model.TInformasiPesertaStrategicMeeting, error)
	FirstOrCreate() (*model.TInformasiPesertaStrategicMeeting, error)
	FindByPage(offset int, limit int) (result []*model.TInformasiPesertaStrategicMeeting, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITInformasiPesertaStrategicMeetingDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tInformasiPesertaStrategicMeetingDo) Debug() ITInformasiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Debug())
}

func (t tInformasiPesertaStrategicMeetingDo) WithContext(ctx context.Context) ITInformasiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tInformasiPesertaStrategicMeetingDo) ReadDB() ITInformasiPesertaStrategicMeetingDo {
	return t.Clauses(dbresolver.Read)
}

func (t tInformasiPesertaStrategicMeetingDo) WriteDB() ITInformasiPesertaStrategicMeetingDo {
	return t.Clauses(dbresolver.Write)
}

func (t tInformasiPesertaStrategicMeetingDo) Session(config *gorm.Session) ITInformasiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Session(config))
}

func (t tInformasiPesertaStrategicMeetingDo) Clauses(conds ...clause.Expression) ITInformasiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tInformasiPesertaStrategicMeetingDo) Returning(value interface{}, columns ...string) ITInformasiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tInformasiPesertaStrategicMeetingDo) Not(conds ...gen.Condition) ITInformasiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tInformasiPesertaStrategicMeetingDo) Or(conds ...gen.Condition) ITInformasiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tInformasiPesertaStrategicMeetingDo) Select(conds ...field.Expr) ITInformasiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tInformasiPesertaStrategicMeetingDo) Where(conds ...gen.Condition) ITInformasiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tInformasiPesertaStrategicMeetingDo) Order(conds ...field.Expr) ITInformasiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tInformasiPesertaStrategicMeetingDo) Distinct(cols ...field.Expr) ITInformasiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tInformasiPesertaStrategicMeetingDo) Omit(cols ...field.Expr) ITInformasiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tInformasiPesertaStrategicMeetingDo) Join(table schema.Tabler, on ...field.Expr) ITInformasiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tInformasiPesertaStrategicMeetingDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITInformasiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tInformasiPesertaStrategicMeetingDo) RightJoin(table schema.Tabler, on ...field.Expr) ITInformasiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tInformasiPesertaStrategicMeetingDo) Group(cols ...field.Expr) ITInformasiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tInformasiPesertaStrategicMeetingDo) Having(conds ...gen.Condition) ITInformasiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tInformasiPesertaStrategicMeetingDo) Limit(limit int) ITInformasiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tInformasiPesertaStrategicMeetingDo) Offset(offset int) ITInformasiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tInformasiPesertaStrategicMeetingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITInformasiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tInformasiPesertaStrategicMeetingDo) Unscoped() ITInformasiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tInformasiPesertaStrategicMeetingDo) Create(values ...*model.TInformasiPesertaStrategicMeeting) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tInformasiPesertaStrategicMeetingDo) CreateInBatches(values []*model.TInformasiPesertaStrategicMeeting, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tInformasiPesertaStrategicMeetingDo) Save(values ...*model.TInformasiPesertaStrategicMeeting) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tInformasiPesertaStrategicMeetingDo) First() (*model.TInformasiPesertaStrategicMeeting, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TInformasiPesertaStrategicMeeting), nil
	}
}

func (t tInformasiPesertaStrategicMeetingDo) Take() (*model.TInformasiPesertaStrategicMeeting, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TInformasiPesertaStrategicMeeting), nil
	}
}

func (t tInformasiPesertaStrategicMeetingDo) Last() (*model.TInformasiPesertaStrategicMeeting, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TInformasiPesertaStrategicMeeting), nil
	}
}

func (t tInformasiPesertaStrategicMeetingDo) Find() ([]*model.TInformasiPesertaStrategicMeeting, error) {
	result, err := t.DO.Find()
	return result.([]*model.TInformasiPesertaStrategicMeeting), err
}

func (t tInformasiPesertaStrategicMeetingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TInformasiPesertaStrategicMeeting, err error) {
	buf := make([]*model.TInformasiPesertaStrategicMeeting, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tInformasiPesertaStrategicMeetingDo) FindInBatches(result *[]*model.TInformasiPesertaStrategicMeeting, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tInformasiPesertaStrategicMeetingDo) Attrs(attrs ...field.AssignExpr) ITInformasiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tInformasiPesertaStrategicMeetingDo) Assign(attrs ...field.AssignExpr) ITInformasiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tInformasiPesertaStrategicMeetingDo) Joins(fields ...field.RelationField) ITInformasiPesertaStrategicMeetingDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tInformasiPesertaStrategicMeetingDo) Preload(fields ...field.RelationField) ITInformasiPesertaStrategicMeetingDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tInformasiPesertaStrategicMeetingDo) FirstOrInit() (*model.TInformasiPesertaStrategicMeeting, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TInformasiPesertaStrategicMeeting), nil
	}
}

func (t tInformasiPesertaStrategicMeetingDo) FirstOrCreate() (*model.TInformasiPesertaStrategicMeeting, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TInformasiPesertaStrategicMeeting), nil
	}
}

func (t tInformasiPesertaStrategicMeetingDo) FindByPage(offset int, limit int) (result []*model.TInformasiPesertaStrategicMeeting, count int64, err error) {
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

func (t tInformasiPesertaStrategicMeetingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tInformasiPesertaStrategicMeetingDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tInformasiPesertaStrategicMeetingDo) Delete(models ...*model.TInformasiPesertaStrategicMeeting) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tInformasiPesertaStrategicMeetingDo) withDO(do gen.Dao) *tInformasiPesertaStrategicMeetingDo {
	t.DO = *do.(*gen.DO)
	return t
}
