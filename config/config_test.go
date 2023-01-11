package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	tests := []struct {
		name    string
		envVars map[string]string
		//TODO: compare cfg with expected cfg
		wantErr bool
	}{
		{
			name: "All env vars set returns no error, correct cfg",
			envVars: map[string]string{
				"DB_HOST":     "localhost",
				"DB_PORT":     "8080",
				"DB_USER":     "root",
				"DB_NAME":     "test",
				"DB_PASSWORD": "password",
				"HTTP_PORT":   "8080",
			},
		},
		{
			name: "Missing env var returns error",
			envVars: map[string]string{
				"DB_HOST":     "localhost2",
				"DB_PORT":     "5050",
				"DB_USER":     "admin",
				"DB_NAME":     "myName",
				"DB_PASSWORD": "",
			},
			wantErr: true,
		},
		{
			name:    "File empty returns error",
			envVars: map[string]string{},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for k, v := range test.envVars {
				os.Setenv(k, v)
			}
			defer func() {
				for k := range test.envVars {
					os.Unsetenv(k)
				}
			}()

			cfg, err := New()

			if test.wantErr {
				assert.Error(t, err)
			} else {
				//TODO: compare cfg with expected cfg
				t.Log("cfg: ", cfg)
				assert.NoError(t, err)
			}
		})
	}
}
