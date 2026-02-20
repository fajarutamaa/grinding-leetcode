# 🧠 Grinding LeetCode

Repository ini berisi kumpulan solusi algoritma yang dikerjakan sebagai latihan rutin untuk meningkatkan kemampuan problem-solving, pemahaman struktur data, dan persiapan technical interview.

Setiap problem disimpan di folder tersendiri lengkap dengan **solusi Go**, **unit test**, dan **README penjelasan** yang mencakup pendekatan, kompleksitas, dan key takeaway.

---

## 📁 Struktur Project

```
grinding-leetcode/
├── algorithms/
│   ├── two-sum/
│   │   ├── two_sum.go
│   │   ├── two_sum_test.go
│   │   └── README.md
│   ├── add-two-numbers/
│   │   ├── add_two_numbers.go
│   │   ├── add_two_numbers_test.go
│   │   └── README.md
│   └── sum-primes/
│       ├── sum_prime.go
│       ├── sum_prime_test.go
│       └── README.md
├── go.mod
└── README.md
```

---

## 🗺️ Navigasi Algoritma

| # | Problem | Difficulty | Topik | Link |
|---|---------|------------|-------|------|
| 1 | Two Sum | 🟢 Easy | Array, Hash Map | [→ Lihat](./algorithms/two-sum/README.md) |
| 2 | Add Two Numbers | 🟡 Medium | Linked List | [→ Lihat](./algorithms/add-two-numbers/README.md) |
| – | Sum Primes | 🟢 Easy | Math, Number Theory | [→ Lihat](./algorithms/sum-primes/README.md) |

---

## 📖 Ringkasan Problem

### 🔵 [Two Sum](./algorithms/two-sum/)

> LeetCode #1 · Easy · Array & Hash Map

Diberikan array integer dan sebuah `target`, kembalikan **indeks dua elemen** yang jika dijumlahkan menghasilkan `target`.

- **Pendekatan optimal:** Hash Map — O(n) time, O(n) space
- **Alternatif:** Brute Force O(n²) atau Two Pointers O(n) jika array sudah sorted

```
Input:  nums = [2, 7, 11, 15], target = 9
Output: [0, 1]  →  nums[0] + nums[1] = 2 + 7 = 9 ✅
```

---

### 🟡 [Add Two Numbers](./algorithms/add-two-numbers/)

> LeetCode #2 · Medium · Linked List

Diberikan dua linked list yang merepresentasikan angka dalam urutan terbalik, jumlahkan keduanya dan kembalikan hasilnya sebagai linked list baru.

- **Pendekatan:** Simulasi penjumlahan digit per digit dengan carry — O(max(m,n)) time & space
- **Konsep kunci:** dummy node sebagai anchor + tracking carry

```
Input:  l1 = 2→4→3  (342),  l2 = 5→6→4  (465)
Output: 7→0→8  (807)  →  342 + 465 = 807 ✅
```

---

### 🔵 [Sum Primes](./algorithms/sum-primes/)

> Math · Easy · Number Theory

Hitung **jumlah semua bilangan prima** dari 2 hingga `limit` (inklusif).

- **Pendekatan dasar:** Trial Division — O(n√n) time, O(1) space
- **Optimasi:** Sieve of Eratosthenes — O(n log log n) time, O(n) space (untuk `limit` besar)

```
Input:  limit = 10
Output: 17  →  2 + 3 + 5 + 7 = 17 ✅
```

---

## ⚙️ Cara Menjalankan

### Prasyarat

- [Go](https://golang.org/dl/) versi 1.21+

### Menjalankan semua test

```bash
go test ./...
```

### Menjalankan test untuk satu problem

```bash
# Contoh: two-sum
go test ./algorithms/two-sum/

# Dengan output verbose
go test -v ./algorithms/two-sum/
```

### Menjalankan test dengan coverage

```bash
go test -cover ./...
```

---

## 🗂️ Kategori & Topik

| Topik | Problem |
|-------|---------|
| **Array** | [Two Sum](./algorithms/two-sum/) |
| **Hash Map** | [Two Sum](./algorithms/two-sum/) |
| **Linked List** | [Add Two Numbers](./algorithms/add-two-numbers/) |
| **Math / Number Theory** | [Sum Primes](./algorithms/sum-primes/) |

---

## 📐 Konvensi Penamaan

| Item | Format | Contoh |
|------|--------|--------|
| Folder problem | kebab-case | `two-sum/` |
| File Go | snake_case | `two_sum.go` |
| File test | snake_case + `_test` | `two_sum_test.go` |
| Package Go | snake_case | `package twosum` |

---

## 💡 Cara Berkontribusi / Menambahkan Problem Baru

1. Buat folder baru di `algorithms/` dengan format kebab-case, contoh: `algorithms/longest-substring/`
2. Buat tiga file di dalamnya:
   - `solution.go` — implementasi solusi
   - `solution_test.go` — unit test mencakup test case normal, edge case, dan case besar
   - `README.md` — penjelasan problem, pendekatan, kompleksitas, dan key takeaway
3. Tambahkan entry baru di tabel **Navigasi Algoritma** pada README ini

---

## 🎯 Tujuan Repository

- ✅ Membangun kebiasaan problem-solving yang konsisten
- ✅ Memahami trade-off antara time complexity dan space complexity
- ✅ Menguasai pola algoritma umum (sliding window, two pointers, BFS/DFS, DP, dll.)
- ✅ Menulis kode Go yang bersih, testable, dan well-documented
- ✅ Persiapan technical interview

---

> *"The only way to get better at algorithms is to solve them, one problem at a time."*
