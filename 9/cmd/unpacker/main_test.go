package main

import (
	"errors"
	"testing"
)

func TestUnpackString(t *testing.T) {
	var testCases = []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{"Single letter", "a", "a", nil},
		{"Multiple letters", "abcd", "abcd", nil},
		{"Single cyrillic letter", "э", "э", nil},
		{"Multiple cyrillic letters", "азбука", "азбука", nil},
		{"Latin and cyrillic letters", "alphabet азбука", "alphabet азбука", nil},
		{"Letters with space", "a b c d ", "a b c d ", nil},
		{"Empty string", "", "", nil},
		{"Single letter with 1", "a1", "a", nil},
		{"Multiple letters with multiple numbers", "a4bc2d5e", "aaaabccddddde", nil},
		{"Single letter with big number", "a10", "aaaaaaaaaa", nil},
		{"Multiple letters with multiple big numbers", "a123b1234", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", nil},
		{"Single cyrillic letter with digit", "ф1", "ф", nil},
		{"Multiple cyrillic letters with digit", "дф1", "дф", nil},
		{"Single cyrillic letter with big number", "ф12", "фффффффффффф", nil},
		{"Only digits", "45", "", ErrOnlyDigits},
		{"Single escape-sequence in the end", "abc\\1", "abc1", nil},
		{"Single escape-sequence in the beginning", "\\1abc", "1abc", nil},
		{"Single escape-sequence in the middle", "ab\\1cd", "ab1cd", nil},
		{"Multiple escape-sequences", "\\1a\\2b\\3c\\4d\\5", "1a2b3c4d5", nil},
		{"Multiple escape-sequences in a row", "a\\1\\2\\3", "a123", nil},
		{"Single escape-sequence with a single digit", "a\\12", "a11", nil},
		{"Single escape-sequence with a big digit", "a\\123", "a11111111111111111111111", nil},
		{"Single letter with zero", "a0", "", nil},
		{"Multiple letters with zero", "abc0", "ab", nil},
		{"Only zero", "0", "", ErrOnlyDigits},
		{"Single escape-sequence with zero", "a\\0", "a0", nil},
		{"Single escape-sequence with digit zero", "a\\10", "a", nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := unpackString(tc.input)
			if res != tc.output {
				t.Errorf("got '%s', want '%s'", res, tc.output)
			}

			if !errors.Is(err, tc.err) {
				t.Errorf("unexpected error: got %v, want %v", err, tc.err)
			}
		})
	}
}

func TestGetFullNumber(t *testing.T) {
	var testCases = []struct {
		name   string
		input  string
		output int
		err    error
	}{
		{"Single digit", "1", 1, nil},
		{"Big digit", "1234", 1234, nil},
		{"Single digit with single letter", "1a", 1, nil},
		{"Single digit with multiple letters", "1abcd", 1, nil},
		{"Big digit with single letter", "1234a", 1234, nil},
		{"Big digit with multiple letters", "1234abcd", 1234, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := getFullNumber([]rune(tc.input))
			if res != tc.output {
				t.Errorf("got '%d', want '%d'", res, tc.output)
			}

			if !errors.Is(err, tc.err) {
				t.Errorf("unexpected error: got %v, want %v", err, tc.err)
			}
		})
	}
}
