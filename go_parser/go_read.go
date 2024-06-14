// Estrutura e tipos de dados
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"log"
)

type Address struct {
	Type    string `json:"type"`
	Address string `json:"address"`
}

type Item struct {
	ItemID   int     `json:"item_id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
}

type Order struct {
	OrderID int      `json:"order_id"`
	Items   []Item   `json:"items"`
	Total   float64  `json:"total"`
}

type Settings struct {
	Theme                string `json:"theme"`
	Notifications        bool   `json:"notifications"`
	NewsletterSubscribed bool   `json:"newsletter_subscribed"`
}

type User struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Addresses []Address `json:"addresses"`
	Orders    []Order   `json:"orders"`
	Settings  Settings  `json:"settings"`
}

type UserProfile struct {
	User User `json:"user"`
}

func main() {
	file, err := os.Open("cart.json")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	byteValue, err := io.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	var userProfile UserProfile
	err = json.Unmarshal(byteValue, &userProfile)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User ID: %d\n", userProfile.User.ID)
	fmt.Printf("User Name: %s\n", userProfile.User.Name)
	fmt.Printf("User Email: %s\n", userProfile.User.Email)

	for _, address := range userProfile.User.Addresses {
		fmt.Printf("Address(%s): %s\n", address.Type, address.Address)
	}

	fmt.Println("Orders:")
	for _, order := range userProfile.User.Orders {
		fmt.Printf("ID do Order: %d, total: %.2f\n", order.OrderID, order.Total)

	}

	fmt.Println("Items:")
	for _, order := range userProfile.User.Orders {
		for _, item := range order.Items {
			fmt.Printf("ID do Item: %d, name: %s, price: %.2f\n", item.ItemID, item.Name, item.Price)
		}
	fmt.Println("Settings:")
	fmt.Printf ("Theme: %s\n", userProfile.User.Settings.Theme)
	fmt.Printf("Notifications: %t\n", userProfile.User.Settings.Notifications)
	fmt.Printf("NewsletterSubscribed: %t\n", userProfile.User.Settings.NewsletterSubscribed)

	}
}
