package dump

import (
	"strings"
	"testing"
)

func TestUnicodeFileNames(t *testing.T) {
	f := &DumpFormatter{}
	used := make(map[string]bool)
	for _, name := range []string{"客戶", "訂單", "顧客", "注文", "café", "caf", "Customer", "customer", "_users", "users", "order detail", "order_detail", "CON", "中文", "!!!", "???", strings.Repeat("中文", 100)} {
		file := f.uniqueFileName("tables", name, used)
		if len(file) > 255 || file == ".sql" {
			t.Fatalf("invalid filename %q", file)
		}
	}
	if len(used) != 17 {
		t.Fatalf("expected 17 unique paths, got %d", len(used))
	}
	if f.sanitizeFileName("客戶") != "客戶" || f.sanitizeFileName("訂單") != "訂單" {
		t.Fatal("Chinese identifiers were not preserved")
	}
	if f.sanitizeFileName("CON") != "_con" {
		t.Fatal("reserved Windows name was not escaped")
	}
}
