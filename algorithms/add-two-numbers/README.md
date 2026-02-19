# Add Two Numbers

## Apa itu Add Two Numbers?

Add Two Numbers adalah problem LeetCode #2. Diberikan dua **linked list** yang merepresentasikan bilangan integer (disimpan dalam urutan terbalik / *reverse order*), jumlahkan keduanya dan kembalikan hasilnya sebagai linked list baru.

## Contoh

```
Input:
  l1 = 2 -> 4 -> 3   (merepresentasikan angka 342)
  l2 = 5 -> 6 -> 4   (merepresentasikan angka 465)

Output:
  7 -> 0 -> 8         (merepresentasikan angka 807)

Alasan: 342 + 465 = 807
```

---

## Struktur Data

```go
type ListNode struct {
    Val  int
    Next *ListNode
}
```

---

## Solusi — Simulasi Penjumlahan dengan Carry

Iterasi kedua linked list sekaligus, jumlahkan digit per digit sambil track nilai **carry** (sisa bagi 10).

```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    curr := dummy
    carry := 0

    for l1 != nil || l2 != nil || carry != 0 {
        sum := carry

        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }

        carry = sum / 10
        curr.Next = &ListNode{Val: sum % 10}
        curr = curr.Next
    }

    return dummy.Next
}
```

**Cara kerja:**
```
l1: 2 -> 4 -> 3
l2: 5 -> 6 -> 4

Step 1: 2+5=7,      carry=0 → node(7)
Step 2: 4+6=10,     carry=1 → node(0)
Step 3: 3+4+1=8,    carry=0 → node(8)

Result: 7 -> 0 -> 8 ✅
```

---

## Kompleksitas

| | Nilai |
|---|---|
| Time | O(max(m, n)) |
| Space | O(max(m, n)) |

m dan n adalah panjang masing-masing linked list.

---

## Contoh Lengkap dengan main()

```go
package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    curr := dummy
    carry := 0

    for l1 != nil || l2 != nil || carry != 0 {
        sum := carry

        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }

        carry = sum / 10
        curr.Next = &ListNode{Val: sum % 10}
        curr = curr.Next
    }

    return dummy.Next
}

func makeList(nums []int) *ListNode {
    dummy := &ListNode{}
    curr := dummy
    for _, n := range nums {
        curr.Next = &ListNode{Val: n}
        curr = curr.Next
    }
    return dummy.Next
}

func printList(node *ListNode) {
    for node != nil {
        fmt.Print(node.Val)
        if node.Next != nil {
            fmt.Print(" -> ")
        }
        node = node.Next
    }
    fmt.Println()
}

func main() {
    l1 := makeList([]int{2, 4, 3}) // 342
    l2 := makeList([]int{5, 6, 4}) // 465

    result := addTwoNumbers(l1, l2)
    printList(result) // 7 -> 0 -> 8  (807)
}
```

---

## Edge Cases

| Kasus | Contoh Input | Output |
|-------|-------------|--------|
| Panjang berbeda | `[9,9]` + `[1]` | `[0,0,1]` |
| Carry di akhir | `[9,9,9]` + `[1]` | `[0,0,0,1]` |
| Salah satu kosong | `[0]` + `[7,3]` | `[7,3]` |

---

## Perbandingan dengan Two Sum

| | Two Sum | Add Two Numbers |
|---|---|---|
| Input | Array of int | Linked List |
| Output | Indeks | Linked List |
| Struktur data kunci | Hash Map | Dummy node + carry |
| Kesulitan | Easy | Medium |

---

## Key Takeaway

> Gunakan **dummy node** sebagai anchor hasil linked list.  
> Selalu handle **carry** bahkan setelah kedua list habis — bisa ada sisa digit di akhir.