# TOP-SECRET-SPLIT GET

Para este servicio necesitaremos 3 satellites los cuales obtendremos de DynamoDB

# Solicitud de ejemplo

```
curl --location 'https://dmzfktok87.execute-api.us-east-2.amazonaws.com/v1/topsecret_split'
```

# Respuesta exitosa 200

```
{
    "X": -185.12361,
    "Y": 310.74167,
    "message": "este es un mensaje secreto"
}
```

![200 ok](image-4.png)


# Respuesta error 404

```
Error: [RequestId: f6b3eee0-08eb-4303-819b-7e314f84d0a1][Error: Satellites are not enough]
```

![404 error](image-5.png)