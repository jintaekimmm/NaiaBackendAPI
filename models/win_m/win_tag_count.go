package win_m

type WTagCount struct {
	Tag   string `json:"tag"`
	Count int64  `json:"count"`
}

// getTagCountQuery 검색에 사용할 Query를 생성한다
func (w *WTagCount) getTagCountQuery() map[string]interface{} {
	query := map[string]interface{}{
		"size": 0,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{"match_all": map[string]string{}},
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
					"field": "tag",
				},
			},
		},
	}

	return query
}

// resultToResponseModel 검색 결과를 응답 모델 값으로 변환한다
func (w *WTagCount) resultToResponseModel(r map[string]interface{}) (resp []WTagCount, err error) {
	// 검색 결과를 반환 모델로 변환한다
	// 인터페이스를 반환하므로 직접 타입 어설션을 해야한다
	// aggregations 값을 변환한다
	aggs := r["aggregations"].(map[string]interface{})["countByWord"].(map[string]interface{})
	buckets := aggs["buckets"].([]interface{})

	if len(buckets) <= 0 {
		return []WTagCount{}, nil
	}

	// 버킷 내용을 모델로 변환한다
	// 태그(Key)별 카운트로 전체 카운트 대비 퍼센트를 계산한다
	var total int64
	for _, val := range buckets {
		doc := val.(map[string]interface{})
		resp = append(resp, WTagCount{
			Tag:   doc["key"].(string),
			Count: int64(doc["doc_count"].(float64)),
		})
		total += int64(doc["doc_count"].(float64))
	}

	resp = append(resp, WTagCount{
		Tag:   "all",
		Count: total,
	})

	return resp, nil
}

func (w *WTagCount) Counts() ([]WTagCount, error) {
	query := w.getTagCountQuery()

	r, err := search(query)
	if err != nil {
		return []WTagCount{}, err
	}

	return w.resultToResponseModel(r)
}
