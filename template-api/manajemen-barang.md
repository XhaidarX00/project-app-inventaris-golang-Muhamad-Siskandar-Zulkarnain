### GET /api/items
tambahkan paramter pada url untuk fitur search dan filter
- Response Sukses
```json
{
    "success": true,
    "page": 1,
    "limit": 2,
    "total_items": 20,
    "total_pages": 2,
    "data": [
        {
            "id": 1,
            "name": "Laptop Lenovo Thinkpad",
            "category": "Elektronik",
            "photo_url": "/images/laptop.png",
            "price": 10000000,
            "purchase_date": "2023-01-15",
            "total_usage_days": 300
        },
        {
            "id": 2,
            "name": "Kursi Ergonomis",
            "category": "Furniture",
            "photo_url": "/images/kursi.png",
            "price": 2000000,
            "purchase_date": "2022-11-10",
            "total_usage_days": 500
        }
    ]
}
```

- Response Gagal
```json
{
    "success": false,
    "message": "Gagal mengambil data barang inventaris"
}
```

### POST /api/items
- Request body (kusus request ini menggunakan format form pada content-type)
```json
{
    "name": "Printer Canon",
    "category_id": 1,
    "photo_url": "/images/printer.png",
    "price": 1500000,
    "purchase_date": "2024-01-05"
}
```

- Response Sukses
```json
{
    "success": true,
    "message": "Barang berhasil ditambahkan",
    "data": {
        "id": 3,
        "name": "Printer Canon",
        "category": "Elektronik",
        "photo_url": "/images/printer.png",
        "price": 1500000,
        "purchase_date": "2024-01-05",
        "total_usage_days": 0
    }
}
```

- Response Gagal (misalnya, kategori tidak ditemukan)
```json
{
    "success": false,
    "message": "Kategori tidak ditemukan"
}
```

### GET /api/items/{id}
- Response Sukses (Status Code: 200)
```json
{
    "success": true,
    "status_code": 200,
    "data": {
        "id": 1,
        "name": "Laptop Lenovo Thinkpad",
        "category": "Elektronik",
        "photo_url": "/images/laptop.png",
        "price": 10000000,
        "purchase_date": "2023-01-15",
        "total_usage_days": 300
    }
}
```

- Response Gagal (Status Code: 404) (misalnya, jika barang dengan id tidak ditemukan)
```json
{
    "success": false,
    "status_code": 404,
    "message": "Barang tidak ditemukan"
}
```

### PUT /api/items/{id}
- Request body
```json
{
    "name": "Laptop Lenovo Thinkpad Updated",
    "category_id": 1,
    "photo_url": "/images/laptop_updated.png",
    "price": 10500000,
    "purchase_date": "2023-01-15"
}
```

- Response Sukses (Status Code: 200)
```json
{
    "success": true,
    "status_code": 200,
    "message": "Barang berhasil diperbarui",
    "data": {
        "id": 1,
        "name": "Laptop Lenovo Thinkpad Updated",
        "category": "Elektronik",
        "photo_url": "/images/laptop_updated.png",
        "price": 10500000,
        "purchase_date": "2023-01-15",
        "total_usage_days": 300
    }
}
```

- Response Gagal (Status Code: 404) (misalnya, jika barang tidak ditemukan)
```json
{
    "success": false,
    "status_code": 404,
    "message": "Barang tidak ditemukan"
}
```

### DELETE /api/items/{id}
- Response Sukses (Status Code: 200)
```json
{
    "success": true,
    "status_code": 200,
    "message": "Barang berhasil dihapus"
}
```

- Response Gagal (Status Code: 404)
```json
{
    "success": false,
    "status_code": 404,
    "message": "Barang tidak ditemukan"
}
```