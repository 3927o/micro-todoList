package dao

import (
	"github.com/CocaineCong/micro-todoList/app/task/internal/repository/db/model"
)

func migration() {
	// 自动迁移模式
	_db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.Task{})
}
