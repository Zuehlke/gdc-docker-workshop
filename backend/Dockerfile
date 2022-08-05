FROM golang:1.18-alpine
WORKDIR /app

COPY . .

RUN go mod download

RUN go install

ENV AUTH_TOKEN="123"
ENV RELEASES_URL="http://localhost:8080"

EXPOSE 8081
ENTRYPOINT ["workshop-backend"]