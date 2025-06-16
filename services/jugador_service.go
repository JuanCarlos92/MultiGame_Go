package services

import (
	"encoding/json"
	"errors"
    "juego/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type JugadorService struct {
    DB *gorm.DB
}

func NewJugadorService(db *gorm.DB) *JugadorService {
    return &JugadorService{DB: db}
}

func (service *JugadorService) RegisterJugador(jugador *models.Jugador) (*models.Jugador, error) {
    var existing models.RegisterInput
    if err := service.DB.Where("Email = ?", jugador.Email).First(&existing).Error; err == nil {
        return nil, errors.New("el correo ya está registrado")
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(jugador.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, errors.New("no se pudo procesar la contraseña")
    }
    jugador.Password = string(hashedPassword)

    if err := service.DB.Create(jugador).Error; err != nil {
        return nil, err
    }

    return &models.Jugador{
        ID:    jugador.ID,
        Name:  jugador.Name,
        Email: jugador.Email,
        Password: jugador.Password, // No enviar la contraseña en la respuesta
    }, nil
}

// service.go
func (service *JugadorService) LoginJugador(email, password string) ([]byte, int, error) {
    var jugador models.Jugador
    err := service.DB.Where("Email = ?", email).First(&jugador).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            resp := map[string]string{"error": "correo o contraseña incorrectos"}
            jsonResp, _ := json.Marshal(resp)
            return jsonResp, 401, nil
        }
        resp := map[string]string{"error": "error interno del servidor"}
        jsonResp, _ := json.Marshal(resp)
        return jsonResp, 500, nil
    }

    if err := bcrypt.CompareHashAndPassword([]byte(jugador.Password), []byte(password)); err != nil {
        resp := map[string]string{"error": "correo o contraseña incorrectos"}
        jsonResp, _ := json.Marshal(resp)
        return jsonResp, 401, nil
    }

    resp := map[string]interface{}{
        "id":    jugador.ID,
        "name":  jugador.Name,
        "email": jugador.Email,
    }
    jsonResp, err := json.Marshal(resp)
    if err != nil {
        respErr := map[string]string{"error": "error interno del servidor"}
        jsonResp, _ = json.Marshal(respErr)
        return jsonResp, 500, nil
    }

    return jsonResp, 200, nil
}
