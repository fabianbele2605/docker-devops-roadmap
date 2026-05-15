# Semana 2 — Dockerfiles Profesionales

## Lo que aprendiste

### Multi-stage builds

Técnica para reducir el tamaño final de una imagen usando múltiples etapas en el mismo Dockerfile. Solo la última etapa termina en la imagen final.

```dockerfile
# Etapa 1: compilar (imagen grande con compilador)
FROM golang:1.24 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o api .

# Etapa 2: ejecutar (imagen mínima, solo el binario)
FROM scratch
COPY --from=builder /app/api .
EXPOSE 8080
CMD ["./api"]
```

### Resultado

| Imagen | Tamaño |
|--------|--------|
| `api-go:v1` (sin multi-stage) | 987 MB |
| `api-go:v2` (con multi-stage) | 8 MB |

### Conceptos clave

- **`AS builder`** — nombra una etapa para referenciarla después
- **`COPY --from=builder`** — copia archivos de una etapa anterior
- **`scratch`** — imagen vacía, sin OS ni librerías
- **`CGO_ENABLED=0`** — compila sin dependencias a librerías C (binario estático)
- **`GOOS=linux`** — compila explícitamente para Linux

### Por qué CGO_ENABLED=0 es necesario con scratch

`golang:1.24` incluye Debian con glibc — el binario encuentra las librerías que necesita ahí.  
`scratch` no tiene nada — ni OS, ni librerías. El binario debe ser completamente autónomo.  
`CGO_ENABLED=0` produce un binario estático que no depende de ninguna librería externa.

## Proyecto completado

API en Go con imagen optimizada de 8MB.

```bash
docker build -t api-go:v2 .
docker run -d -p 8082:8080 api-go:v2
curl http://localhost:8082/health
# {"status": "ok"}
```

## Próximo — Semana 3

- Docker Compose
- Múltiples servicios coordinados
- API Go + PostgreSQL
- Redes personalizadas entre contenedores
