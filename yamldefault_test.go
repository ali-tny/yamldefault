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
		osArgs   []string
		expected string
	}{
		{[]string{"file"}, "test"},
		{[]string{"file", "location"}, "test"},
		{[]string{"file", "location", "new"}, "new"},
		{[]string{"file", "location", "other"}, "other"},
	}

	for _, table := range tables {
		locations := map[string]string{"test": "test", "test_2": "test_2"}
		config := NotesConfig{"test", locations}
		out := overwriteDefaultLocation(config, table.osArgs)
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
}
