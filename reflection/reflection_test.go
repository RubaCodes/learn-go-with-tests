package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	// implementng table tesing to have framework to test multiple cases
	cases := []struct {
		Name          string
		Input         any
		ExpectedCalls []string
	}{
		{
			Name:          "Struct with one string field",
			Input:         struct{ Name string }{"Chris"},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "Struct with two string field",
			Input: struct {
				Name string
				City string
			}{"Chris", "London"},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "Struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{"Chris", 33},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "Struct with nested field",
			Input: Person{
				"Chris", Profile{33, "London"},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "Pointer to things",
			Input: &Person{
				"Chris", Profile{33, "London"},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "Slices",
			Input: []Profile{
				{33, "London"},
				{34, "Rome"},
			},
			ExpectedCalls: []string{"London", "Rome"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"maps",
			map[string]string{
				"Cow":   "Moo",
				"Sheep": "Baa",
			},
			[]string{"Moo", "Baa"},
		},
	}

	for _, v := range cases {
		t.Run(v.Name, func(t *testing.T) {
			var got []string

			walk(v.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, v.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, v.ExpectedCalls)
			}
		})
	}
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})
	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
