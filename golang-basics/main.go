package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string  `gorm:"type:VARCHAR(30); null"`
	LastName  string  `gorm:"size:100; default:'Smith"`
	Email     string  `gorm:"unique; not null"`
	Address   Address `gorm:"foreignKey:UserId "`
	Books     []Book  `gorm:"many2many:user_books"`
}

type Address struct {
	gorm.Model
	UserId int
	Name   string
}
type Book struct {
	gorm.Model
	Title string
}

func main() {
	db, error := gorm.Open(mysql.Open("root:Spartan11713@/go_basics"), &gorm.Config{})

	if error != nil {
		panic("could not connect with the db")
	}
	// db.Migrator().DropTable(&User{}, &Address{}, &Book{})
	//db.AutoMigrate(&User{}, &Address{}, &Book{}) //this way we create a table in the db

	//if automigrate fails migrator should work better
	// db.Migrator().DropTable(&User{}, &Address{}, &Book{})
	// db.Migrator().CreateTable(&User{}, &Address{}, &Book{})

	// create a new user
	// user := User{
	// 	FirstName: "Manuel",
	// 	LastName:  "Escobedo",
	// 	Email:     "manuelfx117@email.com",
	// }
	//db.Create(&user)

	// update a user
	// user := User{
	// 	Id:        1,
	// 	FirstName: "Manuel",
	// 	LastName:  "another last name",
	// 	Email:     "manuelfx117@email.com",
	// }
	// db.Updates(&user)

	// delete the user
	// user := User{
	// 	Id: 1,
	// }
	// db.Delete(user)

	// querying records
	// user := User{}
	// db.First(&user)
	// fmt.Println(user)

	// userLast := User{}
	// db.Last(&userLast)
	// fmt.Println(userLast)

	// userEmail := User{}
	// db.Where("email", "third@email.com").First(&userEmail)
	//fmt.Println(userEmail)

	// user := User{
	// 	Model: gorm.Model{
	// 		CreatedAt: time.Now(),
	// 	},
	// }
	// fmt.Println(user)

	//create data with new foreing key

	// user := User{
	// 	FirstName: "Jose Manuel",
	// 	LastName:  "Escobedo Lopez",
	// 	Email:     "manuelfx2@mail.com",
	// }
	// db.Create(&user)

	// address := Address{
	// 	UserId: int(user.ID),
	// 	Name:   "main str",
	// }

	// db.Create(&address)

	// 	user := User{
	// 		FirstName: "Carlos",
	// 		LastName:  "Hernandez",
	// 		Email:     "carlitos@email.com",
	// 		Address: Address{
	// 			Name: "other street",
	// 		},
	// 	}

	// 	db.Create(&user)

	//create new user with book

	user := User{
		FirstName: "Carlos",
		LastName:  "Hernandez",
		Email:     "carlitos2@email.com",
		Address: Address{
			Name: "other street",
		},
		Books: []Book{
			{
				Title: "some book title",
			},
		},
	}

	db.Create(&user)

}
