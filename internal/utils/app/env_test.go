package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDevelopment(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want bool
	}{
		{
			name: "APP_ENV is development",
			env:  "development",
			want: true,
		}, {
			name: "APP_ENV is not development",
			env:  "production",
			want: false,
		}, {
			name: "APP_ENV is empty",
			env:  "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("APP_ENV", tt.env)
			assert.Equal(t, tt.want, IsDevelopment())
		})
	}
}

func TestIsProduction(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want bool
	}{
		{
			name: "APP_ENV is production",
			env:  "production",
			want: true,
		}, {
			name: "APP_ENV is development",
			env:  "development",
			want: false,
		}, {
			name: "APP_ENV is empty",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("APP_ENV", tt.env)
			assert.Equal(t, tt.want, IsProduction())
		})
	}
}
