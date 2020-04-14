package main

import (
	"log"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User - user
type User struct {
	ID         uint       `gorm:"primary_key" json:"id"`
	Email      string     `gorm:"type:varchar(50); not null; unique" json:"email"`
	Username   string     `gorm:"type:varchar(30); not null; unique" json:"username"`
	LastLogin  time.Time  `gorm:"null" json:"last_login"`
	LoginCount int        `gorm:"not null; default:0" json:"login_count"`
	Password   string     `gorm:"type:varchar(100); not null; default:''"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `sql:"index" json:"deleted_at"`
}

// EncryptPassword - encrypt password
func (user *User) EncryptPassword(password string) bool {
	blob := []byte(password)
	hash := hashAndSalt(blob)
	user.Password = hash
	return true
}

func hashAndSalt(blob []byte) string {
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(blob, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

// UserDetail - user detail
type UserDetail struct {
	user         User      `gorm:"ForeignKey:UserID"`
	UserID       uint      `gorm:"not null;unique" json:"user_id"`
	Status       string    `gorm:"not null; default:'deactivated'"`
	Notification bool      `gorm:"not null; default:1" json:"notificaiton"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// ProfileImage - profile image of user
type ProfileImage struct {
	user           User      `gorm:"ForeignKey:UserID; unique"`
	UserID         uint      `json:"user_id"`
	FileName       string    `gorm:"type:varchar(70); not null" json:"file_name"`
	FileOriginName string    `gorm:"type:varchar(150); not null" json:"file_origin_name"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// ConfirmHash - stores hashes for sign up confirm
type ConfirmHash struct {
	user       User      `gorm:"ForeignKey:UserID; unique"`
	UserID     uint      `json:"user_id"`
	Hash       string    `gorm:"type:char(6); not null"`
	ExpireDate time.Time `gorm:"not null" json:"expire_date"`
	CreatedAt  time.Time `json:"created_at"`
}

// SetHash - generates hash string to confirm sign up
func (ch *ConfirmHash) SetHash() {
	now := time.Now()
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		digits
	length := 6
	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	for i := 2; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	ch.Hash = string(buf) // E.g. "3i[g0|)z"
	ch.ExpireDate = now.Add(time.Minute * 30)
}
