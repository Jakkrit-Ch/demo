package entity

import (
	"fmt"
	"log"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func MGDatabase() {

	// Set connection options
	clientOptions := options.Client().ApplyURI("mongodb+srv://gpio:jacker1342@cluster0.cx6mkn3.mongodb.net/?retryWrites=true&w=majority")
	

	// Connect to MongoDB
	ctx := context.Background()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the MongoDB server to verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// // Close the connection
	// err = client.Disconnect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Disconnected from MongoDB!")


	


}

func SetupDatabase() {

	// clientOptions := options.Client().ApplyURI("mongodb+srv://gpio:jacker1342@cluster0.cx6mkn3.mongodb.net/?retryWrites=true&w=majority")
	

	// // Connect to MongoDB
	// ctx := context.Background()

	// database, err := mongo.Connect(ctx, clientOptions)
	database, err := gorm.Open(sqlite.Open("GPIO.db"), &gorm.Config{})
	// database := options.Client().ApplyURI("mongodb+srv://gpio:jacker1342@cluster0.cx6mkn3.mongodb.net/?retryWrites=true&w=majority")

	if err != nil {
		panic("failed to connect database")
	}

	// database.Database(User, options.Database())

	// Migrate the schema

	database.AutoMigrate(
		&User{},
	)

	db = database

	password, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	db.Model(&User{}).Create(&User{
		Firstname: "Jakkrit",
		Lastname:  "Chaiwan",
		Email:     "jackerchaiwan@gmail.com",
		Tel:       "0610255279",
		Room:      "Admin",
		Password:  string(password),
		Role:      "admin",
		Path:      "admin",
	})

	db.Model(&User{}).Create(&User{
		Firstname: "User",
		Lastname:  "01",
		Email:     "user01@gmail.com",
		Tel:       "0951677644",
		Room:      "A11",
		Password:  string(password),
		Role:      "user",
		Path:      "11",
	})

	var jakkrit User
	db.Raw("SELECT * FROM users WHERE email = ?", "jackerchaiwan@gmail.com").Scan(&jakkrit)

	var user01 User
	db.Raw("SELECT * FROM users WHERE email = ?", "user01@gmail.com").Scan(&user01)

}
