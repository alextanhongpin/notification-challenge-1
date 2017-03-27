package common

import (
	"github.com/alextanhongpin/notification-challenge/model"
)

// UpdateCache takes the old cache, compare with the current data,
// and return a new cache
func UpdateCache(oldCache map[int]bool, similar []model.Repository, notification []model.Repository) map[int]bool {
	var cache map[int]bool
	cache = make(map[int]bool)
	similar = append(similar, notification...)

	if len(similar) == 0 {
		return oldCache
	}
	for _, v := range similar {
		cache[v.ID] = true
	}
	return cache
}

// MakeNotificationPayload compares the data with the cache
// and only return new items that will be used as the notification
// payload
func MakeNotificationPayload(cache map[int]bool, repositories []model.Repository) []model.Repository {
	var notificationPayload []model.Repository
	if repositories == nil || len(repositories) == 0 {
		return notificationPayload
	}
	for _, v := range repositories {
		_, ok := cache[v.ID]
		if !ok {
			notificationPayload = append(notificationPayload, v)
		}
	}
	return notificationPayload
}

// GetSimilarData checks if the cache already contains the data, and
// returns it
func GetSimilarData(cache map[int]bool, repositories []model.Repository) []model.Repository {
	var similarData []model.Repository
	if repositories == nil || len(repositories) == 0 {
		return similarData
	}
	for _, v := range repositories {
		_, ok := cache[v.ID]
		if ok {
			similarData = append(similarData, v)
		}
	}
	return similarData
}
