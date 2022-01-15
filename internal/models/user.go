package models

type User struct {
	ID        int    `header:"ID"`
	Password  string `header:"Password"`
	UserName  string `header:"Username"`
	FirstName string `header:"First Name"`
	LastName  string `header:"Last Name"`
	Email     string `header:"Email"`
	IsPrivate bool   `gorm:"'IsPrivate' boolean"`
	//IsActive  bool   `gorm:"'Active' boolean"`
	//Cars      []Car  `gorm:"many2many:cars_users;"`
}
