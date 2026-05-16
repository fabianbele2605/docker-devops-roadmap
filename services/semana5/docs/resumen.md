# Semana 5 — Microservicios

## Lo que aprendiste

### Arquitectura de microservicios
Cada servicio tiene una responsabilidad única y se despliega de forma independiente:

```
go_api (Go/Gin)     → CRUD de usuarios, PostgreSQL + Redis
rust_api (Rust/Axum) → Autenticación, /login
```

### Dos lenguajes, un stack
- **Go + Gin + pgx:** para servicios que necesitan acceso a datos, alto throughput
- **Rust + Axum:** para servicios críticos donde el rendimiento y seguridad son prioritarios
- Ambos se comunican a través de la red interna de Docker

### Dockerfile de Rust — patrón de cache de dependencias
```dockerfile
# Paso 1: compilar dependencias con dummy main (se cachea)
RUN mkdir src && echo "fn main() {}" > src/main.rs && cargo build --release

# Paso 2: copiar código real y forzar recompilación
COPY src ./src
RUN touch src/main.rs && cargo build --release
```
**Por qué `touch`:** sin él, Cargo no detecta el cambio de timestamp y ejecuta el binario vacío del paso 1.

### Axum — servidor HTTP en Rust
```rust
#[tokio::main]
async fn main() {
    let app = Router::new()
        .route("/health", get(health))
        .route("/login", post(login));

    let listener = tokio::net::TcpListener::bind("0.0.0.0:8081").await.unwrap();
    axum::serve(listener, app).await.unwrap();
}
```

### Docker Compose con múltiples servicios
```yaml
services:
  go_api:
    depends_on:
      db:
        condition: service_healthy
      redis:
        condition: service_started

  rust_api:
    depends_on:
      db:
        condition: service_healthy
```

### Lección importante — Exit 0 sin output
Si un contenedor sale con código 0 y sin logs, el binario que se ejecutó no es el real. En Rust, verificar el patrón de cache del Dockerfile con `touch`.

## Proyecto completado

4 servicios corriendo en red interna:

```bash
docker compose up -d --build
curl http://localhost:8082/health          # {"status":"ok"}
curl http://localhost:8081/health          # ok
curl -X POST http://localhost:8082/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Fabian","email":"fabian@test.com"}'
curl -X POST http://localhost:8081/login \
  -H "Content-Type: application/json" \
  -d '{"username":"fabian","password":"1234"}'
# {"token":"token_for_fabian"}
```

## Próximo — Semana 6

- Traefik como reverse proxy
- Eliminar puertos expuestos directos
- Enrutar por path: `/users` → go_api, `/login` → rust_api
- Un solo punto de entrada al sistema
