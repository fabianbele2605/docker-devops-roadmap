# Semana 3 — Docker Compose + PostgreSQL

## Lo que aprendiste

### Docker Compose fundamentals
- **Por qué Compose:** levantar múltiples contenedores con un solo comando en vez de correr cada `docker run` manualmente
- **Servicios:** cada contenedor es un servicio con nombre propio
- **Redes personalizadas:** los contenedores se comunican por nombre de servicio (`db`, `api`) — no por IP
- **`depends_on` con healthcheck:** garantiza que PostgreSQL esté listo antes de que la API intente conectarse
- **Variables de entorno:** configuración externalizada del contenedor (`DATABASE_URL`, credenciales)

### Volúmenes persistentes
- Sin volumen los datos desaparecen al hacer `docker compose down`
- El volumen se declara en **dos lugares**: dentro del servicio (montaje) y a nivel raíz (definición)
- PostgreSQL guarda sus datos en `/var/lib/postgresql/data`

### Comandos dominados
```bash
docker compose up -d          # levanta en background
docker compose up -d --build  # rebuild + levanta
docker compose down           # para y elimina contenedores (los volúmenes sobreviven)
docker compose ps             # estado de los servicios
docker compose logs -f        # logs en tiempo real
```

### docker-compose.yml — patrones clave
```yaml
services:
  db:
    image: postgres:16-alpine
    volumes:
      - postgres_data:/var/lib/postgresql/data   # persistencia
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U user -d dbname"]
      interval: 5s
      retries: 5
    networks:
      - my_network

  api:
    build:
      context: ./api-go
    depends_on:
      db:
        condition: service_healthy   # espera a que DB esté healthy
    networks:
      - my_network

volumes:
  postgres_data:     # declaración a nivel raíz

networks:
  my_network:
```

### Go — CRUD con pgx v5
```go
// Crear tabla al arrancar
conn.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS users (...)`)

// Insertar y obtener id generado
conn.QueryRow(ctx, "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id",
    user.Name, user.Email).Scan(&user.ID)

// Leer múltiples filas
rows, _ := conn.Query(ctx, "SELECT id, name, email FROM users")
defer rows.Close()
for rows.Next() {
    rows.Scan(&u.ID, &u.Name, &u.Email)
}
```

### Lección importante — tipos en pgx v5
- pgx v5 es estricto con tipos: PostgreSQL `SERIAL` es `int4`, no compatible con `uint` de Go
- Usar `int` en el struct, no `uint`

### Lección importante — SERIAL y gaps
- Los ids del SERIAL no se revierten aunque el insert falle
- Es comportamiento normal de las secuencias en PostgreSQL — no es un bug

## Proyecto completado

API Go + PostgreSQL con persistencia real:

```bash
docker compose up -d
curl -X POST http://localhost:8082/users \
  -H "Content-Type: application/json" \
  -d '{"name":"fabian","email":"fabian@test.com"}'
# {"id":1,"name":"fabian","email":"fabian@test.com"}

docker compose down && docker compose up -d
curl http://localhost:8082/users
# datos siguen ahí — volumen funcionando
```

## Pendiente para Semana 4

- Agregar **Redis** como cache para `GET /users`
- **Migraciones** reales con `golang-migrate` en vez de `CREATE TABLE IF NOT EXISTS`
- **Backups** de PostgreSQL desde Docker
- Endpoint `DELETE /users/:id` y `PUT /users/:id`