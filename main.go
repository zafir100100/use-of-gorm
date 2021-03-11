package main

import (
	// "fmt"

	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	// "github.com/jinzhu/inflection"
)

// Product -
type Product struct {
	// gorm.Model
	Code  string
	Price uint
}

// Test1 =
type Test1 struct {
	// gorm.Model
	Nn  string
	Dd string
}

func main() {

	dsn := "host=localhost user=postgres password=postgres dbname=test1 port=5432 sslmode=disable TimeZone=Asia/Dhaka"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	// db.AutoMigrate(&Product{})

	// Create
	db.Create(&Test1{Nn: "D42", Dd: "2020-02-02"})
	var t1 []Test1
	db.Find(&t1)
	fmt.Println(t1)
	//   // Read
	//   var product Product
	// //   db.First(&product, 1) // find product with integer primary key
	//   db.First(&product, "code = ?", "D42") // find product with code D42

	// var products []Product
	//   result := db.Find(&products)
	//   fmt.Println(result.Error)


	//   // Update - update product's price to 200
	//   db.Model(&product).Update("Price", 200)
	//   // Update - update multiple fields
	//   db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//   db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	//   // Delete - delete product
	//   db.Delete(&product, 1)
}
