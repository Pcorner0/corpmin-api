package controllers

import (
	"errors"
	"net/http"

	"github.com/Pcorner0/corpmin-api/database"
	"github.com/Pcorner0/corpmin-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RepoOperaciones struct {
	DB *gorm.DB
}

func NewRepoOperaciones() *RepoOperaciones {
	db := database.InitDB()
	db.AutoMigrate(&models.Operaciones{})
	return &RepoOperaciones{DB: db}
}

// CreateCliente Dar de alta a una nueva operacion
func (repository *RepoOperaciones) CreateOperation(c *gin.Context) {
	var Operacion models.Operaciones

	c.BindJSON(&Operacion)
	err := models.CreateOperation(repository.DB, &Operacion)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Operacion)
}

// GetEmployee get operation by idcliente
func (repository *RepoOperaciones) GetOperationCustomer(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var Operation []models.Operaciones
	err := models.GetOperationCustomer(repository.DB, &Operation, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Operation)
}

// GetEmployee get operation by idempleado
func (repository *RepoOperaciones) GetOperationEmployee(c *gin.Context) {
	id, _ := c.Params.Get("idvendedor")
	var Operation []models.Operaciones
	err := models.GetOperationEmployee(repository.DB, &Operation, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Operation)
}

// GetEmployees Get Employee
func (repository *RepoOperaciones) GetOperations(c *gin.Context) {
	var Operation []models.Operaciones
	err := models.GetOperations(repository.DB, &Operation)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Operation)
}

// GetActualOperations Get operations
func (repository *RepoOperaciones) GetActualOperations(c *gin.Context) {
	var Operation []models.Operaciones
	err := models.GetActualOperations(repository.DB, &Operation)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Operation)
}

func (repository *RepoOperaciones) GetActualOperation(c *gin.Context) {
	var Operation []models.Operaciones

	id, _ := c.Params.Get("id")

	err := models.GetActualOperation(repository.DB, &Operation, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Operation)
}
