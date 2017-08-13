package insight_test

import (
	"testing"
	"time"
)

func TestTransactionsByBlock(t *testing.T) {
	summaries, err := c.BlockSummaries(1, time.Now())
	if err != nil {
		t.Error(err)
	}
	_, err = c.TransactionsByBlock(summaries.Blocks[0].Hash, 0)
	if err != nil {
		t.Error(err)
	}

}
