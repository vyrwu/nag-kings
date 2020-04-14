package swagger

import (
	"reflect"
	"testing"
	"time"
)

func Test_calculateYearsRuledFromString(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "1000-1500",
			want: 500,
		},
		{
			name: "1000-",
			want: time.Now().Year() - 1000,
		},
		{
			name: "1000",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := calculateYearsRuledFromString(tt.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateYearsRuledFromString() = got %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateLongestRulingKing(t *testing.T) {
	tests := []struct {
		name     string
		argument map[King]int
		want     *KingsStatisticsLongestRulingKing
	}{
		{
			name: "single longest ruling",
			argument: map[King]int{
				King{ID: 3}: 100,
				King{ID: 2}: 50,
				King{ID: 1}: 1,
			},
			want: &KingsStatisticsLongestRulingKing{
				King:       []King{King{ID: 3}},
				YearsRuled: 100,
			},
		}, {
			name: "two longest ruling",
			argument: map[King]int{
				King{ID: 3}: 100,
				King{ID: 1}: 1,
			},
			want: &KingsStatisticsLongestRulingKing{
				King:       []King{King{ID: 3}},
				YearsRuled: 100,
			},
		}, {
			name: "same years ruled",
			argument: map[King]int{
				King{ID: 2}: 1,
			},
			want: &KingsStatisticsLongestRulingKing{
				King:       []King{King{ID: 2}},
				YearsRuled: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cgot := <-calculateLongestRulingKing(tt.argument)

			if !reflect.DeepEqual(cgot, tt.want) {
				t.Errorf("calculateLongestRulingKing() = got %v, want %v", cgot, tt.want)
			}
		})
	}
}

func Test_calculateLongestRulingHouse(t *testing.T) {
	tests := []struct {
		name     string
		argument map[string]int
		want     *KingsStatisticsLongestRulingHouse
	}{
		{
			name: "single longest ruling",
			argument: map[string]int{
				"house3": 100,
				"house2": 50,
				"house1": 1,
			},
			want: &KingsStatisticsLongestRulingHouse{
				HouseName:  []string{"house3"},
				YearsRuled: 100,
			},
		}, {
			name: "two longest ruling",
			argument: map[string]int{
				"house2": 100,
				"house1": 1,
			},
			want: &KingsStatisticsLongestRulingHouse{
				HouseName:  []string{"house2"},
				YearsRuled: 100,
			},
		}, {
			name: "same years ruled",
			argument: map[string]int{
				"house1": 1,
			},
			want: &KingsStatisticsLongestRulingHouse{
				HouseName:  []string{"house1"},
				YearsRuled: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cgot := <-calculateLongestRulingHouse(tt.argument)

			if !reflect.DeepEqual(cgot, tt.want) {
				t.Errorf("calculateLongestRulingHouse() = got %v, want %v", cgot, tt.want)
			}
		})
	}
}

func Test_calculateMostCommonFirstName(t *testing.T) {
	tests := []struct {
		name     string
		argument map[string]int
		want     []string
	}{
		{
			name: "one most common",
			argument: map[string]int{
				"Jim":  100,
				"Joe":  10,
				"John": 1,
			},
			want: []string{"Jim"},
		}, {
			name: "default case",
			argument: map[string]int{
				"Jim":  100,
				"John": 1,
			},
			want: []string{"Jim"},
		}, {
			name: "default case",
			argument: map[string]int{
				"Jim": 1,
			},
			want: []string{"Jim"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cgot := <-calculateMostCommonFirstName(tt.argument)

			if !reflect.DeepEqual(cgot, tt.want) {
				t.Errorf("calculateMostCommonFirstName() = got %v, want %v", cgot, tt.want)
			}
		})
	}
}
