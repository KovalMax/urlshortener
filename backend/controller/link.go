package controller

import (
    "net/http"
    "net/url"

    jsonIterator "github.com/json-iterator/go"
    "github.com/labstack/echo"

    "github.com/KovalMax/urlshortener/dto"
    "github.com/KovalMax/urlshortener/service"
)

func GetLink(ctx echo.Context) error {
    shortLink := ctx.Param("alias")
    if len(shortLink) < 1 {
        ctx.Logger().Warn("Wrong request")
        return ctx.String(http.StatusBadRequest, "Wrong request")
    }

    result := service.GetLinkService().GetLinkByHash(shortLink)
    if result.Err != nil {
        ctx.Logger().Warn("Not found")
        return ctx.String(http.StatusNotFound, "Not found")
    }

    return ctx.Redirect(http.StatusSeeOther, result.Val)
}

func CreateLink(ctx echo.Context) error {
    link := new(dto.Link)
    err := jsonIterator.NewDecoder(ctx.Request().Body).Decode(&link)
    if err != nil {
        ctx.Logger().Warn("Error decode:", err)
        return ctx.String(http.StatusBadRequest, "Malformed request")
    }

    if err := checkUrl(link); err != nil {
        ctx.Logger().Warn("Invalid url parsed: ", err)
        return ctx.String(http.StatusBadRequest, "Url is invalid")
    }

    result := service.GetLinkService().CreateNewLink(link)
    if result.Err != nil {
        ctx.Logger().Warn("Error with create", result.Err)
        return ctx.String(http.StatusInternalServerError, "error")
    }

    return ctx.JSON(http.StatusCreated, result)
}

func checkUrl(linkDto *dto.Link) error {
    parsedUrl, err := url.Parse(linkDto.Url)
    if err != nil {
        return err
    }

    if parsedUrl.Scheme == "" {
        parsedUrl.Scheme = "http"
        linkDto.Url = parsedUrl.String()
    }

    return err
}
