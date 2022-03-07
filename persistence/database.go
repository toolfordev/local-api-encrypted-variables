package persistence

import "github.com/tmdgo/extendedgorm"

type ToolForDevDatabase struct {
	extendedgorm.ExtendedDB
}

func (database *ToolForDevDatabase) Init() {
	database.Connect("TOOLFORDEV")
}
