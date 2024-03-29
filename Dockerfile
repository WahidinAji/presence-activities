FROM golang:1.18.4-buster

# ARG for get CERT_RUL
ARG CERT_URL

WORKDIR /usr/app

# RUN curl --create-dirs -o ./.postgresql/root.crt -O ${CERT_URL}

COPY . .

RUN go mod download

RUN go build main.go

CMD [ "./main" ]