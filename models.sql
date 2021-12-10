CREATE TABLE IF NOT EXISTS Civil (
    id serial NOT NULL,
    Nombres VARCHAR(150) NOT NULL,
    Apellidos VARCHAR(150) NOT NULL,
    DocuIde int(150) NOT NULL UNIQUE,
    Direccion varchar(256) NULL,
    email VARCHAR(150) NULL UNIQUE,
    Nacionalidad VARCHAR(256) NULL,
    EstadoCivil VARCHAR(150) NULL,
    Genero VARCHAR(150) NULL,
    FechaNacimiento datetime NULL,
    LugarNacimiento VARCHAR(150) NULL,
    FechaFallecimiento datetime NULL,
    LugarFallecimiento VARCHAR(150) NULL,
    CausasFallecimiento VARCHAR(150) NULL,
    Profesion VARCHAR(150) NULL,
    NroPasaporte int(150) NULL UNIQUE,
    NombresApellidosPadre VARCHAR(150) NULL,
    DocuIdePadre int(150) NULL UNIQUE,
    ProfesionPadre VARCHAR(150) NULL,
    ResidenciaPadre VARCHAR(150) NULL,
    NombresApellidosMadre VARCHAR(150) NULL,
    DocuIdeMadre int(150) NULL UNIQUE,
    ProfesionMadre VARCHAR(150) NULL,
    ResidenciaMadre VARCHAR(150) NULL,
    NombresApellidoshijo1 VARCHAR(150) NULL,
    DocuIdehijo1 int(150) NULL UNIQUE,
    EdadHijo1 int(150) NULL,
    ActaNaciNro1 int(150) NULL UNIQUE,
    FechaNaci1 datetime NULL,
    NombresApellidoshijo2 VARCHAR(150) NULL,
    DocuIdehijo2 int(150) NULL UNIQUE,
    EdadHijo2 int(150) NULL,
    ActaNaciNro2 int(150) NULL UNIQUE,
    FechaNaci2 datetime NULL,
    NombresApellidoshijo3 VARCHAR(150) NULL,
    DocuIdehijo3 int(150) NULL UNIQUE,
    EdadHijo3 int(150) NULL,
    ActaNaciNro3 int(150) NULL UNIQUE,
    FechaNaci3 datetime NULL,
    NombresApellidoshijo4 VARCHAR(150) NULL,
    DocuIdehijo4 int(150) NULL UNIQUE,
    EdadHijo4 int(150) NULL,
    ActaNaciNro4 int(150) NULL UNIQUE,
    FechaNaci4 datetime NULL,
    Comunidad VARCHAR(150) NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    PRIMARY KEY(DocuIde)
);

CREATE TABLE IF NOT EXISTS Documento (
    id serial NOT NULL,
    Codigo int NOT NULL UNIQUE,
    Fecha datetime NOT NULL,
    Tipo VARCHAR(150) NOT NULL,
    Descripcion VARCHAR(150) NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
	PRIMARY KEY(Codigo),
    CONSTRAINT fk_Documento_Tipo FOREIGN KEY(Tipo) REFERENCES Solicitud(TipoDocumento)
);

CREATE TABLE IF NOT EXISTS Registrador (
    id serial NOT NULL,
    Nombres VARCHAR(150) NOT NULL,
    Apellidos VARCHAR(150) NOT NULL,
    DocuIde int(150) NOT NULL UNIQUE,
    Oficina VARCHAR(256) NOT NULL,
    Resolucion VARCHAR(150) NOT NULL,
    Fecha datetime NOT NULL,
    Resolucion1 VARCHAR(150) NOT NULL,
    Fecha1 datetime NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    PRIMARY KEY(DocuIde)
);

CREATE TABLE IF NOT EXISTS Solicitud (
    id serial NOT NULL,
    Numero int(150) NOT NULL UNIQUE,
    NumeroActa int(150) NOT NULL UNIQUE,
    FechaActa datetime NOT NULL,
    DocuIdeCivil int(150) NOT NULL UNIQUE,
    DocuIdeRegistrador int(150) NOT NULL UNIQUE,
    CodigoDocumento int(150) NOT NULL UNIQUE,
    EstadoDocumento text NOT NULL,
    TipoDocumento VARCHAR(150) NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    PRIMARY KEY(TipoDocumento),
    FOREIGN KEY (DocuIdeCivil) REFERENCES Civil(DocuIde),
    FOREIGN KEY (CodigoDocumento) REFERENCES Documento(Codigo),
    FOREIGN KEY (DocuIdeRegistrador) REFERENCES Registrador(DocuIde)

);