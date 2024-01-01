package main

import (
	"bufio"
	"os"

    "io"
)

func main() {
    // copy STDIN to STDOUT
    reader := io.Reader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()

    buffer := make([]byte, 1024 * 16)
    for {
        n, readErr := reader.Read(buffer)
        _, writeErr := writer.Write(buffer[:n])

        if readErr == io.EOF || readErr == io.ErrUnexpectedEOF {
            break
        } else if writeErr != nil {
            panic(writeErr)
        }
    }
}
