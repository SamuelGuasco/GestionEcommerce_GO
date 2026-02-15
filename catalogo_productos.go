package main

import "errors"

type CatalogoProductos struct {
	productos map[int]Producto // id -> Producto
}

func NuevoCatalogoProductos() *CatalogoProductos {
	return &CatalogoProductos{
		productos: make(map[int]Producto),
	}
}

func (c *CatalogoProductos) Agregar(p Producto) error {
	id := p.ID()
	if id <= 0 {
		return errors.New("producto inválido")
	}
	if _, existe := c.productos[id]; existe {
		return errors.New("ya existe un producto con ese ID")
	}
	c.productos[id] = p
	return nil
}

func (c *CatalogoProductos) Obtener(id int) (Producto, bool) {
	p, ok := c.productos[id]
	return p, ok
}
