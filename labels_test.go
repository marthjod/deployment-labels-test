package labels

import (
	"fmt"
	"strings"
	"testing"

	"sigs.k8s.io/yaml"
)

const (
	path     = "testdata/deployment.yml"
	labelKey = "mylabel"
)

var tests = []struct {
	labelValue        string
	expectedRendering string
}{
	{
		labelValue:        "abcd1234",
		expectedRendering: fmt.Sprintf("%s: abcd1234", labelKey),
	},
	{
		labelValue:        "12345678",
		expectedRendering: fmt.Sprintf(`%s: "12345678"`, labelKey),
	},
}

func TestLabels(t *testing.T) {
	depl, err := getDeployment(path)
	if err != nil {
		t.Error(err)
	}
	for _, tt := range tests {
		t.Run(tt.labelValue, func(t *testing.T) {
			depl.ObjectMeta.Labels[labelKey] = tt.labelValue

			res, err := yaml.Marshal(depl)
			if err != nil {
				t.Error(err)
			}
			if !strings.Contains(string(res), tt.expectedRendering) {
				t.Errorf("expected '%s', got\n%s", tt.expectedRendering, res)
			}
		})
	}
}
