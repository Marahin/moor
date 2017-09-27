FROM scratch
ADD ca-certificates.crt /etc/ssl/certs/
ADD moor.bin /
CMD ["/moor.bin"]
