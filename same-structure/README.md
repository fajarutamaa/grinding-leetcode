# Same Structure

## Apa itu Same Structure?

Same Structure adalah problem mengecek apakah dua string memiliki **komposisi karakter yang sama** — yaitu karakter yang sama dengan jumlah kemunculan yang sama, tanpa memperhatikan urutan.

## Contoh
```
Input:  s1 = "abc", s2 = "bca"
Output: true

Alasan: Keduanya punya a×1, b×1, c×1
```
```
Input:  s1 = "aab", s2 = "bba"
Output: false

Alasan: s1 punya a×2, b×1 — s2 punya b×2, a×1
```

---

## Solusi — Frequency Counter

Hitung frekuensi tiap karakter di `s1` menggunakan map, lalu kurangi frekuensi untuk setiap karakter di `s2`. Jika ada yang turun di bawah 0, strukturnya berbeda.
```go
func SameStructure(s1, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }
    counter := make(map[rune]int)

    for _, ch := range s1 {
        counter[ch]++
    }

    for _, ch := range s2 {
        counter[ch]--
        if counter[ch] < 0 {
            return false
        }
    }
    return true
}
```

**Cara kerja:**
```
s1 = "abc", s2 = "bca"

Step 1 — hitung s1:
  counter = {a:1, b:1, c:1}

Step 2 — kurangi dari s2:
  'b' → counter[b] = 0 ✅
  'c' → counter[c] = 0 ✅
  'a' → counter[a] = 0 ✅

Result: true ✅
```
```
s1 = "aab", s2 = "bba"

Step 1 — hitung s1:
  counter = {a:2, b:1}

Step 2 — kurangi dari s2:
  'b' → counter[b] = 0 ✅
  'b' → counter[b] = -1 ❌ → return false

Result: false ✅
```

---

## Kompleksitas

| | Nilai |
|---|---|
| Time | O(n) |
| Space | O(k) |

n adalah panjang string, k adalah jumlah karakter unik.

---

## Edge Cases

| Kasus | Input | Output |
|-------|-------|--------|
| String kosong | `s1 = ""`, `s2 = ""` | `true` |
| Panjang berbeda | `s1 = "ab"`, `s2 = "abc"` | `false` |
| Karakter sama semua | `s1 = "aaa"`, `s2 = "aaa"` | `true` |
| Anagram sempurna | `s1 = "listen"`, `s2 = "silent"` | `true` |
| Karakter berbeda | `s1 = "abc"`, `s2 = "xyz"` | `false` |

---

## Key Takeaway

> Early return dengan cek panjang string sebelum iterasi — dua string dengan panjang berbeda **tidak mungkin** memiliki struktur yang sama.  
> Satu map sudah cukup: increment untuk `s1`, decrement untuk `s2`. Jika ada nilai negatif, langsung return false tanpa perlu iterasi penuh.