/*
 * Copyright (c) 2019 Mars Lee. All rights reserved.
 */

// modify by huayulei_2003@hotmail.com

package model

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

type SqlDriver interface {
	SetDsn(dsn string, options ...interface{})
	GetDsn() string
	Connect() error
	ReadTablesColumns(table string) []Column
	GetTables() []string
	GetDriverType() string
}

type Convert struct {
	ModelPath   string // save path
	Style       string // tab key save like gorm ,orm ,bee orm......
	PackageName string // go package name

	TablePrefix  map[string]string   //if table exists prefix
	TableColumn  map[string][]Column //key is table , value is Column list
	IgnoreTables []string            // ignore tables
	Tables       []string            // all tables

	Driver SqlDriver // impl SqlDriver instance

	initOrm bool
}

//get real gen tables as []string
func (convert *Convert) getGenTables() []string {
	tables := make([]string, 0)
	convert.Tables = convert.Driver.GetTables()
	for _, table := range convert.Tables {
		isIgnore := false
		for _, ignore := range convert.IgnoreTables {
			if table == ignore {
				isIgnore = true
				break
			}
		}

		if !isIgnore {
			tables = append(tables, table)
		}
	}

	return tables
}

//set table prefix
//if exists
//replace prefix to empty string
func (convert *Convert) SetTablePrefix(table, prefix string) {
	convert.TablePrefix[table] = prefix
}

// set model save path
func (convert *Convert) SetModelPath(path string) {
	_, err := os.Stat(path)

	if err != nil {
		if os.IsNotExist(err) {
			panic(fmt.Sprintf("path not exists with error：%v", err))
		}
		log.Println(fmt.Sprintf("path error：%v", err))
	}

	convert.ModelPath = path
}

// set model save path
func (convert *Convert) SetIgnoreTables(table ...string) {
	convert.IgnoreTables = append(convert.IgnoreTables, table...)
}

// set model save path
func (convert *Convert) SetPackageName(name string) {
	convert.PackageName = name
}

//run
//1. connect
//2. getTable
//3. getColumns
//4. build
//5. write file
func (convert *Convert) Run(skipTables []string, isClient bool) {
	err := convert.Driver.Connect()
	if err != nil {
		panic(err)
	}

	for _, tableRealName := range convert.getGenTables() {
		prefix, ok := convert.TablePrefix[tableRealName]
		if ok {
			tableRealName = tableRealName[len(prefix):]
		}
		tableName := tableRealName

		if len(tableName) < 0 {
			continue
		}

		isSkip := false
		for _, s := range skipTables {
			if s == strings.ToLower(tableName) {
				isSkip = true
			}
		}

		if isSkip {
			continue
		}

		tableName = CamelCase(tableName, prefix, true)

		columns := convert.Driver.ReadTablesColumns(tableRealName)
		content := convert.build(tableName, tableRealName, prefix, columns)
		convert.writeModel(tableRealName, content) //写文件
	}

	if isClient {
		convert.writeInit()
	}
}

