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

func newTLog(db *gorm.DB, opts ...gen.DOOption) tLog {
	_tLog := tLog{}

	_tLog.tLogDo.UseDB(db, opts...)
	_tLog.tLogDo.UseModel(&model.TLog{})

	tableName := _tLog.tLogDo.TableName()
	_tLog.ALL = field.NewAsterisk(tableName)
	_tLog.ID = field.NewInt32(tableName, "id")
	_tLog.CIDLogin = field.NewInt32(tableName, "c_id_login")
	_tLog.CNik = field.NewString(tableName, "c_nik")
	_tLog.CNamaLengkap = field.NewString(tableName, "c_nama_lengkap")
	_tLog.CUserName = field.NewString(tableName, "c_user_name")
	_tLog.CTglLogin = field.NewTime(tableName, "c_tgl_login")
	_tLog.CTglLogout = field.NewTime(tableName, "c_tgl_logout")

	_tLog.fillFieldMap()

	return _tLog
}

type tLog struct {
	tLogDo

	ALL          field.Asterisk
	ID           field.Int32
	CIDLogin     field.Int32
	CNik         field.String
	CNamaLengkap field.String
	CUserName    field.String
	CTglLogin    field.Time
	CTglLogout   field.Time

	fieldMap map[string]field.Expr
}

func (t tLog) Table(newTableName string) *tLog {
	t.tLogDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tLog) As(alias string) *tLog {
	t.tLogDo.DO = *(t.tLogDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tLog) updateTableName(table string) *tLog {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt32(table, "id")
	t.CIDLogin = field.NewInt32(table, "c_id_login")
	t.CNik = field.NewString(table, "c_nik")
	t.CNamaLengkap = field.NewString(table, "c_nama_lengkap")
	t.CUserName = field.NewString(table, "c_user_name")
	t.CTglLogin = field.NewTime(table, "c_tgl_login")
	t.CTglLogout = field.NewTime(table, "c_tgl_logout")

	t.fillFieldMap()

	return t
}

func (t *tLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tLog) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 7)
	t.fieldMap["id"] = t.ID
	t.fieldMap["c_id_login"] = t.CIDLogin
	t.fieldMap["c_nik"] = t.CNik
	t.fieldMap["c_nama_lengkap"] = t.CNamaLengkap
	t.fieldMap["c_user_name"] = t.CUserName
	t.fieldMap["c_tgl_login"] = t.CTglLogin
	t.fieldMap["c_tgl_logout"] = t.CTglLogout
}

func (t tLog) clone(db *gorm.DB) tLog {
	t.tLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tLog) replaceDB(db *gorm.DB) tLog {
	t.tLogDo.ReplaceDB(db)
	return t
}

type tLogDo struct{ gen.DO }

type ITLogDo interface {
	gen.SubQuery
	Debug() ITLogDo
	WithContext(ctx context.Context) ITLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITLogDo
	WriteDB() ITLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITLogDo
	Not(conds ...gen.Condition) ITLogDo
	Or(conds ...gen.Condition) ITLogDo
	Select(conds ...field.Expr) ITLogDo
	Where(conds ...gen.Condition) ITLogDo
	Order(conds ...field.Expr) ITLogDo
	Distinct(cols ...field.Expr) ITLogDo
	Omit(cols ...field.Expr) ITLogDo
	Join(table schema.Tabler, on ...field.Expr) ITLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITLogDo
	Group(cols ...field.Expr) ITLogDo
	Having(conds ...gen.Condition) ITLogDo
	Limit(limit int) ITLogDo
	Offset(offset int) ITLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITLogDo
	Unscoped() ITLogDo
	Create(values ...*model.TLog) error
	CreateInBatches(values []*model.TLog, batchSize int) error
	Save(values ...*model.TLog) error
	First() (*model.TLog, error)
	Take() (*model.TLog, error)
	Last() (*model.TLog, error)
	Find() ([]*model.TLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TLog, err error)
	FindInBatches(result *[]*model.TLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITLogDo
	Assign(attrs ...field.AssignExpr) ITLogDo
	Joins(fields ...field.RelationField) ITLogDo
	Preload(fields ...field.RelationField) ITLogDo
	FirstOrInit() (*model.TLog, error)
	FirstOrCreate() (*model.TLog, error)
	FindByPage(offset int, limit int) (result []*model.TLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tLogDo) Debug() ITLogDo {
	return t.withDO(t.DO.Debug())
}

func (t tLogDo) WithContext(ctx context.Context) ITLogDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tLogDo) ReadDB() ITLogDo {
	return t.Clauses(dbresolver.Read)
}

func (t tLogDo) WriteDB() ITLogDo {
	return t.Clauses(dbresolver.Write)
}

func (t tLogDo) Session(config *gorm.Session) ITLogDo {
	return t.withDO(t.DO.Session(config))
}

func (t tLogDo) Clauses(conds ...clause.Expression) ITLogDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tLogDo) Returning(value interface{}, columns ...string) ITLogDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tLogDo) Not(conds ...gen.Condition) ITLogDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tLogDo) Or(conds ...gen.Condition) ITLogDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tLogDo) Select(conds ...field.Expr) ITLogDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tLogDo) Where(conds ...gen.Condition) ITLogDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tLogDo) Order(conds ...field.Expr) ITLogDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tLogDo) Distinct(cols ...field.Expr) ITLogDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tLogDo) Omit(cols ...field.Expr) ITLogDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tLogDo) Join(table schema.Tabler, on ...field.Expr) ITLogDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITLogDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tLogDo) RightJoin(table schema.Tabler, on ...field.Expr) ITLogDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tLogDo) Group(cols ...field.Expr) ITLogDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tLogDo) Having(conds ...gen.Condition) ITLogDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tLogDo) Limit(limit int) ITLogDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tLogDo) Offset(offset int) ITLogDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITLogDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tLogDo) Unscoped() ITLogDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tLogDo) Create(values ...*model.TLog) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tLogDo) CreateInBatches(values []*model.TLog, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tLogDo) Save(values ...*model.TLog) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tLogDo) First() (*model.TLog, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TLog), nil
	}
}

func (t tLogDo) Take() (*model.TLog, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TLog), nil
	}
}

func (t tLogDo) Last() (*model.TLog, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TLog), nil
	}
}

func (t tLogDo) Find() ([]*model.TLog, error) {
	result, err := t.DO.Find()
	return result.([]*model.TLog), err
}

func (t tLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TLog, err error) {
	buf := make([]*model.TLog, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tLogDo) FindInBatches(result *[]*model.TLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tLogDo) Attrs(attrs ...field.AssignExpr) ITLogDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tLogDo) Assign(attrs ...field.AssignExpr) ITLogDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tLogDo) Joins(fields ...field.RelationField) ITLogDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tLogDo) Preload(fields ...field.RelationField) ITLogDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tLogDo) FirstOrInit() (*model.TLog, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TLog), nil
	}
}

func (t tLogDo) FirstOrCreate() (*model.TLog, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TLog), nil
	}
}

func (t tLogDo) FindByPage(offset int, limit int) (result []*model.TLog, count int64, err error) {
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

func (t tLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tLogDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tLogDo) Delete(models ...*model.TLog) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tLogDo) withDO(do gen.Dao) *tLogDo {
	t.DO = *do.(*gen.DO)
	return t
}