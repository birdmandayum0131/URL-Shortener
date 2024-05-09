FROM golang:1.22.3-bullseye
LABEL author="Bird"
# Create a non-root user and set the home directory
RUN groupadd -r Servers && useradd -r -g Servers -m -d /home/urlserver urlserver
COPY ./bin/urlshortener /home/urlserver/urlshortener
# Set the working directory
WORKDIR /home/urlserver
# Switch to the non-root user
USER urlserver
EXPOSE 8000
CMD ["./urlshortener"]