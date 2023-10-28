package service

import (
	"FP-BDS-Sanbercode-Go-50-anggi/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetTotalTransaction(id int, host string) int {
	url := fmt.Sprintf("http://http://%v/promo/%v", host, id)
	response, err := http.Get(url)
	if err != nil{
		panic(err)
	}
	responseData, err := io.ReadAll(response.Body) 
	if err != nil{
		panic(err)
	}
	var promo models.Promotion

	err = json.Unmarshal([]byte(responseData), &promo)
	if err != nil{
		panic(err)
	}

	url2 := fmt.Sprintf("http://http://%v/product/%v", host, id)
	responseProduct, err := http.Get(url2)
	if err != nil{
		panic(err)
	}
	responseProductData, err := io.ReadAll(responseProduct.Body)
	if err != nil{
		panic(err)
	}
	var product models.Products
	err = json.Unmarshal(responseProductData, &product)
	if err != nil{
		panic(err)
	}

	if promo.IsActive == 1{
		return int(product.Price) - (promo.Promo * int(product.Price)/ 100)
	}
	return int(product.Price)
}

