package set_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/micnncim/go-set"
)

func equal(t *testing.T) func(*set.Set[int], *set.Set[int]) bool {
	t.Helper()

	return func(s1, s2 *set.Set[int]) bool {
		return s1.Equal(s2)
	}
}

func TestSetDelete(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    *set.Set[int]
		v    []int
		want *set.Set[int]
	}{
		{
			name: "delete int",
			s:    set.New(1, 2, 3),
			v:    []int{1, 2},
			want: set.New(3),
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.s.Delete(tt.v...)

			if diff := cmp.Diff(tt.want, tt.s, cmp.Comparer(equal(t))); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}

func TestSetDifference(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    *set.Set[int]
		t    *set.Set[int]
		want *set.Set[int]
	}{
		{
			name: "difference int 1",
			s:    set.New(1, 2, 3),
			t:    set.New(1, 2, 4, 5),
			want: set.New(3),
		},
		{
			name: "difference int 2",
			s:    set.New(1, 2, 4, 5),
			t:    set.New(1, 2, 3),
			want: set.New(4, 5),
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(tt.want, tt.s.Difference(tt.t), cmp.Comparer(equal(t))); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}

func TestSetEqual(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    *set.Set[int]
		t    *set.Set[int]
		want bool
	}{
		{
			name: "equal",
			s:    set.New(1, 2),
			t:    set.New(2, 1),
			want: true,
		},
		{
			name: "not equal",
			s:    set.New(1, 2),
			t:    set.New(1, 3),
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(tt.want, tt.s.Equal(tt.t)); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}

func TestSetHas(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    *set.Set[int]
		v    int
		want bool
	}{
		{
			name: "has",
			s:    set.New(1),
			v:    1,
			want: true,
		},
		{
			name: "has no int",
			s:    set.New(1),
			v:    2,
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(tt.want, tt.s.Has(tt.v)); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}

func TestSetHasAll(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    *set.Set[int]
		v    []int
		want bool
	}{
		{
			name: "has all",
			s:    set.New(1, 2, 3),
			v:    []int{1, 2},
			want: true,
		},
		{
			name: "not has all",
			s:    set.New(1, 2),
			v:    []int{1, 2, 3},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(tt.want, tt.s.HasAll(tt.v...)); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}

func TestSetHasAny(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    *set.Set[int]
		v    []int
		want bool
	}{
		{
			name: "has any",
			s:    set.New(1, 2),
			v:    []int{1, 3},
			want: true,
		},
		{
			name: "not has any",
			s:    set.New(1, 2),
			v:    []int{3},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(tt.want, tt.s.HasAny(tt.v...)); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}

func TestSetInsert(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    *set.Set[int]
		v    int
		want *set.Set[int]
	}{
		{
			name: "insert int",
			s:    set.New(1),
			v:    2,
			want: set.New(1, 2),
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.s.Insert(tt.v)

			if diff := cmp.Diff(tt.want, tt.s); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}

func TestSetIntersection(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    *set.Set[int]
		t    *set.Set[int]
		want *set.Set[int]
	}{
		{
			name: "intersection int with same len",
			s:    set.New(1, 2, 3),
			t:    set.New(2, 3, 5),
			want: set.New(2, 3),
		},
		{
			name: "intersection int with different len",
			s:    set.New(1, 2, 3, 4),
			t:    set.New(2, 3, 5),
			want: set.New(2, 3),
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(tt.want, tt.s.Intersection(tt.t), cmp.Comparer(equal(t))); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}

func TestSetIsSuperset(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    *set.Set[int]
		t    *set.Set[int]
		want bool
	}{
		{
			name: "is superset",
			s:    set.New(1, 2, 3),
			t:    set.New(1, 2),
			want: true,
		},
		{
			name: "is not superset",
			s:    set.New(1, 2),
			t:    set.New(1, 2, 3),
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(tt.want, tt.s.IsSuperset(tt.t)); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}

func TestSetLen(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    *set.Set[int]
		want int
	}{
		{
			name: "len",
			s:    set.New(1, 2),
			want: 2,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(tt.want, tt.s.Len()); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}

func TestSetPopAny(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		s        *set.Set[int]
		want     int
		wantBool bool
		wantSet  *set.Set[int]
	}{
		{
			name:     "pop",
			s:        set.New(1, 2),
			want:     1,
			wantBool: true,
			wantSet:  set.New(2),
		},
		{
			name:     "no pop",
			s:        set.New[int](),
			want:     0,
			wantBool: false,
			wantSet:  set.New[int](),
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, ok := tt.s.PopAny()

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.wantBool, ok); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.wantSet, tt.s); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}

func TestSetString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    *set.Set[int]
		want string
	}{
		{
			name: "string of int",
			s:    set.New(1),
			want: "[1]",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(tt.want, tt.s.String()); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}

func TestSetValues(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    *set.Set[int]
		want []int
	}{
		{
			name: "valid int set",
			s:    set.New(1, 2),
			want: []int{1, 2},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(tt.want, tt.s.Values()); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}

func TestSetUnion(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    *set.Set[int]
		t    *set.Set[int]
		want *set.Set[int]
	}{
		{
			name: "union",
			s:    set.New(1, 2),
			t:    set.New(2, 3),
			want: set.New(1, 2, 3),
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(tt.want, tt.s.Union(tt.t), cmp.Comparer(equal(t))); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}
