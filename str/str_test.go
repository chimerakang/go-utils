package str

import (
	"reflect"
	"testing"
)

func TestStr2UintArr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		wantIds []uint
	}{
		{
			name: "case1",
			args: args{str: "1,2,3,4,5"},
			wantIds: []uint{
				1, 2, 3, 4, 5,
			},
		},
		{
			name: "case2",
			args: args{str: "-1,2,-3,4,5"},
			wantIds: []uint{
				0, 2, 0, 4, 5,
			},
		},
		{
			name: "case3",
			args: args{str: "1"},
			wantIds: []uint{
				1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIds := Str2UintArr(tt.args.str); !reflect.DeepEqual(gotIds, tt.wantIds) {
				t.Errorf("Str2UintArr() = %v, want %v", gotIds, tt.wantIds)
			}
		})
	}
}

func TestCamelCaseLowerFirst(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{str: ""},
			want: "",
		},
		{
			name: "case2",
			args: args{str: "a_bc_def"},
			want: "aBcDef",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelCaseLowerFirst(tt.args.str); got != tt.want {
				t.Errorf("CamelCaseFirstLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeStr2Base64(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{str: "ok"},
			want: "b2s=",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeStr2Base64(tt.args.str); got != tt.want {
				t.Errorf("EncodeStr2Base64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeStrFromBase64(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{str: "b2s="},
			want: "ok",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeStrFromBase64(tt.args.str); got != tt.want {
				t.Errorf("DecodeStrFromBase64() = %v, want %v", got, tt.want)
			}
		})
	}
}
