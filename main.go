package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
    // copy STDIN to STDOUT
    reader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()

    for {
        byte, err := reader.ReadByte()
        if err == io.EOF {
            break
        } else if err != nil {
            panic(err)
        }

        writer.WriteByte(byte)
    }
}
