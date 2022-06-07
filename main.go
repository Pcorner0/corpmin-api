package main

import (
	"os"

	"github.com/Pcorner0/corpmin-api/controllers"
	"github.com/Pcorner0/corpmin-api/database"
	"github.com/Pcorner0/corpmin-api/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	database.Connect()

	r := gin.Default()

	r.Use(cors.Default())

	public := r.Group("/api")

	userRepo := controllers.NewRepoUsers()
	public.POST("/register", userRepo.RegistrarUser)
	public.POST("/login", userRepo.LogIn)
	public.GET("/logout", userRepo.LogOut)

	vendedoresRepo := controllers.NewRepoVendedor()
	public.POST("/empleados", vendedoresRepo.NewVendedor)
	public.GET("/empleados", vendedoresRepo.GetEmployees)
	public.GET("/empleados/:idempleado", vendedoresRepo.GetVendedor)
	public.PUT("/empleados/:idempleado", vendedoresRepo.UpdateEmployee)
	public.DELETE("/empleados/:idempleado", vendedoresRepo.DeleteEmpleado)

	clienteRepo := controllers.NewRepoClientes()
	//Prospectos
	public.GET("/prospectos", clienteRepo.GetProspectos)
	public.GET("/prospectos/:id", clienteRepo.GetProspecto)

	//Clientes
	public.POST("/clientes", clienteRepo.CreateCliente)
	public.GET("/clientes", clienteRepo.GetClientes)
	public.GET("/clientes/:id", clienteRepo.GetCliente)

	//Operaciones
	operacionesRepo := controllers.NewRepoOperaciones()
	public.POST("/operaciones", operacionesRepo.CreateOperation)
	public.GET("/operaciones", operacionesRepo.GetOperations)
	public.GET("/operaciones/actual", operacionesRepo.GetActualOperations)
	public.GET("/operaciones/actual/:id", operacionesRepo.GetActualOperation)
	public.GET("/operaciones/:id", operacionesRepo.GetOperationCustomer)
	public.GET("/operaciones/empleado/:idempleado", operacionesRepo.GetOperationEmployee)

	private := r.Group("/api/admin")
	private.Use(middlewares.JwtAuthMiddleware())
	private.GET("/user", userRepo.CurrenteUser)

	r.Run(":"+port)
}
