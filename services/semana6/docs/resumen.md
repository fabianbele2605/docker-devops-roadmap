# Semana 6 — Traefik Reverse Proxy

## Lo que aprendiste

### ¿Por qué un reverse proxy?
En producción no se exponen puertos individuales de cada servicio. Traefik recibe todo el tráfico en el puerto 80/443 y enruta internamente:

```
internet
    ↓
Traefik :80
    ├── PathPrefix(/users)  → go_api:8080
    ├── PathPrefix(/health) → go_api:8080
    └── PathPrefix(/login)  → rust_api:8081
```

### Traefik vs Nginx
| | Traefik | Nginx |
|---|---|---|
| Config | Labels en contenedores | Archivo manual |
| Recarga | Automática (sin restart) | Manual |
| Discovery | Lee Docker en tiempo real | Estático |

### Labels — cómo cada servicio se registra en Traefik
```yaml
labels:
  - "traefik.http.routers.go_api.entrypoints=web"
  - "traefik.http.routers.go_api.rule=PathPrefix(`/users`) || PathPrefix(`/health`)"
  - "traefik.http.services.go_api.loadbalancer.server.port=8080"
```
- **router.rule:** condición de enrutamiento (path, host, headers)
- **loadbalancer.server.port:** puerto interno del contenedor destino

### Socket Proxy — buena práctica de seguridad
Montar `/var/run/docker.sock` directamente en un contenedor da acceso root al host. La solución es interponer un proxy que solo permite operaciones de lectura:

```
Traefik → socket-proxy:2375 → /var/run/docker.sock → Docker daemon
```

```yaml
socket-proxy:
  image: tecnativa/docker-socket-proxy
  environment:
    CONTAINERS: 1   # solo permite leer contenedores
  volumes:
    - /var/run/docker.sock:/var/run/docker.sock:ro
```

### docker-compose.yml — estructura completa
```yaml
traefik:
  image: traefik:v2.10
  command:
    - "--providers.docker.endpoint=tcp://socket-proxy:2375"
    - "--entrypoints.web.address=:80"
  depends_on:
    - socket-proxy

socket-proxy:
  image: tecnativa/docker-socket-proxy
  environment:
    CONTAINERS: 1
  volumes:
    - /var/run/docker.sock:/var/run/docker.sock:ro

go_api:
  # sin "ports:" — ya no se expone directamente
  labels:
    - "traefik.http.routers.go_api.rule=PathPrefix(`/users`)"
    - "traefik.http.services.go_api.loadbalancer.server.port=8080"
```

## Limitación conocida — Docker Desktop en Windows

El socket de Docker Desktop en Windows + WSL2 tiene una capa de seguridad adicional que bloquea el acceso desde contenedores (HTTP 400 en `/version`). **En un VPS Linux esto funciona correctamente.**

Diagnóstico aplicado:
1. Verificar permisos del socket: `docker exec traefik ls -la /var/run/docker.sock`
2. Revisar logs del proxy: `docker compose logs socket-proxy`
3. El error `400` en `GET /v1.24/version` confirma el bloqueo de Docker Desktop

## Próximo — Semana 7

- Prometheus para métricas
- Exponer `/metrics` desde Go y Rust
- Scraping automático de contenedores
- Alertas y monitoreo
