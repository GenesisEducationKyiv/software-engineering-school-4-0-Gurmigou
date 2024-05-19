package initializer

import "se-school-case/model"

func AutoMigrateDatabase() {
	var err error
	err = DB.AutoMigrate(&model.User{}, &model.Rate{})
	if err != nil {
		return
	}
}
