package service

import (
	"errors"
	"testing"
)

type fakeExpenseStorage struct {
	expenses []Expense
	readErr  error
	writeErr error
	written  []Expense
}

func (f *fakeExpenseStorage) Read() ([]Expense, error) {
	if f.readErr != nil {
		return nil, f.readErr
	}
	return f.expenses, nil
}

func (f *fakeExpenseStorage) Write(expenses []Expense) error {
	if f.writeErr != nil {
		return f.writeErr
	}
	f.written = append([]Expense(nil), expenses...)
	return nil
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		name           string
		expenses       []Expense
		id             string
		amount         int
		description    string
		wantUpdated    bool
		wantWriteCalls bool
	}{
		{
			name: "Updates matching expense",
			expenses: []Expense{
				{ID: "expense-1"},
				{ID: "expense-2"},
				{ID: "expense-3"},
			},
			id:             "expense-2",
			amount:         22,
			description:    "test description updated",
			wantUpdated:    true,
			wantWriteCalls: true,
		},
		{
			name: "Returns false when expense doesn't exist",
			expenses: []Expense{
				{ID: "expense-1"},
			},
			id:             "expense-2",
			amount:         22,
			description:    "test description updated",
			wantUpdated:    false,
			wantWriteCalls: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &fakeExpenseStorage{
				expenses: tt.expenses,
			}
			svc := NewService(store)

			updated, err := svc.Update(tt.id, tt.description, tt.amount)

			if err != nil {
				t.Fatalf("Update() returned unexpected error %v", err)
			}

			if updated != tt.wantUpdated {
				t.Errorf(
					"Update() updated = %v, want %v",
					updated,
					tt.wantUpdated,
				)
			}

			if !tt.wantWriteCalls {
				if store.written != nil {
					t.Errorf("Update() unexpectedly wrote: %#v", store.written)
				}
				return
			}

			if store.written == nil {
				t.Fatal("Update() did not write updated expenses")
			}

			if len(store.written) != len(tt.expenses) {
				t.Fatalf(
					"Update() wrote %d expenses, want %d",
					len(store.written),
					len(tt.expenses),
				)
			}

			for _, expense := range store.written {
				if expense.ID != tt.id {
					continue
				}
				if expense.Description != tt.description {
					t.Errorf(
						"Update() description = %q, want %q",
						expense.Description,
						tt.description,
					)
				}
				if expense.Amount != tt.amount {
					t.Errorf(
						"Update() amount = %d, want %d",
						expense.Amount,
						tt.amount,
					)
				}
				if expense.Date.IsZero() {
					t.Error("Update() did not set the expense date")
				}
				break
			}

			if store.written[0].ID != "expense-1" || store.written[2].ID != "expense-3" {
				t.Error("Update() changed an expense other than the requested one")
			}
		})
	}
}

func TestUpdateValidatesInput(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		description string
		amount      int
	}{
		{
			name:        "Rejects empty ID",
			description: "description",
			amount:      100,
		},
		{
			name:   "Rejects empty description",
			id:     "expense-1",
			amount: 100,
		},
		{
			name:        "Rejects invalid amount",
			id:          "expense-1",
			description: "description",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &fakeExpenseStorage{
				expenses: []Expense{{ID: "expense-1"}},
			}
			svc := NewService(store)

			updated, err := svc.Update(tt.id, tt.description, tt.amount)

			if updated {
				t.Error("Update() updated = true for invalid input")
			}
			if err == nil {
				t.Error("Update() error = nil for invalid input")
			}
			if store.written != nil {
				t.Errorf("Update() unexpectedly wrote: %#v", store.written)
			}
		})
	}
}

func TestUpdateReturnsReadError(t *testing.T) {
	wantErr := errors.New("read failed")
	store := &fakeExpenseStorage{
		readErr: wantErr,
	}
	svc := NewService(store)

	updated, err := svc.Update("expense-1", "description", 100)

	if updated {
		t.Error("Update() updated = true when reading failed")
	}
	if !errors.Is(err, wantErr) {
		t.Fatalf("Update() error = %v, want %v", err, wantErr)
	}
}

