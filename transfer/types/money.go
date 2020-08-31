package types

type MT int

const (
	MoneyCNY MT = 0
	MoneyUSD
	MoneyHKD
	MoneySGD
)

func ConvertMoneyType(in string) MT {
	switch in {
	case "CNY":
		return MoneyCNY
	case "USD":
		return MoneyUSD
	case "HKD":
		return MoneyHKD
	case "SGD":
		return MoneySGD
	}
	return MoneyCNY
}

func ConvertMoneyTypeByCode(code int) MT {
	switch code {
	case 156:
		return MoneyCNY
	case 840:
		return MoneyUSD
	case 344:
		return MoneyHKD
	case 702:
		return MoneySGD
	}
	return MoneyCNY
}
