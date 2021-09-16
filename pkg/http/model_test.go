package http

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzbuzz_Key(t *testing.T) {
	var keyTest = []struct {
		in  Fizzbuzz
		out string
	}{
		{Fizzbuzz{1, 2, 3, "linux", "arch"}, "1,2,3,linux,arch"},
		{Fizzbuzz{43, 21, 1, "php", "header"}, "43,21,1,php,header"},
		{Fizzbuzz{13, 21, 12, "cobol", "erlang"}, "13,21,12,cobol,erlang"},
	}
	for i, tt := range keyTest {
		t.Run(fmt.Sprintf("Key %d", i), func(t *testing.T) {
			assert.Equal(t, tt.in.Key(), tt.out)
		})
	}
}

func TestFizzbuzz_IsValid(t *testing.T) {
	var fizzbuzzTest = []struct {
		in  Fizzbuzz
		out error
	}{
		{Fizzbuzz{1, 2, 3, "linux", "arch"}, nil},
		{Fizzbuzz{-42, 21, 1, "", ""}, errInteger},
		{Fizzbuzz{0, 21, 12, "cobol", "erlang"}, errInteger},
		{Fizzbuzz{0, 21, 12, "cobol", "erlang"}, errInteger},
		{Fizzbuzz{21, 21, 0, "cobol", "erlang"}, errInteger},
		{Fizzbuzz{21, 21, 12, "", "erlang"}, errString},
		{Fizzbuzz{21, 21, 12, "", ""}, errString},
		{Fizzbuzz{21, 21, 12, "a", "a"}, nil},
	}
	for i, tt := range fizzbuzzTest {
		t.Run(fmt.Sprintf("Key %d", i), func(t *testing.T) {
			assert.Equal(t, tt.in.IsValid(), tt.out)
		})
	}
}

func TestFizzbuzz_Compute(t *testing.T) {
	var fizzbuzzTest = []struct {
		in  Fizzbuzz
		out []string
	}{
		{
			Fizzbuzz{3, 5, 10, "a", "b"},
			[]string{
				"1", "2", "a", "4", "b", "a", "7", "8", "a", "b",
			},
		},
		{
			Fizzbuzz{3, 5, 15, "a", "b"},
			[]string{
				"1", "2", "a", "4", "b", "a", "7", "8", "a", "b", "11", "a", "13", "14", "ab",
			},
		},
	}
	for i, tt := range fizzbuzzTest {
		t.Run(fmt.Sprintf("Key %d", i), func(t *testing.T) {
			assert.Equal(t, tt.in.Compute(), tt.out)
		})
	}
}
