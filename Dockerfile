FROM alpine

WORKDIR /app/
ADD ./app /app/

ENTRYPOINT ["./app"]

#
#JWT_SECRET=200Lab.io;
#MYSQL_GORM_DB_TYPE=mysql;
#root:my-secret-pw@tcp(127.0.0.1:3306)/g09-mysql?charset=utf8mb4&parseTime=True&loc=Local