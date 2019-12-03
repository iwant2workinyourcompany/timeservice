# Собирает образ для контейнера с сервисом получения времени.
FROM golang:1.9

#Определяем рабочую директорию
WORKDIR /opt/go/time-service

#Определяем рабочий порт
EXPOSE 8080

#Копируем сорцы
COPY time-service.go .

#Билдим
RUN go build

#Определяем как запускать контейнер
ENTRYPOINT [ "/opt/go/time-service/time-service" ]

