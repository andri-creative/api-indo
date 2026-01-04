<div align="center">

![API Indo Logo](/api_indo_logo_1767542720552.png)

# 🇮🇩 API Indo - Indonesian Regional Data API

API RESTful sederhana untuk mendapatkan data wilayah Indonesia (Provinsi, Kabupaten/Kota, Kecamatan, dan Kelurahan/Desa) yang dibangun dengan **Go** dan **SQLite/LibSQL**.

</div>

## 📋 Daftar Isi

- [Fitur](#-fitur)
- [Teknologi](#-teknologi)
- [Instalasi](#-instalasi)
- [Konfigurasi](#-konfigurasi)
- [Menjalankan Server](#-menjalankan-server)
- [API Endpoints](#-api-endpoints)
- [Contoh Penggunaan](#-contoh-penggunaan)
- [Struktur Project](#-struktur-project)

## ✨ Fitur

- ✅ Mendapatkan semua data provinsi
- ✅ Mendapatkan kabupaten/kota berdasarkan provinsi
- ✅ Mendapatkan kecamatan berdasarkan kabupaten/kota
- ✅ Mendapatkan kelurahan/desa berdasarkan kecamatan
- ✅ Import data dari CSV
- ✅ Response format JSON
- ✅ Lightweight & Fast (menggunakan SQLite/LibSQL)

## 🛠 Teknologi

- **Go** 1.25.4
- **SQLite3** / **LibSQL** (Turso Database)
- **godotenv** - Environment variable management
- **net/http** - HTTP server bawaan Go

## 📦 Instalasi

### 1. Clone Repository

```bash
git clone https://github.com/andri-creative/api-indo.git
cd api-indo-golang
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Setup Database

Pastikan Anda memiliki file database SQLite atau konfigurasi LibSQL/Turso yang sudah terisi data wilayah Indonesia.

## ⚙️ Konfigurasi

### 1. Copy file `.env.example` menjadi `.env`

```bash
cp .env.example .env
```

### 2. Edit file `.env` sesuai konfigurasi Anda

```env
# Server Configuration
APP_PORT=8080
HOST=localhost

# JWT Configuration (if using authentication)
JWT_SECRET=your_jwt_secret_key_here

# Environment
ENV=development
```

> **Note:** Jika menggunakan LibSQL/Turso, sesuaikan konfigurasi database di file `database/libsql.go`

## 🚀 Menjalankan Server

```bash
go run main.go
```

Server akan berjalan di `http://localhost:8080` (atau port yang Anda tentukan di `.env`)

Output:
```
🚀 Server running on :8080
```

## 📡 API Endpoints

### 1. **Get All Provinces**

Mendapatkan semua data provinsi di Indonesia.

```
GET /provinces
```

**Response:**
```json
[
  {
    "id": "11",
    "name": "ACEH"
  },
  {
    "id": "12",
    "name": "SUMATERA UTARA"
  }
]
```

---

### 2. **Get Regencies by Province**

Mendapatkan kabupaten/kota berdasarkan ID provinsi.

```
GET /regencies?province_id={province_id}
```

**Parameters:**
- `province_id` (required) - ID provinsi

**Example:**
```
GET /regencies?province_id=11
```

**Response:**
```json
[
  {
    "id": "1101",
    "province_id": "11",
    "name": "KABUPATEN SIMEULUE"
  },
  {
    "id": "1102",
    "province_id": "11",
    "name": "KABUPATEN ACEH SINGKIL"
  }
]
```

---

### 3. **Get Districts by Regency**

Mendapatkan kecamatan berdasarkan ID kabupaten/kota.

```
GET /districts?regency_id={regency_id}
```

**Parameters:**
- `regency_id` (required) - ID kabupaten/kota

**Example:**
```
GET /districts?regency_id=1101
```

**Response:**
```json
[
  {
    "id": "1101010",
    "regency_id": "1101",
    "name": "TEUPAH SELATAN"
  },
  {
    "id": "1101020",
    "regency_id": "1101",
    "name": "SIMEULUE TIMUR"
  }
]
```

---

### 4. **Get Villages by District**

Mendapatkan kelurahan/desa berdasarkan ID kecamatan.

```
GET /villages?district_id={district_id}
```

**Parameters:**
- `district_id` (required) - ID kecamatan

**Example:**
```
GET /villages?district_id=1101010
```

**Response:**
```json
[
  {
    "id": "1101010001",
    "district_id": "1101010",
    "name": "LATIUNG"
  },
  {
    "id": "1101010002",
    "district_id": "1101010",
    "name": "LABUHAN BAJAU"
  }
]
```

---

### 5. **Import Data from CSV**

Endpoint untuk import data dari file CSV.

```
POST /import/simple
```

## 💡 Contoh Penggunaan

### Menggunakan cURL

#### 1. Get All Provinces
```bash
curl http://localhost:8080/provinces
```

#### 2. Get Regencies by Province ID
```bash
curl "http://localhost:8080/regencies?province_id=11"
```

#### 3. Get Districts by Regency ID
```bash
curl "http://localhost:8080/districts?regency_id=1101"
```

#### 4. Get Villages by District ID
```bash
curl "http://localhost:8080/villages?district_id=1101010"
```

### Menggunakan JavaScript (Fetch API)

```javascript
// Get all provinces
fetch('http://localhost:8080/provinces')
  .then(response => response.json())
  .then(data => console.log(data));

// Get regencies by province_id
fetch('http://localhost:8080/regencies?province_id=11')
  .then(response => response.json())
  .then(data => console.log(data));

// Get districts by regency_id
fetch('http://localhost:8080/districts?regency_id=1101')
  .then(response => response.json())
  .then(data => console.log(data));

// Get villages by district_id
fetch('http://localhost:8080/villages?district_id=1101010')
  .then(response => response.json())
  .then(data => console.log(data));
```

### Menggunakan Python (requests)

```python
import requests

# Get all provinces
response = requests.get('http://localhost:8080/provinces')
provinces = response.json()
print(provinces)

# Get regencies by province_id
response = requests.get('http://localhost:8080/regencies', params={'province_id': '11'})
regencies = response.json()
print(regencies)

# Get districts by regency_id
response = requests.get('http://localhost:8080/districts', params={'regency_id': '1101'})
districts = response.json()
print(districts)

# Get villages by district_id
response = requests.get('http://localhost:8080/villages', params={'district_id': '1101010'})
villages = response.json()
print(villages)
```

## 📁 Struktur Project

```
api-indo-golang/
├── database/
│   └── libsql.go          # Database connection setup
├── handlers/
│   ├── province.go        # Province handlers
│   ├── regency.go         # Regency handlers
│   ├── district.go        # District handlers
│   ├── village.go         # Village handlers
│   └── import_simple.go   # CSV import handler
├── models/
│   ├── province.go        # Province model
│   ├── regency.go         # Regency model
│   ├── district.go        # District model
│   └── village.go         # Village model
├── .env.example           # Environment variables template
├── .gitignore             # Git ignore file
├── go.mod                 # Go module dependencies
├── go.sum                 # Go module checksums
├── main.go                # Application entry point
└── README.md              # Documentation
```

## 🤝 Kontribusi

Kontribusi selalu diterima! Silakan buat **Pull Request** atau buka **Issue** untuk saran dan perbaikan.

## 📄 Lisensi

Project ini bersifat open source dan bebas digunakan untuk keperluan apapun.

## 👨‍💻 Author

**Andri Creative**
- GitHub: [@andri-creative](https://github.com/andri-creative)

---

⭐ Jangan lupa berikan **star** jika project ini bermanfaat!
