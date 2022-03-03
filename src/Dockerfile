FROM scratch
ADD connector ./connector
ADD ca-certificates.crt /etc/ssl/certs/
COPY sample.pdf ./sample.pdf
EXPOSE 9000
ENTRYPOINT ["/caDbConnector"]