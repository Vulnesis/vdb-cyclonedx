package cyclonedx

import (
	"os"
	"strings"
)

// ── Which brand this process speaks for ──────────────────────────────────────
//
// The same library serves two brands: Vulnetix (vulnetix.com, the default) and
// Vulnesis (vulnes.is). The brand is read from VULNETIX_BRAND (VULNESIS_BRAND is
// an alias) — "vulnetix" or "vulnesis", in any case; anything else, or nothing,
// is Vulnetix, so a consumer that never sets it gets exactly the output it
// always has.
//
// Only links a document carries at run time follow the brand. Identifiers —
// tool names, the vulnetix:* property vocabulary, the tool group, purls — do
// not, because consumers key on them.

const (
	brandEnvVar      = "VULNETIX_BRAND"
	brandAliasEnvVar = "VULNESIS_BRAND"

	// vulnesisSiteURL is VulnetixSiteURL under the Vulnesis brand.
	vulnesisSiteURL = "https://www.vulnes.is"
)

// brandIsVulnesis reports whether the environment opts into the Vulnesis brand.
// It is read on every call, so a process that sets the variable before emitting
// a document is honoured without any initialisation order to get right.
func brandIsVulnesis() bool {
	v := os.Getenv(brandEnvVar)
	if strings.TrimSpace(v) == "" {
		v = os.Getenv(brandAliasEnvVar)
	}
	return strings.EqualFold(strings.TrimSpace(v), "vulnesis")
}

// siteURL is the product website for the current brand: VulnetixSiteURL by
// default, its vulnes.is equivalent under Vulnesis.
func siteURL() string {
	if brandIsVulnesis() {
		return vulnesisSiteURL
	}
	return VulnetixSiteURL
}
