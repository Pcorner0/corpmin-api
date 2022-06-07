package controllers

import (
	"github.com/Pcorner0/corpmin-api/database"
	"github.com/Pcorner0/corpmin-api/models"

	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type RepoVendedores struct {
	DB *gorm.DB
}

func NewRepoVendedor() *RepoVendedores {
	db := database.InitDB()
	db.AutoMigrate(&models.Vendedores{})
	return &RepoVendedores{DB: db}
}

var validate = validator.New()

// NewEmployee Dar de alta a un Employee
func (repository *RepoVendedores) NewVendedor(c *gin.Context) {
	var Employee models.Vendedores
	c.BindJSON(&Employee)

	err := models.CreateVendedor(repository.DB, &Employee)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Employee)
}

// GetEmployee get Employee by id
func (repository *RepoVendedores) GetVendedor(c *gin.Context) {
	id, _ := c.Params.Get("idempleado")

	var Vendedor models.Vendedores
	err := models.GetVendedor(repository.DB, &Vendedor, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Vendedor)
}

// GetEmployees Get Employee
func (repository *RepoVendedores) GetEmployees(c *gin.Context) {
	var Employee []models.Vendedores
	id, _ := c.Params.Get("idempleado")
	err := models.GetVendedores(repository.DB, &Employee, id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Employee)
}

// UpdateEmployee Update Employee
func (repository *RepoVendedores) UpdateEmployee(c *gin.Context) {
	var Vendedor models.Vendedores
	id, _ := c.Params.Get("idvendedor")
	err := models.GetVendedor(repository.DB, &Vendedor, id)

	c.JSON(http.StatusOK, Vendedor)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

}

// DeleteEmpleado Delete Empleado
func (repository *RepoVendedores) DeleteEmpleado(c *gin.Context) {
	var Vendedor models.Vendedores
	id, _ := c.Params.Get("idvendedor")
	err := models.DeleteEmployee(repository.DB, &Vendedor, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})

}
