package main

import "errors"

/*
Producto representa un artículo disponible en el sistema de e-commerce.

Se aplica encapsulación:
- Los atributos son privados (minúscula).
- Solo se accede a ellos mediante métodos públicos (getters).
*/
type Producto struct {
	id     int
	nombre string
	precio float64
}

/*
NuevoProducto actúa como constructor.
Valida los datos antes de crear el objeto.
Si los valores no son correctos, retorna error.
*/
func NuevoProducto(id int, nombre string, precio float64) (Producto, error) {
	if id <= 0 {
		return Producto{}, errors.New("id inválido")
	}
	if nombre == "" {
		return Producto{}, errors.New("nombre inválido")
	}
	if precio <= 0 {
		return Producto{}, errors.New("precio inválido")
	}

	return Producto{
		id:     id,
		nombre: nombre,
		precio: precio,
	}, nil
}

/*
GETTERS:
Permiten acceder a los atributos privados sin exponerlos directamente.
Esto protege la integridad del objeto.
*/

// ID devuelve el identificador del producto.
func (p Producto) ID() int {
	return p.id
}

// Nombre devuelve el nombre del producto.
func (p Producto) Nombre() string {
	return p.nombre
}

// Precio devuelve el precio actual del producto.
func (p Producto) Precio() float64 {
	return p.precio
}

/*
CambiarPrecio permite modificar el precio de forma controlada.
Se usa puntero para modificar el objeto original.
*/
func (p *Producto) CambiarPrecio(nuevo float64) error {
	if nuevo <= 0 {
		return errors.New("precio inválido")
	}
	p.precio = nuevo
	return nil
}
