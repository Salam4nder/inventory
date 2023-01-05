package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewConfig(t *testing.T) {
	tests := []struct {
		name    string
		envVars map[string]string
		//TODO: compare cfg with expected cfg
		wantErr bool
	}{
		{
			name: "All env vars set returns no error, correct cfg",
			envVars: map[string]string{
				"DB_PORT":     "8080",
				"DB_HOST":     "localhost",
				"DB_USER":     "root",
				"DB_PASSWORD": "password",
				"DB_NAME":     "test",
			},
		},
		{
			name: "Missing env var returns error",
			envVars: map[string]string{
				"DB_PORT":     "8080",
				"DB_HOST":     "localhost",
				"DB_USER":     "root",
				"DB_PASSWORD": "",
				"DB_NAME":     "test",
			},
			wantErr: true,
		},
		{
			name:    "Empty env var returns error",
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

			cfg, err := NewConfig()
			if test.wantErr {
				assert.Error(t, err)
				assert.Nil(t, cfg)
			} else {
				//TODO: compare cfg with expected cfg
				assert.NoError(t, err)
				assert.NotNil(t, cfg)
			}
		})
	}
}
