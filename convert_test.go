package convert

import (
	"math"
	"testing"
)

type test struct {
	name string
	req float64
	want float64
}

const float32EqualityThreshold = 1e-6

func TestFtoC(t *testing.T) {
	tests := []test{
		{
			name: "freezing",
			req: 32,
			want: 0,
		},
		{
			name: "boiling",
			req: 212,
			want: 100,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {	
			res := FtoC(test.req)
			if math.Abs(res - test.want) > float32EqualityThreshold {
				t.Errorf("FtoC(%v) = %v but want %v", test.req, res, test.want)
			}
		})
	}
}
func TestCtoF(t *testing.T) {
	tests := []test{
		{
			name: "freezing",
			req: 0,
			want: 32,
		},
		{
			name: "boiling",
			req: 100,
			want: 212,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {	
			res := CtoF(test.req)
			if math.Abs(res - test.want) > float32EqualityThreshold {
				t.Errorf("CotF(%v) = %v but want %v", test.req, res, test.want)
			}
		})
	}
}

func TestKtoC(t *testing.T) {
	tests := []test{
		{
			name: "freezing",
			req: 273.15,
			want: 0,
		},
		{
			name: "boiling",
			req: 373.15,
			want: 100,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {	
			res := KtoC(test.req)
			if math.Abs(res - test.want) > float32EqualityThreshold {
				t.Errorf("KtoC(%v) = %v but want %v", test.req, res, test.want)
			}
		})
	}
}


func TestCtoK(t *testing.T) {
	tests := []test{
		{
			name: "freezing",
			req: 0,
			want: 273.15,
		},
		{
			name: "boiling",
			req: 100,
			want: 373.15,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {	
			res := CtoK(test.req)
			if math.Abs(res - test.want) > float32EqualityThreshold {
				t.Errorf("CtoK(%v) = %v but want %v", test.req, res, test.want)
			}
		})
	}
}


func TestKtoF(t *testing.T) {
	tests := []test{
		{
			name: "freezing",
			req: 273.15,
			want: 32,
		},
		{
			name: "boiling",
			req: 373.15,
			want: 212,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {	
			res := KtoF(test.req)
			if math.Abs(res - test.want) > float32EqualityThreshold {
				t.Errorf("KtoF(%v) = %v but want %v", test.req, res, test.want)
			}
		})
	}
}


func TestFtoK(t *testing.T) {
	tests := []test{
		{
			name: "freezing",
			req: 32,
			want: 273.15,
		},
		{
			name: "boiling",
			req: 212,
			want: 373.15,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {	
			res := FtoK(test.req)
			if math.Abs(res - test.want) > float32EqualityThreshold {
				t.Errorf("FtoK(%v) = %v but want %v", test.req, res, test.want)
			}
		})
	}
}

