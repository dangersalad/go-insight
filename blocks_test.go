package insight_test

import (
	"github.com/dangersalad/go-insight"
	"testing"
	"time"
)

var c = insight.NewClient("https://insight.bitpay.com/api")

func TestBlockSummaries(t *testing.T) {
	_, err := c.BlockSummaries(1, time.Now())
	if err != nil {
		t.Fail()
	}
}
