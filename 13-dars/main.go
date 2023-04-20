package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/icrowley/fake"
)

type Product struct {
	Name string
	Price int
}
type User struct {
	Name string
	balance int
	Products []Product
}

func cashier(u User,ch chan int){
	sum:=0
	for _,p:=range u.Products{
        sum+=p.Price
    }
	ch<-sum
} 

func main() {
	rand.Seed(time.Now().UnixNano())
	ch:=make(chan int)

	var products=[]Product{}
	for i := 0; i < 50; i++ {
		product:=Product{Name: fake.Product(),Price: rand.Intn(100)}
		products = append(products, product)
	}
	
	users:=[]User{}
	var length int=len(products)
	
	for i := 0; i < 100; i++ {
		userProducts:=[]Product{}
		
		
		for i := 0; i < 20; i++ {
			randNum:= rand.Intn(length-1)
			
			userProduct:=Product{Name:products[randNum].Name,Price: rand.Intn(100)}
			userProducts = append(userProducts, userProduct)
		}
		user:=User{Name: faker.FirstName(),balance: rand.Intn(1000), Products: userProducts}
		users = append(users, user)
	}
	

	for i := 0; i <= len(users)-3; i+=3 {
		go cashier(users[i],ch)
		u1:=<-ch
		go cashier(users[i+1],ch)
		u2:=<-ch
		go cashier(users[i+2],ch)
		u3:=<-ch
		fmt.Println(u1,u2,u3)
	}
}