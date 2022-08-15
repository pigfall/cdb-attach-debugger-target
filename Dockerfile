FROM golang:1.18 as Builder

WORKDIR /app

COPY . .

RUN go build -o tmpgo .

FROM koitown/ubuntu:1.0

COPY --from=Builder /app/tmpgo .

CMD ["/tmpgo"]
