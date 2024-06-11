FROM alpine:3.15.0
# copy over the binary from the first stage
COPY chatgpt-adapter /chatgpt-adapter/chatgpt-adapter
# Print the contents of the root directory (/) to check initial files
RUN echo "Contents of root directory:" && ls -l /

# Print the contents of /chatgpt-adapter directory after copying the binary
RUN echo "Contents of /chatgpt-adapter directory after copying the binary:" && ls -l /chatgpt-adapter

#COPY config.yaml /chatgpt-adapter/config.yaml
WORKDIR "/chatgpt-adapter"
ENTRYPOINT [ "/chatgpt-adapter/chatgpt-adapter" ]
