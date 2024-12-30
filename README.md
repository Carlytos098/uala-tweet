# Uala-Tweet

Uala-Tweet es una plataforma de microblogging similar a Twitter, construida con un enfoque en la simplicidad y la eficiencia. Permite a los usuarios publicar tweets, seguir a otros usuarios y ver su línea de tiempo.

## Tabla de Contenidos

- [Asumptions](#assumptions)
- [Documentación Adicional](#documentación-adicional)
- [Características](#características)
- [Tecnologías Usadas](#tecnologías-usadas)
- [Requisitos Previos](#requisitos-previos)
- [Instalación y Configuración](#instalación-y-configuración)
- [Levantamiento del Proyecto](#levantamiento-del-proyecto)
- [Datos Pre-cargados](#datos-pre-cargados)
- [Uso de la API](#uso-de-la-api)
- [Ejemplos de cURL](#ejemplos-de-curl)
- [API Desplegada en Render](#api-desplegada-en-render)
- [Contribución](#contribución)
- [Licencia](#licencia)

## Asumptions

Para comprender mejor las decisiones de diseño y las limitaciones de la aplicación, consulta el archivo [ASSUMPTIONS.txt](ASSUMPTIONS.txt). Este documento describe las suposiciones que se hicieron durante el desarrollo y los aspectos pendientes a mejorar.

## Documentación Adicional

Para obtener información más detallada sobre la arquitectura y el diseño de **Uala-Tweet**, consulta nuestra [Wiki](https://github.com/Carlytos098/uala-tweet/wiki) donde encontrarás documentos sobre la arquitectura, RFCs de cambios, y tipos de datos en los endpoints.

## Características

- Registro y autenticación de usuarios.
- Publicación de tweets.
- Permitir a los usuarios seguir a otros usuarios.
- Recuperación de la línea de tiempo del usuario.

## Tecnologías Usadas

- **Lenguaje**: Go (Golang)
- **Framework**: HTTP estándar de Go
- **Almacenamiento**: Repositorio en memoria (Memory Repository)

## Requisitos Previos

Asegúrate de tener instalado:

- [Go](https://golang.org/dl/) (versión 1.23.2 o superior).

## Instalación y Configuración

1. Clona este repositorio en tu máquina local:
   ```bash
   git clone https://github.com/tu_usuario/uala-tweet.git
   cd uala-tweet
   ```

2. Asegúrate de que tienes Go instalado y el entorno correctamente configurado.

## Levantamiento del Proyecto

Para ejecutar la aplicación, sigue estos pasos:

1. Ve al directorio del proyecto:
   ```bash
   cd uala-tweet
   ```

2. Compila y ejecuta la aplicación:
   ```bash
   go run main.go
   ```

3. La aplicación se ejecutará en `http://localhost:8080`.

## Datos Pre-cargados

La aplicación ya viene con algunos usuarios y tweets pre-cargados para que puedas probar las funcionalidades de manera inmediata. Los usuarios precargados son:

- **Usuarios**:
  - user1
  - user2
  - user3
- **Usuario "user2"** tiene un tweet: "Este es un tweet del user2."

## Uso de la API

La API tiene varias rutas disponibles para interactuar con la aplicación:

- **Publicar un Tweet**
  - **Endpoint**: `POST /tweets`
  - **Middleware**: Requiere autenticación.
  - **Cuerpo de la solicitud**: JSON que contenga el contenido del tweet.

- **Seguir a un Usuario**
  - **Endpoint**: `POST /follow`
  - **Middleware**: Requiere autenticación.
  - **Cuerpo de la solicitud**: JSON que especifique el usuario a seguir.

- **Obtener la Línea de Tiempo**
  - **Endpoint**: `GET /timeline`
  - **Middleware**: Requiere autenticación.
  - **Respuesta**: Lista de tweets del usuario y sus seguidos.

## Ejemplos de cURL

A continuación, se presentan ejemplos de cómo utilizar `cURL` para interactuar con la API:

### 1. Publicar un Tweet

```bash
curl --location 'http://localhost:8080/tweets' \
--header 'Content-Type: application/json' \
--header 'User-ID: user2' \
--data '{"content": "Tweet de user2"}'
```

**Explicación**:
- **Método**: POST
- **URL**: `http://localhost:8080/tweets`
- **Headers**:
  - `Content-Type: application/json`: Indica que el cuerpo de la solicitud es un JSON.
  - `User-ID: user2`: Identifica al usuario que está publicando el tweet.
- **Cuerpo de la solicitud**: Un JSON que contiene el contenido del tweet.
- **Descripción**: Este comando publica un nuevo tweet en la API en nombre del usuario "user2".

### 2. Seguir a un Usuario

```bash
curl --location --request POST 'http://localhost:8080/follow?followID=user3' \
--header 'User-ID: user1'
```

**Explicación**:
- **Método**: POST
- **URL**: `http://localhost:8080/follow?followID=user3`
- **Headers**:
  - `User-ID: user1`: Identifica al usuario que está realizando la acción de seguir.
- **Descripción**: Este comando permite al usuario "user1" seguir al usuario "user3". El ID del usuario a seguir se pasa como parámetro en la URL.

### 3. Obtener la Línea de Tiempo

```bash
curl --location 'http://localhost:8080/timeline' \
--header 'User-ID: user1'
```

**Explicación**:
- **Método**: GET
- **URL**: `http://localhost:8080/timeline`
- **Headers**:
  - `User-ID: user1`: Indica qué usuario solicita su línea de tiempo.
- **Descripción**: Este comando recupera la línea de tiempo (tweets) del usuario "user1" y de las personas a las que sigue.

## API Desplegada en Render

La API también está disponible públicamente a través de Render. Aquí están los ejemplos de `cURL` apuntando a la API desplegada:

### 1. Obtener la Línea de Tiempo

```bash
curl --location 'https://uala-tweet-2.onrender.com/timeline' \
--header 'User-ID: user1'
```

### 2. Publicar un Tweet

```bash
curl --location 'https://uala-tweet-2.onrender.com/tweets' \
--header 'Content-Type: application/json' \
--header 'User-ID: user2' \
--data '{"content": "Otro Tweet de prueba para el user2"}'
```

### 3. Seguir a un Usuario

```bash
curl --location --request POST 'https://uala-tweet-2.onrender.com/follow?followID=user3' \
--header 'User-ID: user1'
```

## Contribución

Las contribuciones son bienvenidas. Si deseas contribuir a este proyecto:

1. Realiza un fork del repositorio.
2. Crea una nueva rama para tus cambios:
   ```bash
   git checkout -b mi-nueva-funcionalidad
   ```
3. Realiza tus cambios y añade tus archivos:
   ```bash
   git add .
   ```
4. Confirma tus cambios:
   ```bash
   git commit -m "Descripción de mis cambios"
   ```
5. Sube tus cambios:
   ```bash
   git push origin mi-nueva-funcionalidad
   ```
6. Abre un Pull Request en GitHub.

## Licencia

Este proyecto está bajo la Licencia MIT. Para más detalles, consulta el archivo [LICENSE](LICENSE).
