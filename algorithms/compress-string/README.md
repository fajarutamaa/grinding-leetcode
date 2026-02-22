# Compress String

## Apa itu Compress String?

Compress String adalah problem mengompresi string dengan mengganti karakter yang berulang secara berurutan menjadi **karakter + jumlah kemunculannya**. Jika hasil kompresi tidak lebih pendek dari string asli, kembalikan string asli.

## Contoh
```
Input:  "aabcccccaaa"
Output: "a2b1c5a3"

Alasan: a×2, b×1, c×5, a×3 → lebih pendek dari aslinya
```
```
Input:  "abc"
Output: "abc"

Alasan: "a1b1c1" (6 karakter) >= "abc" (3 karakter) → kembalikan aslinya
```

---

## Solusi — Run-Length Encoding

Iterasi string sambil menghitung karakter yang sama secara berurutan. Setiap kali karakter berganti, tambahkan karakter sebelumnya beserta hitungannya ke hasil kompresi.
```go
func CompressString(str string) string {
    if len(str) == 0 {
        return str
    }

    compressed := ""
    count := 1

    for i := 1; i < len(str); i++ {
        if str[i] == str[i-1] {
            count++
        } else {
            compressed += string(str[i-1]) + strconv.Itoa(count)
            count = 1
        }
    }

    compressed += string(str[len(str)-1]) + strconv.Itoa(count)

    if len(compressed) >= len(str) {
        return str
    }

    return compressed
}
```

**Cara kerja:**
```
str = "aabcccccaaa"

Iterasi:
  i=1: str[1]='a' == str[0]='a' → count=2
  i=2: str[2]='b' != str[1]='a' → compressed="a2", count=1
  i=3: str[3]='c' != str[2]='b' → compressed="a2b1", count=1
  i=4: str[4]='c' == str[3]='c' → count=2
  i=5: str[5]='c' == str[4]='c' → count=3
  i=6: str[6]='c' == str[5]='c' → count=4
  i=7: str[7]='c' == str[6]='c' → count=5
  i=8: str[8]='a' != str[7]='c' → compressed="a2b1c5", count=1
  i=9: str[9]='a' == str[8]='a' → count=2
  i=10: str[10]='a' == str[9]='a' → count=3

Tambah karakter terakhir:
  compressed = "a2b1c5a3"

len("a2b1c5a3")=8 < len("aabcccccaaa")=11 → return "a2b1c5a3" ✅
```
```
str = "abc"

Iterasi:
  i=1: 'b' != 'a' → compressed="a1", count=1
  i=2: 'c' != 'b' → compressed="a1b1", count=1

Tambah karakter terakhir:
  compressed = "a1b1c1"

len("a1b1c1")=6 >= len("abc")=3 → return "abc" ✅
```

---

## Kompleksitas

| | Nilai |
|---|---|
| Time | O(n) |
| Space | O(n) |

n adalah panjang string.

---

## Edge Cases

| Kasus | Input | Output |
|-------|-------|--------|
| String kosong | `""` | `""` |
| Semua karakter sama | `"aaaa"` | `"a4"` |
| Semua karakter unik | `"abcd"` | `"abcd"` |
| Kompresi lebih panjang | `"abc"` | `"abc"` |
| Satu karakter | `"a"` | `"a"` |
| Kompresi lebih pendek | `"aabcccccaaa"` | `"a2b1c5a3"` |

---

## Key Takeaway

> Karakter terakhir selalu ditambahkan **setelah loop** karena loop hanya menulis karakter saat terjadi pergantian — karakter terakhir tidak pernah "diganti".  
> Bandingkan panjang hasil kompresi dengan string asli sebelum return — kompresi hanya berguna jika benar-benar mempersingkat string.