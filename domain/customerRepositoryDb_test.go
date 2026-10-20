package domain

import "testing"

func TestDatabaseStatus(t *testing.T) {
	tests := []struct {
		name   string
		status CustomerStatus
		want   string
		ok     bool
	}{
		{name: "active", status: CustomerStatusActive, want: "1", ok: true},
		{name: "inactive", status: CustomerStatusInactive, want: "0", ok: true},
		{name: "unsupported", status: CustomerStatus("unknown")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, ok := databaseStatus(test.status)
			if got != test.want || ok != test.ok {
				t.Errorf("databaseStatus(%q) = (%q, %t), want (%q, %t)", test.status, got, ok, test.want, test.ok)
			}
		})
	}
}

func TestCustomerFromDBRecord(t *testing.T) {
	tests := []struct {
		name       string
		record     customerDBRecord
		wantStatus CustomerStatus
		ok         bool
	}{
		{name: "active", record: customerDBRecord{Status: "1"}, wantStatus: CustomerStatusActive, ok: true},
		{name: "inactive", record: customerDBRecord{Status: "0"}, wantStatus: CustomerStatusInactive, ok: true},
		{name: "unsupported", record: customerDBRecord{Status: "2"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			customer, ok := customerFromDBRecord(test.record)
			if ok != test.ok {
				t.Fatalf("customerFromDBRecord() ok = %t, want %t", ok, test.ok)
			}
			if ok && customer.Status != test.wantStatus {
				t.Errorf("customer status = %q, want %q", customer.Status, test.wantStatus)
			}
		})
	}
}
