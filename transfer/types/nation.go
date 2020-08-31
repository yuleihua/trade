package types

type NT int

const (
	NationCN NT = 0
	NationUS
	NationHK
	NationSG
)

const DefaultNation = "CN"

func ConvertNationType(in string) NT {
	switch in {
	case "CN":
		return NationCN
	case "US":
		return NationUS
	case "HK":
		return NationHK
	case "SG":
		return NationSG
	}
	return NationCN
}

func ConvertNationTypeByCode(code int) NT {
	switch code {
	case 86:
		return NationCN
	case 1:
		return NationUS
	case 852:
		return NationHK
	case 65:
		return NationSG
	}
	return NationCN
}
