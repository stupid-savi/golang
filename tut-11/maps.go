package main

import (
	"fmt"
	"maps"
)

type AnimeList []string

type Customer struct {
	id        string
	name      string
	animeList AnimeList
}

type CustomersList []Customer

func main() {
	// map -> hash, object, dictionary
	var users = make(map[string]string)
	fmt.Println("users>>>>>>>>>>>>>>>>>>>>>", users)
	users["go"] = "Golang"
	users["js"] = "JavaScript"
	fmt.Println("users>>>>>>>>>>>>>>", users["go"], users["js"])
	fmt.Println(users["py"], "..........")
	// if a key doesn't exist in the map then  it will return the zeroed value i.e in case of int = 0, bool = false & string = ""

	var customers = make(map[string]Customer)
	customers["123456"] = Customer{
		id:        "123456",
		name:      "savi",
		animeList: AnimeList{"Naruto", "One Piece"},
	}

	fmt.Println(customers)
	// var customerList CustomersList
	// customerList = append(customerList, Customer{
	// 	id:        "123456",
	// 	name:      "savi",
	// 	animeList: AnimeList{"Naruto", "One Piece"},
	// })
	customerList := CustomersList{{
		id:        "1",
		name:      "savi",
		animeList: AnimeList{"Naruto", "One Piece"},
	}, {
		id:        "2",
		name:      "kavi",
		animeList: AnimeList{"Naruto", "One Piece"},
	}, {
		id:        "3",
		name:      "havi",
		animeList: AnimeList{"Naruto", "One Piece"},
	}}
	customerList = append(customerList, Customer{
		id:        "4",
		name:      "tavi",
		animeList: AnimeList{"Naruto", "One Piece"},
	})
	fmt.Println(customerList)
	fmt.Println(len(customerList))

	delete(users, "go")
	clear(users) // clear the map
	fmt.Println(users)

	m := map[string]int{"1": 1, "2": 2}
	fmt.Println(m)

	notOk, ok := m["5"]

	if ok {
		fmt.Println("ok", ok)
	} else if notOk == 0 {
		fmt.Println("notok", notOk)
	}

	m1 := map[string]int{"1 ": 1}
	m2 := map[string]int{"1": 1}

	println(maps.Equal(m1, m2))

}
