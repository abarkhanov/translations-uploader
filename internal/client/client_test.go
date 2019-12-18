package client

import "testing"

func TestAPIClient_getURI(t *testing.T) {
	apiKey := "qweqwezzzz"
	apiHost := "http://example.com"
	type args struct {
		orgID     string
		emailType string
		token     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Default translation",
			args: struct {
				orgID     string
				emailType string
				token     string
			}{
				orgID:     "00000000-0000-0000-0000-000000000000",
				emailType: "trip-confirmed",
				token:     "pickup-time",
			},
			want: "http://example.com/v1/translations/email-templates%7C00000000-0000-0000-0000-000000000000%7Ctrip-confirmed%7Cpickup-time",
		},
		{name: "Custom translation",
			args: struct {
				orgID     string
				emailType string
				token     string
			}{
				orgID:     "10000000-1234-7777-9999-000000000000",
				emailType: "preauth-failed",
				token:     "passenger-value",
			},
			want: "http://example.com/v1/translations/email-templates%7C10000000-1234-7777-9999-000000000000%7Cpreauth-failed%7Cpassenger-value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Init(apiKey, apiHost)
			if got := a.getURI(tt.args.orgID, tt.args.emailType, tt.args.token); got != tt.want {
				t.Errorf("getURI() = %v, want %v", got, tt.want)
			}
		})
	}
}
