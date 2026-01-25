package main

type InventarioItem struct {
	ProductoID int
	Stock      int
}

// Función pura: crea un item de inventario
func CrearItemInventario(productoID int, stock int) InventarioItem {
	return InventarioItem{
		ProductoID: productoID,
		Stock:      stock,
	}
}

// Función pura: aumenta stock y devuelve el item actualizado
func AumentarStock(item InventarioItem, cantidad int) InventarioItem {
	if cantidad < 0 {
		return item // no cambia si es inválido
	}
	item.Stock += cantidad
	return item
}

// Función pura: reduce stock si alcanza, devuelve (itemActualizado, ok)
func ReducirStock(item InventarioItem, cantidad int) (InventarioItem, bool) {
	if cantidad < 0 {
		return item, false
	}
	if item.Stock < cantidad {
		return item, false
	}
	item.Stock -= cantidad
	return item, true
}
