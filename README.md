# Praktikum gRPC Polyglot

## Deskripsi

Praktikum ini membahas komunikasi antar-service menggunakan gRPC dengan pendekatan polyglot, yaitu penggunaan bahasa pemrograman yang berbeda dalam satu sistem.

Pada proyek ini, server dibuat menggunakan **Go**, sedangkan client dibuat menggunakan **Python** dan **Node.js**. Komunikasi dilakukan menggunakan gRPC dan Protocol Buffers (Protobuf).

## Tujuan Praktikum

- Memahami konsep komunikasi inter-service menggunakan gRPC.
- Membuat server gRPC menggunakan bahasa Go.
- Menggunakan Protocol Buffers untuk mendefinisikan layanan dan struktur data.
- Menghubungkan client Python dan Node.js ke server gRPC.
- Menguji pertukaran data antara server dan client yang berbeda bahasa pemrograman.

## Teknologi yang Digunakan

- Go
- Python
- Node.js
- gRPC
- Protocol Buffers (Protobuf)
- Visual Studio Code
- Git dan GitHub

## Struktur Proyek

```text
grpc-polyglot/
├── .gitignore
├── README.md
├── go.mod
├── go.sum
├── pb/
│   ├── user.pb.go
│   └── user_grpc.pb.go
├── server/
│   └── main.go
├── python-client/
│   └── client.py
├── node-client/
│   ├── client.js
│   └── package.json
└── Laporan_Praktikum_gRPC_Polyglot_Raya_Deswita_Dengan_Screenshot.docx
```

*Catatan: sesuaikan struktur di atas dengan nama folder dan file yang benar-benar ada pada proyek kamu.*

## Hasil Pengujian

Berdasarkan screenshot pengujian client yang tersedia, client Python dan Node.js menampilkan data profil berikut:

- Nama: Budi Santoso
- Email: budi@example.com

Hasil tersebut menunjukkan bahwa kedua client berhasil menampilkan respons data pada pengujian yang didokumentasikan.

## Kesimpulan

Praktikum ini memberikan pemahaman mengenai komunikasi antar-service menggunakan gRPC dengan bahasa pemrograman yang berbeda. Penggunaan Protocol Buffers membantu mendefinisikan format data dan layanan agar dapat digunakan oleh server Go serta client Python dan Node.js.

## Identitas Mahasiswa

- **Nama:** Raya Deswita
- **NIM:** 2024903430037
- **Kelas:** TRKJ 3C
- **Program Studi:** Teknologi Rekayasa Komputer Jaringan
- **Institusi:** Politeknik Negeri Lhokseumawe
