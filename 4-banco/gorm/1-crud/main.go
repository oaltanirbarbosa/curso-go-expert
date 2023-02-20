package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryId int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// // create
	// db.Create(&Product{
	// 	Name:  "Notebook4",
	// 	Price: 2000.0,
	// })

	// //create batch
	// products := []Product{
	// 	{Name: "Notebook1", Price: 1000.00},
	// 	{Name: "Mouse", Price: 50.00},
	// 	{Name: "keboard", Price: 100.00},
	// }
	// db.Create(&products)

	// //select one
	// var product Product
	// // db.First(&product, 2)
	// // fmt.Println(product)
	// //where
	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product)

	// //select all
	// var products []Product
	// db.Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// //limit e paginate
	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// //where
	// var products []Product
	// db.Where("price > ?", 100).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// //like
	// var products []Product
	// db.Where("name LIKE ?", "%book%").Where("price > ?", 1000).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// //update
	// var p Product
	// db.First(&p, 1)
	// p.Name = "New Mouse"
	// db.Save(&p)

	// var p2 Product
	// db.First(&p2, 1)
	// fmt.Println(p2.Name)

	// //delete
	// db.Delete(&p2)

}
