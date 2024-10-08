package iteration

import "strings"

func Repeat(char string, numberTimes int) string {
    trimmed := strings.Trim(char, " ")
    firstChar := strings.Split(trimmed, "")

    var repeated string

    for i := 0; i < numberTimes; i++ {
        repeated += firstChar[0]
    }

    return repeated
}
