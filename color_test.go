package color

import "testing"

func TestColorize(t *testing.T) {
	expected := "\033[31mtest\033[0m"
	result := Colorize(_Red, "test")
	if result != expected {
		t.Errorf("Colorize() failed, expected %q but got %q", expected, result)
	}
}

func TestColorizef(t *testing.T) {
	expected := "\033[32mtest 123\033[0m"
	result := Colorizef(_Green, "test %d", 123)
	if result != expected {
		t.Errorf("Colorizef() failed, expected %q but got %q", expected, result)
	}
}

func TestBlackf(t *testing.T) {
	expected := "\033[30mtest 123\033[0m"
	result := Blackf("test %d", 123)
	if result != expected {
		t.Errorf("Blackf() failed, expected %q but got %q", expected, result)
	}
}

func TestRedf(t *testing.T) {
	expected := "\033[31mtest 123\033[0m"
	result := Redf("test %d", 123)
	if result != expected {
		t.Errorf("Redf() failed, expected %q but got %q", expected, result)
	}
}

func TestGreenf(t *testing.T) {
	expected := "\033[32mtest 123\033[0m"
	result := Greenf("test %d", 123)
	if result != expected {
		t.Errorf("Greenf() failed, expected %q but got %q", expected, result)
	}
}

func TestYellowf(t *testing.T) {
	expected := "\033[33mtest 123\033[0m"
	result := Yellowf("test %d", 123)
	if result != expected {
		t.Errorf("Yellowf() failed, expected %q but got %q", expected, result)
	}
}

func TestBluef(t *testing.T) {
	expected := "\033[34mtest 123\033[0m"
	result := Bluef("test %d", 123)
	if result != expected {
		t.Errorf("Bluef() failed, expected %q but got %q", expected, result)
	}
}

func TestMagentaf(t *testing.T) {
	expected := "\033[35mtest 123\033[0m"
	result := Magentaf("test %d", 123)
	if result != expected {
		t.Errorf("Magentaf() failed, expected %q but got %q", expected, result)
	}
}

func TestCyanf(t *testing.T) {
	expected := "\033[36mtest 123\033[0m"
	result := Cyanf("test %d", 123)
	if result != expected {
		t.Errorf("Cyanf() failed, expected %q but got %q", expected, result)
	}
}

func TestWhitef(t *testing.T) {
	expected := "\033[37mtest 123\033[0m"
	result := Whitef("test %d", 123)
	if result != expected {
		t.Errorf("Whitef() failed, expected %q but got %q", expected, result)
	}
}

func TestBlack(t *testing.T) {
	expected := "\033[30mtest\033[0m"
	if result := Black("test"); result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestRed(t *testing.T) {
	expected := "\033[31mtest\033[0m"
	if result := Red("test"); result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestGreen(t *testing.T) {
	expected := "\033[32mtest\033[0m"
	if result := Green("test"); result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestYellow(t *testing.T) {
	expected := "\033[33mtest\033[0m"
	if result := Yellow("test"); result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestBlue(t *testing.T) {
	expected := "\033[34mtest\033[0m"
	if result := Blue("test"); result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestMagenta(t *testing.T) {
	expected := "\033[35mtest\033[0m"
	if result := Magenta("test"); result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestCyan(t *testing.T) {
	expected := "\033[36mtest\033[0m"
	if result := Cyan("test"); result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestWhite(t *testing.T) {
	expected := "\033[37mtest\033[0m"
	if result := White("test"); result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestSuccess(t *testing.T) {
	expected := "\033[32msuccess\033[0m"
	actual := Success("success")
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestError(t *testing.T) {
	expected := "\033[31merror\033[0m"
	actual := Error("error")
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestWarning(t *testing.T) {
	expected := "\033[33mwarning\033[0m"
	actual := Warning("warning")
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestInfo(t *testing.T) {
	expected := "\033[34minfo\033[0m"
	actual := Info("info")
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestDebug(t *testing.T) {
	expected := "\033[36mdebug\033[0m"
	result := Debug("debug")
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestFatal(t *testing.T) {
	expected := "\033[35mfatal\033[0m"
	result := Fatal("fatal")
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestSuccessf(t *testing.T) {
	expected := "\033[32mtest: 123\033[0m"
	result := Successf("test: %d", 123)
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestErrorf(t *testing.T) {
	expected := "\033[31mtest: 123\033[0m"
	result := Errorf("test: %d", 123)
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestWarningf(t *testing.T) {
	expected := "\033[33mtest: 123\033[0m"
	result := Warningf("test: %d", 123)
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestInfof(t *testing.T) {
	expected := "\033[34mtest: 123\033[0m"
	result := Infof("test: %d", 123)
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestDebugf(t *testing.T) {
	expected := "\033[36mdebug message\033[0m"
	result := Debugf("debug message")

	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}

func TestFatalf(t *testing.T) {
	expected := "\033[35mfatal message\033[0m"
	result := Fatalf("fatal message")

	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
