package cdn

import (
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/cdn"
	"testing"
)

func DescribeContentBlockTasks(t *testing.T) {
	resp, err := DefaultInstance.DescribeContentBlockTasks(&cdn.DescribeContentBlockTasksRequest{
		TaskType:  "block_url",
		StartTime: testStartTime,
		EndTime:   testEndTime,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Result.Data)
}
