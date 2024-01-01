package main

import (
	"bufio"
	"os"
)

func main() {
    // copy STDIN to STDOUT
    reader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()

    scanner := bufio.NewScanner(reader)

    for {
        ok := scanner.Scan()

        if !ok {
            break
        }

        slice := scanner.Bytes()

        writer.Write(slice)
    }

    err := scanner.Err()
    if err != nil {
        panic(err)
    }
}
