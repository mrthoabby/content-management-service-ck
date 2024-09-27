# content-management-service-ck
Microservice for managing content creation, editing, and deletion within the Cirosky ecosystem. Supports organizing sections and pages with a focus on modular, scalable, and maintainable architecture following Google's Go coding standards. Implements robust unit and integration tests to ensure code quality and high coverage

![Go](https://img.shields.io/badge/Go-1.23-blue?logo=go)


# Cirosky - Microservicio de Gestión de Contenido

[Repositorio de Cirosky](https://github.com/mrthoabby/cirosky)

Cirosky es un servicio diseñado para almacenar y organizar escritos, ideas, proyectos y documentación, facilitando la gestión de información tanto personal como colaborativa. Este microservicio se integra en el ecosistema de Cirosky, proporcionando una solución robusta para la creación, edición y eliminación de secciones y páginas de contenido.

## Guía de Estilo de Código

Este proyecto sigue las recomendaciones de estilo de código sugeridas por Google para Go. Puedes consultar la [Guía de Estilo de Go de Google](https://google.github.io/styleguide/go/decisions#naming) para obtener más información.

## Casos de Prueba

Se implementarán pruebas unitarias y de integración para garantizar la funcionalidad del microservicio. El objetivo es mantener una cobertura de código igual o superior al 80%. Se utilizarán los siguientes tipos de pruebas:

- **Pruebas Unitarias:** Para funciones y métodos individuales.
- **Pruebas de Integración:** Para validar la interacción entre diferentes componentes del sistema.

## Lista de Tareas del Proyecto

- [x] Crear Estructura del Proyecto
- [x] Definir Modelo de Dominio para `Section` y `Page`
- [x] Crear Interfaces (Ports)
- [ ] Implementar Adapters
- [ ] Implementar Servicios de `Section` y `Page`
- [ ] Configurar Routers y Handlers
- [ ] Configurar Logging
- [ ] Implementar `CreateSection` y `CreatePage`
- [ ] Implementar `GetAllSections` con Páginas Anidadas
- [ ] Implementar Búsqueda de Secciones y Páginas
- [ ] Implementar Actualización de Nombres de Sección y Página
- [ ] Implementar Eliminación de Secciones y Páginas
- [ ] Implementar Guardado de Contenido de Página
- [ ] Configurar Grafana con Dashboard Básico
- [ ] Implementar Alertas Básicas en Prometheus
- [ ] Implementar Documentación de la API
- [ ] Pruebas de Integración
- [ ] Refactorización y Optimización

## Colaboración

Si deseas contribuir a este proyecto, sigue estos pasos:

1. Haz un fork del repositorio.
2. Crea una nueva rama para tu característica (`git checkout -b feature/nueva-caracteristica`).
3. Realiza tus cambios y confirma (`git commit -m 'Agrega nueva característica'`).
4. Envía tus cambios a tu repositorio (`git push origin feature/nueva-caracteristica`).
5. Abre un Pull Request para revisar tus cambios.

¡Gracias por tu interés en colaborar en Cirosky!

