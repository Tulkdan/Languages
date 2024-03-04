package iteration

const repeatedCount = 5

func Repeat(char string) string {
    var repeated string

    for i := 0; i < repeatedCount; i++ {
        repeated += char
    }

    return repeated
}
