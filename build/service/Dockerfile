FROM cloud AS base

FROM scratch

ARG service

COPY --from=base /cloud/${service} /cloud/service

ENTRYPOINT [ "/cloud/service" ]