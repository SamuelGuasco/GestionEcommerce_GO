package main

type Usuario struct {
	ID     int
	Nombre string
	Email  string
	Activo bool
}

// Función pura: crea y devuelve un usuario
func CrearUsuario(id int, nombre, email string) Usuario {
	return Usuario{
		ID:     id,
		Nombre: nombre,
		Email:  email,
		Activo: true,
	}
}

// Función pura: cambia estado y devuelve una copia modificada
func DesactivarUsuario(u Usuario) Usuario {
	u.Activo = false
	return u
}

// Función pura: valida email (simple) sin efectos secundarios
func EmailValido(email string) bool {
	for _, ch := range email {
		if ch == '@' {
			return true
		}
	}
	return false
}
