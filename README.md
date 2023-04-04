# Balance JSON API
## Test task for the position of trainee golang backend developer

---
### 1. Requirements

* **PostgreSQL:** `17.2`
* **Go:** `20.0`

## docker-compose
Server is ready immediately after containers are up
```shell
 make compose        #docker-compose up
```

## Dockerfile 
```shell
make docker-build   #docker build -t balance-api .
make docker-run     #docker run balance-api
```

### 2. Application building

```shell
git clone https://github.com/Kl1ck9r/Balance.git 
cd BalanceAPI
make build 
make run 
```

## Tests
* [x] Unit 

```shell
make test 
```
---

### `/get` - Get User Balance

* Method: `GET`

```json5
Request :
{
  "id_user": 2 
}
```

```json5
Response :
{
  "ok": true,
  "currency": "RUB", 
  "balance": "8300" 
}
```

```json5
Request :
{
  "id_user":  2, 
  "currency": "Dollars"
}
```
```json5
Response :
{
    "ok": true,
    "currency": "Dollars",
    "balance": "67.61"
}
```


```json5
Request :
{
  "id_user":  2, 
  "currency": "EURO"
}
```

```json5
Response :
{
    "ok": true,
    "currency": "EUR",
    "balance": "64.05"
}
```


#### Possible error:

```json5
{
    "ok": false,
    "error": "Not Found user :no rows in result set",
    "status": 404
}
```

### `/descrease` - Descrease user balance 

* Method: `POST`

```json5
Request :
{
   "id_user":1,
   "amount": "200"
}
```
```json5
Response :
{
    "balance": "4700",
    "currency": "RUB",
    "description": "Amount writter from account:200"
}
```

#### Possible error:
```json5
{
    "ok": false,
    "error": "User not found: no rows in result set ",
    "status": 404
}
```


### `/replenish/balance` -  replenish user balance 

* Method: `POST`

```json5
Request :
{
   "user_id": 3,
   "balance": "4500",
   "currency":"RUB"
}
```

```json5
Response :
{
    "balance": "4500",
    "currency": "RUB",
    "description": "replenishment amount: 4500"
}
```
#### Possible error:

```json5
{
    "ok": false,
    "error": "The replenishment amount cannot be less than zero"
    "status": 400
}
```


### `/transaction` - Transaction between users 
* Method: `POST`

```json5
Request :
{
   "to_id": 3,
   "from_id": 1,
   "amount": "2290"
}
```

```json5
Response :
{
    "to_id_balance": "3920",
    "from_id_balance": "5080",
    "description": "user balance after transaction: 3920"
}
```
#### Possible error:

```json5
{
    "ok": false,
    "error": "[DB ERROR]: no rows in result set",
    "status": 400
}
```


### `/delete` -  Delete User Balance

* Method: `DELETE`

```json5
Request :
{
   "user_id":3
}
```

```json5
Response :
{
    "user_id": 3,
    "balance": "0",
    "description": "User with this id: 3 Success deleted"
}
```

#### Possible error:

```json5
{
    "ok": false,
    "error": "Failed delete user [DB]:Wrong enter user id",
    "status": 400
}
```

### `/get/list/transactions` -  Get User Transactions

* Method: `Get`

```json5
Request :
{
    "user_id":1,
    "limit":2
}
```

```json5
Response :
[
    {
        "to_id": 1,
        "from_id": 2,
        "amount": "100",
        "descrption": "User with: 2 sent 100 to 1"
    },
    {
        "to_id": 1,
        "from_id": 3,
        "amount": "100",
        "descrption": "Amount has been transffered: 100 fromID: 3 toID 1"
    }
]

```
#### Possible error:

```json5
{
    "ok": false,
    "error": "Error get list transactions:User id cannot is negative",
    "status": 400
}
```
