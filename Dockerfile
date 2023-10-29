ARG APP_NAME=peculi

# Build stage
FROM alpine:latest as build
RUN apk add --no-cache go

ARG APP_NAME
ENV APP_NAME=$APP_NAME

RUN mkdir /usr/src
RUN mkdir /usr/src/app
RUN mkdir /usr/exec

WORKDIR /usr/src/app

COPY . .

WORKDIR /usr/src/app/cmd/$APP_NAME

RUN go build -o /usr/exec/$APP_NAME

# Production stage
FROM alpine:latest as production
ARG APP_NAME
ENV APP_NAME=$APP_NAME
WORKDIR /root/
COPY --from=build /usr/exec/$APP_NAME ./
CMD ./$APP_NAME
