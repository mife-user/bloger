FROM alpine:latest

WORKDIR /app

COPY bloger .

COPY config/dev.yml ./configs/dev.yml

RUN mkdir -p /app /app/logs

EXPOSE 3000

# 运行应用
CMD ["./bloger"]