//build content with table info
func (convert *Convert) build(tableName, tableRealName, prefix string, columns []Column) (content string) {
	depth := 1
	format := GetFormat(convert.Style)
	lcName := LcFirst(tableName)

	content += "package model" + "\n\n" //写包名
	//content += format.AutoImport(tableName)
	content += fmt.Sprintf("import %s \"%s\"\n\n", "client", convert.PackageName)
	content += fmt.Sprintf("var %sName = \"%s\"\n", lcName, LcFirst(tableName))
	content += "type " + tableName + " struct {\n"

	primaryKey := ""
	var primaryColumns Column
	for _, v := range columns {
		var comment string
		if v.ColumnComment != "" {
			comment = fmt.Sprintf(" // %s", v.ColumnComment)
		}
		content += fmt.Sprintf("%s%s %s %s%s\n",
			Tab(depth), v.GetGoColumn(prefix, true), v.GetGoType(), v.GetTag(format), comment)

		if v.IsPrimaryKey() {
			primaryKey = v.ColumnName
			primaryColumns = v
		}

	}

	content += Tab(depth-1) + "}\n\n"

	shortName := strings.ToLower(tableName)[:1]

	if primaryKey != "" {
		content += fmt.Sprintf("// get primary key name \nfunc (%s *%s) %s() string {\n",
			shortName, tableName, "GetKey")
		content += fmt.Sprintf("%sreturn \"%s\"\n",
			Tab(depth), primaryKey)
		content += "}\n\n"

		content += fmt.Sprintf("// get primary key in model\nfunc (%s *%s) %s() %s {\n",
			shortName, tableName, "GetKeyProperty", primaryColumns.GetGoType())
		content += fmt.Sprintf("%sreturn %s.%s\n",
			Tab(depth), shortName, CamelCase(primaryKey, prefix, true))
		content += "}\n\n"

		content += fmt.Sprintf("// set primary key \nfunc (%s *%s) %s(id %s) {\n",
			shortName, tableName, "SetKeyProperty", primaryColumns.GetGoType())
		content += fmt.Sprintf("%s %s.%s = id\n",
			Tab(depth), shortName, CamelCase(primaryKey, prefix, true))
		content += "}\n\n"
	}
	content += fmt.Sprintf("// get table name\nfunc (%s *%s) %s() string {\n",
		shortName, tableName, "TableName")
	content += fmt.Sprintf("%sreturn %s\n",
		Tab(depth), lcName+"Name")
	content += "}\n\n"

	content += convert.buildCurd(tableName, format)
	return content
}

func (convert *Convert) buildKey(tableName string, format Format) string {
	return ""
}

func (convert *Convert) buildCurd(tableName string, format Format) string {
	content := ""
	tpl := format.GetFuncTemplate(convert.Style)
	if tpl != "" {
		shortName := strings.ToLower(tableName)[:1]

		tpl = strings.Replace(tpl, "{{entry}}", LcFirst(tableName), -1)
		tpl = strings.Replace(tpl, "{{object}}", tableName, -1)
		tpl = strings.Replace(tpl, "{{shortName}}", shortName, -1)
		content += tpl
		convert.initOrm = true
	}

	return content
}

//write file
func (convert *Convert) writeInit() {
	if convert.initOrm {
		format := GetFormat(convert.Style)
		tpl := format.GetInitTemplate(convert.Style)

		if tpl != "" {
			tpl = strings.Replace(tpl, "{{package}}", "model", -1)
			tpl = strings.Replace(tpl, "{{dsn}}", convert.Driver.GetDsn(), -1)

			log.Printf("write init file start\n")
			filePath := fmt.Sprintf("%s/%s.go", convert.ModelPath, "dbclient")
			f, err := os.Create(filePath)
			if err != nil {
				log.Println("Can not write file" + filePath)
				return
			}

			defer func() {
				_ = f.Close()
			}()

			_, err = f.WriteString(tpl)
			if err != nil {
				log.Println("Can not write file" + filePath)
				return
			}

			cmd := exec.Command("gofmt", "-w", filePath)
			_ = cmd.Run()
			log.Printf("write init file  success\n")
		}
	}

}

//write file
func (convert *Convert) writeModel(name, content string) {
	log.Printf("write model file %s start\n", name)
	filePath := fmt.Sprintf("%s/%s.go", convert.ModelPath, name)
	f, err := os.Create(filePath)
	if err != nil {
		log.Println("Can not write file" + filePath)
		return
	}

	defer func() {
		_ = f.Close()
	}()

	_, err = f.WriteString(content)
	if err != nil {
		log.Println("Can not write file" + filePath)
		return
	}

	cmd := exec.Command("gofmt", "-w", filePath)
	_ = cmd.Run()
	log.Printf("write model file %s success\n", name)
}

func (convert *Convert) SetStyle(name string) {
	convert.Style = name
}

func (convert *Convert) GetStyle() string {
	if convert.Style == "" {
		return "default"
	}

	return convert.Style
}

func GetDriver(dir, driver, dsn, style, packageName string) *Convert {
	convert := &Convert{}
	convert.SetPackageName(packageName)
	convert.SetModelPath(dir)

	switch driver {
	case "mysql":
		convert.Driver = &MysqlToGo{}
		convert.Driver.SetDsn(dsn)
		convert.SetStyle(style)
	default:
		panic(fmt.Sprintf("do not support this driver: %v\n", driver))
	}

	return convert
}
