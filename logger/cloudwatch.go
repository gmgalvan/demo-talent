package logger

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type Logger struct {
    cwLogger *cloudwatchlogs.CloudWatchLogs
    isDebug  bool
}

func NewLogger(isDebug bool) *Logger {
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-east-1"),
    })
    if err != nil {
        log.Fatal("Error creating AWS session:", err)
    }

    return &Logger{
        cwLogger: cloudwatchlogs.New(sess),
        isDebug:  isDebug,
    }
}

func (l *Logger) Log(logGroupName, logStreamName, message string) {
    if l.isDebug {
        log.Printf("[DEBUG] %s", message)
        return
    }

    // Check if the log group exists
    _, err := l.cwLogger.DescribeLogGroups(&cloudwatchlogs.DescribeLogGroupsInput{
        LogGroupNamePrefix: aws.String(logGroupName),
    })
    if err != nil {
        if aerr, ok := err.(awserr.Error); ok && aerr.Code() == cloudwatchlogs.ErrCodeResourceNotFoundException {
            // Log group doesn't exist, create it
            _, err = l.cwLogger.CreateLogGroup(&cloudwatchlogs.CreateLogGroupInput{
                LogGroupName: aws.String(logGroupName),
            })
            if err != nil {
                fmt.Println("Error creating log group:", err)
                os.Exit(1)
            }
        } else {
            fmt.Println("Error describing log groups:", err)
            os.Exit(1)
        }
    }

    timestamp := time.Now().UnixNano() / int64(time.Millisecond)

    _, err = l.cwLogger.PutLogEvents(&cloudwatchlogs.PutLogEventsInput{
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