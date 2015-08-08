package path

import (
	"go/build"
	"os"
	"path/filepath"
	"testing"
)

func TestSunplateDir(t *testing.T) {
	p := filepath.ToSlash(filepath.Join(build.Default.GOPATH, "src", spImp))

	if v := SunplateDir(); v != p {
		t.Errorf(`Incorrect result. Expected "%s", got "%s".`, p, v)
	}

	hp := filepath.ToSlash(filepath.Join(p, "generation", "handlers"))
	if v := SunplateDir("generation", "handlers"); v != hp {
		t.Errorf(`Incorrect result. Expected "%s", got "%s".`, hp, v)
	}
}

func TestSunplateImport(t *testing.T) {
	p := spImp

	if v := SunplateImport(); v != p {
		t.Errorf(`Incorrect result. Expected "%s", got "%s".`, p, v)
	}

	hp := filepath.ToSlash(filepath.Join(p, "generation", "handlers"))
	if v := SunplateImport("generation", "handlers"); v != hp {
		t.Errorf(`Incorrect result. Expected "%s", got "%s".`, hp, v)
	}
}

func TestWorkingDir(t *testing.T) {
	p := filepath.ToSlash(filepath.Join(build.Default.GOPATH, "src"))
	os.Chdir(p)

	if v := WorkingDir(); v != p {
		t.Errorf(`Incorrectly detected working directory. Expected "%s", got "%s"`, p, v)
	}
}

func TestAbsoluteImport_AbsImportArgument(t *testing.T) {
	if v := AbsoluteImport(spImp); v != "github.com/anonx/sunplate" {
		t.Errorf(`Incorrect result. Expected "%s", got "%s".`, spImp, v)
	}

	if v := AbsoluteImport(""); v != "" {
		t.Errorf(`Empty input: empty output expected, got "%s".`, v)
	}
}

func TestAbsoluteImport_OutsideGOPATH(t *testing.T) {
	p := filepath.ToSlash(filepath.Join(build.Default.GOPATH, "../"))
	os.Chdir(p)

	defer func() {
		if err := recover(); err == nil {
			t.Errorf(`Start outside of $GOPATH: "%s", panic expected.`, p)
		}
	}()
	AbsoluteImport("./strings")
}

func TestAbsoluteImport(t *testing.T) {
	os.Chdir(filepath.Join(build.Default.GOPATH, "src", spImp))

	d := "./generation"
	exp := filepath.ToSlash(filepath.Join(spImp, d))
	if v := AbsoluteImport(d); v != exp {
		t.Errorf(`Incorrect result. Expected "%s", got "%s".`, exp, v)
	}

	d = "/string"
	exp = "string"
	if v := AbsoluteImport(d); v != exp {
		t.Errorf(`Incorrect result. Expected "%s", got "%s".`, exp, v)
	}
}

func TestPrefixless(t *testing.T) {
	p := "/anonx/sunplate"
	if v := Prefixless(filepath.ToSlash(filepath.Join("github.com", p)), "github.com"); v != p {
		t.Errorf(`Inncorrect result. Expected "%s", got %s.`, p, v)
	}
}

var spImp = "github.com/anonx/sunplate"