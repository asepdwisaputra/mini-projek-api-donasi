# Mini Proyek API Donasi
## Deskripsi Proyek
Proyek Mini API Donasi adalah bagian dari pengalaman belajar saya dengan metode Project-Based Learning. Proyek ini merupakan salah satu topik belajar dalam program MSIB Kemendikbud, yang secara khusus dalam kasus saya, melibatkan kolaborasi dengan mitra Alterra Academy. Tujuan dari proyek ini adalah untuk memahami dasar-dasar pengembangan API dan mengaplikasikannya dalam konteks penggalangan dana atau donasi.

## Teknologi yang Digunakan
Dalam proyek ini, saya menggunakan sejumlah teknologi dan alat penting. Berikut adalah daftar teknologi yang digunakan dalam proyek ini:
1. `Go (Golang)`: Go adalah bahasa pemrograman yang digunakan untuk mengembangkan aplikasi server-side, termasuk API dalam proyek ini.
2. `Echo`: Echo adalah framework web berperforma tinggi untuk Go yang digunakan untuk membangun API dengan mudah dan efisien.
3. `GORM`: GORM adalah ORM (Object-Relational Mapping) untuk Go yang digunakan untuk berinteraksi dengan database MySQL.
4. `JWT(JSON Web Tokens)`: JWT digunakan untuk otentikasi dan otorisasi pengguna pada API.
5. `JSON`: Format data JSON digunakan untuk berkomunikasi dengan backend dan menyimpan data dalam aplikasi.
6. `Postman`: Postman digunakan untuk menguji dan memvalidasi endpoint API yang telah dibuat.
7. `MySQL`: MySQL adalah basis data relasional yang digunakan untuk menyimpan dan mengelola data terkait donasi serta informasi pengguna.
8. `db4free`: db4free adalah penyedia database MySQL gratis yang digunakan untuk pengembangan proyek.
9. `Git` dan GitHub: Git digunakan untuk mengelola versi proyek, dan GitHub digunakan sebagai platform kolaborasi dan penyimpanan repositori proyek.
10. `Draw.io`: Draw.io adalah alat untuk membuat diagram alur dan arsitektur proyek yang membantu dalam perancangan API.

## Cara Menjalankan Proyek
Untuk menjalankan proyek ini di komputer Anda, ikuti langkah-langkah berikut:

### Kloning Repositori
Klon repositori ini ke komputer Anda menggunakan perintah `git clone`.
```bash
git clone https://github.com/asepdwisaputra/mini-projek-api-donasi.git
```

### Instalasi Dependencies
Navigasikan ke direktori proyek dan instal semua dependensi dengan perintah berikut:
```bash
go mod tidy
```

### Konfigurasi Database
Konfigurasi koneksi ke database MySQL dalam `config.go`.

### Konfigurasi Tambahan
Membuat Secret Key JWT dalam `constants/constant.go`.
```go
package constants

const SECRET_JWT = "<key-password>"

```

### Menjalankan Aplikasi
Jalankan server dengan perintah berikut:

```bash
go run main.go
```

### Menggunakan Aplikasi
Buka browser atau aplikasi Postman untuk mengakses API melalui alamat `http://localhost:8000`.

## Dokumentai Lainnya
- **ERD** https://drive.google.com/file/d/1wNnIh4S79mLTHfl9iUVV1SqgTVYH9Ev6/view?usp=sharing

## Kontribusi
Saya menyambut kontribusi dan saran dari para kontributor. Jika Anda ingin berkontribusi, silakan buat pull request atau laporkan masalah (issue) di repositori ini.

## Kontak
Jika Anda memiliki pertanyaan atau memerlukan bantuan lebih lanjut, Anda dapat menghubungi saya melalui email di [asepputra3003@gmail.com]. Atau kunjungi media sosial saya di https://linktr.ee/4sep

## Terima kasih
Terima kasih sudah mengunjungi proyek Mini API Donasi ini. Semoga proyek ini membantu Anda dalam belajar dan pengembangan lebih lanjut!
