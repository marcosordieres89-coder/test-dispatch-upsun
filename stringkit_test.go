package stringkit

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"single", "a", "a"},
		{"ascii", "hello", "olleh"},
		{"multibyte", "日本語", "語本日"},
		{"accents", "áé", "éá"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.in); got != tt.want {
				t.Errorf("Reverse(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"empty", "", true},
		{"single", "x", true},
		{"simple", "racecar", true},
		{"mixed case", "RaceCar", true},
		{"with punctuation", "A man, a plan, a canal: Panama", true},
		{"not a palindrome", "hello", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.in); got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestWordCount(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int
	}{
		{"empty", "", 0},
		{"whitespace only", "   \t\n ", 0},
		{"one word", "hello", 1},
		{"two words", "hello world", 2},
		{"extra spaces", "  hello   world  ", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordCount(tt.in); got != tt.want {
				t.Errorf("WordCount(%q) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}

func TestCapitalize(t *testing.T) {
	cases := map[string]string{"": "", "go": "Go", "Go": "Go", "élan": "Élan", "a b": "A b"}
	for in, want := range cases {
		if got := Capitalize(in); got != want {
			t.Errorf("Capitalize(%q) = %q, want %q", in, got, want)
		}
	}
}
