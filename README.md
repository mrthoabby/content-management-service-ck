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

# Project Task Checklist

## Tareas a Realizar

### Funcionalidades de **GET**

- [ ] **Obtener todas las secciones y sus páginas**
  - Endpoint: `GET /sections`
  - **Tareas**:
    - [ ] Implementar la lógica para listar todas las secciones en el servicio de dominio.
    - [ ] Implementar la lógica para listar las páginas de cada sección.
    - [ ] Crear el controlador para manejar la solicitud y formatear la respuesta.

- [ ] **Obtener el contenido de una página específica**
  - Endpoint: `GET /sections/{sectionId}/pages/{pageId}`
  - **Tareas**:
    - [ ] Implementar la lógica para buscar una sección por su `sectionId`.
    - [ ] Implementar la lógica para buscar la página por su `pageId`.
    - [ ] Crear el controlador que devuelva el contenido de la página.

- [ ] **Buscar secciones y páginas**
  - Endpoint: `GET /sections/search?query={query}`
  - **Tareas**:
    - [ ] Implementar la lógica de búsqueda en el servicio de dominio.
    - [ ] Crear un método en el repositorio para realizar la búsqueda en la base de datos.
    - [ ] Crear el controlador que maneje la solicitud de búsqueda y devuelva los resultados.

### Funcionalidades de **POST**

- [ ] **Crear una nueva sección**
  - Endpoint: `POST /sections`
  - **Tareas**:
    - [ ] Implementar la lógica para agregar una nueva sección en el servicio de dominio.
    - [ ] Crear el método en el repositorio para persistir la nueva sección.
    - [ ] Crear el controlador que maneje la solicitud y devuelva la sección creada.

- [ ] **Crear una nueva página dentro de una sección**
  - Endpoint: `POST /sections/{sectionId}/pages`
  - **Tareas**:
    - [ ] Implementar la lógica para agregar una página en una sección existente.
    - [ ] Crear el método en el repositorio para persistir la nueva página.
    - [ ] Crear el controlador que maneje la solicitud y devuelva la página creada.

- [ ] **Guardar contenido en una página existente**
  - Endpoint: `POST /sections/{sectionId}/pages/{pageId}/content`
  - **Tareas**:
    - [ ] Implementar la lógica para agregar contenido a una página existente en el servicio de dominio.
    - [ ] Crear el método en el repositorio para actualizar el contenido de la página.
    - [ ] Crear el controlador que maneje la solicitud y devuelva la confirmación de la actualización.

### Funcionalidades de **PUT**

- [ ] **Modificar el nombre de una sección**
  - Endpoint: `PUT /sections/{sectionId}`
  - **Tareas**:
    - [ ] Implementar la lógica para actualizar el nombre de una sección en el servicio de dominio.
    - [ ] Crear el método en el repositorio para actualizar la sección en la base de datos.
    - [ ] Crear el controlador que maneje la solicitud y devuelva la sección actualizada.

- [ ] **Modificar el nombre de una página**
  - Endpoint: `PUT /sections/{sectionId}/pages/{pageId}`
  - **Tareas**:
    - [ ] Implementar la lógica para actualizar el nombre de una página en el servicio de dominio.
    - [ ] Crear el método en el repositorio para actualizar la página en la base de datos.
    - [ ] Crear el controlador que maneje la solicitud y devuelva la página actualizada.

### Funcionalidades de **DELETE**

- [ ] **Eliminar una sección y todas sus páginas**
  - Endpoint: `DELETE /sections/{sectionId}`
  - **Tareas**:
    - [ ] Implementar la lógica para eliminar una sección y sus páginas en el servicio de dominio.
    - [ ] Crear el método en el repositorio para eliminar la sección y sus páginas en la base de datos.
    - [ ] Crear el controlador que maneje la solicitud y devuelva la confirmación de la eliminación.

- [ ] **Eliminar una página específica**
  - Endpoint: `DELETE /sections/{sectionId}/pages/{pageId}`
  - **Tareas**:
    - [ ] Implementar la lógica para eliminar una página específica en el servicio de dominio.
    - [ ] Crear el método en el repositorio para eliminar la página en la base de datos.
    - [ ] Crear el controlador que maneje la solicitud y devuelva la confirmación de la eliminación.

- [ ] Implementar Alertas Básicas en Prometheus
- [ ] Implementar Documentación de la API
- [ ] Pruebas de Integración
- [ ] Refactorización y Optimización
- [ ] Dejar documentado que el manejo de errores va orientado a panics, con un error handler comun que trabaja de acuerdo al error generado el es quien determina cuando lanzar o no un panic, que será controlado por un middleware central
- [ ] No olvidar hacer las pruebas unitarias y de integración
- [ ] Tratar de lanzar un panic en cada capa para asegurarme de que el middeware está trabanado como se debe
## Colaboración

Si deseas contribuir a este proyecto, sigue estos pasos:

1. Haz un fork del repositorio.
2. Crea una nueva rama para tu característica (`git checkout -b feature/nueva-caracteristica`).
3. Realiza tus cambios y confirma (`git commit -m 'Agrega nueva característica'`).
4. Envía tus cambios a tu repositorio (`git push origin feature/nueva-caracteristica`).
5. Abre un Pull Request para revisar tus cambios.

¡Gracias por tu interés en colaborar en Cirosky!

