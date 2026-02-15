package main

import "fmt"

func mostrarMenu() {
	fmt.Println("\n===== MENÚ E-COMMERCE =====")
	fmt.Println("1. Ejecutar demo (crear usuario, producto, pedido, descontar stock)")
	fmt.Println("2. Salir")
	fmt.Print("Opción: ")
}

func demoPedido() {
	fmt.Println("\nSistema de Gestión E-commerce iniciado")

	// Usuario
	u := CrearUsuario(1, "Samuel", "saguascoc@uide.edu.ec")
	fmt.Println("Email válido?", EmailValido(u.Email()))

	// Catálogo + Inventario
	catalogo := NuevoCatalogoProductos()
	inv := NuevoInventario()

	p1, err := NuevoProducto(101, "Producto Demo", 9.99)
	if err != nil {
		fmt.Println("Error producto:", err)
		return
	}

	// Agregar a catálogo e inventario
	if err := catalogo.Agregar(p1); err != nil {
		fmt.Println("Error agregando al catálogo:", err)
		return
	}
	if err := inv.SetStock(101, 10); err != nil {
		fmt.Println("Error seteando stock:", err)
		return
	}

	// Pedido
	pedido, err := NuevoPedido(1, u.ID())
	if err != nil {
		fmt.Println("Error creando pedido:", err)
		return
	}

	producto, ok := catalogo.Obtener(101)
	if !ok {
		fmt.Println("Producto no encontrado")
		return
	}

	// OJO: ahora se usa producto.ID() y producto.Precio()
	det, err := NuevoDetallePedido(producto.ID(), 2, producto.Precio())
	if err != nil {
		fmt.Println("Error detalle:", err)
		return
	}

	if err := pedido.AgregarDetalle(det, inv); err != nil {
		fmt.Println("Error agregando detalle:", err)
		return
	}

	fmt.Println("Pedido total:", pedido.Total())
	fmt.Println("Stock restante producto 101:", inv.ObtenerStock(101))
}

func main() {
	var opcion int

	for {
		mostrarMenu()
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			demoPedido()
		case 2:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción inválida")
		}
	}
}
