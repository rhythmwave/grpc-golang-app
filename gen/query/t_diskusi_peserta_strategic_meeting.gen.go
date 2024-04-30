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

func newTDiskusiPesertaStrategicMeeting(db *gorm.DB, opts ...gen.DOOption) tDiskusiPesertaStrategicMeeting {
	_tDiskusiPesertaStrategicMeeting := tDiskusiPesertaStrategicMeeting{}

	_tDiskusiPesertaStrategicMeeting.tDiskusiPesertaStrategicMeetingDo.UseDB(db, opts...)
	_tDiskusiPesertaStrategicMeeting.tDiskusiPesertaStrategicMeetingDo.UseModel(&model.TDiskusiPesertaStrategicMeeting{})

	tableName := _tDiskusiPesertaStrategicMeeting.tDiskusiPesertaStrategicMeetingDo.TableName()
	_tDiskusiPesertaStrategicMeeting.ALL = field.NewAsterisk(tableName)
	_tDiskusiPesertaStrategicMeeting.CNik = field.NewString(tableName, "c_nik")
	_tDiskusiPesertaStrategicMeeting.CBidang = field.NewString(tableName, "c_bidang")
	_tDiskusiPesertaStrategicMeeting.CNikPenilai = field.NewString(tableName, "c_nik_penilai")
	_tDiskusiPesertaStrategicMeeting.CNilai = field.NewInt32(tableName, "c_nilai")
	_tDiskusiPesertaStrategicMeeting.CKeterangan = field.NewString(tableName, "c_keterangan")
	_tDiskusiPesertaStrategicMeeting.CLastUpdate = field.NewTime(tableName, "c_last_update")

	_tDiskusiPesertaStrategicMeeting.fillFieldMap()

	return _tDiskusiPesertaStrategicMeeting
}

type tDiskusiPesertaStrategicMeeting struct {
	tDiskusiPesertaStrategicMeetingDo

	ALL         field.Asterisk
	CNik        field.String
	CBidang     field.String
	CNikPenilai field.String
	CNilai      field.Int32 // 1-10
	CKeterangan field.String
	CLastUpdate field.Time

	fieldMap map[string]field.Expr
}

func (t tDiskusiPesertaStrategicMeeting) Table(newTableName string) *tDiskusiPesertaStrategicMeeting {
	t.tDiskusiPesertaStrategicMeetingDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tDiskusiPesertaStrategicMeeting) As(alias string) *tDiskusiPesertaStrategicMeeting {
	t.tDiskusiPesertaStrategicMeetingDo.DO = *(t.tDiskusiPesertaStrategicMeetingDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tDiskusiPesertaStrategicMeeting) updateTableName(table string) *tDiskusiPesertaStrategicMeeting {
	t.ALL = field.NewAsterisk(table)
	t.CNik = field.NewString(table, "c_nik")
	t.CBidang = field.NewString(table, "c_bidang")
	t.CNikPenilai = field.NewString(table, "c_nik_penilai")
	t.CNilai = field.NewInt32(table, "c_nilai")
	t.CKeterangan = field.NewString(table, "c_keterangan")
	t.CLastUpdate = field.NewTime(table, "c_last_update")

	t.fillFieldMap()

	return t
}

func (t *tDiskusiPesertaStrategicMeeting) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tDiskusiPesertaStrategicMeeting) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 6)
	t.fieldMap["c_nik"] = t.CNik
	t.fieldMap["c_bidang"] = t.CBidang
	t.fieldMap["c_nik_penilai"] = t.CNikPenilai
	t.fieldMap["c_nilai"] = t.CNilai
	t.fieldMap["c_keterangan"] = t.CKeterangan
	t.fieldMap["c_last_update"] = t.CLastUpdate
}

func (t tDiskusiPesertaStrategicMeeting) clone(db *gorm.DB) tDiskusiPesertaStrategicMeeting {
	t.tDiskusiPesertaStrategicMeetingDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tDiskusiPesertaStrategicMeeting) replaceDB(db *gorm.DB) tDiskusiPesertaStrategicMeeting {
	t.tDiskusiPesertaStrategicMeetingDo.ReplaceDB(db)
	return t
}

type tDiskusiPesertaStrategicMeetingDo struct{ gen.DO }

