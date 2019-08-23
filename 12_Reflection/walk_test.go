package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"should pass a struct with one string field",
			struct {
				Name string
			}{"Andre"},
			[]string{"Andre"},
		},
		{
			"should pass a struct with two string fields",
			struct {
				Name string
				City string
			}{"Andre", "Botucatu"},
			[]string{"Andre", "Botucatu"},
		},
		{
			"should pass a struct with a string and an int field",
			struct {
				Name string
				Age  int
			}{"Andre", 25},
			[]string{"Andre"},
		},
		{
			"should pass a nested struct",
			Person{
				"Andre",
				Profile{25, "Botucatu"},
			},
			[]string{"Andre", "Botucatu"},
		},
		{
			"should pass a pointer",
			&Person{
				"Andre",
				Profile{25, "Botucatu"},
			},
			[]string{"Andre", "Botucatu"},
		},
		{
			"should pass a slice",
			[]Profile{
				{25, "Botucatu"},
				{21, "Botucatu"},
			},
			[]string{"Botucatu", "Botucatu"},
		},
		{
			"should pass an array",
			[2]Profile{
				{25, "Botucatu"},
				{21, "Botucatu"},
			},
			[]string{"Botucatu", "Botucatu"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var received []string

			walk(test.Input, func(input string) {
				received = append(received, input)
			})

			if !reflect.DeepEqual(received, test.ExpectedCalls) {
				t.Errorf("❌ expected %v received %v", test.ExpectedCalls, received)
			}
		})
	}

	t.Run("should pass a map", func(t *testing.T) {
		testMap := map[string]string{
			"name": "Andre",
			"city": "Botucatu",
		}

		var received []string
		walk(testMap, func(input string) {
			received = append(received, input)
		})

		assertContains(t, received, "Andre")
		assertContains(t, received, "Botucatu")
	})
}

func assertContains(t *testing.T, received []string, value string) {
	contains := false

	for _, x := range received {
		if x == value {
			contains = true
			break
		}
	}

	if !contains {
		t.Errorf("❌ expected %+v to contain %q but it wasn't found", received, value)
	}
}

// func assertInt(t *testing.T, received, expected int) {
// 	t.Helper()
// 	if received != expected {
// 		t.Errorf("❌ expected %d received %d", received, expected)
// 	}
// }

// func assertString(t *testing.T, received string, expected string) {
// 	t.Helper()
// 	if received != expected {
// 		t.Errorf("❌ received %q expected %q", received, expected)
// 	}
// }
