FROM golang:1.17

WORKDIR /backend

COPY . .

RUN go mod download

ARG user
ENV user $user

ARG db
ENV db $db

ARG pass
ENV pass $pass

ARG port
ENV port $port

ARG domain
ENV domain $domain

RUN go build -o backend

EXPOSE 3000

CMD [ "./backend" ]