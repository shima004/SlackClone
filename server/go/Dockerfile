FROM --platform=linux/amd64 golang:1.19 

RUN apt update

WORKDIR /backend

# RUN go mod tidy
RUN go install github.com/cosmtrek/air@latest

CMD ["air" "-c" ".air.toml"]
