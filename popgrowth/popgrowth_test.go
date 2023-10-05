package popgrowth

import "testing"

func TestNbYear(t *testing.T) {
	tc := []struct {
		p0      int
		percent float64
		aug     int
		p       int
		want    int
	}{
		{
			p0:      1000,
			percent: 2,
			aug:     50,
			p:       1200,
			want:    3,
		},
		{
			p0:      1500,
			percent: 5,
			aug:     100,
			p:       5000,
			want:    15,
		},
		{
			p0:      1500000,
			percent: 2.5,
			aug:     10000,
			p:       2000000,
			want:    10,
		},
		{
			p0:      1500000,
			percent: 0.25,
			aug:     -1000,
			p:       2000000,
			want:    151,
		},
	}
	for _, test := range tc {
		if got := NbYear(test.p0, test.percent, test.aug, test.p); got != test.want {
			t.Errorf("want %d, got %d", test.want, got)
		}
	}
}
