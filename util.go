package youtube_go

func Filter(dictionary map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for key, value := range dictionary {
		if value != nil {
			switch v := value.(type) {
			case *int:
				if v != nil {
					result[key] = value
				}
			case *string:
				if v != nil {
					result[key] = value
				}
			// 你可以继续添加其他类型的判断
			default:
				result[key] = value
			}
		}
	}
	return result
}

func Contextualise(clientContext ClientContext, data map[string]interface{}) map[string]interface{} {
	if _, ok := data["context"]; !ok {
		data["context"] = make(map[string]interface{})
	}
	if _, ok := data["context"].(map[string]interface{})["client"]; !ok {
		data["context"].(map[string]interface{})["client"] = make(map[string]interface{})
	}

	clientData := clientContext.Context()
	for key, value := range clientData {
		data["context"].(map[string]interface{})["client"].(map[string]interface{})[key] = value
	}

	return data
}
