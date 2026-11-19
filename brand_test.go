package cyclonedx

import "testing"

func websiteOf(c Component) string {
	for _, r := range c.ExternalReferences {
		if r.Type == "website" {
			return r.URL
		}
	}
	return ""
}

func TestVulnetixToolWebsiteFollowsBrand(t *testing.T) {
	cases := []struct {
		brand, alias, want string
	}{
		{"", "", VulnetixSiteURL},
		{"vulnetix", "", VulnetixSiteURL},
		{"bogus", "", VulnetixSiteURL},
		{"vulnesis", "", "https://www.vulnes.is"},
		{" VulNesis ", "", "https://www.vulnes.is"},
		{"", "vulnesis", "https://www.vulnes.is"},
		{"vulnetix", "vulnesis", VulnetixSiteURL},
	}
	for _, tc := range cases {
		t.Setenv(brandEnvVar, tc.brand)
		t.Setenv(brandAliasEnvVar, tc.alias)
		c := VulnetixTool(ToolSCA, "1.0.0")
		if got := websiteOf(c); got != tc.want {
			t.Errorf("brand=%q alias=%q: website = %q, want %q", tc.brand, tc.alias, got, tc.want)
		}
		// Identifiers never follow the brand.
		if c.Name != ToolSCA || c.Group != VulnetixToolGroup || c.Purl != "pkg:golang/github.com/vulnetix/cli@v1.0.0" {
			t.Errorf("brand=%q alias=%q: identifiers changed: %+v", tc.brand, tc.alias, c)
		}
	}
}

func TestIsVulnetixToolNameAcceptsBothBrands(t *testing.T) {
	for _, name := range []string{"vulnesis", "vulnesis-sca", "Vulnesis Malscan", "Vulnetix Malscan", "vulnetix-sca"} {
		if !IsVulnetixToolName(name) {
			t.Errorf("IsVulnetixToolName(%q) = false", name)
		}
	}
	for _, name := range []string{"vulnesisx", "vulnetixx", "my-vulnesis"} {
		if IsVulnetixToolName(name) {
			t.Errorf("IsVulnetixToolName(%q) = true", name)
		}
	}
}
