package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask25(t *testing.T) {
	s := NewService(NewRegistry(), time.Now)
	merged := s.MergeCampaignSKUs(context.Background(), []string{"old"}, []string{"new", " "})
	require.Equal(t, []string{"old", "new"}, merged)
}
