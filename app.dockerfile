FROM golang:1.19.0-buster as base

ENV USER=app
ENV UID=1001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --shell "/sbin/nologin" \
    --uid "${UID}" \
    "${USER}"

WORKDIR /app
COPY . /app

CMD make build

FROM alpine:3.15

WORKDIR /banner-service
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group
COPY --from=base /app/app /banner-service/app
COPY --from=base /app/config/config_prod.yaml /banner-service/config.yaml

RUN chown -R app:app /banner-service
USER app:app

CMD ["./app", "-config", "/banner-service/config.yaml"]
