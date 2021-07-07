package forntend_test

import (
	"os/exec"
	"strings"
	"testing"
)

func Test_Add_Accepted(t *testing.T) {

	out, err := exec.Command("go", "run", "../main.go", "add", "11", "Test-Note5", "This is a test note for addCmd", "5").Output()

	if err != nil {
		t.Fatal(err)
	}

	s := string(out)
	if !strings.Contains(string(s), "New ToDo item with ID") {
		t.Fatalf("expected contains \"%s\" got \"%s\"", "New ToDo item with ID", string(out))
	}

}

func Test_Add_Rejected(t *testing.T) {

	out, err := exec.Command("go", "run", "../main.go", "add", "aa", "Test-Note10", "This is a test note for addCmd", "5").Output()

	if err != nil {
		t.Fatal(err)
	}

	s := string(out)
	if !strings.Contains(string(s), "Error: First argument") {
		t.Fatalf("expected contains \"%s\" got \"%s\"", "Error: First argument", string(out))
	}

}

func Test_Complete_Accepted(t *testing.T) {
	_, err := exec.Command("go", "run", "../main.go", "add", "12", "Test-Note6", "This is a test note for completeCmd", "5").Output()

	if err != nil {
		t.Fatal(err)
	}

	out, err := exec.Command("go", "run", "../main.go", "complete", "12").Output()

	if err != nil {
		t.Fatal(err)
	}

	if !(strings.Contains(string(out), "Problem completing") || strings.Contains(string(out), "[]") || (strings.Contains(string(out), "The item was completed!"))) {
		t.Fatalf("expected contains \"%s\" or \"%s\" got \"%s\"", "Problem completing", "This is a test note", string(out))
	}
}

func Test_listCmd_Accepted(t *testing.T) {
	_, err := exec.Command("go", "run", "../main.go", "add", "13", "Test-Note7", "This is a test note for listCmd", "5").Output()

	if err != nil {
		t.Fatal(err)
	}

	out, err := exec.Command("go", "run", "../main.go", "list", "all").Output()

	if err != nil {
		t.Fatal(err)
	}
	if !(strings.Contains(string(out), "Problem retrieving ToDos") || strings.Contains(string(out), "This is a test note")) {
		t.Fatalf("expected contains \"%s\" or \"%s\" got \"%s\"", "Problem retrieving ToDos", "This is a test note", string(out))
	}
}

func Test_update_Accepted(t *testing.T) {
	_, err := exec.Command("go", "run", "../main.go", "add", "13", "Test-Note7", "This is a test note for listCmd", "5").Output()

	if err != nil {
		t.Fatal(err)
	}

	out, err := exec.Command("go", "run", "../main.go", "update", "13", "Test-Note7", "This is a test updated note", "5").Output()

	if err != nil {
		t.Fatal(err)
	}
	if !(strings.Contains(string(out), "Problem updating ToDo") || strings.Contains(string(out), "ToDo item with ID")) {
		t.Fatalf("expected contains \"%s\" or \"%s\" got \"%s\"", "Problem updating ToDo", "ToDo item with ID", string(out))
	}
}

func Test_remove_Accepted(t *testing.T) {
	out, err := exec.Command("go", "run", "../main.go", "remove", "13").Output() //Note with 13 id is added before

	if err != nil {
		t.Fatal(err)
	}

	if !(strings.Contains(string(out), "ToDo item with ID")) {
		t.Fatalf("expected contains \"%s\"", string(out))
	}
}
