package util

import (
	"bufio"
	"bytes"
	"net"
)

// DropCR drops a terminal \r from the data.
func DropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

// ScanCRLF is a split function for the bufio scanner
func scanCRLF(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte{'\r', '\n'}); i >= 0 {
		// We have a full newline-terminated line.
		return i + 2, DropCR(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), DropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

// Scan reads and parses the request
func Scan(c net.Conn) (request string, err error) {
	scanner := bufio.NewScanner(c)
	scanner.Split(scanCRLF)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	if len(lines) > 0 {
		return lines[0], nil
	}
	return "", nil
}
