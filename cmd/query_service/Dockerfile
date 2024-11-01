# Usar una imagen base de Go para compilar
FROM golang:1.20 AS builder

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos de módulo desde la raíz
COPY ../go.mod ./
COPY ../go.sum ./
RUN go mod download

# Copiar el resto del código fuente
COPY . .

# Compilar el servicio
RUN go build -o query-service .

# Usar una imagen base de Alpine para el runtime
FROM alpine:latest

# Copiar el binario compilado
WORKDIR /app
COPY --from=builder /app/query-service .

# Exponer el puerto
EXPOSE 8082

# Comando por defecto
CMD ["./query-service"]
