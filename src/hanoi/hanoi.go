package hanoi

import "fmt"

func Hanoi(A, B, C byte, n int) {
    // A, B, C for stack 'A', 'B' and 'C'
    // n for the number of plates.
    // complexity: 2^n -1 = O(2^n)
    if (n==1) {
        fmt.Printf("Move disk %d from %q to %q\n", n, A, C)
    } else {
        Hanoi(A, C, B, n-1)
        fmt.Printf("Move disk %d from %q to %q\n", n, A, C)
        Hanoi(B, A, C, n-1)
    }
}
