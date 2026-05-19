# Docker DevOps Roadmap

12-week hands-on program to become a backend/DevOps engineer using Go, Docker, and modern infrastructure.

## Stack

Go · Rust · Docker · Docker Compose · PostgreSQL · Redis · Traefik · Prometheus · Grafana · GitHub Actions

## Progress

| Week | Topic | Status |
|------|-------|--------|
| 1 | Docker basics — Go API, images, containers | ✅ |
| 2 | Professional Dockerfiles — multi-stage builds, scratch image (987MB → 8MB) | ✅ |
| 3 | Docker Compose — Go API + PostgreSQL, networks, healthchecks, volumes, CRUD | ✅ |
| 4 | Databases & Redis — Redis cache, cache invalidation, cache-aside pattern | ✅ |
| 5 | Microservices — user-service (Go) + auth-service (Rust/Axum), inter-service network | ✅ |
| 6 | Reverse proxy — Traefik + socket-proxy, labels, path routing | ✅ ⚠️ Docker Desktop socket limitation on Windows, works on Linux VPS |
| 7 | Observability — Prometheus, metrics, exporters | ✅ |
| 8 | Grafana & Logs — dashboards, centralized logging | ✅ |
| 9 | CI/CD — GitHub Actions, automated testing | ✅ |
| 10 | Automated deployment — SSH deploy, Docker Registry | ⏳ |
| 11 | Security — secrets, non-root users, hardening | ⏳ |
| 12 | Final project — complete SaaS microservices platform | ⏳ |

## Structure

```
go/
  semana1/        # Docker basics — Go API, simple Dockerfile
  semana2/        # Multi-stage builds — scratch image, 987MB → 8MB
  semana3/        # Docker Compose + PostgreSQL — networks, healthchecks, CRUD
  semana4/        # Redis cache — cache-aside pattern, TTL, invalidation
services/
  semana5/        # Microservices — user-service (Go) + auth-service (Rust/Axum)
  semana6/        # Reverse proxy — Traefik + socket-proxy, labels, path routing
docs/             # Course guide and references
```

## Final Architecture

Traefik → User Service (Go) + Auth Service (Rust) + AI Service  
PostgreSQL · Redis · Prometheus · Grafana · GitHub Actions CI/CD
