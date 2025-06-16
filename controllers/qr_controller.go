package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)
// QRController maneja las operaciones relacionadas con códigos QR
type QRController struct{}

// NewQRController crea una nueva instancia de QRController
func NewQRController() *QRController {
	return &QRController{}
}

// Genera un código QR (UUID)
func (qc *QRController) GenerateQR(c *gin.Context) {
	qrCode := uuid.New().String()
	c.JSON(http.StatusOK, gin.H{
		"qr_code": qrCode,
	})
}

// Verifica si el QR ha sido usado (simulado como pendiente)
func (qc *QRController) CheckQRStatus(c *gin.Context) {
	var input struct {
		QRCode string `json:"qr_code"`
	}

	if err := c.ShouldBindJSON(&input); err != nil || input.QRCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "QR inválido"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "pending"})
}

// Simula login con QR
func (qc *QRController) LoginWithQR(c *gin.Context) {
	var input struct {
		QRData    string `json:"qr_data"`
		Timestamp int64  `json:"timestamp"`
	}

	if err := c.ShouldBindJSON(&input); err != nil || input.QRData == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
// Unirse a una sala de juego mediante un código QR
func (qc *QRController) JoinGameRoom(c *gin.Context) {
	var input struct {
		QRCode string `json:"qr_code"`
		Player string `json:"player"`
	}

	if err := c.ShouldBindJSON(&input); err != nil || input.QRCode == "" || input.Player == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Simula búsqueda de sala con QRCode (esto se debería conectar a tu lógica de salas)
	roomExists := true // Supón que sí existe
	if !roomExists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sala no encontrada"})
		return
	}

	// Simula unir al jugador
	// Aquí podrías agregar lógica para registrar al jugador en la sala

	c.JSON(http.StatusOK, gin.H{
		"message": "Jugador unido a la sala exitosamente",
		"player":  input.Player,
		"room":    input.QRCode,
	})
}