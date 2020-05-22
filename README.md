# go-torch

![build](https://github.com/TorchPing/go-torch/workflows/build/badge.svg?branch=master)

Torch backend rewrote on Golang with ARM support

## Usage

- Ping TCP port
  - `/ping/:host/:port`
  - onSuccess: `{"status":true,"time":975.9993876666667}`
  - onFailure(3 tries): `{"status":false,"time":0}`
- Resolve given addr
  - `/resolve/:host` using default resolver
  - onSuccess: `{"result":["220.181.38.148","39.156.69.79"],"status":true,"time":34.33183533333334}`
  - onFailure(3 tries): `{"result":null,"status":false,"time":0}`

## Download

### Binary 

[Release](releases)

###  Docker 

https://hub.docker.com/repository/docker/neverbehave/go-torch

```bash
docker run -d -p 8080:8080 neverbehave/go-torch
```
|key|required|default|
|---|---|---|
|PORT|false|8080|
|GIN_MODE|false|production|