### GET /api/categories
- Response Sukses (Status Code: 200)
```json
{
    "success": true,
    "status_code": 200,
    "data": [
        {
            "id": 1,
            "name": "Elektronik",
            "description": "Peralatan elektronik kantor"
        },
        {
            "id": 2,
            "name": "Furniture",
            "description": "Peralatan furniture kantor"
        }
    ]
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

- Response Gagal (Status Code: 500)
```json
{
    "success": false,
    "status_code": 500,
    "message": "Gagal mengambil data kategori"
}
```

### POST /api/categories
- Request body
```json
{
    "name": "Peralatan Dapur",
    "description": "Kategori untuk peralatan dapur"
}

```

- Response Sukses (Status Code: 201):
```json
{
    "success": true,
    "status_code": 201,
    "message": "Kategori berhasil ditambahkan",
    "data": {
        "id": 3,
        "name": "Peralatan Dapur",
        "description": "Kategori untuk peralatan dapur"
    }
}
```

- Response Gagal (Status Code: 400) (misalnya, nama kategori kosong)
```json
{
    "success": false,
    "status_code": 400,
    "message": "Nama kategori tidak boleh kosong"
}
```

### GET /api/categories/{id}
- Response Sukses (Status Code: 200)
```json
{
    "success": true,
    "status_code": 200,
    "data": {
        "id": 1,
        "name": "Elektronik",
        "description": "Peralatan elektronik kantor"
    }
}
```

- Response Gagal (Status Code: 404)
```json
{
    "success": false,
    "status_code": 404,
    "message": "Kategori tidak ditemukan"
}
```

### PUT /api/categories/{id}
- Request body
```json
{
    "name": "Peralatan Elektronik",
    "description": "Peralatan elektronik yang digunakan di kantor"
}
```

- Response Sukses (Status Code: 200)
```json
{
    "success": true,
    "status_code": 200,
    "message": "Kategori berhasil diperbarui",
    "data": {
        "id": 1,
        "name": "Peralatan Elektronik",
        "description": "Peralatan elektronik yang digunakan di kantor"
    }
}
```

- Response Gagal (Status Code: 404) (misalnya, jika id kategori tidak ditemukan)
```json
{
    "success": false,
    "status_code": 404,
    "message": "Kategori tidak ditemukan"
}
```

### DELETE /api/categories/{id}
- Response Sukses (Status Code: 200)
```json
{
    "success": true,
    "status_code": 200,
    "message": "Kategori berhasil dihapus"
}
```
- Response Gagal (Status Code: 404)
```json
{
    "success": false,
    "status_code": 404,
    "message": "Kategori tidak ditemukan"
}
```