package main

import "fmt"

func main() {
	fmt.Println("Sistema de Gestión E-commerce iniciado")
	u := CrearUsuario(1, "Samuel", "saguascoc@uide.edu.ec")
	fmt.Println("Email válido?", EmailValido(u.Email))

	item := CrearItemInventario(101, 10)
	item = AumentarStock(item, 5)

	item, ok := ReducirStock(item, 12)
	fmt.Println("Stock actualizado:", item.Stock, "OK:", ok)
}
