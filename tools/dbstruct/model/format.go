package model

import (
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Format struct {
	Framework      string
	TabFormat      string // format must use 3 %s in it, first column name, second property  third json name
	AutoInfo       string
	PropertyFormat PropertyFormat // like size s
	JsonUseCamel   bool
}
type PropertyFormat struct {
	Size    string
	Type    string
	Index   string
	Default string
}

//`gorm:"column:beast_id"`
var BeeFormat Format
var DefaultFormat Format
var GormFormat Format

func init() {
	BeeFormat = Format{
		Framework: "bee",
		TabFormat: "`orm:\"column(%s);%s\" json:\"%s\"`",
		PropertyFormat: PropertyFormat{
			Size:    "size(%d)",
			Type:    "type(%s)",
			Index:   "%s",
			Default: "",
		},
		AutoInfo: "\nimport \"github.com/astaxie/beego/orm\"\n\nfunc init(){\n\torm.RegisterModel(new({{modelName}}))\n}\n\n",
	}
	DefaultFormat = Format{
		Framework: "default",
		TabFormat: "`orm:\"%s;%s\" json:\"%s\"`",
	}
	GormFormat = Format{
		Framework: "gorm",
		PropertyFormat: PropertyFormat{
			Size:    "size:%d",
			Type:    "type:%s",
			Index:   "",
			Default: "default:%s",
		},
		TabFormat: "`gorm:\"column:%s;%s\" json:\"%s\"`",
		//AutoInfo:  "\nimport (\n\t\"fmt\"\n)\n\n",
	}
}

func GetFormat(framework string) Format {
	switch framework {
	case "bee":
		return BeeFormat
	case "gorm":
		return GormFormat
	default:
		return DefaultFormat
	}
}

func (format Format) AutoImport(modelName string) string {
	if format.AutoInfo == "" {
		return ""
	}
	return strings.Replace(format.AutoInfo, "{{modelName}}", modelName, -1)
}

func (format Format) GetTabFormat() string {
	return format.TabFormat
}

func (format Format) GetPropertyFormat() PropertyFormat {
	return format.PropertyFormat
}

func (pf PropertyFormat) GetSizeFormat() string {
	return pf.Size
}

func (pf PropertyFormat) GetIndexFormat() string {
	return pf.Index
}

func (pf PropertyFormat) GetTypeFormat() string {
	return pf.Type
}

func (pf PropertyFormat) GetDefaultFormat() string {
	return pf.Default
}

func (format Format) GetFuncTemplate(t string) string {
	switch t {
	case "gorm":
		return GormTpl
	default:
		return ""
	}
}

func (format Format) GetInitTemplate(t string) string {
	switch t {
	case "gorm":
		return GormInit
	default:
		return ""
	}
}
