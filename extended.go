package main

import "os"
import "bufio"
import "fmt"
import "regexp"

func main() {
    r := bufio.NewReader(os.Stdin)
    e, er := regexp.Compile("died")

    if er != nil {
        os.Exit(0)
    }

    var d = make(map[string]string)
    d["wat"] = "datnil"
    d["who"] = "alsonil"

    for k, v := range d {
        fmt.Println("k:", k, "v:", v)
    }

    for {
        line, err := r.ReadString('\n')

        if err != nil {
            break
        }

        if e.MatchString(line) {
            fmt.Print(line)
        }
    }
}
