package tasks

import "barrage_video_website/cache"

func RestartDailyRank() error {
	return cache.RedisNewClient.Del(cache.Ctx, cache.RankDailyKey).Err()
}
