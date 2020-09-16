package model

//var GormTpl = `
//func Get{{object}}ById(id string)  ({{entry}} *{{object}}) {
//	err := Orm.Model({{entry}}).First({{entry}}, {{entry}}.GetKey() + " = '"+id+"'").GetErrors()
//	if len(err) > 0 {
//		return nil
//	}
//	return
//}
//
//func Get{{object}}One(where string, args... interface{}) ({{entry}} *{{object}}) {
//	err := Orm.Model({{entry}}).First({{entry}}, where, args...).GetErrors()
//	if len(err) > 0 {
//		return nil
//	}
//	return
//}
//
//func Get{{object}}List(page,limit int64, where string, condition interface{}) (list []*{{object}}) {
//	err := Orm.Model({{entry}}).Limit(limit).Offset((page-1) * limit).Find(list, where, condition).GetErrors()
//	if err != nil {
//		return nil
//	}
//	return
//}
//
//func ({{entry}} *{{object}}) Create() []error {
//	return Orm.Model({{entry}}).Create({{entry}}).GetErrors()
//}
//
//func ({{entry}} *{{object}}) Update(update {{object}}) []error  {
//	return Orm.Model({{entry}}).UpdateColumns(update).GetErrors()
//}
//
//func ({{entry}} *{{object}}) Delete()  {
//	Orm.Model({{entry}}).Delete({{entry}})
//}
//`

//var GormInit = `
//package {{package}}
//
//import (
//	_ "github.com/go-sql-driver/mysql"
//	"github.com/jinzhu/gorm"
//)
//
//var Orm *gorm.DB
//
//func init() {
//	db, err := gorm.Open("mysql", "{{dns}}")
//	if err != nil {
//		panic("连接数据库失败")
//	}
//	Orm = db
//}
//`
