package handlers

import (
	"bytes"
	"errors"
	"testing"

	"github.com/mdw-cohort-c/calc-lib"
	"github.com/smarty/assertions/should"
)

func assertError(t *testing.T, actual, target error) {
	t.Helper()
	if !errors.Is(actual, target) {
		t.Errorf("expected %v, got %v", target, actual)
	}
}
func TestHandler_WrongNumberOfArguments(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle(nil)
	should.So(t, err, should.Wrap, errWrongNumberOfArgs)
}
func TestHandler_InvalidFirstArgument(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"INVALID", "3"})
	should.So(t, err, should.Wrap, errInvalidArg)
}
func TestHandler_InvalidSecondArgument(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"3", "INVALID"})
	should.So(t, err, should.Wrap, errInvalidArg)
}
func TestHandler_OutputWriterError(t *testing.T) {
	boink := errors.New("boink")
	writer := &ErringWriter{err: boink}
	handler := NewHandler(writer, &calc.Addition{})

	err := handler.Handle([]string{"3", "4"})

	should.So(t, err, should.Wrap, boink)
	should.So(t, err, should.Wrap, errWriterFailure)
}
func TestHandler_HappyPath(t *testing.T) {
	writer := &bytes.Buffer{}
	handler := NewHandler(writer, &calc.Addition{})

	err := handler.Handle([]string{"3", "4"})

	should.So(t, err, should.BeNil)
	should.So(t, writer.String(), should.Equal, "7")
}

type ErringWriter struct {
	err error
}

func (this *ErringWriter) Write(p []byte) (n int, err error) {
	return 0, this.err
}
