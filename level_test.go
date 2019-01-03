package errors

import (
	"fmt"
	"testing"
)

func TestLevel_String(t *testing.T) {
	ds := []struct {
		level  Level
		string string
	}{
		{level: -1, string: "-1"},
		{level: Level_Debug, string: "Debug"},
		{level: Level_Normal, string: "Normal"},
		{level: Level_Warn, string: "Warn"},
		{level: Level_Error, string: "Error"},
		{level: Level_Panic, string: "Panic"},
		{level: Level_Fatal, string: "Fatal"},
		{level: Level_Fatal + 1, string: fmt.Sprintf("%d", Level_Fatal+1)},
	}

	for idx, d := range ds {
		if d.string != d.level.String() {
			t.Errorf("%d string should be %s, not be %s", idx+1, d.string, d.level.String())
		}
	}
}

func TestLevel_CapitalString(t *testing.T) {
	ds := []struct {
		level  Level
		string string
	}{
		{level: -1, string: "LEVEL(-1)"},
		{level: Level_Debug, string: "DEBUG"},
		{level: Level_Normal, string: "NORMAL"},
		{level: Level_Warn, string: "WARN"},
		{level: Level_Error, string: "ERROR"},
		{level: Level_Panic, string: "PANIC"},
		{level: Level_Fatal, string: "FATAL"},
		{level: Level_Fatal + 1, string: fmt.Sprintf("LEVEL(%d)", Level_Fatal+1)},
	}

	for idx, d := range ds {
		if d.string != d.level.CapitalString() {
			t.Errorf("%d string should be %s, not be %s", idx+1, d.string, d.level.CapitalString())
		}
	}
}

func TestLevel_GetStack(t *testing.T) {
	ds := []struct {
		level      Level
		haveSource bool
		haveStack  bool
	}{
		{level: -1, haveSource: true, haveStack: false},
		{level: Level_Debug, haveSource: true, haveStack: true},
		{level: Level_Normal, haveSource: false, haveStack: false},
		{level: Level_Warn, haveSource: true, haveStack: false},
		{level: Level_Error, haveSource: true, haveStack: true},
		{level: Level_Panic, haveSource: true, haveStack: true},
		{level: Level_Fatal, haveSource: true, haveStack: false},
		{level: Level_Fatal + 1, haveSource: true, haveStack: false},
	}

	for idx, d := range ds {
		caller := d.level.GetCaller(3, 64)
		if caller == nil {
			continue
		}
		if d.haveSource {
			if caller.File == "" || caller.Function == "" || caller.Line == 0 {
				t.Errorf("%d %s should have source", idx, d.level)
			}
		} else {
			if caller.File != "" || caller.Function != "" || caller.Line != 0 {
				t.Errorf("%d %s should not have source", idx, d.level)
			}
		}
		if d.haveStack {
			if len(caller.Stacks) == 0 {
				t.Errorf("%d %s should have stack", idx, d.level)
			}
		} else {
			if len(caller.Stacks) != 0 {
				t.Errorf("%d %s should not have stack", idx, d.level)
			}
		}
	}
}

func TestLevel_Trace(t *testing.T) {
	var level Level = Level_Panic

	func() {
		defer func() {
			if e := recover(); e != nil {
				t.Log(e)
				// panic(e)
			}
		}()
		level.Trace(fmt.Errorf("%s %d", "TestLevel_Trace", level))
	}()

	level = Level_Fatal
	level.Trace(fmt.Errorf("%s %d", "TestLevel_Trace", level))
}
