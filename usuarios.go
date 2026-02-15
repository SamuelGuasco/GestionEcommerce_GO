package main

// Usuario representa una entidad del sistema.
// Sus atributos son privados (minúscula) para aplicar encapsulación.
type Usuario struct {
	id     int
	nombre string
	email  string
	activo bool
}

// CrearUsuario actúa como constructor.
// Inicializa el usuario y por defecto lo deja activo.
func CrearUsuario(id int, nombre, email string) Usuario {
	return Usuario{
		id:     id,
		nombre: nombre,
		email:  email,
		activo: true,
	}
}

/*
GETTERS:

Los getters son métodos que permiten acceder a atributos privados
sin exponerlos directamente.

Esto protege la integridad del objeto y evita modificaciones externas
no controladas.
*/

// ID devuelve el identificador del usuario.
func (u Usuario) ID() int {
	return u.id
}

// Nombre devuelve el nombre del usuario.
func (u Usuario) Nombre() string {
	return u.nombre
}

// Email devuelve el correo electrónico del usuario.
func (u Usuario) Email() string {
	return u.email
}

// Activo indica si el usuario está activo en el sistema.
func (u Usuario) Activo() bool {
	return u.activo
}

// Desactivar cambia el estado del usuario a inactivo.
// Se usa puntero porque modifica el objeto original.
func (u *Usuario) Desactivar() {
	u.activo = false
}

// EmailValido valida que el correo tenga una sola '@'
// y que no esté al inicio ni al final.
func EmailValido(email string) bool {
	arroba := -1
	for i, ch := range email {
		if ch == '@' {
			if arroba != -1 {
				return false
			}
			arroba = i
		}
	}
	return arroba > 0 && arroba < len(email)-1
}
