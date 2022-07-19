# Sprint-Name-Generator
[Sprint Name Generator](https://sprintnamegenerator.com)

Originally built by: 
- [Florian Stadler](https://github.com/flostadler) - Go -> Wasm
- [Wolfgang Ederer](https://github.com/wederer) - Frontend

With support from following contributors:
- [Kris - The Coding Unicorn](https://github.com/guidemetothemoon)

## Usage

### Build the image locally

You will need to run a few commands to get necessary files generated.

1. Set environment variables:

* `GOROOT=[path to Go installation folder]`
* `GOOS=js`
* `GOARCH=wasm`

2. Run below commands:

``` bash
go build -o web/generator.wasm cmd/web
cp $GOROOT/misc/wasm/wasm_exec.js web/wasm_exec.js
cd web
npm install
npm run build --if-present

docker build -t sprint-name-generator -f build/package/Dockerfile .
```

3. Start the container: 

`docker run -dp 3000:80 sprint-name-generator:latest`

### Run latest image from Docker Hub
``` bash
docker pull flostadler/sprint-name-generator:[image_tag]
docker run -dp 3006:80  flostadler/sprint-name-generator:[image_tag]
```

## Images

Original prebuilt image available at [Dockerhub](https://hub.docker.com/r/flostadler/name-generator)
