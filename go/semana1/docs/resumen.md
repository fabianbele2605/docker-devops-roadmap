# Semana 1 — Docker Básico

## Lo que aprendiste

### Docker fundamentals
- **Imagen vs contenedor:** la imagen es la plantilla inmutable, el contenedor es la instancia en ejecución
- **Por qué Docker es liviano:** no incluye kernel, lo comparte con el host
- **Ciclo de vida:** `run` → `stop` → `start` → `rm`
- **Flags importantes:**
  - `-d` — corre en background (detached)
  - `-it` — modo interactivo con terminal
  - `--rm` — elimina el contenedor al salir
  - `-p host:contenedor` — mapeo de puertos
  - `--name` — nombra el contenedor

### Comandos dominados
```bash
docker run
docker ps / docker ps -a
docker stop / docker start / docker rm
docker images
docker logs
docker build -t nombre:tag .
```

### Dockerfile
```dockerfile
FROM golang:1.24       # imagen base
WORKDIR /app           # directorio de trabajo
COPY . .               # copia archivos
RUN go build -o api .  # compila durante el build
EXPOSE 8080            # documenta el puerto
CMD ["./api"]          # comando de arranque
```

### Go — estructura básica
```go
package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
    // w = respuesta, r = petición
}

func main() {
    http.HandleFunc("/ruta", handler)
    http.ListenAndServe(":8080", nil)
}
```

## Proyecto completado

API en Go con endpoint `/health` contenerizada en Docker.

```bash
docker build -t api-go:v1 .
docker run -d -p 8082:8080 api-go:v1
curl http://localhost:8082/health
# {"status": "ok"}
```

## Pendiente para Semana 2

- La imagen pesa **987MB** — demasiado para producción
- Solución: **multi-stage builds** con `AS builder` + imagen `scratch`
- Objetivo: reducir a ~10MB