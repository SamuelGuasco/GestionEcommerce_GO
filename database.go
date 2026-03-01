package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// ---------------------------
// CONEXIÓN A BASE DE DATOS
// ---------------------------

func conectarBD() (*sql.DB, error) {
	user := getEnv("DB_USER", "root")
	pass := getEnv("DB_PASS", "")
	host := getEnv("DB_HOST", "127.0.0.1")
	port := getEnv("DB_PORT", "3306")
	dbName := getEnv("DB_NAME", "ecommerce_db")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		user, pass, host, port, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("✅ Conexión exitosa a MySQL")
	return db, nil
}

// Función auxiliar para variables de entorno
func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

// ---------------------------
// USUARIOS
// ---------------------------

func guardarUsuario(db *sql.DB, id int, nombre, email string, activo bool) error {
	activoInt := 0
	if activo {
		activoInt = 1
	}

	_, err := db.Exec(`
		INSERT INTO usuarios (id, nombre, email, activo)
		VALUES (?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
			nombre = VALUES(nombre),
			email  = VALUES(email),
			activo = VALUES(activo)
	`, id, nombre, email, activoInt)

	return err
}

// ---------------------------
// PRODUCTOS
// ---------------------------

func guardarProducto(db *sql.DB, id int, nombre string, precio float64) error {
	_, err := db.Exec(`
		INSERT INTO productos (id, nombre, precio)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE
			nombre = VALUES(nombre),
			precio = VALUES(precio)
	`, id, nombre, precio)

	return err
}

// ---------------------------
// PEDIDOS
// ---------------------------

// Crea pedido y devuelve ID generado por MySQL
func crearPedido(db *sql.DB, usuarioID int, total float64) (int64, error) {
	res, err := db.Exec(`
		INSERT INTO pedidos (usuario_id, total)
		VALUES (?, ?)
	`, usuarioID, total)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// Inserta un detalle individual
func crearDetallePedido(db *sql.DB, pedidoID int64, productoID int, cantidad int, precio float64) error {
	subtotal := float64(cantidad) * precio

	_, err := db.Exec(`
		INSERT INTO pedido_detalles (pedido_id, producto_id, cantidad, precio_unitario, subtotal)
		VALUES (?, ?, ?, ?, ?)
	`, pedidoID, productoID, cantidad, precio, subtotal)

	return err
}

// Guarda pedido completo dentro de una TRANSACCIÓN
func guardarPedidoConDetalle(db *sql.DB, usuarioID int, productoID int, cantidad int, precio float64, total float64) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	// 1️⃣ Crear pedido
	res, err := tx.Exec(`
		INSERT INTO pedidos (usuario_id, total)
		VALUES (?, ?)
	`, usuarioID, total)
	if err != nil {
		return err
	}

	pedidoID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// 2️⃣ Crear detalle
	subtotal := float64(cantidad) * precio

	_, err = tx.Exec(`
		INSERT INTO pedido_detalles (pedido_id, producto_id, cantidad, precio_unitario, subtotal)
		VALUES (?, ?, ?, ?, ?)
	`, pedidoID, productoID, cantidad, precio, subtotal)
	if err != nil {
		return err
	}

	// 3️⃣ Commit
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
