# GO_Standard_Library

Kumpulan contoh penggunaan package standar Go yang lengkap dan terstruktur. Proyek ini cocok untuk belajar, referensi, dan eksperimen fitur-fitur penting dari bahasa Go.

## Struktur Direktori

- **main.go** — Menjalankan seluruh demo package.
- **package_fmt/** — Contoh penggunaan package fmt.
- **package_math/** — Contoh penggunaan package math.
- **package_reflect/** — Contoh penggunaan package reflect.
- **package_regexp/** — Contoh penggunaan package regexp.
- **package_encoding/** — Contoh encoding/decoding (base64, json, xml, gob, csv).
- **package_slices/** — Contoh manipulasi slice modern.
- **package_path/** — Contoh manipulasi path.
- **package_io/** — Contoh IO, bufio, dan manipulasi file.

Setiap package memiliki file `penjelasan.txt` untuk dokumentasi dan file `run_all.go` untuk menjalankan seluruh demo di package tersebut.

## Cara Menjalankan

1. Pastikan sudah terinstall Go (minimal Go 1.21 agar mendukung package slices).
2. Clone/download repo ini.
3. Jalankan:

```bash
cd GO_Standard_Library
go run main.go
```

## Kontribusi & Lisensi

- Bebas digunakan untuk belajar dan referensi.
- Kontribusi, perbaikan, dan penambahan contoh sangat terbuka!

---

Dibuat untuk mempermudah eksplorasi fitur-fitur penting Go secara praktis dan terstruktur.
