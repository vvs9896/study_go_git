# 📅 Неделя 13: gRPC и микросервисная архитектура

## 🎯 Цели недели
- Освоить gRPC для межсервисной коммуникации
- Понять принципы микросервисной архитектуры
- Создать distributed систему
- Изучить service discovery и load balancing

## 📚 Теория (Go 80% | Git 20%)

### Go (6 часов)
- **День 1-2**: Protocol Buffers, gRPC основы
- **День 3-4**: gRPC сервисы, streaming, interceptors
- **День 5-6**: Микросервисы, service mesh, distributed tracing
- **Выходные**: Построение микросервисной архитектуры

### Git (1.5 часа)
- **День 2**: Monorepo vs multirepo для микросервисов
- **День 4**: Service versioning strategies
- **День 6**: Deployment orchestration

## 🛠 Практические задания

### Go
1. **gRPC основы**:
   - Protocol Buffers schema
   - Unary и streaming RPC
   - Error handling в gRPC

2. **Микросервисы**:
   - User service, Product service, Order service
   - Inter-service communication
   - Service discovery с Consul

3. **Production-ready features**:
   - Health checks
   - Metrics и monitoring
   - Circuit breaker pattern

### Git
1. **Microservices Git workflow**:
   - Service-specific repositories
   - Shared libraries management
   - Coordinated releases

## 🏠 Домашнее задание

**Go**: E-commerce микросервисная платформа:
- User service (gRPC)
- Product catalog service
- Order processing service
- API Gateway
- Service discovery

## ✅ Критерии успеха
- [ ] Работающие gRPC сервисы
- [ ] Межсервисная коммуникация
- [ ] Service discovery
- [ ] Monitoring и health checks

**Время на неделю**: ~20-24 часа 