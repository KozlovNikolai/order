package main

import (
	"fmt"
	"order/order/internal/app/auth"
	"order/order/internal/app/order"
)

func main() {
	fmt.Println("I'm main client app.")
	auth.Auth()
	order.Order()
}
