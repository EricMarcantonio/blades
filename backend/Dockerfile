FROM golang:1.17

WORKDIR /backend

COPY . .

RUN go mod download

RUN go build -o backend

EXPOSE 3000

CMD [ "./backend" ]