FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

# RUN go get -u github.com/gorilla/mux
# RUN go get -u github.com/go-sql-driver/mysql
# RUN go get -u github.com/swaggo/http-swagger
# RUN go get -u google.golang.org/api/idtoken

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install golang.org/x/tools/cmd/godoc@latest

RUN apt-get update && apt-get install -y \
  jq \
  mariadb-client \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

ENTRYPOINT ["tail", "-F", "/dev/null"]
