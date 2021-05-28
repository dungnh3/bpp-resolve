FROM golang:1.15 AS build-env

WORKDIR /build
COPY . .

RUN make build

FROM gcr.io/distroless/base
COPY --from=build-env /build/bin ./
COPY --from=build-env /build/config ./config
COPY --from=build-env /build/migrations ./migrations

CMD ["./runtime", "server"]