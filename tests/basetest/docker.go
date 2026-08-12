/*
Copyright © 2026 Matze
*/
package base

import (
	"testing"
	"time"

	"github.com/mattia37773/mt/helper/docker"
)

func WaitForContainersRunning(t *testing.T, containers []string, timeout time.Duration) {
	t.Helper()

	deadline := time.Now().Add(timeout)

	for _, container := range containers {
		for !docker.IsContainerRunning(container) {
			if time.Now().After(deadline) {
				t.Fatalf("Container '%s' is not running within %v", container, timeout)
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
}
