package services

import "github.com/99-66/NaiaBackendApi/models/win"

// weeklyWordCountQuery 주간 단어개수 검색에 사용할 Query를 생성한다
func weeklyWordCountQuery() map[string]interface{} {
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

// weeklyWordCountResponseModel 검색 결과를 응답 모델 값으로 변환한다
func weeklyWordCountResponseModel(r map[string]interface{}) (resp win.WordCount, err error) {
	// 검색 결과를 반환 모델로 변환한다
	// 인터페이스를 반환하므로 직접 타입 어설션을 해야한다
	// aggregations 값을 변환한다
	aggs := r["aggregations"].(map[string]interface{})["wordsOverTime"].(map[string]interface{})
	buckets := aggs["buckets"].([]interface{})

	if len(buckets) <= 0 {
		return win.WordCount{}, nil
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

	return win.WordCount{
		Date:      date,
		Total:     total,
		SNS:       sns,
		Article:   article,
		Community: community,
	}, nil
}

// WeeklyWordCount 주간 단어개수를 검색하여 반환한다
func WeeklyWordCount() (win.WordCount, error) {
	query := weeklyWordCountQuery()

	r, err := search(query)
	if err != nil {
		return win.WordCount{}, err
	}

	return weeklyWordCountResponseModel(r)
}

type WordCloudList []interface{}

// wordCloudQuery 검색에 사용할 Query를 생성한다
func wordCloudQuery(word []string, count int, filter string) map[string]interface{} {
	var query map[string]interface{}

	if filter == "all" || filter == "" {
		query = map[string]interface{}{
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
						"size":    count,
						"exclude": word,
					},
				},
			},
		}
	} else {
		query = map[string]interface{}{
			"size": 0,
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"must": []map[string]interface{}{
						{
							"match": map[string]interface{}{
								"tag": filter,
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
				"countByWord": map[string]interface{}{
					"terms": map[string]interface{}{
						"field":   "word",
						"size":    30,
						"exclude": word,
					},
				},
			},
		}
	}

	return query
}

// wordCloudResponseModel 검색 결과를 응답 모델 값으로 변환한다
func wordCloudResponseModel(r map[string]interface{}) (resp WordCloudList, err error) {
	// 검색 결과를 반환 모델로 변환한다
	// 인터페이스를 반환하므로 직접 타입 어설션을 해야한다
	// aggregations 값을 변환한다
	aggs := r["aggregations"].(map[string]interface{})["countByWord"].(map[string]interface{})
	buckets := aggs["buckets"].([]interface{})
	// 검색 결과가 없는 경우 빈 모델을 반환한다
	if len(buckets) <= 0 {
		return WordCloudList{}, nil
	}

	// 버킷 내용을 모델로 변환한다
	for _, val := range buckets {
		doc := val.(map[string]interface{})
		resp = append(resp, WordCloudList{doc["key"].(string), doc["doc_count"].(float64)})
	}

	return resp, nil
}

// WordCloud 워드 클라우드에 사용할 데이터를 검색하여 반환한다
func WordCloud(count int, filter string) (WordCloudList, error) {
	words, err := StopWords()

	// words(불용어) 개수가 0이라면 빈 값으로 채워 null_value 에러를 없앤다
	var word []string
	if len(words) > 0 {
		word = words
	} else {
		word = []string{""}
	}

	query := wordCloudQuery(word, count, filter)

	// 검색을 요청한다
	r, err := search(query)
	if err != nil {
		return WordCloudList{}, err
	}

	return wordCloudResponseModel(r)
}
