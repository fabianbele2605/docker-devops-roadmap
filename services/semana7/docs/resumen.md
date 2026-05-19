# Semana 7 — Prometheus

## Lo que aprendiste

### Qué es Prometheus
Sistema de monitoreo que hace "scrape" (recolecta) métricas de tus servicios en intervalos regulares. Guarda los datos en una base de datos de series de tiempo y permite hacer queries con PromQL.

### Patrón de métricas en Go
```go
// Una línea agrega /metrics automáticamente con todas las métricas de Go
r.GET("/metrics", gin.WrapH(promhttp.Handler()))
```

Métricas que obtienes gratis:
- `go_goroutines` — goroutines activas
- `go_memstats_alloc_bytes` — memoria en uso
- `process_cpu_seconds_total` — CPU consumida
- `process_resident_memory_bytes` — RAM del proceso
- `promhttp_metric_handler_requests_total` — cuántas veces Prometheus hizo scrape

### prometheus.yml — configuración de scrape
```yaml
global:
  scrape_interval: 15s   # cada cuánto recolecta métricas

scrape_configs:
  - job_name: 'user-service'
    static_configs:
      - targets: ['go_api:8080']  # hostname = nombre del servicio en Compose
```

### Docker Compose con Prometheus
```yaml
prometheus:
  image: prom/prometheus:v2.53.0
  volumes:
    - ./infrastructure/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml
  ports:
    - "9090:9090"
  networks:
    - noah_network
```

### PromQL — queries básicas
```
go_goroutines                          # goroutines activas
process_resident_memory_bytes          # RAM en bytes
rate(process_cpu_seconds_total[5m])   # tasa de CPU en últimos 5 min
```

### Verificación de targets
En `http://localhost:9090` → Status → Targets:
- Estado `UP` = Prometheus llega al servicio y recolecta métricas
- Estado `DOWN` = el servicio no es alcanzable o el endpoint no existe

## Proyecto completado

Prometheus recolectando métricas reales del user-service:

```bash
docker compose up -d --build

# Ver métricas crudas
curl http://localhost:8082/metrics

# UI de Prometheus
open http://localhost:9090

# Query ejemplo
go_goroutines → 6 goroutines activas
```

## Próximo — Semana 8

- Grafana conectado a Prometheus como datasource
- Dashboards visuales con gráficas en tiempo real
- Alertas cuando una métrica supera un umbral
