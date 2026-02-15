package main

import "errors"

type Inventario struct {
	stock map[int]int
}

func NuevoInventario() *Inventario {
	return &Inventario{stock: make(map[int]int)}
}

func (i *Inventario) SetStock(productoID int, cantidad int) error {
	if productoID <= 0 {
		return errors.New("productoID inválido")
	}
	if cantidad < 0 {
		return errors.New("cantidad inválida")
	}
	i.stock[productoID] = cantidad
	return nil
}

func (i *Inventario) ObtenerStock(productoID int) int {
	return i.stock[productoID]
}

func (i *Inventario) AumentarStock(productoID int, cantidad int) error {
	if productoID <= 0 {
		return errors.New("productoID inválido")
	}
	if cantidad <= 0 {
		return errors.New("cantidad inválida")
	}
	i.stock[productoID] += cantidad
	return nil
}

func (i *Inventario) ReducirStock(productoID int, cantidad int) error {
	if productoID <= 0 {
		return errors.New("productoID inválido")
	}
	if cantidad <= 0 {
		return errors.New("cantidad inválida")
	}
	if i.stock[productoID] < cantidad {
		return errors.New("stock insuficiente")
	}
	i.stock[productoID] -= cantidad
	return nil
}
