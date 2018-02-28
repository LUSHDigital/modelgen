package main

import "testing"

func TestGetOrderFromComment(t *testing.T) {
	type args struct {
		comment string
	}
	tests := []struct {
		name      string
		args      args
		wantOrder int
	}{
		{
			name: "valid comment",
			args: args{
				comment: "modelgen:1",
			},
			wantOrder: 1,
		},
		{
			name: "invalid comment",
			args: args{
				comment: "modelgen:2:something else",
			},
			wantOrder: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOrder := GetOrderFromComment(tt.args.comment); gotOrder != tt.wantOrder {
				t.Errorf("GetOrderFromComment() = %v, want %v", gotOrder, tt.wantOrder)
			}
		})
	}
}
