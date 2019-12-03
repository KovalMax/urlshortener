package service

import (
    "github.com/lithammer/shortuuid"

    "github.com/KovalMax/urlshortener/dto"
    "github.com/KovalMax/urlshortener/redis"
)

type LinkService struct {
    repository redis.Repository
}

func GetLinkService() LinkService {
    return LinkService{redis.GetRepository(redis.GetConfig())}
}

func (s LinkService) GetLinkByHash(hash string) (result dto.GetLinkResult) {
    defer func() {
        err := s.repository.Close()
        if err != nil {
            result = dto.GetLinkResult{Err: err}
        }
    }()

    val, err := s.repository.GetValue(hash)
    if err != nil {
        return dto.GetLinkResult{Err: err}
    }

    return dto.GetLinkResult{Val: val}
}

func (s LinkService) CreateNewLink(link *dto.Link) (result dto.CreateLinkResult) {
    defer func() {
        err := s.repository.Close()
        if err != nil {
            result = dto.CreateLinkResult{Err: err}
        }
    }()

    isLinkExist, err := s.repository.IsKeyExists(link.Url)
    if err != nil {
        return dto.CreateLinkResult{Err: err}
    }

    if isLinkExist {
        return dto.CreateLinkResult{ShortLink: link.Url}
    }

    hash := shortuuid.NewWithNamespace(link.Url)
    isHashExist, err := s.repository.IsKeyExists(hash)
    if err != nil {
        return dto.CreateLinkResult{Err: err}
    }

    if isHashExist {
        return dto.CreateLinkResult{ShortLink: hash}
    }

    err = s.repository.SetValue(hash, link.Url)

    return dto.CreateLinkResult{ShortLink: hash, Err: err}
}

func (s LinkService) UuidAutoShort(link string) string {
    hash := shortuuid.NewWithNamespace(link)

    maxReduce := 0
    for i := 1; i < len(hash); i++ {
        // reduce hash by i symbols each iteration and check existence
        isExist, _ := s.repository.IsKeyExists(hash[i:])

        // till we'll find existing one
        if isExist {
            break
        }
        // If not exists, increase maxReduce var
        maxReduce = i
    }

    return hash[maxReduce:]
}
