package Registrador

import "time"


type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nombres    string      `json:"Nombres,omitempty"`
    Apellidos    string      `json:"Apellidos,omitempty"`
    DocuIde    uint      `json:"DocuIde,omitempty"`
    Oficina    string      `json:"Oficina,omitempty"`
    Resolucion    uint      `json:"Resolucion,omitempty"`
    Fecha    uint      `json:"Fecha,omitempty"`
    Resolucion    uint      `json:"Resolucion,omitempty"`
    Fecha    uint      `json:"Fecha,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
}