// dup1 prints the text of each line that appears more than
// once in the standard input, preceeded by its count
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        line := input.Text()
        if len(line) == 0 {
            break
        }
        counts[line]++
    }
    // Note: ignoring potential errors from input Err()
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
