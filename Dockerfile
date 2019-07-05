FROM golang:1.12.6-stretch AS appbuild
ADD main.go /src/
RUN cd /src && go build -o app

FROM golang:1.12.6-stretch AS kustomizebuild
ADD https://github.com/kubernetes-sigs/kustomize/releases/download/v3.0.0/kustomize_3.0.0_linux_amd64 /app/kustomize
RUN chmod 755 /app/kustomize

FROM scratch
WORKDIR /app
COPY --from=appbuild /src/app /app/
COPY --from=kustomizebuild /app/kustomize /app/
ENTRYPOINT ["/app/app"]