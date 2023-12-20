package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{}
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int{}
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, exists := units[unit]
	if !exists {
		return false
	}
	bill[item] += units[unit]
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	amount, itemExists := bill[item]
	reduce, unitExists := units[unit]
	if !itemExists {
		return false
	}
	if !unitExists {
		return false
	}
	if amount-reduce < 0 {
		return false
	}

	if amount-reduce == 0 {
		delete(bill, item)
	} else {
		bill[item] -= reduce
	}
	return true
	//_, unitExists := units[unit]
	/* 	if !itemExists || !unitExists || len(bill)-1 < 0 {
	   		return false
	   	} else if len(bill)-1 == 0 {
	   		delete(bill, item)
	   		return true
	   	} else {
	   		bill[item] = bill[unit] - 1
	   	}
	   	return true */
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	_, itemExists := bill[item]
	if !itemExists {
		return 0, false
	}
	return bill[item], true
}
