package main

import "errors"

type DetallePedido struct {
	ProductoID int
	Cantidad   int
	Precio     float64
}

func NuevoDetallePedido(productoID, cantidad int, precio float64) (DetallePedido, error) {
	if productoID <= 0 {
		return DetallePedido{}, errors.New("productoID inválido")
	}
	if cantidad <= 0 {
		return DetallePedido{}, errors.New("cantidad inválida")
	}
	if precio <= 0 {
		return DetallePedido{}, errors.New("precio inválido")
	}
	return DetallePedido{
		ProductoID: productoID,
		Cantidad:   cantidad,
		Precio:     precio,
	}, nil
}

func (d DetallePedido) Subtotal() float64 {
	return float64(d.Cantidad) * d.Precio
}

type Pedido struct {
	ID        int
	UsuarioID int
	Estado    string
	detalles  []DetallePedido // encapsulado
}

func NuevoPedido(id int, usuarioID int) (*Pedido, error) {
	if id <= 0 {
		return nil, errors.New("id inválido")
	}
	if usuarioID <= 0 {
		return nil, errors.New("usuarioID inválido")
	}
	return &Pedido{
		ID:        id,
		UsuarioID: usuarioID,
		Estado:    "CREADO",
		detalles:  make([]DetallePedido, 0),
	}, nil
}

func (p *Pedido) AgregarDetalle(det DetallePedido, inv InventarioInterface) error {
	if err := inv.ReducirStock(det.ProductoID, det.Cantidad); err != nil {
		return err
	}

	p.detalles = append(p.detalles, det)
	return nil
}

func (p *Pedido) Total() float64 {
	total := 0.0
	for _, d := range p.detalles {
		total += d.Subtotal()
	}
	return total
}

func (p *Pedido) Detalles() []DetallePedido {
	// devolvemos copia para no exponer el slice interno
	copia := make([]DetallePedido, len(p.detalles))
	copy(copia, p.detalles)
	return copia
}

func (p *Pedido) CambiarEstado(nuevo string) error {
	if nuevo == "" {
		return errors.New("estado inválido")
	}
	p.Estado = nuevo
	return nil
}
