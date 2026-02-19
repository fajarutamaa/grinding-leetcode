# Two Sum

## Apa itu Two Sum?

Two Sum adalah salah satu problem algoritma paling populer (LeetCode #1). Diberikan array integer dan sebuah target, cari **dua angka** dalam array yang jika dijumlahkan menghasilkan nilai target tersebut. Kembalikan indeks kedua angka itu.

## Contoh

```
Input:  nums = [2, 7, 11, 15], target = 9
Output: [0, 1]
Alasan: nums[0] + nums[1] = 2 + 7 = 9
```

---

## Pendekatan Solusi

### 1. Brute Force — O(n²)

Bandingkan setiap pasangan elemen.

```go
func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i]+nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return nil
}
```

**Kekurangan:** Lambat untuk array besar karena nested loop.

---

### 2. Hash Map — O(n) ✅ Optimal

Simpan setiap elemen di hash map, cari komplemen (`target - nums[i]`) saat iterasi.

```go
func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
    for i, num := range nums {
        complement := target - num
        if j, ok := seen[complement]; ok {
            return []int{j, i}
        }
        seen[num] = i
    }
    return nil
}
```

**Cara kerja:**
```
nums = [2, 7, 11, 15], target = 9

i=0: num=2, complement=7 → tidak ada di seen → seen={2:0}
i=1: num=7, complement=2 → ada di seen[2]=0 → return [0, 1] ✅
```

---

## Perbandingan Kompleksitas

| Pendekatan | Time | Space |
|------------|------|-------|
| Brute Force | O(n²) | O(1) |
| Hash Map | O(n) | O(n) |

---

## Variasi Problem

- **Two Sum II** — array sudah sorted, gunakan two pointers O(n) O(1)
- **Two Sum III** — design data structure dengan `add` dan `find`
- **3Sum / 4Sum** — ekstensi ke 3 atau 4 angka

---

## Two Pointers (untuk sorted array)

```go
func twoSumSorted(nums []int, target int) []int {
    left, right := 0, len(nums)-1
    for left < right {
        total := nums[left] + nums[right]
        switch {
        case total == target:
            return []int{left, right}
        case total < target:
            left++
        default:
            right--
        }
    }
    return nil
}
```

---

## Contoh Lengkap dengan main()

```go
package main

import "fmt"

func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
    for i, num := range nums {
        if j, ok := seen[target-num]; ok {
            return []int{j, i}
        }
        seen[num] = i
    }
    return nil
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    fmt.Println(twoSum(nums, target)) // [0 1]
}
```

---

## Key Takeaway

> Gunakan **Hash Map** untuk array tidak terurut → O(n).  
> Gunakan **Two Pointers** jika array sudah sorted → O(n) dengan O(1) space.