FROM alpine:latest

RUN mkdir /app

COPY boardApp /app

CMD [ "/app/boardApp"]