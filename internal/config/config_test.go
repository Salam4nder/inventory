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
				"POSTGRES_HOST":     "localhost",
				"POSTGRES_PORT":     "8080",
				"POSTGRES_USER":     "root",
				"POSTGRES_DB":       "test",
				"POSTGRES_PASSWORD": "password",
				"SERVER_HOST":       "localhost",
				"SERVER_PORT":       "8080",
				"SERVER_JWT":        "secret",
				"SERVER_LOG_FILE":   "log.txt",
				"REDIS_HOST":        "localhost",
				"REDIS_PORT":        "8080",
				"REDIS_PASSWORD":    "password",
			},
		},
		{
			name: "Missing env var returns error",
			envVars: map[string]string{
				"POSTGRES_HOST":     "localhost2",
				"POSTGRES_PORT":     "5050",
				"POSTGRES_USER":     "admin",
				"POSTGRES_DB":       "myName",
				"POSTGRES_PASSWORD": "",
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