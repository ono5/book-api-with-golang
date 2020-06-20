FROM postgres:latest

#  /docker-entrypoint-initdb.d/ 内に sql ファイルを配置すると自動で実行
COPY ./docker/postgres/init/*.sql /docker-entrypoint-initdb.d/
