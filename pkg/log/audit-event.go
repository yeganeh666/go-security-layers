package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

var auditLogFilePath = "./logs/audit-log.txt"

func AuditEvent(method, path, username, ip string) {
	logEntry := fmt.Sprintf("Time: %s, Method: %s, Path: %s, Username:%s, IP: %s\n",
		time.Now(), method, path, username, ip)
	if err := appendToFile(auditLogFilePath, logEntry); err != nil {
		log.Printf("Error appending to audit log file: %v\n", err)
	}
}

func appendToFile(filePath, data string) error {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(data); err != nil {
		return err
	}

	return nil
}
