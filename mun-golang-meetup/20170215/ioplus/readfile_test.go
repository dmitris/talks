package ioplus

import (
	"context"
	"testing"
	"time"
)

func TestReadFileWithoutTimeout(t *testing.T) {
	t.Parallel()
	filename := "readfile_test.go"
	_, err := ReadFile(context.Background(), filename) // HL
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadFileWithTimeout(t *testing.T) {
	t.Parallel()
	filename := "readfile.go"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Microsecond) // HL
	defer cancel()
	data, err := ReadFile(ctx, filename)
	if err != context.DeadlineExceeded { // HL
		t.Fatalf("want DeadlineExceeded, got %v", err)
	}
	if len(data) != 0 {
		t.Errorf("want empty data slice on timeout, got %v", data)
	}
}
