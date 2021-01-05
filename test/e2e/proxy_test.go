package e2e

import (
	"strings"
	"testing"
)

func TestHTTPProxy(t *testing.T) {
	logger := Logger{}
	kubectl := Kubectl{t, "kapp-controller", logger}

	logger.Section("inspect controller logs for propogation of proxy env vars", func() {
		// app name must match the app name being deployed in hack/deploy-test.sh
		out := kubectl.Run([]string{"logs", "deployment/kapp-controller"})

		if !strings.Contains(out, "Using http proxy") {
			t.Fatalf("expected log line saying a proxy is in use")
		}
	})

}
