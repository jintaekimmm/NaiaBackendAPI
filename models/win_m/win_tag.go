package win_m

type WTag struct {
	Tag     string  `json:"tag"`
	Percent float64 `json:"percent"`
}

// getTagPercentQuery 검색에 사용할 Query를 생성한다
func (w *WTag) getTagPercentQuery(word string) map[string]interface{} {
	query := map[string]interface{}{
		"size": 0,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"match": map[string]string{
							"word": word,
						},
					},
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
			"countPerTag": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "tag",
					"size":  10,
				},
			},
		},
	}

	return query
}

// resultToResponseModel 검색 결과를 응답 모델 값으로 변환한다
func (w *WTag) resultToResponseModel(r map[string]interface{}) (resp []WTag, err error) {
	// 검색 결과를 반환 모델로 변환한다
	// 인터페이스를 반환하므로 직접 타입 어설션을 해야한다
	// 전체 문서 수를 변환한다
	total := r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)

	// aggregations 값을 변환한다
	aggs := r["aggregations"].(map[string]interface{})["countPerTag"].(map[string]interface{})
	buckets := aggs["buckets"].([]interface{})

	if len(buckets) <= 0 {
		return []WTag{}, nil
	}

	// 버킷 내용을 모델로 변환한다
	// 태그(Key)별 카운트로 전체 카운트 대비 퍼센트를 계산한다
	for _, val := range buckets {
		doc := val.(map[string]interface{})
		resp = append(resp, WTag{
			Tag:     doc["key"].(string),
			Percent: doc["doc_count"].(float64) / total * 100,
		})
	}

	return resp, nil
}

// WordToTagPercent 이슈 Word의 태그별 점유율을 반환한다
func (w *WTag) WordToTagPercent(word string) ([]WTag, error) {
	query := w.getTagPercentQuery(word)

	r, err := search(query)
	if err != nil {
		return nil, err
	}

	return w.resultToResponseModel(r)
}
