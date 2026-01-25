package main

type Producto struct {
	ID     int
	Nombre string
	Precio float64
}

func CrearProducto(id int, nombre string, precio float64) Producto {
	return Producto{id, nombre, precio}
}
