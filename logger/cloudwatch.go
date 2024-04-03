package logger

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

var cwLogger *cloudwatchlogs.CloudWatchLogs

func init() {
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-east-1"),
    })
    if err != nil {
        log.Fatal("Error creating AWS session:", err)
    }

    cwLogger = cloudwatchlogs.New(sess)
}

func LogToCloudWatch(logGroupName, logStreamName, message string) {
    timestamp := time.Now().UnixNano() / int64(time.Millisecond)

    _, err := cwLogger.PutLogEvents(&cloudwatchlogs.PutLogEventsInput{
        LogGroupName:  aws.String(logGroupName),
        LogStreamName: aws.String(logStreamName),
        LogEvents: []*cloudwatchlogs.InputLogEvent{
            {
                Message:   aws.String(message),
                Timestamp: aws.Int64(timestamp),
            },
        },
    })
    if err != nil {
        fmt.Println("Error logging to CloudWatch:", err)
        os.Exit(1)
    }
}