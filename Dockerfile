FROM golang:1.23
WORKDIR /app 
ENV TODO_PORT=7540
COPY go.mod go.sum ./
RUN go mod download
COPY cmd/ cmd/
COPY internal/ internal/
COPY tests/ tests/
COPY web/ web/
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /my_app ./cmd/service
EXPOSE 7540
CMD ["/my_app"] 


