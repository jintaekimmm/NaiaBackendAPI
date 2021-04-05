package services

import (
	"github.com/99-66/NaiaBackendApi/models"
	"github.com/99-66/NaiaBackendApi/repositories"
)

type WINService struct {
	WINRepository repositories.WINRepository
}

func ProvideWINService(w repositories.WINRepository) WINService {
	return WINService{WINRepository: w}
}

// List 현재시간 기준 6시간 전까지의 상위 이슈 30개를 반환한다
func (w *WINService) List() ([]models.WinList, error){
	// List는 맵 인터페이스를 반환한다
	r, err := w.WINRepository.List()
	if err != nil {
		return nil, err
	}

	// 검색 결과를 반환 모델로 변환한다
	// 인터페이스를 반환하므로 직접 타입 어설션을 해야한다
	// aggregations 값을 변환한다
	aggs := r["aggregations"].(map[string]interface{})["countByWord"].(map[string]interface{})
	buckets := aggs["buckets"].([]interface{})
	// 검색 결과가 없는 경우 빈 모델을 반환한다
	if len(buckets) <= 0 {
		return []models.WinList{}, nil
	}

	var winResp []models.WinList
	// 버킷 내용을 모델로 변환한다
	for _, val := range buckets {
		doc := val.(map[string]interface{})
		winResp = append(winResp, models.WinList{
			Word: doc["key"].(string),
			Count: int64(doc["doc_count"].(float64)),
		})
	}

	return winResp, nil
}


// FindWordToTagPercent 이슈 Word의 태그별 점유율을 반환한다
func (w *WINService) FindWordToTagPercent(word string) ([]models.WinTag, error){
	// FindWordToTagPercent는 맵 인터페이스를 반환한다
	r, err := w.WINRepository.FindWordToTagPercent(word)
	if err != nil {
		return nil, err
	}

	// 검색 결과를 반환 모델로 변환한다
	// 인터페이스를 반환하므로 직접 타입 어설션을 해야한다
	// 전체 문서 수를 변환한다
	total := r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)

	// aggregations 값을 변환한다
	aggs := r["aggregations"].(map[string]interface{})["countPerTag"].(map[string]interface{})
	buckets := aggs["buckets"].([]interface{})

	if len(buckets) <= 0 {
		return []models.WinTag{}, nil
	}

	var tagResp []models.WinTag
	// 버킷 내용을 모델로 변환한다
	// 태그(Key)별 카운트로 전체 카운트 대비 퍼센트를 계산한다
	for _, val := range buckets {
		doc := val.(map[string]interface{})
		tagResp = append(tagResp, models.WinTag{
			Tag: doc["key"].(string),
			Percent: doc["doc_count"].(float64) / total * 100,
		})
	}

	return tagResp, nil
}