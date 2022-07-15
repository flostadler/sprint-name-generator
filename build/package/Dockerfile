FROM nginx:stable
COPY web/build /usr/share/nginx/html/
RUN apt-get --only-upgrade install curl
EXPOSE 80