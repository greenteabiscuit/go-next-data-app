FROM golang:1.15.2

COPY . /go/src/backend

WORKDIR /go/src/backend

RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]**
