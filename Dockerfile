# Usa la imagen de Go para compilar el servicio
FROM golang:1.20 as builder

# Configura el directorio de trabajo
WORKDIR /app

# Copia los archivos de tu proyecto al contenedor
COPY . .

# Compila el binario
RUN go mod tidy
RUN go build -o server .

# Usa una imagen más ligera para ejecutar el servicio
FROM gcr.io/distroless/base-debian11

# Copia el binario compilado desde la fase de construcción
COPY --from=builder /app/server /server

# Expone el puerto que usa Cloud Run
EXPOSE 8080

# Ejecuta el binario
CMD ["/server"]
