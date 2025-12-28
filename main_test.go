package main

import "testing"

func TestCalculateTotal(t *testing.T) {
	item_count := 10
	item_price := 2.5
	total_amount := CalculateTotal(item_count, item_price)
	expected_amount := float64(item_count) * item_price
	if total_amount == 0 || (total_amount != expected_amount) {
		t.Errorf("Expected Amount: %f is not equal to total amount: %f", expected_amount, total_amount)
		return
	}
}
