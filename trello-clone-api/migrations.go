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
	{
		ID: "202004141620",
		Migrate: func(tx *gorm.DB) error {
			user := User{}
			userDetail := UserDetail{}
			confirmHash := ConfirmHash{}

			return tx.AutoMigrate(
				&user,
				&userDetail,
				&confirmHash,
			).Error
		},
		Rollback: func(tx *gorm.DB) error {
			user := User{}
			userDetail := UserDetail{}
			confirmHash := ConfirmHash{}

			return tx.DropTableIfExists(
				&user,
				&userDetail,
				&confirmHash,
			).Error
		},
	},
	{
		ID: "202004141630",
		Migrate: func(tx *gorm.DB) error {
			user := User{}
			if err := tx.Model(&user).DropColumn("first_name").Error; err != nil {
				return err
			}
			if err := tx.Model(&user).DropColumn("last_name").Error; err != nil {
				return err
			}
			if err := tx.Model(&user).DropColumn("password_md5").Error; err != nil {
				return err
			}
			if err := tx.Model(&user).DropColumn("password_sha1").Error; err != nil {
				return err
			}
			if err := tx.Model(&user).DropColumn("password_sha512").Error; err != nil {
				return err
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Rollback().Error
		},
	},
	{
		ID: "202004141634",
		Migrate: func(tx *gorm.DB) error {
			user := User{}
			if err := tx.Model(&user).DropColumn("password").Error; err != nil {
				return err
			}
			if err := tx.AutoMigrate(&user).Error; err != nil {
				return err
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Rollback().Error
		},
	},
	{
		ID: "202004141645",
		Migrate: func(tx *gorm.DB) error {
			if err := tx.Model(&User{}).AddUniqueIndex("email", "email").Error; err != nil {
				return err
			}
			if err := tx.Model(&User{}).AddUniqueIndex("username", "username").Error; err != nil {
				return err
			}
			if err := tx.Model(&UserDetail{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").Error; err != nil {
				return err
			}
			if err := tx.Model(&ProfileImage{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").Error; err != nil {
				return err
			}
			if err := tx.Model(&ConfirmHash{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").Error; err != nil {
				return err
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Rollback().Error
		},
	},
	{
		ID: "202004141655",
		Migrate: func(tx *gorm.DB) error {
			userDetail := UserDetail{}
			if err := tx.Model(&userDetail).DropColumn("email").Error; err != nil {
				return err
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Rollback().Error
		},
	},
}