type ITDiskusiPesertaStrategicMeetingDo interface {
	gen.SubQuery
	Debug() ITDiskusiPesertaStrategicMeetingDo
	WithContext(ctx context.Context) ITDiskusiPesertaStrategicMeetingDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITDiskusiPesertaStrategicMeetingDo
	WriteDB() ITDiskusiPesertaStrategicMeetingDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITDiskusiPesertaStrategicMeetingDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITDiskusiPesertaStrategicMeetingDo
	Not(conds ...gen.Condition) ITDiskusiPesertaStrategicMeetingDo
	Or(conds ...gen.Condition) ITDiskusiPesertaStrategicMeetingDo
	Select(conds ...field.Expr) ITDiskusiPesertaStrategicMeetingDo
	Where(conds ...gen.Condition) ITDiskusiPesertaStrategicMeetingDo
	Order(conds ...field.Expr) ITDiskusiPesertaStrategicMeetingDo
	Distinct(cols ...field.Expr) ITDiskusiPesertaStrategicMeetingDo
	Omit(cols ...field.Expr) ITDiskusiPesertaStrategicMeetingDo
	Join(table schema.Tabler, on ...field.Expr) ITDiskusiPesertaStrategicMeetingDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITDiskusiPesertaStrategicMeetingDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITDiskusiPesertaStrategicMeetingDo
	Group(cols ...field.Expr) ITDiskusiPesertaStrategicMeetingDo
	Having(conds ...gen.Condition) ITDiskusiPesertaStrategicMeetingDo
	Limit(limit int) ITDiskusiPesertaStrategicMeetingDo
	Offset(offset int) ITDiskusiPesertaStrategicMeetingDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITDiskusiPesertaStrategicMeetingDo
	Unscoped() ITDiskusiPesertaStrategicMeetingDo
	Create(values ...*model.TDiskusiPesertaStrategicMeeting) error
	CreateInBatches(values []*model.TDiskusiPesertaStrategicMeeting, batchSize int) error
	Save(values ...*model.TDiskusiPesertaStrategicMeeting) error
	First() (*model.TDiskusiPesertaStrategicMeeting, error)
	Take() (*model.TDiskusiPesertaStrategicMeeting, error)
	Last() (*model.TDiskusiPesertaStrategicMeeting, error)
	Find() ([]*model.TDiskusiPesertaStrategicMeeting, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TDiskusiPesertaStrategicMeeting, err error)
	FindInBatches(result *[]*model.TDiskusiPesertaStrategicMeeting, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TDiskusiPesertaStrategicMeeting) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITDiskusiPesertaStrategicMeetingDo
	Assign(attrs ...field.AssignExpr) ITDiskusiPesertaStrategicMeetingDo
	Joins(fields ...field.RelationField) ITDiskusiPesertaStrategicMeetingDo
	Preload(fields ...field.RelationField) ITDiskusiPesertaStrategicMeetingDo
	FirstOrInit() (*model.TDiskusiPesertaStrategicMeeting, error)
	FirstOrCreate() (*model.TDiskusiPesertaStrategicMeeting, error)
	FindByPage(offset int, limit int) (result []*model.TDiskusiPesertaStrategicMeeting, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITDiskusiPesertaStrategicMeetingDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tDiskusiPesertaStrategicMeetingDo) Debug() ITDiskusiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Debug())
}

func (t tDiskusiPesertaStrategicMeetingDo) WithContext(ctx context.Context) ITDiskusiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tDiskusiPesertaStrategicMeetingDo) ReadDB() ITDiskusiPesertaStrategicMeetingDo {
	return t.Clauses(dbresolver.Read)
}

func (t tDiskusiPesertaStrategicMeetingDo) WriteDB() ITDiskusiPesertaStrategicMeetingDo {
	return t.Clauses(dbresolver.Write)
}

func (t tDiskusiPesertaStrategicMeetingDo) Session(config *gorm.Session) ITDiskusiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Session(config))
}

