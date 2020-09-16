// modify by huayulei_2003@hotmail.com

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/yuleihua/tower/tools/dbstruct/model"
)

var (
	outDir      string
	driver      string
	dsn         string
	skipTables  string
	style       string
	packageName string
	isClient    bool
	isHelp      bool
)

func init() {
	flag.StringVar(&outDir, "o", "./db_model", "output-dir name save model file path")
	flag.StringVar(&driver, "d", "mysql", "database driver, default is mysql")
	flag.StringVar(&dsn, "u", "", "connection info names url(dsn)")
	flag.StringVar(&skipTables, "t", "", "skip table names")
	flag.StringVar(&style, "s", "gorm", "use orm style like `bee` `gorm`, `default`")
	flag.StringVar(&packageName, "p", "", "package name")
	flag.BoolVar(&isClient, "c", false, "create db-client or not")
	flag.BoolVar(&isHelp, "h", false, "help")
}

func main() {
	flag.Parse()

	if isHelp {
		flag.Usage()
		os.Exit(0)
	}

	if outDir != "" {
		if _, err := os.Stat(outDir); os.IsNotExist(err) {
			fmt.Printf("create dir[%s]", outDir)
			os.MkdirAll(outDir, 0777)
		}
	}

	if dsn == "" {
		fmt.Println("dsn is null")
		flag.Usage()
		os.Exit(0)
	}

	skips := strings.Split(strings.ToLower(skipTables), ",")

	model.GetDriver(outDir, driver, dsn, style, packageName).Run(skips, isClient)
}
