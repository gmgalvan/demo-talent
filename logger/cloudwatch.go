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

type LogLevel int

// Log level constants
const (
    DEBUG LogLevel = iota
    INFO
    WARN
    ERROR
    FATAL
)

type Logger struct {
    cwLogger *cloudwatchlogs.CloudWatchLogs
    isDebug  bool
    level    LogLevel
}

func NewLogger(isDebug bool, level LogLevel) *Logger {
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-east-1"),
    })
    if err != nil {
        log.Fatal("Error creating AWS session:", err)
    }
    return &Logger{
        cwLogger: cloudwatchlogs.New(sess),
        isDebug:  isDebug,
        level:    level,
    }
}

func (l *Logger) Log(level LogLevel, logGroupName, logStreamName, message string) {
    if l.isDebug || level >= l.level {
        logMessage := fmt.Sprintf("[%s] %s", level, message)
        if l.isDebug {
            log.Printf("[DEBUG] %s", logMessage)
            if level == FATAL {
                os.Exit(1)
            }
            return
        }

        currentTime := time.Now()
        formattedLogStreamName := fmt.Sprintf("%d/%02d/%02d-[%s]", currentTime.Year(), currentTime.Month(), currentTime.Day(), logStreamName)

        // Check if the log group exists
        describeLogGroupsOutput, err := l.cwLogger.DescribeLogGroups(&cloudwatchlogs.DescribeLogGroupsInput{
            LogGroupNamePrefix: aws.String(logGroupName),
        })
        fmt.Println("Checking if log group exists end", err, describeLogGroupsOutput, len(describeLogGroupsOutput.LogGroups))
        if err != nil {
            fmt.Println("Error describing log groups:", err)
            if level == FATAL {
                os.Exit(1)
            }
            return
        }
        if len(describeLogGroupsOutput.LogGroups) == 0 {
            // Log group doesn't exist, create it
            fmt.Println("Creating log group:", logGroupName)
            _, err := l.cwLogger.CreateLogGroup(&cloudwatchlogs.CreateLogGroupInput{
                LogGroupName: aws.String(logGroupName),
            })
            if err != nil {
                fmt.Println("Error creating log group:", err)
                if level == FATAL {
                    os.Exit(1)
                }
                return
            }
        }

        // Check if the log stream exists
        describeLogStreamsOutput, err := l.cwLogger.DescribeLogStreams(&cloudwatchlogs.DescribeLogStreamsInput{
            LogGroupName:        aws.String(logGroupName),
            LogStreamNamePrefix: aws.String(formattedLogStreamName),
        })
        if err != nil {
            fmt.Println("Error describing log streams:", err)
            if level == FATAL {
                os.Exit(1)
            }
            return
        }
        fmt.Println("Checking if log stream exists end", err, describeLogStreamsOutput, len(describeLogStreamsOutput.LogStreams))
        if len(describeLogStreamsOutput.LogStreams) == 0 {
            _, err := l.cwLogger.CreateLogStream(&cloudwatchlogs.CreateLogStreamInput{
                LogGroupName:  aws.String(logGroupName),
                LogStreamName: aws.String(formattedLogStreamName),
            })
            if err != nil {
                fmt.Println("Error creating log stream:", err)
                if level == FATAL {
                    os.Exit(1)
                }
                return
            }
        }

        timestamp := time.Now().UnixNano() / int64(time.Millisecond)
        _, err = l.cwLogger.PutLogEvents(&cloudwatchlogs.PutLogEventsInput{
            LogGroupName:  aws.String(logGroupName),
            LogStreamName: aws.String(formattedLogStreamName),
            LogEvents: []*cloudwatchlogs.InputLogEvent{
                {
                    Message:   aws.String(logMessage),  // Use the logMessage with the log level
                    Timestamp: aws.Int64(timestamp),
                },
            },
        })
        if err != nil {
            fmt.Println("Error logging to CloudWatch:", err)
            if level == FATAL {
                os.Exit(1)
            }
            return
        }

        if level == FATAL {
            os.Exit(1)
        }
    }
}

func (l LogLevel) String() string {
    switch l {
    case DEBUG:
        return "DEBUG"
    case INFO:
        return "INFO"
    case WARN:
        return "WARN"
    case ERROR:
        return "ERROR"
    case FATAL:
        return "FATAL"
    default:
        return "UNKNOWN"
    }
}