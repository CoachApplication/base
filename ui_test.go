package standard

import "testing"

func commonUi() *Ui {
	return NewUi("test.1", "Test", "This is a Test", "In order to use this test do X")
}

func TestUi_Id(t *testing.T) {
	ui := commonUi()
	if ui.Id() != "test.1" {
		t.Error("UI returned the wrong Id")
	}
}

func TestUi_Label(t *testing.T) {
	ui := commonUi()
	if ui.Label() != "Test" {
		t.Error("UI returned the wrong label")
	}
}

func TestUi_Description(t *testing.T) {
	ui := commonUi()
	if ui.Description() != "This is a Test" {
		t.Error("UI returned the wrong description")
	}
}

func TestUi_Help(t *testing.T) {
	ui := commonUi()
	if ui.Help() != "In order to use this test do X" {
		t.Error("UI returned the wrong help")
	}
}

func TestUi_Merge(t *testing.T) {
	ui := commonUi()

	// Test overriding values
	merge := NewUi("test.2", "Test 2", "This is another Test", "In order to use this test do Y")
	ui.Merge(merge.Ui())

	if ui.Id() != "test.2" {
		t.Error("Merged UI returned the wrong Id")
	}
	if ui.Label() != "Test 2" {
		t.Error("Merged UI returned the wrong label")
	}
	if ui.Description() != "This is another Test" {
		t.Error("Merged UI returned the wrong description")
	}
	if ui.Help() != "In order to use this test do Y" {
		t.Error("Merged UI returned the wrong help")
	}

	// Test that empty values get ignored.
	merge = NewUi("test.3", "", "", "")
	ui.Merge(merge.Ui())

	if ui.Id() != "test.3" {
		t.Error("Merged UI returned the wrong Id")
	}
	if ui.Label() != "Test 2" {
		t.Error("Merged UI returned the wrong label")
	}
	if ui.Description() != "This is another Test" {
		t.Error("Merged UI returned the wrong description")
	}
	if ui.Help() != "In order to use this test do Y" {
		t.Error("Merged UI returned the wrong help")
	}
}
