package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Price struct {
	ID     int `gorm:"primaryKey"`
	Code   string
	CodeIn string
	Bid    float64
	gorm.Model
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// var data Resp
	resp_map := make(map[string]map[string]interface{})
	err = json.Unmarshal(body, &resp_map)
	if err != nil {
		panic(err)
	}
	code := resp_map["USDBRL"]["code"].(string)
	code_in := resp_map["USDBRL"]["codein"].(string)
	result := resp_map["USDBRL"]["bid"].(string)
	bid, _ := strconv.ParseFloat(result, 64)

	ctxdb, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()

	dsn := "root:@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Price{})
	// create
	db.WithContext(ctxdb).Create(&Price{
		Code:   code,
		CodeIn: code_in,
		Bid:    bid,
	})

	w.Header().Set("Contet-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
