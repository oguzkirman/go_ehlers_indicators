package go_ehlers_indicators

import (
	"fmt"
	"github.com/MathisWellmann/go_timeseries_generator"
	"testing"
)

func TestInverseFisherTransform(t *testing.T) {
	sine := go_timeseries_generator.SineWave(1024)
	ift := InverseFisherTransform(sine)

	for i := 0; i < len(ift); i++ {
		if ift[i] > 1 || ift[i] < -1 {
			t.Error(ift[i] > 1 || ift[i] < -1)
		}
	}
}

func TestInverseFisherTransformGraph(t *testing.T) {
	sine := go_timeseries_generator.SineWave(1024)
	ift := InverseFisherTransform(sine)

	filename := fmt.Sprintf("img/inverse_fisher_transform.png")
	err := Plt(ift, filename)
	if err != nil {
		t.Error(err)
	}
}
