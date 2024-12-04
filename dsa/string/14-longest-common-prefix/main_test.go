package main

import "testing"

func TestLongstCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "test1",
			args: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "test2",
			args: []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			name: "test3",
			args: []string{"ab", "a"},
			want: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
