package domain

// Event - generic wrapper for passing triggers to handlers
// this was done so S3Event didn't need to pollute every layer
type Event interface{}
