package win_m

type WList struct {
	Word  string `json:"word"`
	Count int64  `json:"count"`
}

// getListQuery 검색에 사용할 Query를 생성한다
func (w *WList) getListQuery(word []string) map[string]interface{} {
	query := map[string]interface{}{
		"size": 0,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{"match_all": map[string]interface{}{}},
				},
				"filter": map[string]interface{}{
					"range": map[string]interface{}{
						"createdAt": map[string]interface{}{
							"gte": "now-3h/H",
							"lte": "now",
						},
					},
				},
			},
		},
		"aggs": map[string]interface{}{
			"countByWord": map[string]interface{}{
				"terms": map[string]interface{}{
					"field":   "word",
					"size":    30,
					"exclude": word,
				},
			},
		},
	}

	return query
}

// resultToResponseModel 검색 결과를 응답 모델 값으로 변환한다
func (w *WList) resultToResponseModel(r map[string]interface{}) (resp []WList, err error) {
	// 검색 결과를 반환 모델로 변환한다
	// 인터페이스를 반환하므로 직접 타입 어설션을 해야한다
	// aggregations 값을 변환한다
	aggs := r["aggregations"].(map[string]interface{})["countByWord"].(map[string]interface{})
	buckets := aggs["buckets"].([]interface{})
	// 검색 결과가 없는 경우 빈 모델을 반환한다
	if len(buckets) <= 0 {
		return []WList{}, nil
	}

	// 버킷 내용을 모델로 변환한다
	for _, val := range buckets {
		doc := val.(map[string]interface{})
		resp = append(resp, WList{
			Word:  doc["key"].(string),
			Count: int64(doc["doc_count"].(float64)),
		})
	}

	return resp, nil
}

func (w *WList) List() ([]WList, error) {
	var stopWord WStopWord
	words, err := stopWord.List()

	// words(불용어) 개수가 0이라면 빈 값으로 채워 null_value 에러를 없앤다
	var word []string
	if len(words) > 0 {
		word = words
	} else {
		word = []string{""}
	}

	query := w.getListQuery(word)

	// 검색을 요청한다
	r, err := search(query)
	if err != nil {
		return []WList{}, err
	}

	return w.resultToResponseModel(r)
}
