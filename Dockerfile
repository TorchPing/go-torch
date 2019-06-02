FROM alpine
COPY builds/torch-linux-amd64 /bin/torch
ENTRYPOINT ["/bin/torch"]