package render

import "testing"

func TestFormatAmount(t *testing.T) {
	tests := []struct {
		name   string
		amount int
		want   string
	}{
		{name: "Whole pounds", amount: 1200, want: "£12.00"},
		{name: "Single digit pence", amount: 1205, want: "£12.05"},
		{name: "Zero", amount: 0, want: "£0.00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatAmount(tt.amount); got != tt.want {
				t.Errorf("FormatAmount(%d) = %q, want %q", tt.amount, got, tt.want)
			}
		})
	}
}
