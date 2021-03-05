FROM golang:1.15.2

COPY . /go/src/backend

WORKDIR /go/src/backend

RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/cosmtrek/air
RUN go get -u github.com/go-delve/delve/cmd/dlv

CMD ["air", "-c", ".air.toml"]**
