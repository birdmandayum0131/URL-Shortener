FROM golang:1.22.3-bullseye
LABEL author="Bird"
# Create a non-root user and set the home directory
RUN groupadd -r Servers && useradd -r -g Servers -m -d /home/urlserver urlserver
# Copy project to the image for building
COPY . /home/urlserver/urlproject
# Build the project and move it to home dir
WORKDIR /home/urlserver/urlproject
RUN ./docker/build.sh && \
    mv ./build/urlshortener /home/urlserver/urlshortener
# Remove the project folder after build
WORKDIR /home/urlserver
RUN rm -r urlproject
# Switch to the non-root user
USER urlserver
EXPOSE 8000
CMD ["./urlshortener"]