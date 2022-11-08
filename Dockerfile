FROM golang as builder
WORKDIR /home/nullwulf/F22/CSC482/nwolf2-server

COPY . .

RUN go get .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o nwolf2-server .

# deployment image
FROM scratch

# copy ca-certificates from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

WORKDIR /bin/

COPY --from=builder /home/nullwulf/F22/CSC482/nwolf2-server/nwolf2-server .
# COPY --from=builder /home/nullwulf/F22/CSC482/nwolf2-server/.env .
# Sets gin to release mode
# Other environment variables loaded in from program
ARG gin_mode=release
ENV GIN_MODE=$gin_mode

CMD [ "./nwolf2-server" ]

EXPOSE 8080
