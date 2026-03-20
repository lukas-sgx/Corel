package agent

import (
	"testing"
	"os"
	"strings"
	"strconv"
)

func TestGenerateIdentity(t *testing.T) {
    result := generateIdentity()
    host, _ := os.Hostname()

    if !strings.HasPrefix(result, host+"-") {
        t.Errorf("Expected prefix %s-, got %s", host, result)
    }

    parts := strings.Split(result, "-")
    if len(parts) < 2 {
        t.Fatalf("Malformed identity string: %s", result)
    }
    
    _, err := strconv.ParseUint(parts[1], 10, 32)
    if err != nil {
        t.Errorf("Suffix is not a valid uint32: %s", parts[1])
    }
}