package necromancy

import (
	"reflect"
	"testing"
)

type Entity struct {
	Id   int
	Name string
}

func TestExtraction(t *testing.T) {
	type args struct {
		entity any
		name   string
	}
	tests := []struct {
		name    string
		args    args
		wantVal any
		wantErr string
	}{
		{"t1", args{Entity{1, "a"}, "Id"}, 1, ""},
		{"t2", args{Entity{2, "b"}, "Name"}, "b", ""},
		{"t3", args{Entity{3, "c"}, "Type"}, "c", "reflect:未找到指定的字段"},
		{"t4", args{4, "Id"}, "d", "reflect:元素非结构体"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVal, err := Extraction(tt.args.entity, tt.args.name)
			if (err != nil) != (tt.wantErr != "") {
				t.Errorf("getFieldValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				if err.Error() != tt.wantErr {
					t.Errorf("getFieldValue() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("Extraction() = %v, want %v", gotVal, tt.wantVal)
			}
		})
	}
}

func BenchmarkExtraction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Extraction(Entity{1, "a"}, "Id")
	}
}
