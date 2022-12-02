FROM golang

COPY ["go.mod", "api.go", "/app/"]

WORKDIR /app

RUN go install

COPY . .

CMD ["go", "run", "cmd/api/main.go"]