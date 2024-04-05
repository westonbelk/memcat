package util

import (
    "os"
    "bufio"
)

func ReadLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if scanner.Err() != nil {
        return nil, err
    }
    return lines, err
}
