FROM golang:alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o k8s-ingress-linklist src/cmd/kubernetes-ingress-linklist/main.go


FROM alpine:edge
WORKDIR /app
COPY --from=build /app/k8s-ingress-linklist .
COPY src/internal/frontend ./internal/frontend
COPY src/web ./web

RUN apk --no-cache add ca-certificates tzdata

EXPOSE 8080

CMD ["./k8s-ingress-linklist"]
