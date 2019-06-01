package convert

import "testing"

type test struct {
	name string
	req float64
	want float64
}
func TestFtoC(t *testing.T) {
	tests := []test{
		{
			name: "frezzing",
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
			if res != test.want {
				t.Errorf("FtoC(%v) = %v but want %v", test.req, res, test.want)
			}
		})
	}
}
func TestCtof(t *testing.T) {
	tests := []test{
		{
			name: "frezzing",
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
			if res != test.want {
				t.Errorf("CotF(%v) = %v but want %v", test.req, res, test.want)
			}
		})
	}
}

func TestKtoC(t *testing.T) {
	tests := []test{
		{
			name: "frezzing",
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
			if res != test.want {
				t.Errorf("KtoC(%v) = %v but want %v", test.req, res, test.want)
			}
		})
	}
}


func TestCtoK(t *testing.T) {
	tests := []test{
		{
			name: "frezzing",
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
			if res != test.want {
				t.Errorf("CtoK(%v) = %v but want %v", test.req, res, test.want)
			}
		})
	}
}


func TestKtof(t *testing.T) {
	tests := []test{
		{
			name: "frezzing",
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
			if res != test.want {
				t.Errorf("KtoF(%v) = %v but want %v", test.req, res, test.want)
			}
		})
	}
}


func TestFtoK(t *testing.T) {
	tests := []test{
		{
			name: "frezzing",
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
			if res != test.want {
				t.Errorf("FtoK(%v) = %v but want %v", test.req, res, test.want)
			}
		})
	}
}

