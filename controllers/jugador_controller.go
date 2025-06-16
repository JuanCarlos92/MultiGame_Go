package controllers

import (
    "fmt"
    "juego/models"
    "juego/services"
    "net/http"

    "github.com/gin-gonic/gin"
)

// JugadorController maneja las operaciones relacionadas con los jugadores
type JugadorController struct {
    JugadorService *services.JugadorService
}

// NewJugadorController crea una nueva instancia de JugadorController
// Esta función recibe un servicio de jugador y lo asigna al controlador
func NewJugadorController(jugadorService *services.JugadorService) *JugadorController {
    return &JugadorController{JugadorService: jugadorService}
}

// RegisterInput es la estructura de entrada para el registro de un jugador
// Aquí no necesitamos la contraseña, ya que no la vamos a enviar al cliente
type RegisterInput struct {
    Name     string `json:"name" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

// Register maneja el registro de un nuevo jugador
// Esta función recibe los datos del jugador, crea un nuevo registro y lo guarda en la base de datos
// No devuelve la contraseña al cliente, solo los datos del jugador sin la contraseña
func (controller *JugadorController) Register(c *gin.Context) {
    var input RegisterInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos no válidos"})
        return
    }

    jugador := models.Jugador{
        Name:     input.Name,
        Email:    input.Email,
        Password: input.Password, // Aquí sí tenemos la contraseña
    }

    createdJugador, err := controller.JugadorService.RegisterJugador(&jugador)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"jugador": createdJugador})
}


// LoginInput es la estructura de entrada para el inicio de sesión de un jugador
// Aquí no necesitamos la contraseña, ya que no la vamos a enviar al cliente
func (controller *JugadorController) Login(c *gin.Context) {
    var loginInput struct {
        Email    string `json:"email" binding:"required"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&loginInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos no válidos"})
        return
    }

    fmt.Printf("DEBUG LOGIN: email=%s, password=%s\n", loginInput.Email, loginInput.Password)

    jugador, _, err := controller.JugadorService.LoginJugador(loginInput.Email, loginInput.Password)
    if err != nil {
        fmt.Printf("DEBUG LOGIN ERROR: %v\n", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"jugador": jugador})
}
