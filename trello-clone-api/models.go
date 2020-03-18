package main

import "time"

// User user
type User struct {
	ID             uint       `gorm:"primary_key" json:"id"`
	Email          string     `gorm:"type:varchar(50); not null; unique" json:"email"`
	FirstName      string     `gorm:"type:varchar(30); not null" json:"first_name"`
	LastName       string     `gorm:"type:varchar(30); null; default:''" json:"last_name"`
	LastLogin      time.Time  `gorm:"null" json:"last_login"`
	LoginCount     int        `gorm:"not null; default:0" json:"login_count"`
	PasswordMD5    string     `gorm:"type:varchar(32); not null; default:''"`
	PasswordSHA1   string     `gorm:"type:varchar(40); not null; default:''"`
	PasswordSHA512 string     `gorm:"type:varchar(128); not null; default:''"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `sql:"index" json:"deleted_at"`
}

// UserDetail user detail
type UserDetail struct {
	user         User      `gorm:"ForeignKey:UserID"`
	UserID       uint      `gorm:"not null;unique" json:"user_id"`
	Notification bool      `gorm:"not null; default:1" json:"notificaiton"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// ProfileImage profile image of user
type ProfileImage struct {
	user           User      `gorm:"ForeignKey:UserID; unique"`
	UserID         uint      `json:"user_id"`
	FileName       string    `gorm:"type:varchar(70); not null" json:"file_name"`
	FileOriginName string    `gorm:"type:varchar(150); not null" json:"file_origin_name"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
