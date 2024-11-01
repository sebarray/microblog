# Documentación de Arquitectura: Microblogging App

## 1. Descripción General de la Arquitectura

La aplicación de microblogging es una solución distribuida y escalable, diseñada para gestionar publicaciones, interacciones entre usuarios, y un timeline personalizado. La arquitectura emplea el patrón CQRS (Command Query Responsibility Segregation) para optimizar las operaciones de escritura y lectura de datos, permitiendo que el sistema responda eficientemente a las solicitudes de diferentes tipos de usuarios.

### Diagrama de Arquitectura (High Level)

- **Microservicios**: Componentes independientes y desplegados en contenedores Kubernetes.
- **Base de Datos**: MongoDB como sistema NoSQL para escalabilidad y flexibilidad en la gestión de datos.
- **Load Balancer y API Gateway**: Para manejar el tráfico de entrada y facilitar el enrutamiento a los servicios correspondientes.

## 2. Componentes Principales

### a) Microservicios

La aplicación está compuesta por varios microservicios en Go, cada uno de los cuales maneja una funcionalidad específica de la aplicación:

- **Microblog-Command Service**: Responsable de la lógica de escritura, como crear publicaciones y gestionar acciones de los usuarios (seguir, dejar de seguir). Este servicio usa el puerto 8081.
- **Microblog-Query Service**: Optimizado para operaciones de lectura, como obtener el timeline de un usuario y consultar publicaciones.

### b) Base de Datos - MongoDB

MongoDB es la base de datos elegida para la capa de persistencia, adecuada para almacenar y gestionar documentos JSON (u objetos BSON) que representan publicaciones y datos de usuarios. MongoDB facilita las operaciones de consulta rápida, lo cual es beneficioso para la generación de timelines y la recuperación de publicaciones.

### c) Kubernetes (K8s)

La aplicación se despliega en un entorno de Kubernetes usando Minikube como entorno de desarrollo local. Cada servicio está en un pod independiente, permitiendo un despliegue y escalado individual. La configuración de Kubernetes incluye:

- **Manifiestos de K8s**: Definen el despliegue y configuración de los pods y servicios.
- **Servicio de ClusterIP**: Asegura la comunicación entre los servicios dentro del clúster.
- **ConfigMap y Secrets**: Almacenan configuraciones y credenciales de forma segura.

### d) Minikube

Minikube permite simular un entorno de Kubernetes en desarrollo, facilitando la configuración y pruebas antes del despliegue en producción. Se usa para correr los pods de cada microservicio y MongoDB en el entorno local.

### e) EC2 (para despliegue en producción)

Para el despliegue en producción, se planea usar una instancia de Amazon EC2, proporcionando un entorno escalable y económico, con capacidad para correr el clúster de Kubernetes y MongoDB en un entorno de producción.

## 3. Patrón CQRS

La aplicación implementa CQRS para dividir las responsabilidades de comandos (escritura) y consultas (lectura). Esto significa que:

- **Command Service** se encarga de todas las operaciones que modifican datos, como crear y eliminar publicaciones.
- **Query Service** gestiona las operaciones de consulta y acceso a datos optimizados para la lectura, como mostrar el timeline de los usuarios.

Esta separación permite escalar cada servicio independientemente según las necesidades de lectura o escritura.

## 4. Despliegue y Configuración

### a) Despliegue en Kubernetes

- **Manifiestos de Despliegue**: Cada servicio tiene un manifiesto de Kubernetes que define sus configuraciones específicas (puerto, volúmenes, variables de entorno, etc.).
- **Autoscaling**: Se puede configurar en Kubernetes para ajustar el número de pods según la carga de tráfico.

### b) Escalabilidad y Alta Disponibilidad

Gracias a Kubernetes, la aplicación soporta la replicación y balanceo de carga de los microservicios. MongoDB, al ser distribuido, también facilita la escalabilidad de la capa de datos.

## 5. Flujo de Datos y Operaciones Clave

- **Creación de Publicaciones**: Los usuarios pueden crear publicaciones a través del Command Service, que luego se almacenan en MongoDB.
- **Timeline del Usuario**: Cuando un usuario solicita su timeline, el Query Service accede a MongoDB para compilar una lista de publicaciones relevantes.
- **Seguir/Dejar de Seguir Usuarios**: Se gestionan a través del Command Service, actualizando la lista de seguidores en la base de datos para determinar qué publicaciones mostrar en el timeline del usuario.




## 6 documentacion funcion 
 https://documenter.getpostman.com/view/16703239/2sAY4xA21S