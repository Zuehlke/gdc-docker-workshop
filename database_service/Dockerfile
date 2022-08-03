FROM maven:3-amazoncorretto-17 as builder
WORKDIR /build

COPY . .

RUN mvn -Dmaven.test.skip clean install

FROM amazoncorretto:17-alpine
WORKDIR /app

ENV DATABASE_HOST="jdbc:mysql://releases-database:3306/releases"
ENV DATABASE_USERNAME="root"
ENV DATABASE_PASSWORD="my-secret-pw"

COPY --from=builder /build/target/database-*.jar /app/app.jar

ENTRYPOINT ["java", "-jar", "app.jar"]