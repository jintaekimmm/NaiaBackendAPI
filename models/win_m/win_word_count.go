package win_m

type WWordCount struct {
	Date      []string `json:"date"`
	Total     []int64  `json:"total"`
	SNS       []int64  `json:"sns"`
	Article   []int64  `json:"article"`
	Community []int64  `json:"community"`
}

// getWeeklyWordCountQuery 검색에 사용할 Query를 생성한다
func (w *WWordCount) getWeeklyWordCountQuery() map[string]interface{} {
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
							"gte": "now-7d/d",
							"lte": "now/d",
						},
					},
				},
			},
		},
		"aggs": map[string]interface{}{
			"wordsOverTime": map[string]interface{}{
				"date_histogram": map[string]interface{}{
					"field":             "createdAt",
					"calendar_interval": "day",
					"format":            "yyyy-MM-dd",
					"time_zone":         "+09:00",
				},
				"aggs": map[string]interface{}{
					"countByWord": map[string]interface{}{
						"terms": map[string]interface{}{
							"field": "tag",
						},
					},
				},
			},
		},
	}

	return query
}

// resultToResponseModel 검색 결과를 응답 모델 값으로 변환한다
func (w *WWordCount) resultToResponseModel(r map[string]interface{}) (resp WWordCount, err error) {
	// 검색 결과를 반환 모델로 변환한다
	// 인터페이스를 반환하므로 직접 타입 어설션을 해야한다
	// aggregations 값을 변환한다
	aggs := r["aggregations"].(map[string]interface{})["wordsOverTime"].(map[string]interface{})
	buckets := aggs["buckets"].([]interface{})

	if len(buckets) <= 0 {
		return WWordCount{}, nil
	}

	// 버킷 내용을 모델로 변환한다
	var date []string
	var total, sns, article, community []int64
	for _, val := range buckets {
		doc := val.(map[string]interface{})
		date = append(date, doc["key_as_string"].(string))
		total = append(total, int64(doc["doc_count"].(float64)))
		// Sub Aggregations
		subBuckets := doc["countByWord"].(map[string]interface{})["buckets"].([]interface{})
		for _, subVal := range subBuckets {
			subDoc := subVal.(map[string]interface{})
			key := subDoc["key"].(string)
			if key == "sns" {
				sns = append(sns, int64(subDoc["doc_count"].(float64)))
			} else if key == "article" {
				article = append(article, int64(subDoc["doc_count"].(float64)))
			} else if key == "community" {
				community = append(community, int64(subDoc["doc_count"].(float64)))
			}
		}
	}

	return WWordCount{
		Date:      date,
		Total:     total,
		SNS:       sns,
		Article:   article,
		Community: community,
	}, nil
}

func (w *WWordCount) Counts() (WWordCount, error) {
	query := w.getWeeklyWordCountQuery()

	r, err := search(query)
	if err != nil {
		return WWordCount{}, err
	}

	return w.resultToResponseModel(r)
}
