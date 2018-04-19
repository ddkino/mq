FROM scratch
# FROM golang:1.10.0-alpine
COPY . /
EXPOSE 4222 6222 8222
# ENV TOPIC=welcome 
# ENV ENV=production
# keep cmd
CMD ["/messaging/events/signup/main", "-topic=welcome", "-env=production"]
# ENTRYPOINT [ "/main" ]

