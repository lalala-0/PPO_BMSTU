FROM postgres:13

# Установка pglogical
RUN apt-get update && apt-get install -y \
    postgresql-server-dev-13 \
    build-essential \
    && apt-get clean

# Устанавливаем расширение pglogical
RUN git clone https://github.com/2ndQuadrant/pglogical.git /tmp/pglogical && \
    cd /tmp/pglogical && \
    make && make install && \
    rm -rf /tmp/pglogical

# Копируем скрипт для инициализации репликации
COPY init-replication.sh /docker-entrypoint-initdb.d/

# Экспонируем стандартный порт PostgreSQL
EXPOSE 5432
