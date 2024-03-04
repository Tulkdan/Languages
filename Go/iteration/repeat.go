package iteration

func Repeat(char string, numberTimes int) string {
    var repeated string

    for i := 0; i < numberTimes; i++ {
        repeated += char
    }

    return repeated
}
