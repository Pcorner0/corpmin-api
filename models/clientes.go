package models

import (
	"gorm.io/gorm"
)

// Cliente Estructura para los datos de clientes
type Clientes struct {
	gorm.Model
	Id                      int    `json:"id"`
	Rfc                     string `json:"rfc" binding:"required,uppercase,max=50"`
	Razonsocial             string `json:"razonsocial" binding:"required,uppercase,max=250"`
	Calle                   string `json:"calle" binding:"uppercase,max=50"`
	Noext                   string `json:"noext" binding:"uppercase,max=50"`
	Noint                   string `json:"noint" binding:"uppercase,max=50"`
	Cp                      string `json:"cp" binding:"required,uppercase,max=20"`
	Colonia                 string `json:"colonia" binding:"uppercase,max=50"`
	Municipio               string `json:"municipio" binding:"uppercase,max=50"`
	Estado                  string `json:"estado" binding:"uppercase,max=50"`
	Contactonombre          string `json:"contactonombre" binding:"required,uppercase,max=50"`
	Contactoprimerapellido  string `json:"contactoprimerapellido" binding:"uppercase,max=50"`
	Contactosegundoapellido string `json:"contactosegundoapellido" binding:"uppercase,max=50"`
	Contactopuesto          string `json:"contactopuesto" binding:"uppercase,max=50"`
	Email                   string `json:"email" binding:"required,email,max=100"`
	Telefonooficina         string `json:"telefonooficina" binding:"max=50"`
	Telefonopersonal        string `json:"telefonopersonal" binding:"max=50"`
	Prospecto               string `json:"prospecto" binding:"uppercase,max=50"`
	Idvendedor              int    `json:"idvendedor" binding:"required"`
	Estatus                 string `json:"estatus" binding:"required"`
}

//registrar un nuevo empleado
func CreateCliente(db *gorm.DB, Cliente *Clientes) (err error) {
	err = db.Create(Cliente).Error
	if err != nil {
		return err
	}
	return nil
}

//Check if client already exists
func CheckCliente(db *gorm.DB, Cliente *Clientes, rfc string) bool {
	var exists bool
	//db.Model(model).Select("count(*) > 0").Where("id = ?", id).Find(&exists).Error
	_ = db.Select("count(*) > 0").Where("rfc = ?", rfc).Find(&exists).Error

	return exists
}

//Obtener cliente por ID
func GetCliente(db *gorm.DB, Cliente *Clientes, id string) (err error) {
	//err = db.Where("idcliente = ?", id).Find(Cliente).Error
	db.Raw("SELECT * FROM clientes WHERE id = ? AND deleted_at IS NULL LIMIT 1", id).Scan(Cliente)
	if err != nil {
		return err
	}
	return nil
}

//Obtener cliente por ID
func GetProspecto(db *gorm.DB, prospecto *Clientes, id string) (err error) {
	//err = db.Where("idcliente = ?", id).Find(Cliente).Error
	db.Raw("SELECT * FROM clientes WHERE id = ? AND estatus != 'CLIENTE' AND deleted_at IS NULL LIMIT 1 ", id).Scan(prospecto)
	if err != nil {
		return err
	}
	return nil
}

//Obtener todos los clientes
func GetClientes(db *gorm.DB, Cliente *[]Clientes) (err error) {
	db.Raw("SELECT * FROM clientes WHERE estatus = 'CLIENTE' AND deleted_at IS NULL").Scan(Cliente)
	if err != nil {
		return err
	}
	return nil
}

//Obtener todos los empleados
func GetProspectos(db *gorm.DB, ClienteS *[]Clientes) (err error) {
	db.Raw("SELECT * FROM clientes WHERE estatus != 'CLIENTE' AND deleted_at IS NULL").Scan(ClienteS)
	if err != nil {
		return err
	}
	return nil
}

//Obtener todos los empleados
func GetClientesID(db *gorm.DB, Cliente *[]Clientes, id string) (err error) {
	db.Raw("SELECT * FROM clientes WHERE idvendedor = ? AND deleted_at IS NULL", id).Scan(Cliente)
	if err != nil {
		return err
	}
	return nil
}

//Actualizar datos de empleado
func UpdateCliente(db *gorm.DB, Cliente *Clientes) (err error) {
	db.Save(Cliente)

	return nil
}

//Eliminar cliente
func DeleteCliente(db *gorm.DB, Cliente *Clientes, id string) (err error) {
	db.Where("id = ?", id).Delete(Cliente)
	return nil
}
