package Documento

import "time"

type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Codigo    uint      `json:"Codigo,omitempty"`
    Fecha    uint      `json:"Fecha,omitempty"`
    Tipo    string      `json:"Tipo,omitempty"`
    Descripcion    string      `json:"Descripcion,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
}