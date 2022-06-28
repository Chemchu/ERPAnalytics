FROM golang:1.18.3-alpine

WORKDIR /erpanalytics

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . ./

RUN go build -o /erp_analytics

EXPOSE 6060
CMD [ "/erp_analytics" ]