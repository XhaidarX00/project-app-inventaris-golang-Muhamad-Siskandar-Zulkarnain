### GET /api/items/investment
- Response Sukses
```json
{
    "success": true,
    "data": {
        "total_investment": 25000000,
        "depreciated_value": 18000000
    }
}
```

- Response Gagal (Status Code: 404)
```json
{
    "success": false,
    "status_code": 404,
    "message": "Data not found"
}
```

- Response Gagal
```json
{
    "success": false,
    "message": "Gagal menghitung total investasi"
}
```

### GET /api/items/investment/{id}
- Response Sukses
```json
{
    "success": true,
    "data": {
        "item_id": 1,
        "name": "Laptop Lenovo Thinkpad",
        "initial_price": 10000000,
        "depreciated_value": 7000000,
        "depreciation_rate": 10
    }
}
``` 

- Response Gagal (misalnya, barang tidak ditemukan)
```json
{
    "success": false,
    "message": "Barang tidak ditemukan"
}
``` 