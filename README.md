# api

shokujinjpのWeb API

## Endpoints

### GET `/`

Endpoint for healthcheck.

#### response

```
{
  "health": "ok"
}
```

### GET `/menu/today`

List of menus that can order today.

#### response

```
[
  {
    "id": "1",
    "name": "麻婆豆腐",
    "price": "600",
    "category": "定食",
    "day_start": "",
    "day_end": "",
    "can_weekday": "",
    "description": "1番"
  },
  ...
]
```


### GET `/menu/all`

List of all menus in [shokujinjp/data](https://github.com/shokujinjp/data)

#### response

```
[
  {
    "id": "1",
    "name": "麻婆豆腐",
    "price": "600",
    "category": "定食",
    "day_start": "",
    "day_end": "",
    "can_weekday": "",
    "description": "1番"
  },
  ...
]
```
