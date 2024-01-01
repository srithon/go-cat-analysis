package main

import (
    "os"
    "io"
)

func main() {
    // copy STDIN to STDOUT
    reader := io.Reader(os.Stdin)
    writer := io.Writer(os.Stdout)

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
