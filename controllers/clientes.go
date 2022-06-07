package controllers

import (
	
	"errors"
	"fmt"
	"net/http"

	"github.com/Pcorner0/corpmin-api/models"
	"github.com/Pcorner0/corpmin-api/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RepoClientes struct {
	DB *gorm.DB
}

func NewRepoClientes() *RepoClientes {
	db := database.InitDB()
	db.AutoMigrate(&models.Clientes{})
	return &RepoClientes{DB: db}
}

// CreateCliente Dar de alta a un nuevo cliente
func (repository *RepoClientes) CreateCliente(c *gin.Context) {
	var Cliente models.Clientes
	c.BindJSON(&Cliente)
	fmt.Println(Cliente.Rfc)
	errCliente := models.CheckCliente(repository.DB, &Cliente, Cliente.Rfc)

	if errCliente == true {
		c.JSON(http.StatusOK, "Cliente previamente registrado")
		return
	}
	err := models.CreateCliente(repository.DB, &Cliente)

	if err != nil {
		c.JSON(http.StatusOK, Cliente)
		return
	}
	c.JSON(http.StatusOK, Cliente)
}

// GetCliente get user by id
func (repository *RepoClientes) GetCliente(c *gin.Context) {
	id, _ := c.Params.Get("id")
	fmt.Println(id)
	var Cliente models.Clientes
	err := models.GetCliente(repository.DB, &Cliente, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Cliente)
}

// GetProspecto get user by id
func (repository *RepoClientes) GetProspecto(c *gin.Context) {
	id, _ := c.Params.Get("id")
	fmt.Println(id)
	var prospecto models.Clientes
	err := models.GetProspecto(repository.DB, &prospecto, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, prospecto)
}

// GetClients Get clientes
func (repository *RepoClientes) GetClientes(c *gin.Context) {
	var Cliente []models.Clientes
	err := models.GetClientes(repository.DB, &Cliente)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Cliente)
}

func (repository *RepoClientes) GetProspectos(c *gin.Context) {
	var prospectos []models.Clientes
	err := models.GetProspectos(repository.DB, &prospectos)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, prospectos)
}

// GetClients Get clientes by idempleado
func (repository *RepoClientes) GetClientesID(c *gin.Context) {
	var Cliente []models.Clientes
	id, _ := c.Params.Get("idvendedor")
	err := models.GetClientesID(repository.DB, &Cliente, id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Cliente)
}

// UpdateCliente Update Cliente
func (repository *RepoClientes) UpdateCliente(c *gin.Context) {
	var Cliente models.Clientes
	id, _ := c.Params.Get("idvendedor")
	err := models.GetCliente(repository.DB, &Cliente, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&Cliente)
	err = models.UpdateCliente(repository.DB, &Cliente)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Cliente)
}

// DeleteCliente Delete Empleado
func (repository *RepoClientes) DeleteCliente(c *gin.Context) {
	var Cliente models.Clientes
	id, _ := c.Params.Get("id")
	err := models.DeleteCliente(repository.DB, &Cliente, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
