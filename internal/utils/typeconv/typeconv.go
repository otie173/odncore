package typeconv

func GetByte(v interface{}) byte {
	if val, ok := v.(byte); ok {
		return val
	}
	return 0
}

func GetUint32(v interface{}) uint32 {
	if val, ok := v.(uint32); ok {
		return val
	}
	return 0
}
func GetPtrUint32(v interface{}) *uint32 {
	if val, ok := v.(uint32); ok {
		return &val
	}
	return nil
}

func GetFloat32(v interface{}) float32 {
	if val, ok := v.(float32); ok {
		return val
	}
	return 0
}

func GetBool(v interface{}) bool {
	if val, ok := v.(bool); ok {
		return val
	} else {
		return false
	}
}

func GetString(v interface{}) string {
	if val, ok := v.(string); ok {
		return val
	}
	return ""
}
