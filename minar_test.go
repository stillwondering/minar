package minar

import "testing"

func TestEquals(t *testing.T) {
	tt := []struct {
		name    string
		a       Minutes
		b       Minutes
		isEqual bool
	}{
		{
			name:    "Test1",
			a:       Minutes{Title: "Hi"},
			b:       Minutes{Title: "Hello"},
			isEqual: false,
		},
		{
			name:    "Test2",
			a:       Minutes{Title: "Hi"},
			b:       Minutes{Title: "Hi"},
			isEqual: true,
		},
		{
			name:    "Test3",
			a:       Minutes{Title: "Hi", Participants: []string{"A"}},
			b:       Minutes{Title: "Hello", Participants: []string{"B"}},
			isEqual: false,
		},
		{
			name:    "Test4",
			a:       Minutes{Title: "Hi", Participants: []string{"A"}},
			b:       Minutes{Title: "Hi", Participants: []string{"A"}},
			isEqual: true,
		},
		{
			name:    "Test5",
			a:       Minutes{Title: "Hi", Participants: []string{"A"}, Topics: []Topic{{Title: "Topic A", Content: "Content A"}}},
			b:       Minutes{Title: "Hi", Participants: []string{"A"}, Topics: []Topic{{Title: "Topic B", Content: "Content B"}}},
			isEqual: false,
		},
		{
			name:    "Test6",
			a:       Minutes{Title: "Hi", Participants: []string{"A"}, Topics: []Topic{{Title: "Topic A", Content: "Content A"}}},
			b:       Minutes{Title: "Hi", Participants: []string{"A"}, Topics: []Topic{{Title: "Topic A", Content: "Content A"}}},
			isEqual: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			isEqual := tc.a.Equals(tc.b)
			if tc.isEqual != tc.a.Equals(tc.b) {
				t.Fatalf("expected isEqual:%v, got isEqual: %v", tc.isEqual, isEqual)
			}

			isEqual = tc.b.Equals(tc.a)
			if tc.isEqual != tc.a.Equals(tc.b) {
				t.Fatalf("expected isEqual:%v, got isEqual: %v", tc.isEqual, isEqual)
			}
		})
	}
}
