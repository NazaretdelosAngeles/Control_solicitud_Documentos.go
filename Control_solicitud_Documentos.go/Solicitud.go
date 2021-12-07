package Solicitud

import "time"

type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Numero    uint      `json:"Numero,omitempty"`
    NumeroActa    uint      `json:"NumeroActa,omitempty"`
    Fecha    uint      `json:"Fecha,omitempty"`
    FechaActa    uint      `json:"FechaActa,omitempty"`
    TipoCivil    string      `json:"TipoCivil,omitempty"`
    CodigoCivil    uint      `json:"CodigoCivil,omitempty"`
    DocuIdeRegistrador    uint      `json:"DocuIdeRegistrador,omitempty"`
    CodigoDocumento    uint      `json:"CodigoDocumento,omitempty"`
    EstadoDocumento    string      `json:"EstadoDocumento,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
}