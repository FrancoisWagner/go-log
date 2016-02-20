package testutils_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/opentracing/basictracer-go"
	"github.com/opentracing/basictracer-go/testutils"
)

func TestInMemoryRecorderSpans(t *testing.T) {
	recorder := testutils.NewInMemoryRecorder()
	var apiRecorder basictracer.SpanRecorder = recorder
	span := basictracer.RawSpan{
		Context:   basictracer.Context{},
		Operation: "test-span",
		Start:     time.Now(),
		Duration:  -1,
	}
	apiRecorder.RecordSpan(span)
	if len(recorder.GetSpans()) != 1 {
		t.Fatal("No spans recorded")
	}
	if !reflect.DeepEqual(recorder.GetSpans()[0], span) {
		t.Fatal("Span not recorded")
	}
}
