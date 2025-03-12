package chatik

type User struct {
	ID       int    `gorm:"primarykey" json:"id"`
	Username     string `gorm:"type:varchar(100);not null" json:"username"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password string `gorm:"type:text;not null" json:"password"`
}


