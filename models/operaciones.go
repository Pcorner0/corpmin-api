package models

import (
	"gorm.io/gorm"
	"time"
)

// Operaciones Accion Estructura para operaciones nuevas
type Operaciones struct {
	Id            int       `json:"id" binding:"required"`
	RFC           string    `json:"rfc" binding:"required,uppercase,max=50"`
	Razonsocial   string    `json:"razonsocial" binding:"required,uppercase,max=250"`
	Actividad     string    `json:"actividad" binding:"required,uppercase,max=50"`
	Estatus       string    `json:"estatus" binding:"required,uppercase,max=50"`
	Idvendedor    int       `json:"idvendedor" binding:"required"`
	Observaciones string    `json:"observaciones" binding:"uppercase,max=250"`
	CreatedAt     time.Time `json:"created_at" gorm:"primaryKey;"`
}

// CreateOperation registrar una nuevo empleado
func CreateOperation(db *gorm.DB, Operation *Operaciones) (err error) {
	err = db.Create(Operation).Error
	if err != nil {
		return err
	}
	return nil
}

// GetOperationCustomer Obtener operaciones por IDcliente
func GetOperationCustomer(db *gorm.DB, Operation *[]Operaciones, id string) (err error) {
	db.Raw("SELECT * FROM operaciones WHERE id = ? ORDER BY created_at DESC", id).Scan(Operation)
	if err != nil {
		return err
	}
	return nil
}

// GetEmployee Obtener empleado por IDempleado
func GetOperationEmployee(db *gorm.DB, Operation *[]Operaciones, id string) (err error) {
	db.Raw("SELECT * FROM operaciones WHERE id = ? ORDER BY created_at DESC", id).Scan(Operation)
	if err != nil {
		return err
	}
	return nil
}

// GetOperations Obtener todas las operaciones
func GetOperations(db *gorm.DB, Operation *[]Operaciones) (err error) {
	db.Raw("SELECT * FROM operaciones ORDER BY created_at DESC").Scan(Operation)
	if err != nil {
		return err
	}
	return nil
}

// GetOperations Obtener todas las operaciones
func GetActualOperations(db *gorm.DB, Operation *[]Operaciones) (err error) {
	db.Raw("SELECT * FROM operaciones INNER JOIN ( SELECT id AS idSecondary, MAX(created_at) AS maxsign FROM operaciones GROUP BY idSecondary) ms ON operaciones.id = ms.idSecondary AND created_at = maxsign").Scan(Operation)
	if err != nil {
		return err
	}
	return nil
}

// GetOperations Obtener todas las operaciones
func GetActualOperation(db *gorm.DB, Operation *[]Operaciones, id string) (err error) {
	db.Raw("SELECT * FROM operaciones INNER JOIN ( SELECT id AS idSecondary, MAX(created_at) AS maxsign FROM operaciones GROUP BY idSecondary) ms ON operaciones.id = ms.idSecondary AND created_at = maxsign WHERE operaciones.id= ?", id).Scan(Operation)
	if err != nil {
		return err
	}
	return nil
}
