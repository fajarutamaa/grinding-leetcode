# Sum Primes

## Apa itu Sum Primes?

Sum Primes adalah problem menghitung **jumlah semua bilangan prima** dari 2 hingga `limit` (inklusif). Bilangan prima adalah bilangan yang hanya habis dibagi 1 dan dirinya sendiri.

## Contoh

```
Input:  limit = 10
Output: 17

Alasan: 2 + 3 + 5 + 7 = 17
```

```
Input:  limit = 1
Output: 0

Alasan: Tidak ada bilangan prima ≤ 1
```

---

## Solusi — Trial Division

Untuk setiap angka `i` dari 2 hingga `limit`, cek apakah `i` adalah bilangan prima dengan mencoba pembagi dari `2` hingga `√i`.

```go
func SumPrimes(limit int) int {
    if limit < 2 {
        return 0
    }
    result := 0

    for i := 2; i <= limit; i++ {
        if isPrime(i) {
            result += i
        }
    }
    return result
}

func isPrime(num int) bool {
    if num < 2 {
        return false
    }

    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            return false
        }
    }
    return true
}
```

**Cara kerja:**
```
limit = 10

i=2:  prima ✅ → result = 2
i=3:  prima ✅ → result = 5
i=4:  4%2==0  ❌
i=5:  prima ✅ → result = 10
i=6:  6%2==0  ❌
i=7:  prima ✅ → result = 17
i=8:  8%2==0  ❌
i=9:  9%3==0  ❌
i=10: 10%2==0 ❌

Result: 17 ✅
```

---

## Kompleksitas

| | Nilai |
|---|---|
| Time | O(n√n) |
| Space | O(1) |

n adalah nilai `limit`.

---

## Edge Cases

| Kasus | Input | Output |
|-------|-------|--------|
| Limit di bawah 2 | `limit = 1` | `0` |
| Limit tepat 2 | `limit = 2` | `2` |
| Limit negatif | `limit = -5` | `0` |
| Limit prima | `limit = 7` | `17` |

---

## Optimasi — Sieve of Eratosthenes

Trial division cukup untuk `limit` kecil. Untuk `limit` besar, gunakan **Sieve of Eratosthenes**:

```go
func SumPrimesSieve(limit int) int {
    if limit < 2 {
        return 0
    }

    sieve := make([]bool, limit+1)
    for i := range sieve {
        sieve[i] = true
    }
    sieve[0], sieve[1] = false, false

    for i := 2; i*i <= limit; i++ {
        if sieve[i] {
            for j := i * i; j <= limit; j += i {
                sieve[j] = false
            }
        }
    }

    result := 0
    for i, isPrime := range sieve {
        if isPrime {
            result += i
        }
    }
    return result
}
```

| | Trial Division | Sieve of Eratosthenes |
|---|---|---|
| Time | O(n√n) | O(n log log n) |
| Space | O(1) | O(n) |
| Cocok untuk | limit kecil | limit besar |

---

## Key Takeaway

> Cukup cek pembagi hingga **√n** — jika `n` punya faktor lebih besar dari √n, pasti ada faktor pasangannya yang lebih kecil dari √n.  
> Untuk kasus performa tinggi, trade space dengan **Sieve** untuk time complexity yang jauh lebih baik.