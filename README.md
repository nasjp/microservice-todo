# microservice-todo

## dependencies

- Go
- Docker
- Kubernetes
- gRPC

## wake up

```sh
make up
```

## how to use

```sh
curl -s -X POST -d "{\"title\":\"buy phone\"}" (make endpoint)/create | jq .
{
  "id": "71d6573e-778c-45d2-8232-8e3b6533e945",
  "title": "buy phone"
}

curl -s (make endpoint)/list | jq .
[
  {
    "id": "2aa7d540-97b5-460b-8afb-f9c1261c9021",
    "title": "buy pc"
  },
  {
    "id": "71d6573e-778c-45d2-8232-8e3b6533e945",
    "title": "buy phone"
  }
]
```
