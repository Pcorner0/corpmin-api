package models

import (
	"gorm.io/gorm"
	"time"
)

// Empleados Estructura para los datos de empleados
type Vendedores struct {
	Idvendedor      int            `json:"idvendedor" gorm:"primaryKey; autoIncrement=1000"`
	Nombre          string         `json:"nombre" binding:"required,uppercase,max=50"`
	Primerapellido  string         `json:"primerapellido" binding:"required,uppercase,max=50"`
	Segundoapellido string         `json:"segundoapellido" binding:"uppercase,max=50"`
	Telefono        string         `json:"telefono" binding:"required"`
	Email           string         `json:"email" binding:"required,email,max=100"`
	Rol             string         `json:"rol"`
	Oficina         string         `json:"oficina"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}

// CreateEmployee registrar un nuevo empleado
func CreateVendedor(db *gorm.DB, Vendedor *Vendedores) (err error) {
	err = db.Create(Vendedor).Error
	if err != nil {
		return err
	}
	return nil
}

// GetEmployee Obtener empleado por ID
func GetVendedor(db *gorm.DB, Vendedor *Vendedores, id string) (err error) {
	db.Raw("SELECT * FROM vendedores WHERE idvendedor = ? AND deleted_at IS NULL LIMIT 1", id).Scan(Vendedor)
	if err != nil {
		return err
	}
	return nil
}

// GetEmployees Obtener todos los empleados
func GetVendedores(db *gorm.DB, Vendedor *[]Vendedores, id string) (err error) {

	db.Raw("SELECT * FROM vendedores WHERE deleted_at IS NULL").Scan(Vendedor)
	//err = db.Find(Employee).Error
	if err != nil {
		return err
	}
	return nil
}

func (E *Vendedores) UpdateVendedor(db *gorm.DB, id string) (*Vendedores, error) {
	// Sección de validación de contreña
	db = db.Debug().
		Model(&Vendedores{}).
		Where("idvendedor = ?", id).
		Take(&Vendedores{}).UpdateColumns(
		map[string]interface{}{
			"nombre":          E.Nombre,
			"primerapellido":  E.Primerapellido,
			"segundoapellido": E.Segundoapellido,
			"telefono":        E.Telefono,
			"email":           E.Email,
			"rol":             E.Rol,
			"oficina":         E.Oficina,
		},
	)
	if db.Error != nil {
		return &Vendedores{}, db.Error
	}

	// Visualizar al empleado actualizado
	err := db.Debug().
		Model(&Vendedores{}).
		Where("idvendedor = ?", id).
		Take(&E).Error
	if err != nil {
		return &Vendedores{}, err
	}
	return E, nil
}

// DeleteEmployee Eliminar empleado
func DeleteEmployee(db *gorm.DB, Employee *Vendedores, id string) (err error) {
	db.Where("idvendedor = ?", id).Delete(Employee)
	return nil
}
