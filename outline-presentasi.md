# Proyek RestfulAPI: Sistem Inventaris Kantor

## Slide 1: Gambaran Umum Proyek
- API REST untuk sistem inventaris kantor
- Diimplementasikan dalam Golang
- Menggunakan database PostgreSQL

## Slide 2: Fitur Utama
- Manajemen Kategori
- Manajemen Barang Inventaris
- Pelacakan Penggantian Barang
- Pelaporan Investasi dan Depresiasi
- Fitur Tambahan (Penyaringan, Paginasi)

## Slide 3: Endpoint API - Kategori
- GET /api/categories
- POST /api/categories
- GET /api/categories/{id}
- PUT /api/categories/{id}
- DELETE /api/categories/{id}

## Slide 4: Endpoint API - Barang Inventaris
- GET /api/items
- POST /api/items
- GET /api/items/{id}
- PUT /api/items/{id}
- DELETE /api/items/{id}

## Slide 5: Endpoint API - Fitur Khusus
- GET /api/items/replacement-needed
- GET /api/items/investment
- GET /api/items/investment/{id}

## Slide 6: Persyaratan Teknis
- Bahasa pemrograman Golang
- Pola repository
- PostgreSQL dengan driver github.com/lib/pq
- go-chi untuk routing
- Implementasi paginasi
- Validasi input
- Format JSON untuk pertukaran data
- Kode status HTTP yang sesuai

## Slide 7: Fitur Tambahan
- Login, registrasi, dan manajemen sesi

## Slide 8: Template Respons API
- Tautan ke repositori GitHub:
  https://github.com/XhaidarX00/project-app-inventaris-golang-Muhamad-Siskandar-Zulkarnain.git


## Slide 9: Arsitektur Proyek
- [Link Diagram Flowchart : https://mermaid.live/view#pako:eNqFlNtymzAQhl9Fo2ufwPGhZKYzccCH-EQMuSn2hWI2tmZAUCE6TY3fvTLCrZzaKReLf_Ttr5F2vQe8TULAFt5xku6Rb68Zks9DMM8jQjeoXv-KBgeXRnSPlilwktGjQlQcnIhi5PgFeqwwh4VpQpk46oS79CRiBz6JX8ke2UQQNCA831xAL5JxAiekoiQuFm1n5vhOgYbBmKR5pgMqPpZYk6S0uSUCdgmnkBVoFDzErzRCHsQ5QdNqZfNJTpOGBRpXaTYIIl-386iAWG4zudhGHo2w3Q249H-69P88gUMakS3EwESdAYQgDaaVgcpELvAoRzbdESbozY3ZD8jEyaZAsyrfTwSJ0ILKcqNJuS5r_H-D8hDzykNl25ByyKhMv__rdHkyFe3rlVqcu-N82XqH2P_c-PKMVzegwSo6N4vrqh77UFTnWp2eFXrlFMOb9quqRT_4D6_5exV7ZYNR-e_zA5cnGWRoBd9zeaf3aEo5jaXM0oRlVcZYsUpMdPGki6kuZrqY62Khi6UuXF0862KlC08XKvrlp5fAgwgyOVpwDcfAY0JDOX0OJ2aNxV72-Bpb8mcIbySPxBqv2VGiJBeJ98622BI8hxrmSb7bY-uNRJlUeRrKq7YpkVMsPiMpYd-SJP4DSY2tA_6JrXrfbLTNlmH0TbPXktGo4Xdsde4arXavddfrdttG3_jSP9bwr9LBaJhGSz5mr9vpdjodo1_DINsi4XM1PcshevwNBwufoA]

## Slide 10: Skema Database
- [File Datanase ada didirectory]