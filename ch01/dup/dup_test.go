package dup

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestDup(t *testing.T) {
	tests := []struct {
		name    string
		osInput []string
		want    string
	}{
		{
			"file without duplicate line",
			[]string{"cmd", "./text_freedom_poetry.txt"},
			"there are no duplicate lines!\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()
			os.Args = tt.osInput
			rescueStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			Dup()

			w.Close()
			out, _ := ioutil.ReadAll(r)
			fmt.Println(string(out))
			os.Stdout = rescueStdout
			fmt.Println(string(out))
			if string(out) != tt.want {
				t.Errorf("Expected %s, got %s", tt.want, out)
			}
		})
	}
}
