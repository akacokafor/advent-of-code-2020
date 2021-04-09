package main

import (
	"errors"
	"testing"
)

func TestGetRowCount(t *testing.T) {
	inputs := []struct {
		Name      string
		In        string
		Out       int
		WantError bool
		Err       error
	}{
		{
			Name:      "Test FBFBBFF succeeds",
			In:        "FBFBBFF",
			Out:       44,
			WantError: false,
			Err:       nil,
		},
		{
			Name:      "Test BFFFBBF succeeds",
			In:        "BFFFBBF",
			Out:       70,
			WantError: false,
			Err:       nil,
		},
		{
			Name:      "Test FFFBBBF succeeds",
			In:        "FFFBBBF",
			Out:       14,
			WantError: false,
			Err:       nil,
		},
		{
			Name:      "Test BBFFBBF succeeds",
			In:        "BBFFBBF",
			Out:       102,
			WantError: false,
			Err:       nil,
		},
	}

	for _, v := range inputs {
		t.Run(v.Name, func(t *testing.T) {
			result, err := getRowCount(v.In)
			if v.WantError && err == nil {
				t.Fatalf("Expected Error %v, Got Nil", v.Err)
			}

			if v.WantError && err != nil {
				if !errors.Is(err, v.Err) {
					t.Fatalf("Expected Error %v, Got %v", v.Err, err)
				}
			}

			if v.Out != result {
				t.Fatalf("Expected result %v, Got %v", v.Out, result)
			}
		})
	}
}

func TestGetColumnCount(t *testing.T) {
	inputs := []struct {
		Name      string
		In        string
		Out       int
		WantError bool
		Err       error
	}{
		{
			Name:      "Test RRR succeeds",
			In:        "RRR",
			Out:       7,
			WantError: false,
			Err:       nil,
		},
		{
			Name:      "Test RLL succeeds",
			In:        "RLL",
			Out:       4,
			WantError: false,
			Err:       nil,
		},
	}

	for _, v := range inputs {
		t.Run(v.Name, func(t *testing.T) {
			result, err := getColumnCount(v.In)
			if v.WantError && err == nil {
				t.Fatalf("Expected Error %v, Got Nil", v.Err)
			}

			if v.WantError && err != nil {
				if !errors.Is(err, v.Err) {
					t.Fatalf("Expected Error %v, Got %v", v.Err, err)
				}
			}

			if v.Out != result {
				t.Fatalf("Expected result %v, Got %v", v.Out, result)
			}
		})
	}
}

func TestDecodeBoardingPass(t *testing.T) {
	inputs := []struct {
		Name      string
		In        string
		Row       int
		Column    int
		SeatID    int
		WantError bool
		Err       error
	}{
		{
			Name:      "Test BFFFBBFRRR succeeds",
			In:        "BFFFBBFRRR",
			Row:       70,
			Column:    7,
			SeatID:    567,
			WantError: false,
			Err:       nil,
		},
		{
			Name:      "Test FFFBBBFRRR succeeds",
			In:        "FFFBBBFRRR",
			Row:       14,
			Column:    7,
			SeatID:    119,
			WantError: false,
			Err:       nil,
		},
		{
			Name:      "Test BBFFBBFRLL succeeds",
			In:        "BBFFBBFRLL",
			Row:       102,
			Column:    4,
			SeatID:    820,
			WantError: false,
			Err:       nil,
		},
	}

	for _, v := range inputs {
		t.Run(v.Name, func(t *testing.T) {
			row, col, seatId, err := decodeBoardingPass(v.In)
			if v.WantError && err == nil {
				t.Fatalf("Expected Error %v, Got Nil", v.Err)
			}

			if v.WantError && err != nil {
				if !errors.Is(err, v.Err) {
					t.Fatalf("Expected Error %v, Got %v", v.Err, err)
				}
			}

			if v.Row != row {
				t.Errorf("Expected row %v, Got %v", v.Row, row)
			}

			if v.Column != col {
				t.Errorf("Expected column %v, Got %v", v.Column, col)
			}

			if v.SeatID != seatId {
				t.Errorf("Expected seat id %v, Got %v", v.SeatID, seatId)
			}
		})
	}
}
