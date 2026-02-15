package main

type InventarioInterface interface {
	ObtenerStock(productoID int) int
	SetStock(productoID int, cantidad int) error
	ReducirStock(productoID int, cantidad int) error
	AumentarStock(productoID int, cantidad int) error
}
