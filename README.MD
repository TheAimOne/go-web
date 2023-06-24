Initial Commit


## Local Database
```
    docker-compose -f file.yml up -d
```

## Event endpoint

```

curl --location --request POST 'localhost:8080/events' \
--header 'Content-Type: application/json' \
--data-raw '{
    "groupId": "663b0627-d5a8-4b74-8db9-4090c8df50f1",
    "creatorId":"8fd11aa0-5234-42a9-9095-637cb7dff592",
    "name": "FootBall sa",
    "type": "F",
    "params": {
        "total": 10,
        "standby": "5"
    }
}'

```