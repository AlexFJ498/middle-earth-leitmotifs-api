# ----------- Etapa build -----------
FROM golang:1.24.4-alpine AS build

# Instalar git y otras herramientas necesarias
RUN apk add --no-cache git

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos al contenedor
COPY . .

# Compilar el binario
RUN CGO_ENABLED=0 go build -o /bin/middle-earth-leitmotifs-api ./cmd/api/main.go

# ----------- Etapa prod -----------
FROM scratch AS prod

# Copiar el binario ya compilado
COPY --from=build /bin/middle-earth-leitmotifs-api /middle-earth-leitmotifs-api

# Comando que se ejecutará al iniciar el contenedor
ENTRYPOINT ["/middle-earth-leitmotifs-api"]