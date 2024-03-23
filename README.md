# addllog
A simplistic logging framework for files and cloud.

## include
`import github.com/infni/addllog/log`

## usage

```
    // import golog "log"

	logFilename := fmt.Sprintf("application_%s.log", time.Now().Format("20060102150405"))

	logFilePointer, logError := os.Create(logFilename)
	if logError != nil {
		fmt.Printf("Failed to open log file '%v'. Error: %v\n", logFilename, logError)
		return false
	}

	logger := log.NewFileLogger(golog.New(logFilePointer, "", 0))
	defer logFilePointer.Close()

	addlInfo := log.AddlInfo{
		"cfg":     "url-string-example",
		"version": "1.0.0-1234",
	}

    logger.Log("Applicatoin is running", "application-started", addlInfo)

    logger.Log("example", null, null)

    logger.LogCritical(fmt.Sprintf(" : gRPC Listener aborted : %v", err.Error()), "grpc-failure", addlInfo)
```
