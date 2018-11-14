# Sprint-Name-Generator
Simple sprint name generator using the name-generator of the docker engine

## Building
Run `docker image build -t sprint-name-generator .`

## Usage
Start the container: `docker run -p 80 -d sprint-name-generator`

To generate a sprint name, browse to `localhost:"PORT"/generate`