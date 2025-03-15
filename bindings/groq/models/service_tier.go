package models

type ServiceTier string

const (
	// ServiceTierOnDemand is the default service tier.
	ServiceTierOnDemand ServiceTier = "on_demand"
	// ServiceTierAuto automatically select the highest tier available within the rate limits of your organization.
	ServiceTierAuto ServiceTier = "auto"
	// ServiceTierFlex uses the flex tier, which will succeed or fail quickly.
	ServiceTierFlex ServiceTier = "flex"
)
