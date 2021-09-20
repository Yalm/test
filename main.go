package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yalm/cloud-messaging/database"
	"github.com/yalm/cloud-messaging/models"
	"github.com/yalm/cloud-messaging/repositories"
	"github.com/yalm/cloud-messaging/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func captureClick(eventRepository repositories.EventRepository, id primitive.ObjectID) {
	event := models.Event{
		LinkId:    id,
		Name:      "dynamic_link_open",
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}
	eventRepository.InsertOne(&event)
}

func renderDynamicLink(
	dynamicLinkRepository repositories.DynamicLinkRepository,
	eventRepository repositories.EventRepository,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		suffix := c.Param("suffix")
		userAgent := c.Request.Header.Get("User-Agent")

		if utils.IsRenderParam(suffix) {
			c.HTML(http.StatusOK, "dynamic_link.tmpl", gin.H{
				"socialTitle":        c.Request.URL.Query().Get("socialTitle"),
				"socialDescription":  c.Request.URL.Query().Get("socialDescription"),
				"socialImageUrl":     c.Request.URL.Query().Get("socialImageUrl"),
				"androidPackageName": c.Request.URL.Query().Get("androidPackageName"),
				"androidUrl":         c.Request.URL.Query().Get("androidUrl"),
				"link":               c.Request.URL.Query().Get("link"),
				"iosUrl":             c.Request.URL.Query().Get("iosUrl"),
			})
			return
		}

		dynamicLink, err := dynamicLinkRepository.FindByHostnameAndSuffix(c.Request.Host, suffix)

		if err != nil {
			c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		if dynamicLink == nil {
			c.String(http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}

		if utils.IsBot(userAgent) && dynamicLink.DynamicLinkInfo.NavigationInfo.EnableForcedRedirect == false {
			c.Redirect(http.StatusFound, utils.GenerateSocialTag(dynamicLink))
		} else {
			link, errUtm := utils.AddUtm(dynamicLink)

			if errUtm != nil {
				c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
				return
			}

			go captureClick(eventRepository, dynamicLink.ID)
			c.Redirect(http.StatusFound, link)
		}
	}
}

func main() {
	database := database.ConnectToDatabase(os.Getenv("MONGODB_URI"))
	repository := repositories.Create(database)
	dynamicLinkRepository := repository.GetDynamicLinkRepository()
	eventRepository := repository.GetEventRepository()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.StaticFile("/favicon.ico", "resources/favicon.ico")
	r.GET("/:suffix", renderDynamicLink(dynamicLinkRepository, eventRepository))
	r.Run()
}
