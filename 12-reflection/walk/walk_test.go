package walk

import (
	"reflect"
	"slices"
	"testing"
)

type User struct {
	Name string
}

type Organization struct {
	Name string
}

type UserWithAge struct {
	Name string
	Age  int
}

type UserWithLastName struct {
	FirstName string
	LastName  string
}

type UserWithOrganization struct {
	Name         string
	Organization Organization
}

func TestWalk(t *testing.T) {
	cases := []struct {
		name      string
		input     any
		wantCalls []string
	}{
		{
			name:      "struct with one string field",
			input:     User{Name: "Evgeny"},
			wantCalls: []string{"Evgeny"},
		},
		{
			name: "struct with two string fields",
			input: UserWithLastName{
				FirstName: "Evgeny",
				LastName:  "Markov",
			},
			wantCalls: []string{"Evgeny", "Markov"},
		},
		{
			name: "struct with int and string fields",
			input: UserWithAge{
				Name: "Evgeny",
				Age:  27,
			},
			wantCalls: []string{"Evgeny"},
		},
		{
			name: "struct with nested structs",
			input: UserWithOrganization{
				Name: "Evgeny",
				Organization: Organization{
					Name: "Yandex",
				},
			},
			wantCalls: []string{"Evgeny", "Yandex"},
		},
		{
			name: "pointers to things",
			input: &UserWithOrganization{
				Name: "Evgeny",
				Organization: Organization{
					Name: "Yandex",
				},
			},
			wantCalls: []string{"Evgeny", "Yandex"},
		},
		{
			name: "slices",
			input: []UserWithOrganization{
				{
					Name: "Evgeny",
					Organization: Organization{
						Name: "Yandex",
					},
				},
				{
					Name: "Arkady",
					Organization: Organization{
						Name: "Tinkoff",
					},
				},
			},
			wantCalls: []string{"Evgeny", "Yandex", "Arkady", "Tinkoff"},
		},
		{
			name: "arrays",
			input: [2]UserWithOrganization{
				{
					Name: "Evgeny",
					Organization: Organization{
						Name: "Yandex",
					},
				},
				{
					Name: "Arkady",
					Organization: Organization{
						Name: "Tinkoff",
					},
				},
			},
			wantCalls: []string{"Evgeny", "Yandex", "Arkady", "Tinkoff"},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			gotCalls := make([]string, 0)

			walk(test.input, func(value string) {
				gotCalls = append(gotCalls, value)
			})

			if !reflect.DeepEqual(gotCalls, test.wantCalls) {
				t.Errorf("got %v, want %v, input %v", gotCalls, test.wantCalls, test.input)
			}
		})
	}

	t.Run("maps", func(t *testing.T) {
		input := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		gotCalls := make([]string, 0)
		walk(input, func(value string) {
			gotCalls = append(gotCalls, value)
		})

		assertContains(t, gotCalls, "Moo")
		assertContains(t, gotCalls, "Baa")
	})

	t.Run("channels", func(t *testing.T) {
		aChannel := make(chan UserWithOrganization)

		go func() {
			aChannel <- UserWithOrganization{
				Name: "Evgeny",
				Organization: Organization{
					Name: "Yandex",
				},
			}
			aChannel <- UserWithOrganization{
				Name: "Arkady",
				Organization: Organization{
					Name: "Tinkoff",
				},
			}
			close(aChannel)
		}()

		gotCalls := make([]string, 0)
		wantCalls := []string{"Evgeny", "Yandex", "Arkady", "Tinkoff"}

		walk(aChannel, func(value string) {
			gotCalls = append(gotCalls, value)
		})

		if !reflect.DeepEqual(gotCalls, wantCalls) {
			t.Errorf("got %v, want %v", gotCalls, wantCalls)
		}
	})

	t.Run("functions", func(t *testing.T) {
		aFunction := func() (UserWithOrganization, UserWithOrganization) {
			return UserWithOrganization{
					Name: "Evgeny",
					Organization: Organization{
						Name: "Yandex",
					},
				},
				UserWithOrganization{
					Name: "Arkady",
					Organization: Organization{
						Name: "Tinkoff",
					},
				}
		}

		gotCalls := make([]string, 0)
		wantCalls := []string{"Evgeny", "Yandex", "Arkady", "Tinkoff"}

		walk(aFunction, func(value string) {
			gotCalls = append(gotCalls, value)
		})

		if !reflect.DeepEqual(gotCalls, wantCalls) {
			t.Errorf("got %v, want %v", gotCalls, wantCalls)
		}
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()

	contains := slices.Contains(haystack, needle)

	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
