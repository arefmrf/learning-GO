package domain

func DefaultCapacity() Capacity {
	return Capacity{
		"male":         0,
		"female":       0,
		"base_person":  0,
		"extra_person": 0,
	}
}

func DefaultPlaceInfo() map[string]map[string]any {
	return map[string]map[string]any{
		"country":  {},
		"province": {},
		"city":     {},
	}
}

func DefaultHostInfo() map[string]any {
	return map[string]any{
		"draft": map[string]any{
			"capacity": DefaultCapacity(),
			"prices": map[string]int{
				"price":       0,
				"extra_price": 0,
			},
			"dailies": map[string]int64{
				"start": 0,
				"end":   0,
			},
			"reject_reason": map[string]any{},
		},
	}
}
