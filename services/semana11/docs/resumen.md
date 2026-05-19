# Semana 11 — Seguridad: Secrets y Variables de Entorno

## Lo que aprendiste

### El problema
Credenciales hardcodeadas en el código son un riesgo de seguridad:
```yaml
# MAL — expone credenciales en el repositorio
POSTGRES_PASSWORD: noah123
DATABASE_URL: postgresql://noah:noah123@db:5432/noah_db
```

### La solución — archivo `.env`
```env
POSTGRES_USER=noah
POSTGRES_PASSWORD=noah123
POSTGRES_DB=noah_db
DATABASE_URL=postgres://noah:noah123@db:5432/noah_db
REDIS_URL=redis://redis:6379
```

Docker Compose carga `.env` automáticamente si está en la misma carpeta.

### docker-compose.yml con variables
```yaml
services:
  db:
    environment:
      POSTGRES_USER: ${POSTGRES_USER}
      POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
      POSTGRES_DB: ${POSTGRES_DB}
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U ${POSTGRES_USER} -d ${POSTGRES_DB}"]

  go_api:
    environment:
      DATABASE_URL: ${DATABASE_URL}
      REDIS_URL: ${REDIS_URL}
```

### Regla fundamental — `.gitignore`
```
.env
```

El `.env` NUNCA debe subirse a Git. Contiene credenciales reales.

### En producción
En un VPS real, creas el `.env` directamente en el servidor:
```bash
ssh user@vps
nano /app/.env  # creas el archivo manualmente en el servidor
```

Nunca vía Git. Nunca en el código.

## Proyecto completado

```bash
# Las credenciales viven en .env, no en el código
docker compose up -d

curl -X POST http://localhost:8082/users \
  -H "Content-Type: application/json" \
  -d '{"name":"fabian","email":"fabian@test.com"}'
# {"id":1,"name":"fabian","email":"fabian@test.com"}
```

## Próximo — Semana 12

- Proyecto final completo
- Todo el stack junto con deploy real en VPS
- CI/CD + SSL + monitoreo en producción
