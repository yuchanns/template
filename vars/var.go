package vars

import (
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
)

var DB *gorm.DB

var AsyncIDChan chan string

var SnowflakeNode *snowflake.Node

// 一个不知道指向哪里的服务
var ServerUnknown string
