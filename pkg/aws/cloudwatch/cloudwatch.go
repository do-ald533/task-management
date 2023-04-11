package cloudwatch

import (
	"goapi/pkg/aws/awshelper"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type CWLogs struct {
	LOG_GROUP  string
	LOG_STREAM string
	service    *cloudwatchlogs.CloudWatchLogs
}

func Init() *CWLogs {
	service := new()

	return &CWLogs{
		LOG_GROUP:  os.Getenv("LOG_GROUP"),
		LOG_STREAM: os.Getenv("LOG_STREAM"),
		service:    service,
	}
}

func new() *cloudwatchlogs.CloudWatchLogs {
	session := awshelper.NewAwsSession()
	cwlogs := cloudwatchlogs.New(session)

	return cwlogs
}

// https://docs.aws.amazon.com/sdk-for-go/api/service/cloudwatchlogs/#CloudWatchLogs.DescribeLogStreams
// https://docs.aws.amazon.com/sdk-for-go/api/service/cloudwatchlogs/#DescribeLogStreamsInput
// https://docs.aws.amazon.com/sdk-for-go/api/service/cloudwatchlogs/#CloudWatchLogs
// https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeLogStreams.html
// https://docs.aws.amazon.com/sdk-for-go/api/service/cloudwatchlogs/#DescribeLogStreamsOutput
// https://docs.aws.amazon.com/sdk-for-go/api/service/cloudwatchlogs/#LogStream
func (cwl *CWLogs) nextSequenceToken() string {
	input := cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName:        &cwl.LOG_GROUP,
		LogStreamNamePrefix: &cwl.LOG_STREAM,
	}

	resp, err := cwl.service.DescribeLogStreams(&input)
	if err != nil || len(resp.LogStreams) == 0 {
		return ""
	}
	nextTokenAddr := resp.LogStreams[0].UploadSequenceToken
	if nextTokenAddr == nil {
		return ""
	}
	return *nextTokenAddr
}

// https://docs.aws.amazon.com/sdk-for-go/api/service/cloudwatchlogs/#CloudWatchLogs.PutLogEvents
// https://docs.aws.amazon.com/sdk-for-go/api/service/cloudwatchlogs/#PutLogEventsInput
// https://docs.aws.amazon.com/sdk-for-go/api/service/cloudwatchlogs/#InputLogEvent
func (cwl *CWLogs) Log(message *string) (*cloudwatchlogs.PutLogEventsOutput, error) {
	unixTimeInMil := time.Now().UnixNano() / int64(time.Millisecond)

	event := cloudwatchlogs.InputLogEvent{Message: message, Timestamp: &unixTimeInMil}
	events := []*cloudwatchlogs.InputLogEvent{&event}

	nextToken := cwl.nextSequenceToken()

	data := cloudwatchlogs.PutLogEventsInput{
		LogEvents:     events,
		LogGroupName:  &cwl.LOG_GROUP,
		LogStreamName: &cwl.LOG_STREAM,
	}

	if nextToken != "" {
		data.SetSequenceToken(nextToken)
	}

	resp, err := cwl.service.PutLogEvents(&data)
	return resp, err
}
