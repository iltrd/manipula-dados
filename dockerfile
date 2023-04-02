# Imagem base
FROM golang:1.17.2-alpine3.14

# Define a pasta de trabalho
WORKDIR /app

# Copia os arquivos necessários
COPY . .

# Instala as dependências do projeto
RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go install -v ./...

# Compila o serviço
RUN go build -o main .

# Define a imagem final
FROM alpine:3.14

# Instala as dependências do banco de dados
RUN apk add --no-cache postgresql-client

# Copia o serviço compilado
COPY --from=0 /app/main .

# Define as variáveis de ambiente
ENV PGHOST=db
ENV PGUSER=postgres
ENV PGPASSWORD=postgres
ENV PGDATABASE=postgres

# Expõe a porta utilizada pelo serviço
EXPOSE 8080

# Inicia o serviço
CMD ["./main"]
