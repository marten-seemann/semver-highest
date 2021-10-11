package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInvalidTarget(t *testing.T) {
	_, err := GetHighestBefore("v0.1", "v0.1.0,v0.1.1", false)
	require.EqualError(t, err, "target version is not a semver version: v0.1")
}

func TestTargetAlreadyReleased(t *testing.T) {
	_, err := GetHighestBefore("v0.1.0", "v0.1.0,v0.1.1", false)
	require.EqualError(t, err, "target version v0.1.0 already exists")
}

func TestHighestBefore(t *testing.T) {
	v, err := GetHighestBefore("v0.2.0", "v0.1.0,v0.1.1,v0.1.2-alpha,v0.2.2,v0.3.0,v1.0.0", false)
	require.NoError(t, err)
	require.Equal(t, v, "v0.1.1")
}

func TestHighestBeforeIncludePrerelease(t *testing.T) {
	v, err := GetHighestBefore("v0.2.0", "v0.1.0,v0.1.1,v0.1.2-alpha.1,v0.2.2,v0.3.0,v1.0.0", true)
	require.NoError(t, err)
	require.Equal(t, v, "v0.1.2-alpha.1")
}

func TestHighestBeforeIncludePrereleaseForPrerelease(t *testing.T) {
	v, err := GetHighestBefore("v0.1.2-alpha.3", "v0.1.0,v0.1.1,v0.1.2-beta.1,v0.2.2", true)
	require.NoError(t, err)
	require.Equal(t, v, "v0.1.1")
}
