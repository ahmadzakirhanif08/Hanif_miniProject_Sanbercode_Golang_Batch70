# 📚 Mini Project: Golang Book API

[![Build Status](https://img.shields.io/badge/Status-Completed-brightgreen)](https://railway.app/project-link)
[![Language](https://img.shields.io/badge/Language-Golang-blue)](https://go.dev/)
[![Framework](https://img.shields.io/badge/Framework-Gin_Gonic-red)](https://github.com/gin-gonic/gin)
[![Database](https://img.shields.io/badge/Database-PostgreSQL/MySQL-informational)]()

Mini Project ini dikembangkan sebagai bagian dari Quiz Bootcamp Intensif Golang
untuk mengelola data buku dan kategori dengan fokus pada implementasi RESTful API, Basic Authentication, dan deployment ke Railway.

## 💡 Kegunaan Proyek

Proyek ini menyediakan API *backend* lengkap untuk sistem katalog buku sederhana, memungkinkan operasi *Create, Read, Update,* dan *Delete* (CRUD) pada entitas Kategori dan Buku[.

Fitur utama meliputi:
1.  **Basic Authentication Middleware** untuk mengamankan semua *endpoint* API.
2.  **Validasi Bisnis Khusus:**
    * Validasi `release_year` antara **1980 dan 2024**.
    * Konversi otomatis `total_page` menjadi `thickness` (**tebal** jika > 100, **tipis** jika < 100) saat input buku.
3.  Koneksi ke **Relational Database** (PostgreSQL/MySQL) menggunakan `sql-migrate`.
4.  Deployment ke platform **Railway**

## 🛠️ Setup Project (Lokal)

### Prerequisites
* Go (Golang)
* PostgreSQL atau MySQL
* Git
* File konfigurasi database sudah diatur di `config/database.go`.

### Langkah-langkah Instalasi

1.  **Clone Repository:**
    ```bash
    git clone [Link Repository Anda]
    cd golang-book-api
    ```

2.  **Inisialisasi Modules dan Dependencies:**
    ```bash
    go mod tidy
    #  Dependencies utama: [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin), [github.com/lib/pq](https://github.com/lib/pq), [github.com/rubenv/sql-migrate](https://github.com/rubenv/sql-migrate)
    ```

3.  **Jalankan Migrasi Database:**
    Pastikan database Anda berjalan dan *connection string* di `config/database.go` sudah benar. Migrasi akan dijalankan otomatis saat aplikasi pertama kali dijalankan oleh `main.go`.

4.  **Jalankan Aplikasi:**
    ```bash
    go run main.go
    # Server akan berjalan di localhost:8080
    ```

## 🔐 Cara Penggunaan (Basic Authentication)

 Semua *endpoint* di bawah ini memerlukan *header* **Authorization**  dengan skema **Basic Auth**.

Gunakan salah satu kredensial yang telah diatur di `routes/routes.go`:
* **Kredensial 1:** `admin:supersecretpassword`
* **Kredensial 2:** `developer:devpass`

**Contoh Header:**

## 🗺️ List Path API yang Tersedia

Base URL: `http://localhost:8080/api`

---

### 1. Kategori API (`/api/categories`) 

| Path | Method | Kegunaan | Keterangan |
| :--- | :--- | :--- | :--- |
| `/categories` |`GET`| Menampilkan seluruh kategori. | Memerlukan Basic Auth. |
| `/categories` |`POST`| Menambahkan kategori | Memerlukan Basic Auth. |
| `/categories/:id`|`GET`| Menampilkan detail kategori. | Memerlukan Basic Auth. |
| `/categories/:id`|`DELETE`| Menghapus kategori. | Memerlukan Basic Auth.|
| `/categories/:id/books` |`GET`|  Menampilkan buku berdasarkan kategori| Memerlukan Basic Auth. |

---

###  2. Buku API (`/api/books`) 

| Path | Method | Kegunaan | Keterangan |
| :--- | :--- | :--- | :--- |
|`/books` | `GET` |  Menampilkan seluruh buku. | Memerlukan Basic Auth. |
| `/books` |  `POST` |  Menambahkan buku. | Memerlukan Basic Auth.  Terdapat validasi `release_year` dan konversi `total_page` ke `thickness`. |
| `/books/:id` |  `GET`  |  Menampilkan detail buku. | Memerlukan Basic Auth. |
| `/books/:id` |  `DELETE`  |  Menghapus buku. | Memerlukan Basic Auth. |

## 🚀 Deployment

 Proyek ini telah dideploy ke **Railway**.

[Link Deployment Railway Anda]
