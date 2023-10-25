package interface_implementations_test

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

// AKdmsaodaodkqkwe s-a crapat
// {"msg":"AKdmsaodaodkqkwe s-a crapat ", "level": "info"}
// printate in standard out
type Logger interface {
	LogErr(tags map[string]string, message string) string
	LogInfo(tags map[string]string, message string) string
	LogWarn(tags map[string]string, message string) string
}

func TestLogging(t *testing.T) {
	type input struct {
		fn      func(tags map[string]string, message string) string
		tags    map[string]string
		message string
	}
	type testData struct {
		name     string
		input    input
		expected string
	}

	var l Logger
	for _, td := range []testData{
		{
			name: "info message",
			input: input{
				fn:      l.LogInfo,
				tags:    map[string]string{"action": "handleStudentEvents", "eventType": "studentAdded"},
				message: "Finished parsing messages",
			},
			expected: `{"level":"info","message":"Finished parsing messages","action":"handleStudentEvents","eventType":"studentAdded"}`,
		},
		{
			name: "warn message",
			input: input{
				fn:      l.LogWarn,
				tags:    map[string]string{"action": "handleStudentEvents"},
				message: "Received a message of an unhandled type",
			},
			expected: `{"level":"warn","message":"Received a message of an unhandled type","action":"handleStudentEvents"}`,
		},
		{
			name: "error message",
			input: input{
				fn:      l.LogErr,
				tags:    map[string]string{"action": "handleStudentEvents"},
				message: "Could not parse message",
			},
			expected: `{"level":"error","message":"Could not parse message","action":"handleStudentEvents"}`,
		},
	} {
		t.Run(td.name, func(t *testing.T) {
			gotJsonString := td.input.fn(td.input.tags, td.input.message)
			gotParsed := map[string]any{}
			require.NoError(t, json.Unmarshal([]byte(gotJsonString), &gotParsed))

			expectedParsed := map[string]any{}
			require.NoError(t, json.Unmarshal([]byte(td.expected), &expectedParsed))

			require.Equal(t, expectedParsed, gotParsed)
		})
	}
}
