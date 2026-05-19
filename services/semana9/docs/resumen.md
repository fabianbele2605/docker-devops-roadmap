# Semana 9 — GitHub Actions CI/CD

## Lo que aprendiste

### Qué es GitHub Actions
Sistema de CI/CD integrado en GitHub. Ejecuta workflows automáticamente en respuesta a eventos (push, PR, etc.) en máquinas virtuales de GitHub.

### Estructura del workflow
```yaml
name: CI

on:
  push:
    branches: [ main ]   # se dispara en cada push a main

jobs:
  build:
    runs-on: ubuntu-latest   # máquina virtual de GitHub
    steps:
      - uses: actions/checkout@v4          # descarga el código
      - uses: docker/login-action@v3       # login a Docker Hub
      - uses: docker/build-push-action@v5  # build + push de imagen
```

### Secrets en GitHub
Nunca pongas credenciales en el código. Se guardan en:
**Settings → Secrets and variables → Actions**

```yaml
username: ${{ secrets.DOCKERHUB_USERNAME }}
password: ${{ secrets.DOCKERHUB_TOKEN }}
```

### Token de Docker Hub
- Usar **Read & Write** — no más permisos de los necesarios
- Nunca usar tu password directo — siempre tokens revocables
- Se puede revocar sin cambiar el password

### Flujo completo CI/CD
```
git push → GitHub Actions trigger
         → checkout código
         → login Docker Hub
         → docker build
         → docker push fabianrobles26/user-service:latest
         → ✅ imagen disponible en Docker Hub
```

## Resultado

```
fabianrobles26/user-service:latest
Pushed: automáticamente en cada push a main
```

Cualquier servidor con Docker puede ahora correr la imagen con:
```bash
docker pull fabianrobles26/user-service:latest
docker run -p 8080:8080 fabianrobles26/user-service:latest
```

## Próximo — Semana 10

- Deploy automático a VPS via SSH
- El pipeline no solo construye — también despliega
- `docker compose pull && docker compose up -d` en el servidor remoto