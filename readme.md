# Sprint-Name-Generator
Simple sprint name generator using the name-generator of the docker engine

## Building
Run `docker image build -t sprint-name-generator .`

## Usage
Start the container: `docker run -p 80 -d sprint-name-generator`

To generate a sprint name, browse to `localhost:"PORT"/generate`

## Images

Prebuilt image available at [Dockerhub](https://hub.docker.com/r/flostadler/sprint-name-generator)
