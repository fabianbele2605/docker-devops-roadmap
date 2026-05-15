# Docker DevOps Roadmap

12-week hands-on program to become a backend/DevOps engineer using Go, Docker, and modern infrastructure.

## Stack

Go · Rust · Docker · Docker Compose · PostgreSQL · Redis · Traefik · Prometheus · Grafana · GitHub Actions

## Progress

| Week | Topic | Status |
|------|-------|--------|
| 1 | Docker basics — Go API, images, containers | ✅ |
| 2 | Professional Dockerfiles — multi-stage builds, scratch image (987MB → 8MB) | ✅ |
| 3 | Docker Compose — Go API + PostgreSQL, networks, healthchecks | 🔄 |
| 4 | Databases & Redis — volumes, backups, caching, migrations | ⏳ |
| 5 | Microservices — distributed architecture, internal APIs | ⏳ |
| 6 | Reverse proxy — Traefik, SSL, load balancing | ⏳ |
| 7 | Observability — Prometheus, metrics, exporters | ⏳ |
| 8 | Grafana & Logs — dashboards, centralized logging | ⏳ |
| 9 | CI/CD — GitHub Actions, automated testing | ⏳ |
| 10 | Automated deployment — SSH deploy, Docker Registry | ⏳ |
| 11 | Security — secrets, non-root users, hardening | ⏳ |
| 12 | Final project — complete SaaS microservices platform | ⏳ |

## Structure

```
go/
  semana1/   # Docker basics
  semana2/   # Multi-stage builds
  semana3/   # Docker Compose + PostgreSQL
docs/        # Course guide and references
```

## Final Architecture

Traefik → User Service (Go) + Auth Service (Rust) + AI Service  
PostgreSQL · Redis · Prometheus · Grafana · GitHub Actions CI/CD
