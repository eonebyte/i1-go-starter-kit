# i1-go-starter-kit: GoFiber-Based MVC Innovation: Auto-Migrate, Auth-Protect, & CLI Magic

i1-go-starter-kit adalah sebuah proyek inovatif yang memanfaatkan kekuatan GoFiber, sebuah kerangka kerja mutakhir untuk membangun aplikasi web yang tangguh. Proyek ini fokus pada penyediaan arsitektur MVC (Model-View-Controller) yang komprehensif, menyederhanakan proses pengembangan dengan migrasi database otomatis, mekanisme otentikasi yang diperkuat, dan keunggulan dari antarmuka baris perintah (Command Line Interface/CLI) untuk meningkatkan produktivitas.

## Fitur Utama

- **Auto-Migration**: Sinkronisasi model dengan skema basis data secara otomatis.
- **Perlindungan Autentikasi**: Perlindungan otentikasi bawaan untuk menjaga keamanan aplikasi.
- **Keajaiban CLI**: Membuat kontroler dan model dengan mudah untuk pengembangan yang cepat.

i1-go-starter-kit bertujuan memberikan pengembang landasan yang efisien dan aman untuk membuat aplikasi web, memanfaatkan kemampuan kinerja tinggi dari GoFiber. Baik Anda seorang pengembang berpengalaman maupun yang baru memulai, i1 menawarkan lingkungan yang kaya fitur untuk mempercepat usaha pengembangan web Anda.

## Cara Menggunakan

### Clone Proyek

1. Untuk meng-clone proyek ini, jalankan perintah berikut di terminal:

```
git clone https://github.com/eonebyte/i1-go-starter-kit.git
```
2. Setelah mengunduh proyek, pastikan untuk menjalankan langkah-langkah berikut untuk mengelola dependensi dan memulai proyek.
### Buat folder template pada folder static
```
static/template/assets
static/template/dist
```
### Pastikan dependensi yang diperlukan terinstal dengan menjalankan perintah berikut:

```
go mod tidy
```

## Dokumentasi CLI

### Generate Controller or Model Files

- **Usage:** `go run main.go [namaFile] [tyOfFile]`
- **Description:** Perintah ini menghasilkan file kontroler atau model berdasarkan argumen yang diberikan.
  - `[namaFile]`: Menentukan nama file yang akan dihasilkan.
  - `[tyOfFile]`: Mendefinisikan jenis file yang akan dihasilkan. Pilihan: controllers, models.

### Contoh Penggunaan:

1. Generate Kontroler:
   ```
   go run main.go user_controller controllers
   ```
3. Generate Model:
   ```
   go run main.go user models
   ```