func TestUpdateReturnsWriteError(t *testing.T) {
	wantErr := errors.New("write failed")
	store := &fakeExpenseStorage{
		expenses: []Expense{{ID: "expense-1"}},
		writeErr: wantErr,
	}
	svc := NewService(store)

	updated, err := svc.Update("expense-1", "description", 100)

	if updated {
		t.Error("Update() updated = true when persistence failed")
	}
	if !errors.Is(err, wantErr) {
		t.Fatalf("Update() error = %v, want %v", err, wantErr)
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name           string
		expenses       []Expense
		id             string
		wantDeleted    bool
		wantIDs        []string
		wantWriteCalls bool
	}{
		{
			name: "Deletes matching expense",
			expenses: []Expense{
				{ID: "expense-1"},
				{ID: "expense-2"},
				{ID: "expense-3"},
			},
			id:             "expense-2",
			wantDeleted:    true,
			wantIDs:        []string{"expense-1", "expense-3"},
			wantWriteCalls: true,
		},
		{
			name: "Returns false when expense doesn't exist",
			expenses: []Expense{
				{ID: "expense-1"},
			},
			id:             "expense-2",
			wantDeleted:    false,
			wantWriteCalls: false,
		},
		{
			name:           "Handles an empty expense list",
			expenses:       nil,
			id:             "expense-1",
			wantDeleted:    false,
			wantWriteCalls: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &fakeExpenseStorage{
				expenses: tt.expenses,
			}
			svc := NewService(store)

			deleted, err := svc.Delete(tt.id)
			if err != nil {
				t.Fatalf("Delete() returned unexpected error: %v", err)
			}

			if deleted != tt.wantDeleted {
				t.Errorf(
					"Delete() deleted = %v, want %v",
					deleted,
					tt.wantDeleted,
				)
			}

			if !tt.wantWriteCalls {
				if store.written != nil {
					t.Errorf("Delete() unexpectedly wrote: %#v", store.written)
				}
			}

			if len(store.written) != len(tt.wantIDs) {
				t.Fatalf(
					"Delete() wrote %d expenses, want %d",
					len(store.written),
					len(tt.wantIDs),
				)
			}

			for i, wantID := range tt.wantIDs {
				if store.written[i].ID != wantID {
					t.Errorf(
						"written expense %d has ID %q, want %q",
						i,
						store.written[i].ID,
						wantID,
					)
				}
			}
		})
	}
}

func TestDeleteReturnsReadError(t *testing.T) {
	wantErr := errors.New("read failed")
	store := &fakeExpenseStorage{
		readErr: wantErr,
	}
	svc := NewService(store)
	deleted, err := svc.Delete("expense-1")

	if deleted {
		t.Error("Delete() deleted = true, want false")
	}

	if !errors.Is(err, wantErr) {
		t.Fatalf("Delete() error = %v, want %v", err, wantErr)
	}
}

func TestDeleteReturnsWriteError(t *testing.T) {
	wantErr := errors.New("write failed")
	store := &fakeExpenseStorage{
		expenses: []Expense{{ID: "expense-1"}},
		writeErr: wantErr,
	}
	svc := NewService(store)

	deleted, err := svc.Delete("expense-1")

	if deleted {
		t.Error("Delete() deleted = true when persistence failed")
	}

	if !errors.Is(err, wantErr) {
		t.Fatalf("Delete() error = %v, want %v", err, wantErr)
	}
}

func TestSummary(t *testing.T) {
	tests := []struct {
		name     string
		expenses []Expense
		want     int
	}{
		{
			name: "Sums all expenses",
			expenses: []Expense{
				{Amount: 1205},
				{Amount: 795},
				{Amount: 99},
			},
			want: 2099,
		},
		{
			name:     "Returns zero for an empty list",
			expenses: nil,
			want:     0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &fakeExpenseStorage{expenses: tt.expenses}
			svc := NewService(store)

			got, err := svc.Summary()
			if err != nil {
				t.Fatalf("Summary() returned unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("Summary() = %d, want %d", got, tt.want)
			}
			if store.written != nil {
				t.Errorf("Summary() unexpectedly wrote: %#v", store.written)
			}
		})
	}
}

func TestSummaryReturnsReadError(t *testing.T) {
	wantErr := errors.New("read failed")
	store := &fakeExpenseStorage{readErr: wantErr}
	svc := NewService(store)

	sum, err := svc.Summary()

	if sum != 0 {
		t.Errorf("Summary() = %d when reading failed, want 0", sum)
	}
	if !errors.Is(err, wantErr) {
		t.Fatalf("Summary() error = %v, want %v", err, wantErr)
	}
}
