package other

import (
	"testing"
)

func TestIsIPv4(t *testing.T) {
	tests := []struct {
		name string
		ip   string
		want bool
	}{
		{
			name: "Valid IPv4 address",
			ip:   "192.168.0.1",
			want: true,
		},
		{
			name: "Invalid IPv4 address",
			ip:   "256.60.124.1",
			want: false,
		},
		{
			name: "Invalid IPv4 address with leading zeros",
			ip:   "192.168.001.001",
			want: false,
		},
		{
			name: "Invalid IPv4 address with extra dots",
			ip:   "192.168..1",
			want: false,
		},
		{
			name: "Empty string",
			ip:   "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIPv4(tt.ip); got != tt.want {
				t.Errorf("isIPv4() = %v, want %v", got, tt.want)
			}
		})
	}
}
