FROM golang:1.16-alpine AS builder


# This will download all certificates (ca-certificates) and builds it in a
# single file under /etc/ssl/certs/ca-certificates.crt (update-ca-certificates)
# I also add git so that we can download with `go mod download` and
# tzdata to configure timezone in final image
RUN apk update \
    && apk upgrade \
    && apk add --no-cache ca-certificates openssl shadow \
    && update-ca-certificates 2>/dev/null || true

RUN adduser -D admin 

RUN apk add sudo openssh 
#&& rc-update  \ 
#add sshd \
#&& rc-status \
#&& /etc/init.d/sshd start

#RUN ssh-keygen -A 

# Ubicarse en el directorio /build.
WORKDIR /build

# Copiar y descargar las dependencias usando el comando mod.
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código del proyecto al contendor
COPY . .

# Establecer las variables de ambiente para el contenedor y construir el código de go 
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o apiserver .

FROM scratch
#FROM alpine

# Copiar del directorio /build el archivo compilado y el archivo .env de configuración
COPY --from=builder ["/build/apiserver", "/build/.env", "/"]

# This line will copy all certificates to final image
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Determinar el puerto de publicación
EXPOSE 8080

# Definir el comando a ejecutar cada vez que se inicie el contenedor
ENTRYPOINT ["/apiserver"]
