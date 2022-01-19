package startup

import (
	"github.com/bwmarrin/snowflake"
	"github.com/yuchanns/template/vars"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// RegisterVars 注册全局变量
func RegisterVars() {
	var err error
	vars.DB, err = gorm.Open(sqlite.Open("template.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	vars.AsyncIDChan = make(chan string, 10)
	vars.SnowflakeNode, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	vars.ServerUnknown = "http://unknown"
}
