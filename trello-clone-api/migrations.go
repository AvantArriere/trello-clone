package main

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

// Migrations set history of migrations
var Migrations = []*gormigrate.Migration{
	{
		ID: "202003160100",
		Migrate: func(tx *gorm.DB) error {
			user := User{}
			userDetail := UserDetail{}
			profileImage := ProfileImage{}

			return tx.AutoMigrate(
				&user,
				&userDetail,
				&profileImage,
			).Error
		},
		Rollback: func(tx *gorm.DB) error {
			user := User{}
			userDetail := UserDetail{}
			profileImage := ProfileImage{}

			return tx.DropTableIfExists(
				&user,
				&userDetail,
				&profileImage,
			).Error
		},
	},
}