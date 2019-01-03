package errors

import (
	"runtime"
	"strconv"
	"strings"
)

// Get the absolute path split by src path.
func GetFrameABSFileName(file string) string {
	if splits := strings.Split(file, "src/"); len(splits) == 2 {
		return splits[1]
	}

	return file
}

// Get pc caller stacks. Callers fills the slice pc with the return program counters of function invocations
// on the calling goroutine's stack.
func GetCallerStacks(pc []uintptr) (stacks []string) {

	frames := runtime.CallersFrames(pc)

	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		if frame.Line == 0 {
			continue
		}

		buffer := strings.Builder{}

		buffer.WriteString(frame.Function)
		buffer.WriteString("\n\t")
		buffer.WriteString(GetFrameABSFileName(frame.File))
		buffer.WriteByte(':')
		buffer.WriteString(strconv.FormatInt(int64(frame.Line), 10))

		stacks = append(stacks, buffer.String())
	}

	return
}

// Takes a slice of PC values returned by Callers and
// prepares to return file/function/line information.
func GetCallerSource(pc []uintptr) (file string, function string, line int) {
	frames := runtime.CallersFrames(pc)

	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		if frame.Line == 0 {
			continue
		}
		return GetFrameABSFileName(frame.File), frame.Function, frame.Line
	}

	return
}

// New caller, skip and deep is callers. The argument skip is the number of stack frames
// to skip before recording in pc, with 0 identifying the frame for Callers itself and
// 1 identifying the caller of Callers. If stack is true then get call stacks.
func NewCaller(skip, deep int, stack bool) *Caller {

	pc := make([]uintptr, deep)
	runtime.Callers(skip, pc)

	frames := runtime.CallersFrames(pc)

	caller := &Caller{}
	first := true

	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		if frame.Line == 0 {
			continue
		}

		if first {
			first = false
			caller.Function = frame.Function
			caller.File = frame.File
			caller.Line = int64(frame.Line)
		}

		if stack {
			buffer := strings.Builder{}
			buffer.WriteString(frame.Function)
			buffer.WriteString("\n\t")
			buffer.WriteString(GetFrameABSFileName(frame.File))
			buffer.WriteByte(':')
			buffer.WriteString(strconv.FormatInt(int64(frame.Line), 10))

			caller.Stacks = append(caller.Stacks, buffer.String())
		}

	}

	return caller
}
