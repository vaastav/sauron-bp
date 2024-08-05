package sauron

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

// SauronLogger implements the backend.Logger interface for Sauron
type SauronLogger struct {
	msg_chan    chan string
	outfile     string
	BUFFER_SIZE int64
}

// Infinite loop function that processes messages
func (l *SauronLogger) msgFunc() {
	f, err := os.Create(l.outfile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	counter := int64(0)
	for {
		select {
		case msg := <-l.msg_chan:
			// For now write message to file
			w.WriteString(msg)
			counter += 1
			// Buffer writing to file
			if counter >= l.BUFFER_SIZE {
				w.Flush()
				counter = 0
			}
			// TODO: Generate embedding for the message
			// TODO: Send message to otel collector or log ingestor like Loki
		}
	}
}

// Handle the generated message
// RIght now only sends it to the consumer thread
func (l *SauronLogger) handleMsg(ctx context.Context, msg string) {
	log.Println(msg)
	l.msg_chan <- msg
}

// Implements backend.Logger
func (l *SauronLogger) Debug(ctx context.Context, format string, args ...any) (context.Context, error) {
	msg := "[DEBUG]:" + fmt.Sprintf(format, args...)
	go l.handleMsg(ctx, msg)
	return ctx, nil
}

// Implements backend.Logger
func (l *SauronLogger) Info(ctx context.Context, format string, args ...any) (context.Context, error) {
	msg := "[INFO]:" + fmt.Sprintf(format, args...)
	go l.handleMsg(ctx, msg)
	return ctx, nil
}

// Implements backend.Logger
func (l *SauronLogger) Warn(ctx context.Context, format string, args ...any) (context.Context, error) {
	msg := "[WARN]:" + fmt.Sprintf(format, args...)
	go l.handleMsg(ctx, msg)
	return ctx, nil
}

// Implements backend.Logger
func (l *SauronLogger) Error(ctx context.Context, format string, args ...any) (context.Context, error) {
	msg := "[ERROR]:" + fmt.Sprintf(format, args...)
	go l.handleMsg(ctx, msg)
	return ctx, nil
}

// Implements backend.Logger
func (l *SauronLogger) Logf(ctx context.Context, opts backend.LogOptions, format string, args ...any) (context.Context, error) {
	msg := "[" + opts.Level.String() + "]" + fmt.Sprintf(format, args...)
	go l.handleMsg(ctx, msg)
	return ctx, nil
}

// Initialize the logger
func (l *SauronLogger) initLogger(buffer_size int64, outfile string) {
	l.msg_chan = make(chan string, l.BUFFER_SIZE)
	l.outfile = outfile
	go l.msgFunc()
}

// Returns a new logger object
func NewSauronLogger(ctx context.Context, buffer_size_string string, outfile string) (*SauronLogger, error) {
	l := &SauronLogger{}
	buffer_size, err := strconv.ParseInt(buffer_size_string, 10, 64)
	if err != nil {
		return nil, err
	}
	l.initLogger(buffer_size, outfile)
	backend.SetDefaultLogger(l)
	return l, nil
}
