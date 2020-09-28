# Go Clean Architechture

[![Build Status](https://github.com/prongbang/goclean/workflows/Go/badge.svg)](https://github.com/prongbang/goclean/actions)
[![Codecov](https://img.shields.io/codecov/c/github/prongbang/goclean.svg)](https://codecov.io/gh/prongbang/goclean) 
[![Go Report Card](https://goreportcard.com/badge/github.com/prongbang/goclean)](https://goreportcard.com/report/github.com/prongbang/goclean)

<img src="http://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg">
<center>ภาพจาก blog.cleancoder.com</center>

### Swagger Generate

```
$ make swaggen
```

- http://localhost:1323/swagger/index.html


### REST API

```
$ make run
```

#### ADD

- Request

```
POST http://localhost:1323/api/v1/promotion
```

Body

```
{
    "id": 1,
    "code": "sd-promo",
    "name": "Sunday promotion",
    "priority": 4,
    "exclusive": false,
    "used": 0,
    "couponBased": false,
    "rules": [],
    "actions": [],
    "createdAt": "2017-02-28T12:05:12+0100",
    "updatedAt": "2017-02-28T12:05:13+0100",
    "channels": [],
    "_links": {
        "self": {
            "href": "\/api\/v1\/promotions\/sd-promo"
        }
    }
}
```

#### GET ALL

- Request

```
GET http://localhost:1323/api/v1/promotion

```

#### GET BY ID

- Request

```
GET http://localhost:1323/api/v1/promotion/1

```
