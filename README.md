import "github.com/gkyh/string"

func main() {
    fmt.Println(stringx.Parse(123))
    fmt.Println(stringx.Parse(123.23))

    fmt.Println(stringx.Int("123"))
    fmt.Println(stringx.Float64("123.23"))
}
