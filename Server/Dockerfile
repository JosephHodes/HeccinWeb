# NOT SUBFOLDER
# ======================
#  GO FIRST STAGE
# ======================

FROM golang:latest as builder
USER ${USER}
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
ENV GO111MODULE="on" \
  GOARCH="amd64" \
  GOOS="linux" \
  CGO_ENABLED="0"
RUN apt-get autoclean \
  && apt-get autoremove \
  && apt-get clean \
  && apt-get remove

# ======================
#  GO FINAL STAGE
# ======================

FROM builder
COPY --from=builder . ./
RUN apt-get update \
  && apt-get install -y \
  build-essential \
  wget \
  unzip \
  ca-certificates \
  xdg-utils \
  python3 \
  g++ \
  make \
  vim \
  && go build -o main .
EXPOSE 4000
ENTRYPOINT ["./main"]