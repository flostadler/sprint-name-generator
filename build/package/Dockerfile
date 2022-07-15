FROM nginx:latest
COPY web/build /usr/share/nginx/html/
RUN apk add --update --no-cache 'libcurl>=7.84.0' && \
    apk add --update --no-cache 'curl>=7.84.0' && \
    apk add --update --no-cache 'libdb5.3>=5.3.28+dfsg1-0.9'
EXPOSE 80