package productsubscription

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sourcegraph/sourcegraph/enterprise/cmd/llm-proxy/internal/dotcom"
)

func TestNewActor(t *testing.T) {
	type args struct {
		s               dotcom.ProductSubscriptionState
		devLicensesOnly bool
	}
	tests := []struct {
		name        string
		args        args
		wantEnabled bool
	}{
		{
			name: "not dev only",
			args: args{
				dotcom.ProductSubscriptionState{
					LlmProxyAccess: dotcom.ProductSubscriptionStateLlmProxyAccessLLMProxyAccess{
						LLMProxyAccessFields: dotcom.LLMProxyAccessFields{
							Enabled: true,
							RateLimit: &dotcom.LLMProxyAccessFieldsRateLimitLLMProxyRateLimit{
								Limit:           10,
								IntervalSeconds: 10,
							},
						},
					},
				},
				false,
			},
			wantEnabled: true,
		},
		{
			name: "dev only, not a dev license",
			args: args{
				dotcom.ProductSubscriptionState{
					LlmProxyAccess: dotcom.ProductSubscriptionStateLlmProxyAccessLLMProxyAccess{
						LLMProxyAccessFields: dotcom.LLMProxyAccessFields{
							Enabled: true,
							RateLimit: &dotcom.LLMProxyAccessFieldsRateLimitLLMProxyRateLimit{
								Limit:           10,
								IntervalSeconds: 10,
							},
						},
					},
				},
				true,
			},
			wantEnabled: false,
		},
		{
			name: "dev only, is a dev license",
			args: args{
				dotcom.ProductSubscriptionState{
					LlmProxyAccess: dotcom.ProductSubscriptionStateLlmProxyAccessLLMProxyAccess{
						LLMProxyAccessFields: dotcom.LLMProxyAccessFields{
							Enabled: true,
							RateLimit: &dotcom.LLMProxyAccessFieldsRateLimitLLMProxyRateLimit{
								Limit:           10,
								IntervalSeconds: 10,
							},
						},
					},
					ActiveLicense: &dotcom.ProductSubscriptionStateActiveLicenseProductLicense{
						Info: &dotcom.ProductSubscriptionStateActiveLicenseProductLicenseInfo{
							Tags: []string{"dev"},
						},
					},
				},
				true,
			},
			wantEnabled: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			act := NewActor(nil, "", tt.args.s, tt.args.devLicensesOnly)
			assert.Equal(t, act.AccessEnabled, tt.wantEnabled)
		})
	}
}
