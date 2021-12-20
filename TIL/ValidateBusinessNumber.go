package main

    import (
        "errors"
        "fmt"
    )

    func ValidateBusinessNumber(x string) bool {
        if len(x) == 10 {
            key := "137137135"
            sum := 0
            for i:=0; i <= 8; i++ {
                sum += int(x[i]-48) * int(key[i]-48)
            }
            num := 10 - (sum + (int(x[8]-48) * int(key[8]-48))/10)%10
            if int(x[9]-48) == num {
                fmt.Println("올바른 사업자번호 입니다.")
                return true
            }
            fmt.Println("잘못된 사업자번호입니다.")
            return false
        }
        fmt.Println("잘못된 사업자번호입니다.")
        return false
    }
