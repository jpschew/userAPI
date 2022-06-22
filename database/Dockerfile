# for mac
FROM arm64v8/mysql
# for windows
# FROM mysql

COPY create-local-db.sql /tmp

CMD [ "mysqld", "--init-file=/tmp/create-local-db.sql" ]