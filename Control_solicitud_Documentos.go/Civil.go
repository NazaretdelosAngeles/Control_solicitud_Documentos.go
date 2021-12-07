package Civil

import "time"

type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nombres    string      `json:"Nombres,omitempty"`
    Apellidos    string      `json:"Apellidos,omitempty"`
    DocuIde    uint      `json:"DocuIde,omitempty"`
    Direccion    string      `json:"Direccion,omitempty"`
    Email    string      `json:"Email,omitempty"`
    Nacionalidad    string      `json:"Nacionalidad,omitempty"`
    EstadoCivil    string      `json:"EstadoCivil,omitempty"`
    Genero    string      `json:"Genero,omitempty"`
    FechaNacimiento    uint      `json:"FechaNacimiento,omitempty"`
    LugarNacimiento    string      `json:"LugarNacimiento,omitempty"`
    FechaFallecimiento    uint      `json:"FechaFallecimiento,omitempty"`
    LugarFallecimiento    string      `json:"LugarNacimiento,omitempty"`
    CausasFallecimiento    string      `json:"LugarNacimiento,omitempty"`
    Profesion    string      `json:"Profesion,omitempty"`
    NroPasaporte    uint      `json:"NroPasaporte,omitempty"`
    NombresApellidosPadre    string      `json:"NombresApellidosPadre,omitempty"`
    DocuIdePadre    uint      `json:"DocuIdePadre,omitempty"`
    ProfesionPadre    string      `json:"ProfesionPadre,omitempty"`
    ResidenciaPadre    string      `json:"ResidenciaPadre,omitempty"`
    NombresApellidosMadre    string      `json:"NombresApellidosMadre,omitempty"`
    DocuIdeMadre    uint      `json:"DocuIdeMadre,omitempty"`
    ProfesionMadre    string      `json:"ProfesionMadre,omitempty"`
    ResidenciaMadre    string      `json:"ResidenciaMadre,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    NombresApellidoshijo1    string      `json:"NombresApellidoshijo1,omitempty"`
    DocuIdehijo1    uint      `json:"DocuIdehijo1,omitempty"`
    EdadHijo1    uint      `json:"EdadHijo1,omitempty"`
    ActaNaciNro1    uint      `json:"ActaNaciNro1,omitempty"`
    FechaNaci1    uint      `json:"FechaNaci1,omitempty"`
    NombresApellidoshijo2    string      `json:"NombresApellidoshijo2,omitempty"`
    DocuIdehijo2    uint      `json:"DocuIdehijo2,omitempty"`
    EdadHijo2    uint      `json:"EdadHijo2,omitempty"`
    ActaNaciNro2    uint      `json:"ActaNaciNro2,omitempty"`
    FechaNaci2    uint      `json:"FechaNaci2,omitempty"`
    NombresApellidoshijo3    string      `json:"NombresApellidoshijo3,omitempty"`
    DocuIdehijo3    uint      `json:"DocuIdehijo3,omitempty"`
    EdadHijo3    uint      `json:"EdadHijo3,omitempty"`
    ActaNaciNro3    uint      `json:"ActaNaciNro3,omitempty"`
    FechaNaci3    uint      `json:"FechaNaci3,omitempty"`
    NombresApellidoshijo4    string      `json:"NombresApellidoshijo4,omitempty"`
    DocuIdehijo4    uint      `json:"DocuIdehijo4,omitempty"`
    EdadHijo4    uint      `json:"EdadHijo4,omitempty"`
    ActaNaciNro4    uint      `json:"ActaNaciNro4,omitempty"`
    FechaNaci4    uint      `json:"FechaNaci4,omitempty"`
    Comunidad    string      `json:"Comunidad,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
  }