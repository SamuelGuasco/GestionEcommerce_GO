package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func mostrarMenu() {
	fmt.Println("\n===== MENÚ E-COMMERCE =====")
	fmt.Println("1. Ejecutar demo (crear usuario, producto, pedido, descontar stock)")
	fmt.Println("2. Salir")
	fmt.Print("Opción: ")
}

// ---------------------------
// CONEXIÓN A MYSQL
// ---------------------------

func abrirDB() (*sql.DB, error) {
	user := "root"
	pass := "Beyond12345." // tu contraseña con punto
	host := "127.0.0.1"
	port := "3306"
	dbName := "ecommerce_db"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func guardarUsuarioDB(db *sql.DB, u Usuario) error {
	_, err := db.Exec(`
		INSERT INTO usuarios (id, nombre, email, activo)
		VALUES (?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
			nombre = VALUES(nombre),
			email = VALUES(email),
			activo = VALUES(activo);
	`, u.ID(), u.Nombre(), u.Email(), u.Activo())

	return err
}

func guardarProductoDB(db *sql.DB, p Producto) error {
	_, err := db.Exec(`
		INSERT INTO productos (id, nombre, precio)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE
			nombre = VALUES(nombre),
			precio = VALUES(precio);
	`, p.ID(), p.Nombre(), p.Precio())

	return err
}

// Guarda pedido + detalles en una transacción
func guardarPedidoYDetallesDB(db *sql.DB, pedido *Pedido) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	// 1) Pedido (upsert)
	_, err = tx.Exec(`
		INSERT INTO pedidos (id, usuario_id, fecha, total)
		VALUES (?, ?, NOW(), ?)
		ON DUPLICATE KEY UPDATE
			usuario_id = VALUES(usuario_id),
			fecha = NOW(),
			total = VALUES(total);
	`, pedido.ID, pedido.UsuarioID, pedido.Total())
	if err != nil {
		return err
	}

	// 2) Limpiamos detalles anteriores de ese pedido (para evitar duplicados si corres varias veces)
	_, err = tx.Exec(`DELETE FROM pedido_detalles WHERE pedido_id = ?;`, pedido.ID)
	if err != nil {
		return err
	}

	// 3) Insert detalles (ojo: columnas reales son precio_unitario y subtotal)
	for _, det := range pedido.Detalles() {
		_, err = tx.Exec(`
			INSERT INTO pedido_detalles (pedido_id, producto_id, cantidad, precio_unitario, subtotal)
			VALUES (?, ?, ?, ?, ?);
		`, pedido.ID, det.ProductoID, det.Cantidad, det.Precio, det.Subtotal())
		if err != nil {
			return err
		}
	}

	// 4) Commit
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// ---------------------------
// DEMO
// ---------------------------

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

	// ---------------------------
	// GUARDAR EN MYSQL
	// ---------------------------

	db, err := abrirDB()
	if err != nil {
		fmt.Println("Error conectando a MySQL:", err)
		return
	}
	defer db.Close()

	if err := guardarUsuarioDB(db, u); err != nil {
		fmt.Println("Error guardando usuario:", err)
		return
	}

	if err := guardarProductoDB(db, p1); err != nil {
		fmt.Println("Error guardando producto:", err)
		return
	}

	if err := guardarPedidoYDetallesDB(db, pedido); err != nil {
		fmt.Println("Error guardando pedido/detalles:", err)
		return
	}

	fmt.Println("✅ Datos guardados en MySQL correctamente")
}

// ---------------------------
// MAIN
// ---------------------------

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
