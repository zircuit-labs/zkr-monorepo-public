package bindgen

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var generator = BindGenGeneratorRemote{}

func TestRemoveDeploymentSalt(t *testing.T) {
	t.Parallel()
	for _, tt := range removeDeploymentSaltTests {

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, _ := generator.removeDeploymentSalt(tt.deploymentData, tt.deploymentSalt)
			require.Equal(t, tt.expected, got)
		})
	}
}

func TestRemoveDeploymentSaltFailures(t *testing.T) {
	t.Parallel()
	for _, tt := range removeDeploymentSaltTestsFailures {

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := generator.removeDeploymentSalt(tt.deploymentData, tt.deploymentSalt)
			require.Equal(t, err.Error(), tt.expectedError)
		})
	}
}
