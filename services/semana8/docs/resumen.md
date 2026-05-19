# Semana 8 — Grafana

## Lo que aprendiste

### Qué es Grafana
Herramienta de visualización que se conecta a fuentes de datos (Prometheus, PostgreSQL, etc.) y permite crear dashboards con gráficas en tiempo real. No almacena datos — solo los visualiza.

### Conexión Grafana → Prometheus
Dentro de Docker Compose, Grafana encuentra Prometheus por nombre de servicio:
```
http://prometheus:9090
```

### Docker Compose con Grafana
```yaml
grafana:
  image: grafana/grafana:11.0.0
  volumes:
    - grafana_data:/var/lib/grafana   # persiste dashboards y config
  ports:
    - "3000:3000"
  networks:
    - noah_network
```

El volumen es crítico — sin él pierdes todos los dashboards al reiniciar.

### Datasource
- Connections → Data sources → Add → Prometheus
- URL: `http://prometheus:9090`
- Save & test → "Successfully queried the Prometheus API"

### Paneles creados
| Panel | Métrica | Tipo | Unidad |
|-------|---------|------|--------|
| Goroutines | `go_goroutines` | Stat | — |
| Memoria RAM | `process_resident_memory_bytes` | Stat | bytes(SI) |

### Tipos de panel más útiles
- **Stat** — valor único grande, ideal para métricas puntuales
- **Time series** — gráfica histórica, ideal para ver tendencias
- **Gauge** — barra de progreso, ideal para porcentajes

### Credenciales por defecto
- usuario: `admin`
- password: `admin`

## Proyecto completado

Stack completo de observabilidad:

```bash
docker compose up -d --build

# Prometheus recolecta métricas
http://localhost:9090

# Grafana visualiza en tiempo real
http://localhost:3000
# Dashboard "Goroutines" → 6 goroutines, 30.4 MB RAM
```

## Próximo — Semana 9

- GitHub Actions para CI/CD
- Build automático de imágenes Docker
- Tests automáticos en cada push
- Pipeline completo de integración continua