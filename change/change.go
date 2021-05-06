package change

//RetrieveChange eh foda
func RetrieveChange(value float64) (total int64, values []float64) {
	money_notes := []float64{50, 10, 5, 1}

	totalChange := 0

	for index, element := range money_notes {
		if value < element {
			continue
		}

		if value%element != 0 {

		}
	}

	return 2
}

/*
cédulas de R$100,00, R$50,00, R$10,00, R$5,00 e R$1,00;
moedas de R$0,50, R$0,10, R$0,05 e R$0,01.
*/
