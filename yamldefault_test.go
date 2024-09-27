package main

import (
	"os/user"
	"reflect"
	"testing"
)

func TestExpandPath(t *testing.T) {
	usr, _ := user.Current()
	home_dir := usr.HomeDir
	tables := []struct {
		path     string
		expected string
	}{
		{"this/path", "this/path"},
		{"this/path/", "this/path/"},
		{"path", "path"},
		{"path/", "path/"},
		{"~", home_dir},
		{"~/", home_dir},
		{"~/test", home_dir + "/test"},
	}

	for _, table := range tables {
		out := expandPath(table.path)
		if out != table.expected {
			t.Errorf(
				"Expansion of %s was incorrect, got: %s, want: %s.",
				table.path,
				out,
				table.expected,
			)
		}
	}
}

func TestOverwriteDefaultLocation(t *testing.T) {
	tables := []struct {
		name     string
		osArgs   []string
		expected string
		wantErr  bool
	}{
		{"Default location", []string{"file"}, "test", false},
		{"No change", []string{"file", "location"}, "test", false},
		{"Valid new location", []string{"file", "location", "test_2"}, "test_2", false},
		{"Invalid new location", []string{"file", "location", "invalid"}, "", true},
	}

	for _, table := range tables {
		t.Run(table.name, func(t *testing.T) {
			locations := map[string]string{"test": "test", "test_2": "test_2"}
			config := NotesConfig{"test", locations}
			out, err := overwriteDefaultLocation(config, table.osArgs)

			if table.wantErr {
				if err == nil {
					t.Errorf("Expected an error, but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if !reflect.DeepEqual(out.Locations, locations) {
					t.Errorf(
						"Overwrite altered locations: got: %s, expected no change (%s)",
						out.Locations,
						locations,
					)
				}
				if out.DefaultLocation != table.expected {
					t.Errorf(
						"Overwrite was incorrect, got: %s, want: %s.", out.DefaultLocation, table.expected,
					)
				}
			}
		})
	}
}
