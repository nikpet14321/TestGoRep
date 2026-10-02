package main

import (
	"testing"
	"errors"
)

func TestGetUTFLength(t *testing.T) {
	cases := []struct {
		name string
		inp []byte
		wantlen int
		wanter error
	}{
		{
			name: "zero length",
			inp: []byte(""),
			wantlen: 0,
			wanter: nil,
		},
		{
			name: "Hello! length",
			inp: []byte("Hello!"),
			wantlen: 6,
			wanter: nil,
		},
		{
			name: "long length",
			inp: []byte("очень много слов, прям очень много, ну вот десяток точно будет"),
			wantlen: 62,
			wanter: nil,
		},
		{
			name: "one length",
			inp: []byte("Я"),
			wantlen: 1,
			wanter: nil,
		},
		{
			name: "error length",
			inp: []byte{0xff, 0xfe, 0xfd},
			wantlen: 0,
			wanter: ErrInvalidUTF8,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			gotlen, goter := GetUTFLength(tc.inp)
			if !errors.Is(goter, tc.wanter) {
				t.Errorf("goter %v; wanter %v", goter, tc.wanter)
			}
			if gotlen != tc.wantlen {
				t.Errorf("gotlen %v; wantlen %v", gotlen, tc.wantlen)
			}
		})
	}
}