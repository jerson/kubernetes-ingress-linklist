FROM golang:alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o k8s-ingress-linklist src/cmd/kubernetes-ingress-linklist/main.go


FROM alpine:edge

RUN apk --no-cache add ca-certificates tzdata

COPY --from=build /app/k8s-ingress-linklist .
COPY src/internal/frontend /app/internal/frontend
COPY src/web /app/web

EXPOSE 8080

CMD ["/app/k8s-ingress-linklist"]
