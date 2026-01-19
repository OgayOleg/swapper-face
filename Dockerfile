FROM golang:1.25.4
WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o faceswapper ./cmd/app/main.go

CMD ["./faceswapper"]
