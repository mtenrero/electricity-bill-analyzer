package parser

import "testing"

func TestTimeZerosNoZeros(t *testing.T) {

	consumptionNoZero := ConsumptionEntry{
		CUPS:    "87126873678TEST",
		Fecha:   "3/7/2019",
		Hora:    "17",
		Consumo: "2,45",
	}
	consumptionNoZeroTime, _ := consumptionNoZero.TimeConsumption()

	consumptionZero := ConsumptionEntry{
		CUPS:    "TEST2",
		Fecha:   "03/07/2019",
		Hora:    "5",
		Consumo: "1,22",
	}
	consumptionZeroTime, _ := consumptionZero.TimeConsumption()

	if consumptionNoZeroTime.Day() != 3 {
		t.Fail()
	}
	if consumptionZeroTime.Day() != 3 {
		t.Fail()
	}

	if consumptionNoZero.WeekDay() != 3 {
		t.Fail()
	}
	if consumptionZero.WeekDay() != 3 {
		t.Fail()
	}
}
