FROM postgres:latest

# Устанавливаем переменные среды (наследуемые из compose, но можно продублировать)
ENV POSTGRES_USER=testuser
ENV POSTGRES_DB=testDB
ENV POSTGRES_PASSWORD=testpassword
ENV POSTGRES_HOST_AUTH_METHOD=trust

# Копируем файлы инициализации в контейнер
COPY init.sql /docker-entrypoint-initdb.d/
COPY load_data.sql /docker-entrypoint-initdb.d/
COPY data/ /docker-entrypoint-initdb.d/data/

# Вызываем стандартный entrypoint
ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["postgres"]
