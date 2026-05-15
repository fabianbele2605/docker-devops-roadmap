# Semana 4 — Redis Cache

## Lo que aprendiste

### Patrón Cache-Aside
El patrón más común para usar cache en APIs:

```
GET /users
  → ¿está en Redis?
      SÍ  → devolver desde Redis (microsegundos)
      NO  → consultar PostgreSQL → guardar en Redis → devolver
```

### Cache Invalidation
Cuando los datos cambian (POST, DELETE, PUT), el cache debe invalidarse para que la próxima lectura traiga datos frescos de PostgreSQL:

```go
// Al insertar un usuario nuevo, borrar el cache
rdb.Del(context.Background(), "users")
```

**Por qué borrar y no actualizar:** con múltiples servicios escribiendo al mismo tiempo, actualizar el cache genera condiciones de carrera. Borrar es atómico y seguro.

### go-redis con ParseURL
```go
// Forma profesional — soporta redis://:password@host:port/db
opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
rdb := redis.NewClient(opt)

// Verificar conexión al arrancar
rdb.Ping(context.Background()).Err()
```

### TTL — Time To Live
```go
rdb.Set(ctx, "users", data, 30*time.Second)
```
El cache expira automáticamente después de 30 segundos. Evita datos eternamente desactualizados.

### Tiempos reales observados
```
6.43ms   ← GET sin cache → PostgreSQL
553µs    ← GET desde Redis  (12x más rápido)
5.21ms   ← POST → inserta + invalida cache
774µs    ← GET reconstruye cache desde PostgreSQL
269µs    ← GET desde Redis  (24x más rápido)
```

### Docker Compose con Redis
```yaml
redis:
  image: redis:7-alpine
  ports:
    - "6379:6379"
  networks:
    - noah_network

api:
  environment:
    REDIS_URL: redis://redis:6379   # hostname = nombre del servicio
  depends_on:
    redis:
      condition: service_started    # Redis no tiene healthcheck oficial
```

## Proyecto completado

API Go + PostgreSQL + Redis con cache invalidation:

```bash
docker compose up -d --build

# Construye cache
curl http://localhost:8082/users

# Inserta e invalida cache automáticamente
curl -X POST http://localhost:8082/users \
  -H "Content-Type: application/json" \
  -d '{"name":"fabian","email":"fabian@test.com"}'

# Primera llamada reconstruye cache desde PostgreSQL
curl http://localhost:8082/users
```

## Próximo — Semana 5

- Separar la API en microservicios independientes
- User Service en Go
- Auth Service en Rust
- Comunicación HTTP entre servicios
- Arquitectura distribuida
