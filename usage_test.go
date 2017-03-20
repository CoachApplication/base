package standard

import (
	"testing"

	api "github.com/james-nesbitt/coach-api"
)

func TestExternalOperationUsage_Allows(t *testing.T) {
	us := (&ExternalOperationUsage{}).Usage()

	if !us.Allows(api.UsageOperationPublicView) {
		t.Error("ExternalOperation usage did not Allow Public Viewing")
	}
	if !us.Allows(api.UsageOperationPublicExecute) {
		t.Error("ExternalOperation usage did not Allow Public Exec")
	}
}

func TestInternalOperationUsage_Allows(t *testing.T) {
	us := (&InternalOperationUsage{}).Usage()

	if us.Allows(api.UsageOperationPublicView) {
		t.Error("InternalOperation usage Allowed Public Viewing")
	}
	if us.Allows(api.UsageOperationPublicExecute) {
		t.Error("InternalOperation usage Allowed Public Exec")
	}
}

func TestOptionalPropertyUsage_Allows(t *testing.T) {
	us := (&OptionalPropertyUsage{}).Usage()

	if !us.Allows(api.UsagePropertyPublicView) {
		t.Error("OptionalPropertyUsage usage did not Allow Public Viewing")
	}
	if !us.Allows(api.UsagePropertyPublicWrite) {
		t.Error("OptionalPropertyUsage usage did not Allow Public Writing")
	}
	if us.Allows(api.UsagePropertyPublicRequired) {
		t.Error("OptionalPropertyUsage usage was marked Required")
	}
}

func TestReadonlyPropertyUsage_Allows(t *testing.T) {
	us := (&ReadonlyPropertyUsage{}).Usage()

	if !us.Allows(api.UsagePropertyPublicView) {
		t.Error("ReadonlyPropertyUsage usage did not Allow Public Viewing")
	}
	if us.Allows(api.UsagePropertyPublicWrite) {
		t.Error("ReadonlyPropertyUsage usage Allowed Public Writing")
	}
}

func TestRequiredPropertyUsage_Allows(t *testing.T) {
	us := (&RequiredPropertyUsage{}).Usage()

	if !us.Allows(api.UsagePropertyPublicView) {
		t.Error("OptionalPropertyUsage usage did not Allow Public Viewing")
	}
	if !us.Allows(api.UsagePropertyPublicWrite) {
		t.Error("OptionalPropertyUsage usage did not Allow Public Writing")
	}
	if !us.Allows(api.UsagePropertyPublicRequired) {
		t.Error("OptionalPropertyUsage usage was not marked Required")
	}
}
