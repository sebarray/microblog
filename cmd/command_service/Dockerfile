# Usa una imagen base de Go
FROM golang:1.20 as builder

# Establece el directorio de trabajo
WORKDIR /app

# Copia los archivos del módulo y descarga las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el código fuente
COPY . .

# Compila la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -o command_service ./cmd/command_service

# Usa una imagen base de Alpine para el contenedor final
FROM alpine:latest

# Copia el binario compilado desde la etapa de construcción
COPY --from=builder /app/command_service .

# Define la variable de entorno para la conexión a MongoDB
ENV MONGO_URI="mongodb://mongo:27017/microblog"

# Expone el puerto en el que la aplicación escuchará
EXPOSE 8081

# Comando para ejecutar la aplicación
CMD ["./command_service"]
