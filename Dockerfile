FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY controllers/*.go ./controllers/
COPY models/*.go ./models/
COPY routes/*.go ./routes/

RUN go build -o /cats

EXPOSE 8080

CMD [ "/cats" ]
