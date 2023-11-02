package resource_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	resource "github.com/telia-oss/github-pr-resource"
)


var key = `-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA0BUezcR7uycgZsfVLlAf4jXP7uFpVh4geSTY39RvYrAll0yh
q7uiQypP2hjQJ1eQXZvkAZx0v9lBYJmX7e0HiJckBr8+/O2kARL+GTCJDJZECpjy
97yylbzGBNl3s76fZ4CJ+4f11fCh7GJ3BJkMf9NFhe8g1TYS0BtSd/sauUQEuG/A
3fOJxKTNmICZr76xavOQ8agA4yW9V5hKcrbHzkfecg/sQsPMmrXixPNxMsqyOMmg
jdJ1aKr7ckEhd48ft4bPMO4DtVL/XFdK2wJZZ0gXJxWiT1Ny41LVql97Odm+OQyx
tcayMkGtMb1nwTcVVl+RG2U5E1lzOYpcQpyYFQIDAQABAoIBAAfUY55WgFlgdYWo
i0r81NZMNBDHBpGo/IvSaR6y/aX2/tMcnRC7NLXWR77rJBn234XGMeQloPb/E8iw
vtjDDH+FQGPImnQl9P/dWRZVjzKcDN9hNfNAdG/R9JmGHUz0JUddvNNsIEH2lgEx
C01u/Ntqdbk+cDvVlwuhm47MMgs6hJmZtS1KDPgYJu4IaB9oaZFN+pUyy8a1w0j9
RAhHpZrsulT5ThgCra4kKGDNnk2yfI91N9lkP5cnhgUmdZESDgrAJURLS8PgInM4
YPV9L68tJCO4g6k+hFiui4h/4cNXYkXnaZSBUoz28ICA6e7I3eJ6Y1ko4ou+Xf0V
csM8VFkCgYEA7y21JfECCfEsTHwwDg0fq2nld4o6FkIWAVQoIh6I6o6tYREmuZ/1
s81FPz/lvQpAvQUXGZlOPB9eW6bZZFytcuKYVNE/EVkuGQtpRXRT630CQiqvUYDZ
4FpqdBQUISt8KWpIofndrPSx6JzI80NSygShQsScWFw2wBIQAnV3TpsCgYEA3reL
L7AwlxCacsPvkazyYwyFfponblBX/OvrYUPPaEwGvSZmE5A/E4bdYTAixDdn4XvE
ChwpmRAWT/9C6jVJ/o1IK25dwnwg68gFDHlaOE+B5/9yNuDvVmg34PWngmpucFb/
6R/kIrF38lEfY0pRb05koW93uj1fj7Uiv+GWRw8CgYEAn1d3IIDQl+kJVydBKItL
tvoEur/m9N8wI9B6MEjhdEp7bXhssSvFF/VAFeQu3OMQwBy9B/vfaCSJy0t79uXb
U/dr/s2sU5VzJZI5nuDh67fLomMni4fpHxN9ajnaM0LyI/E/1FFPgqM+Rzb0lUQb
yqSM/ptXgXJls04VRl4VjtMCgYEAprO/bLx2QjxdPpXGFcXbz6OpsC92YC2nDlsP
3cfB0RFG4gGB2hbX/6eswHglLbVC/hWDkQWvZTATY2FvFps4fV4GrOt5Jn9+rL0U
elfC3e81Dw+2z7jhrE1ptepprUY4z8Fu33HNcuJfI3LxCYKxHZ0R2Xvzo+UYSBqO
ng0eTKUCgYEAxW9G4FjXQH0bjajntjoVQGLRVGWnteoOaQr/cy6oVii954yNMKSP
rezRkSNbJ8cqt9XQS+NNJ6Xwzl3EbuAt6r8f8VO1TIdRgFOgiUXRVNZ3ZyW8Hegd
kGTL0A6/0yAu9qQZlFbaD5bWhQo7eyx63u4hZGppBhkTSPikOYUPCH8=
-----END RSA PRIVATE KEY-----`

func TestSource(t *testing.T) {
	tests := []struct {
		description string
		source      resource.Source
		wantErr     string
	}{
		{
			description: "validate passes",
			source: resource.Source{
				AccessToken: "123456",
				Repository:  "test/test",
			},
		},
		{
			description: "should have an access_token",
			source: resource.Source{
				Repository: "test/test",
			},
			wantErr: "access_token must be set if not using GitHub App authentication",
		},
		{
			description: "should have a repository",
			source: resource.Source{
				AccessToken: "123456",
			},
			wantErr: "repository must be set",
		},
		{
			description: "should support GitHub App authentication",
			source: resource.Source{
				Repository:     "test/test",
				UseGitHubApp:   true,
				PrivateKey:     key,
				ApplicationID:  123456,
				InstallationID: 1,
			},
		},
		{
			description: "requires a private_key GitHub App configuration values",
			source: resource.Source{
				Repository:     "test/test",
				UseGitHubApp:   true,
				ApplicationID:  123456,
				InstallationID: 1,
			},
			wantErr: "private_key should be supplied if using GitHub App authentication",
		},
		{
			description: "requires an application_id and installation_id GitHub App configuration values",
			source: resource.Source{
				Repository:     "test/test",
				UseGitHubApp:   true,
				PrivateKey:     key,
				ApplicationID:  123456,
			},
			wantErr: "application_id and installation_id must be set if using GitHub App authentication",
		},
		{
			description: "should not have an access_token when using GitHub App authentication",
			source: resource.Source{
				Repository:     "test/test",
				UseGitHubApp:   true,
				PrivateKey:     key,
				ApplicationID:  123456,
				InstallationID: 1,
				AccessToken:    "123456",
			},
			wantErr: "access_token is not required when using GitHub App authentication",
		},
		{
			description: "requires v3_endpoint when v4_endpoint is set",
			source: resource.Source{
				AccessToken: "123456",
				Repository:  "test/test",
				V3Endpoint:  "https://github.com/v3",
			},
			wantErr: "v4_endpoint must be set together with v3_endpoint",
		},
		{
			description: "requires v4_endpoint when v3_endpoint is set",
			source: resource.Source{
				AccessToken: "123456",
				Repository:  "test/test",
				V4Endpoint:  "https://github.com/v4",
			},
			wantErr: "v3_endpoint must be set together with v4_endpoint",
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			err := tc.source.Validate()

			if tc.wantErr != "" {
				if err == nil {
					t.Logf("Expected error '%s', got nothing", tc.wantErr)
					t.Fail()
				}
				assert.EqualError(t, err, tc.wantErr, fmt.Sprintf("Expected '%s', got '%s'", tc.wantErr, err))
			}

			if tc.wantErr == "" && err != nil {
				t.Logf("Got an error when none expected: %s", err)
				t.Fail()
			}
		})
	}
}
