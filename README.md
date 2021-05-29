# BPP-Resolve (Backend)

## How to run server application
```shell
chmod +x start.sh
./start.sh
```
or run command
```text
make start
```

#### Server side default value config
```text
- Server: 
    - host: 127.0.0.1 | localhost
    - port: 8080
- Database:
    - host: 127.0.0.1
    - port: 3306
```


## API Interface

#### Health
- Method: `GET`
- URL path: `/health/ready`

#### Place Wager

- Method: `POST`
- URL path: `/wagers`
- Request body:
    ```
    {
        "total_wager_value": <total_wager_value>,
        "odds": <odds>,
        "selling_percentage": <selling_percentage>,
        "selling_price": <selling_price>,
    }
    ```

- Response:
  Header: `HTTP 201`
  Body:
    ```
    {
        "id": <wager_id>,
        "total_wager_value": <total_wager_value>,
        "odds": <odds>,
        "selling_percentage": <selling_percentage>,
        "selling_price": <selling_price>,
        "current_selling_price": <current_selling_price>,
        "percentage_sold": <percentage_sold>,
        "amount_sold": <amount_sold>,
        "placed_at": <placed_at>
    }
    ```
  or

  Header: `HTTP <HTTP_CODE>`
  Body:
    ```
    {
        "error": "ERROR_DESCRIPTION"
    }
    ```


#### Buy wager

- Method: `POST`
- URL path: `/buy/:wager_id`
- Request body:
    ```
    {
        "buying_price": <buying_price>
    }
    ```

- Response:
  Header: `HTTP 201`
  Body:
    ```
    {
        "id": <purchase_id>,
        "wager_id": <wager_id>,
        "buying_price": <buying_price>,
        "bought_at": <bought_at>
    }
    ```
  or

  Header: `HTTP <HTTP_CODE>`
  Body:
    ```
    {
        "error": "ERROR_DESCRIPTION"
    }
    ```


#### Wager list

- Method: `GET`
- URL path: `/wagers?page=:page&limit=:limit`
- Response:
  Header: `HTTP 200`
  Body:
    ```
    [
        {
            "id": <wager_id>,
            "total_wager_value": <total_wager_value>,
            "odds": <odds>,
            "selling_percentage": <selling_percentage>,
            "selling_price": <selling_price>,
            "current_selling_price": <current_selling_price>,
            "percentage_sold": <percentage_sold>,
            "amount_sold": <amount_sold>,
            "placed_at": <placed_at>
        }
        ...
    ]
    ```

## Testing
```text
make test
```