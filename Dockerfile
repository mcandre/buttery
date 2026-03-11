FROM dhi.io/golang:1.26-alpine3.23 AS build
COPY . /src
WORKDIR /src
RUN go install ./...

FROM scratch
COPY --from=build /go/bin/buttery /buttery
ENTRYPOINT ["/buttery"]
