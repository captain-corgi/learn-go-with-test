package reflection

import (
	"reflect"
	"testing"
)

type (
	Person struct {
		Name    string
		Profile Profile
	}
	Profile struct {
		Age  int
		City string
	}
)

func TestWalk(t *testing.T) {
	assertContains := func(t *testing.T, got string, want []string) {
		t.Helper()
		contains := false
		for _, x := range want {
			if x == got {
				contains = true
			}
		}
		if !contains {
			t.Errorf("expected %+v to contain %q but it didn't", want, got)
		}
	}
	aChannel := make(chan Profile)
	go func() {
		aChannel <- Profile{18, "AnhChan18"}
		aChannel <- Profile{22, "AnhChan22"}
		close(aChannel)
	}()

	type args struct {
		x interface{}
	}
	tests := []struct {
		name  string
		args  args
		isMap bool
		want  []string
	}{
		{
			name: "1. TDD",
			args: args{
				x: struct{ Name string }{"Anh"},
			},
			want: []string{"Anh"},
		},
		{
			name: "2. Struct with two string fields",
			args: args{
				x: struct {
					Name string
					City string
				}{"Anh", "Can Tho"},
			},
			want: []string{"Anh", "Can Tho"},
		},
		{
			name: "3. Struct with none string field",
			args: args{
				x: struct {
					Name string
					Age  int
				}{"Anh", 1},
			},
			want: []string{"Anh"},
		},
		{
			name: "4. Struct with nested fields",
			args: args{
				x: Person{"Anh", Profile{18, "Can Tho"}},
			},
			want: []string{"Anh", "Can Tho"},
		},
		{
			name: "5. Input is a pointer",
			args: args{
				x: &Person{"Anh", Profile{18, "Can Tho"}},
			},
			want: []string{"Anh", "Can Tho"},
		},
		{
			name: "6. Input is a slice",
			args: args{
				x: []Person{
					{"Anh18", Profile{18, "Can Tho"}},
					{"Anh22", Profile{22, "TP HCM"}},
				},
			},
			want: []string{"Anh18", "Can Tho", "Anh22", "TP HCM"},
		},
		{
			name: "7. Input is a string",
			args: args{
				x: "Anh",
			},
			want: []string{"Anh"},
		},
		{
			name: "8. Input is a slice",
			args: args{
				x: []Person{
					{"Anh18", Profile{18, "Can Tho"}},
					{"Anh22", Profile{22, "TP HCM"}},
				},
			},
			want: []string{"Anh18", "Can Tho", "Anh22", "TP HCM"},
		},
		{
			name: "9. Input is an array",
			args: args{
				x: [2]Person{
					{"Anh18", Profile{18, "Can Tho"}},
					{"Anh22", Profile{22, "TP HCM"}},
				},
			},
			want: []string{"Anh18", "Can Tho", "Anh22", "TP HCM"},
		},
		{
			name: "10. Input is a map",
			args: args{
				x: map[string]string{
					"Anh18": "Can Tho",
					"Anh22": "TP HCM",
				},
			},
			isMap: true,
			want:  []string{"Can Tho", "TP HCM"},
		},
		{
			name: "11. Input is an array",
			args: args{
				x: [2]Person{
					{"Anh18", Profile{18, "Can Tho"}},
					{"Anh22", Profile{22, "TP HCM"}},
				},
			},
			isMap: true,
			want:  []string{"Anh18", "Can Tho", "Anh22", "TP HCM"},
		},
		{
			name: "12. Input is channel",
			args: args{
				x: aChannel,
			},
			want: []string{"AnhChan18", "AnhChan22"},
		},
		{
			name: "13. Input is function",
			args: args{
				x: func() (p1 Profile, p2 Profile) {
					p1 = Profile{18, "AnhChan18"}
					p2 = Profile{22, "AnhChan22"}
					return
				},
			},
			want: []string{"AnhChan18", "AnhChan22"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string
			Walk(tt.args.x, func(s string) {
				got = append(got, s)
			})
			if tt.isMap {
				for _, g := range got {
					assertContains(t, g, tt.want)
				}
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			}
		})
	}
}
