package imgconv

import (
	"testing"
)

type testProperty struct {
	args args
	want string
}

type args struct {
	testPath   string
	bExtension string
	aExtension string
}

func TestConverter(t *testing.T) {
	tests := map[string]testProperty{
		"normal":                         {args: args{"../testdata/case_0", ".jpg", ".png"}, want: ""},
		"normal2":                        {args: args{"../testdata/case_6", ".jpg", ".png"}, want: ""},
		"permission error (directory)":   {args: args{"../testdata/case_1", ".jpg", ".png"}, want: "walking the path open ../testdata/case_1: permission denied"},
		"permission error (input file)":  {args: args{"../testdata/case_2", ".jpg", ".png"}, want: "open ../testdata/case_2/normal_test.jpg: permission denied"},
		"permission error (output file)": {args: args{"../testdata/case_4", ".jpg", ".png"}, want: "open ../testdata/case_4/normal_test.png: permission denied"},
		"directory not found (file)":     {args: args{"../testdata/case_X", ".jpg", ".png"}, want: "walking the path lstat ../testdata/case_X: no such file or directory"},
		"decode failed":                  {args: args{"../testdata/case_5", ".jpg", ".png"}, want: "../testdata/case_5/coverage.jpg is not a valid file"},
	}
	// err := createTestEnv()
	ic := &ImageConverter{}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := ic.ConvertImg(tt.args.testPath, tt.args.bExtension, tt.args.aExtension)
			if got != nil && got.Error() != tt.want {
				t.Errorf("ConvertImg() = %v, want %v", got.Error(), tt.want)
			} else if got == nil && tt.want != "" {
				t.Errorf("ConvertImg() = nil, want %v", tt.want)
			}
		})
	}
}
