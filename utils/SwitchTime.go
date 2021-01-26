package utils


type Month 		*int

type Year  		*int

type Hour 		*int

type Minute  	*int

type Day  		*int

type Second  	*int


func SwitchForTime(year, month, day, hour, minute, second , year2, month2, day2, hour2, minute2, second2 int, condIF string, switch_type1 string, switch_value int)bool {
	//oldArrayTime := [6]int{year,month,day,hour,minute,second}
	//newArrayTime := [6]int{year2,month2,day2,hour2,minute2,second2}
	// 先计数
	switch switch_type1 {
	case "hour":
		hour2 += switch_value
		break
	case "year":
		year2 += switch_value
		break
	case "month":
		month2 += switch_value
		break
	case "day":
		day2 += switch_value
		break
	case "minute":
		minute2 += switch_value
		break
	case "second":
		second2 += switch_value
		break
	}
	switch condIF {
	case "<":
		switch switch_type1 {
		case "hour":
			if year == year2 {
				if month2 == month {
					if day2 == day {
						if hour < hour2 {
							return true
						}
					}
				}
			}
			break
		case "year":
			if year < year2 {
				return true
			}
			break
		case "month":
			if year2 == year {
				if month < month2 {
					return true
				}
			}
			break
		case "day":
			if year2 == year {
				if month2 == month {
					if day2 < day {
						return true
					}
				}
			}
			break
		case "minute":
			if year2 == year {
				if month2 == month {
					if day2 == day {
						if hour2 == hour {
							if minute < minute2 {
								return true
							}
						}
					}
				}
			}
			break
		case "second":
			if year2 == year {
				if month2 == month {
					if day2 == day {
						if hour2 == hour {
							if minute2 == minute {
								if second < second2 {
									return true
								}
							}
						}
					}
				}
			}
			break
		}
		break
	case ">":
		switch switch_type1 {
		case "hour":
			if year2 == year {
				if month2 == month {
					if day2 == day {
						if hour < hour2 {
							return true
						}
					}
				}
			}
			break
		case "year":
			if year > year2 {
				return true
			}
			break
		case "month":
			if year2 == year {
				if month > month2 {
					return true
				}
			}
			break
		case "day":
			if year2 == year {
				if month2 == month {
					if day > day2 {
						return true
					}
				}
			}
			break
		case "minute":
			if year2 == year {
				if month2 == month {
					if day2 == day {
						if hour2 == hour {
							if minute > minute2 {
								return true
							}
						}
					}
				}
			}
			break
		case "second":
			if year2 == year {
				if month2 == month {
					if day2 == day {
						if hour2 == hour {
							if minute2 == minute {
								if second > second2 {
									return true
								}
							}
						}
					}
				}
			}
			break
		}
		break
	case "<=":
		switch switch_type1 {
		case "hour":
			if year2 == year {
				if month2 == month {
					if day2 == day {
						if hour <= hour2 {
							return true
						}
					}
				}
			}
			break
		case "year":
			if year2 <= year {
				return true
			}
			break
		case "month":
			if year2 == year {
				if month <= month2 {
					return true
				}
			}
			break
		case "day":
			if year2 == year {
				if month2 == month {
					if day <= day2 {
						return true
					}
				}
			}
			break
		case "minute":
			if year2 == year {
				if month2 == month {
					if day2 == day {
						if hour2 == hour {
							if minute <= minute2 {
								return true
							}
						}
					}
				}
			}
			break
		case "second":
			if year2 == year {
				if month2 == month {
					if day2 == day {
						if hour2 == hour {
							if minute2 == minute {
								if second <= second2 {
									return true
								}
							}
						}
					}
				}
			}
			break
		}
		break
	case ">=":
		switch switch_type1 {
		case "hour":
			if year2 == year {
				if month2 == month {
					if day2 == day {
						if hour >= hour2 {
							return true
						}
					}
				}
			}
			break
		case "year":
			if year >= year2 {
				return true
			}
			break
		case "month":
			if year2 == year {
				if month >= month2 {
					return true
				}
			}
			break
		case "day":
			if year2 == year {
				if month2 == month {
					if day >= day2 {
						return true
					}
				}
			}
			break
		case "minute":
			if year2 == year {
				if month2 == month {
					if day2 == day {
						if hour2 == hour {
							if minute >= minute2 {
								return true
							}
						}
					}
				}
			}
			break
		case "second":
			if year2 == year {
				if month2 == month {
					if day2 == day {
						if hour2 == hour {
							if minute2 == minute {
								if second >= second2 {
									return true
								}
							}
						}
					}
				}
			}
			break
		}
		break
	case "==":
		switch switch_type1 {
		case "hour":
			if year2 == year {
				if month2 == month {
					if day2 == day {
						if hour2 == hour {
							return true
						}
					}
				}
			}
			break
		case "year":
			if year == year2 {
				return true
			}
			break
		case "month":
			if year2 == year {
				if month2 == month {
					return true
				}
			}
			break
		case "day":
			if year2 == year {
				if month2 == month {
					if day2 == day {
						return true
					}
				}
			}
			break
		case "minute":
			if year2 == year {
				if month2 == month {
					if day2 == day {
						if hour2 == hour {
							if minute2 == minute {
								return true
							}
						}
					}
				}
			}
			break
		case "second":
			if year2 == year {
				if month2 == month {
					if day2 == day {
						if hour2 == hour {
							if minute2 == minute {
								if second == second2 {
									return true
								}
							}
						}
					}
				}
			}
			break
		}
		break
	case "!=":
		switch switch_type1 {
		case "hour":
			if year2 == year {
				if month2 == month {
					if day2 == day {
						if hour != hour2 {
							return true
						}
					}
				}
			}
			break
		case "year":
			if year != year2 {
				return true
			}
			break
		case "month":
			if year2 == year {
				if month != month2 {
					return true
				}
			}
			break
		case "day":
			if year2 == year {
				if month2 == month {
					if day != day2 {
						return true
					}
				}
			}
			break
		case "minute":
			if year2 == year {
				if month2 == month {
					if day2 == day {
						if hour2 == hour {
							if minute != minute2 {
								return true
							}
						}
					}
				}
			}
			break
		case "second":
			if year2 == year {
				if month2 == month {
					if day2 == day {
						if hour2 == hour {
							if minute2 == minute {
								if second != second2 {
									return true
								}
							}
						}
					}
				}
			}
			break
		}
		break
	}
	return false
}