func (t tDiskusiPesertaStrategicMeetingDo) Clauses(conds ...clause.Expression) ITDiskusiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tDiskusiPesertaStrategicMeetingDo) Returning(value interface{}, columns ...string) ITDiskusiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tDiskusiPesertaStrategicMeetingDo) Not(conds ...gen.Condition) ITDiskusiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tDiskusiPesertaStrategicMeetingDo) Or(conds ...gen.Condition) ITDiskusiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tDiskusiPesertaStrategicMeetingDo) Select(conds ...field.Expr) ITDiskusiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tDiskusiPesertaStrategicMeetingDo) Where(conds ...gen.Condition) ITDiskusiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tDiskusiPesertaStrategicMeetingDo) Order(conds ...field.Expr) ITDiskusiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tDiskusiPesertaStrategicMeetingDo) Distinct(cols ...field.Expr) ITDiskusiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tDiskusiPesertaStrategicMeetingDo) Omit(cols ...field.Expr) ITDiskusiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tDiskusiPesertaStrategicMeetingDo) Join(table schema.Tabler, on ...field.Expr) ITDiskusiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tDiskusiPesertaStrategicMeetingDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITDiskusiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tDiskusiPesertaStrategicMeetingDo) RightJoin(table schema.Tabler, on ...field.Expr) ITDiskusiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tDiskusiPesertaStrategicMeetingDo) Group(cols ...field.Expr) ITDiskusiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tDiskusiPesertaStrategicMeetingDo) Having(conds ...gen.Condition) ITDiskusiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tDiskusiPesertaStrategicMeetingDo) Limit(limit int) ITDiskusiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tDiskusiPesertaStrategicMeetingDo) Offset(offset int) ITDiskusiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tDiskusiPesertaStrategicMeetingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITDiskusiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tDiskusiPesertaStrategicMeetingDo) Unscoped() ITDiskusiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tDiskusiPesertaStrategicMeetingDo) Create(values ...*model.TDiskusiPesertaStrategicMeeting) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tDiskusiPesertaStrategicMeetingDo) CreateInBatches(values []*model.TDiskusiPesertaStrategicMeeting, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tDiskusiPesertaStrategicMeetingDo) Save(values ...*model.TDiskusiPesertaStrategicMeeting) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tDiskusiPesertaStrategicMeetingDo) First() (*model.TDiskusiPesertaStrategicMeeting, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TDiskusiPesertaStrategicMeeting), nil
	}
}

func (t tDiskusiPesertaStrategicMeetingDo) Take() (*model.TDiskusiPesertaStrategicMeeting, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TDiskusiPesertaStrategicMeeting), nil
	}
}

func (t tDiskusiPesertaStrategicMeetingDo) Last() (*model.TDiskusiPesertaStrategicMeeting, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TDiskusiPesertaStrategicMeeting), nil
	}
}

func (t tDiskusiPesertaStrategicMeetingDo) Find() ([]*model.TDiskusiPesertaStrategicMeeting, error) {
	result, err := t.DO.Find()
	return result.([]*model.TDiskusiPesertaStrategicMeeting), err
}

func (t tDiskusiPesertaStrategicMeetingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TDiskusiPesertaStrategicMeeting, err error) {
	buf := make([]*model.TDiskusiPesertaStrategicMeeting, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tDiskusiPesertaStrategicMeetingDo) FindInBatches(result *[]*model.TDiskusiPesertaStrategicMeeting, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tDiskusiPesertaStrategicMeetingDo) Attrs(attrs ...field.AssignExpr) ITDiskusiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tDiskusiPesertaStrategicMeetingDo) Assign(attrs ...field.AssignExpr) ITDiskusiPesertaStrategicMeetingDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tDiskusiPesertaStrategicMeetingDo) Joins(fields ...field.RelationField) ITDiskusiPesertaStrategicMeetingDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tDiskusiPesertaStrategicMeetingDo) Preload(fields ...field.RelationField) ITDiskusiPesertaStrategicMeetingDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tDiskusiPesertaStrategicMeetingDo) FirstOrInit() (*model.TDiskusiPesertaStrategicMeeting, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TDiskusiPesertaStrategicMeeting), nil
	}
}

func (t tDiskusiPesertaStrategicMeetingDo) FirstOrCreate() (*model.TDiskusiPesertaStrategicMeeting, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TDiskusiPesertaStrategicMeeting), nil
	}
}

func (t tDiskusiPesertaStrategicMeetingDo) FindByPage(offset int, limit int) (result []*model.TDiskusiPesertaStrategicMeeting, count int64, err error) {
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

func (t tDiskusiPesertaStrategicMeetingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tDiskusiPesertaStrategicMeetingDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tDiskusiPesertaStrategicMeetingDo) Delete(models ...*model.TDiskusiPesertaStrategicMeeting) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tDiskusiPesertaStrategicMeetingDo) withDO(do gen.Dao) *tDiskusiPesertaStrategicMeetingDo {
	t.DO = *do.(*gen.DO)
	return t
}