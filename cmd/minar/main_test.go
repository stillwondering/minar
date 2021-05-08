package main

import "testing"

func TestRun(t *testing.T) {
	tt := []struct {
		name        string
		cmdline     []string
		expectError bool
	}{
		{
			name:        "Empty cmdline",
			expectError: false,
		},
		{
			name:        "Help flag",
			cmdline:     []string{"-h"},
			expectError: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := run(tc.cmdline)

			gotError := err != nil
			if tc.expectError != gotError {
				t.Errorf("expected error: %v, got: %v", tc.expectError, err)
			}
		})
	}
}

func TestConfigFromCmdline(t *testing.T) {
	tt := []struct {
		name           string
		cmdline        []string
		expectError    bool
		expectedConfig *config
	}{
		{
			name:        "Help flag",
			cmdline:     []string{"-h"},
			expectError: true,
		},
		{
			name:        "Empty command line",
			cmdline:     []string{},
			expectError: false,
			expectedConfig: &config{
				listenAddress: ":8080",
			},
		},
		{
			name:        "Address configured",
			cmdline:     []string{"--address", ":3000"},
			expectError: false,
			expectedConfig: &config{
				listenAddress: ":3000",
			},
		},
		{
			name:        "Address configuration with long format",
			cmdline:     []string{"--address=:4000"},
			expectError: false,
			expectedConfig: &config{
				listenAddress: ":4000",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			cfg, err := configFromCmdline(tc.cmdline)

			gotError := err != nil
			if tc.expectError != gotError {
				t.Fatalf("expected error: %v, got error: %v", tc.expectError, gotError)
			}

			// If we expect an error, we don't need to check anything else.
			if tc.expectError {
				return
			}

			if *tc.expectedConfig != *cfg {
				t.Errorf("expected config: %v, got: %v", tc.expectedConfig, cfg)
			}
		})
	}
}
