package redis

import "encoding/json"

func ToStruct(data interface{}, v interface{}) error {
	bts, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(bts, v)
}

func ConvertToStruct(v, out interface{}) error {
	var bts []byte
	var err error
	switch v := v.(type) {
	case string:
		bts = []byte(v)
	case []byte:
		bts = v
	default:
		bts, err = json.Marshal(v)
		if err != nil {
			return err
		}
	}
	return json.Unmarshal(bts, out)
}

func ToMap(s interface{}) (map[string]interface{}, error) {
	bts, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	hm := make(map[string]interface{})

	if err := json.Unmarshal(bts, hm); err != nil {
		return nil, err
	}
	return hm, nil
}
