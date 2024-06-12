FROM alpine:latest
# copy over the binary from the first stage
COPY chatgpt-adapter /app/chatgpt-adapter
COPY config.yaml /app/config.yaml
WORKDIR "/app"
ENTRYPOINT [ "/app/chatgpt-adapter" ]
