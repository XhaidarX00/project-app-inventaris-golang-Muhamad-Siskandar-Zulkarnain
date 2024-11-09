### GET /api/items/replacement-needed
- Response Sukses
```json
{
    "success": true,
    "data": [
        {
            "id": 1,
            "name": "Laptop Lenovo Thinkpad",
            "category": "Elektronik",
            "purchase_date": "2019-01-15",
            "total_usage_days": 1800,
            "replacement_required": true
        }
    ]
}
```

- Response Gagal
```json
{
    "success": false,
    "message": "Gagal mengambil data barang yang perlu diganti"
}
```