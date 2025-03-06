// Problem
/*
Solve the problem below
*/

package main

import "fmt"

func main() {
    num := 6
    unsolved := true
    big := 1
    for unsolved {
        num++
        fmt.Printf("Current num: %d -- ", num)
        for i := 2; i <= 6; i++ {
            var mod int
            mod = num % i
            if mod != 1 {
                fmt.Printf("%d failed, mod = %d \n", i, mod)
                break
            } else {
                big = mod
            }
        if (big == 6 && (num % 7) == 0) {
            fmt.Printf("The answer is %d:", num)
            unsolved = false
        }
        }
    }
}

// func main() {
//     num := 7
//     unsolved := true
//     big := 1
//     for unsolved {
//         num += 7
//         fmt.Printf("Current num: %d -- ", num)
//         for i := 2; i <= 6; i++ {
//             var mod int
//             mod = num % i
//             if mod != 1 {
//                 fmt.Printf("%d failed, mod = %d \n", i, mod)
//                 break
//             } else {
//                 big = mod
//             }
//         if (big == 6) {
//             fmt.Printf("The answer is %d:", num)
//             unsolved = false
//         }
//         }
//     }
// }

// func main() {
//     num := 1
//     unsolved := true
//     for unsolved {
//         num += 60
//         fmt.Printf("Current num: %d -- ", num)
//         mod := (num % 7)
//         if mod == 0 {
//             fmt.Printf("The answer is %d:", num)
//             unsolved = false
//         } else {
//             fmt.Printf("failed with mod : %d \n", mod)
//         }
//     }
// }