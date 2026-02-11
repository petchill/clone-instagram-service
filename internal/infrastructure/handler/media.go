package handler

import (
	"fmt"
	"net/http"

	mMedia "clone-instagram-service/internal/domain/model/media"
	eUser "clone-instagram-service/internal/domain/model/user/entity"

	"github.com/labstack/echo/v4"
)

type mediaHandler struct {
	mediaService mMedia.MediaService
}

func NewMediaHandler(mediaService mMedia.MediaService) *mediaHandler {
	return &mediaHandler{
		mediaService,
	}
}

func (h *mediaHandler) RegisterRoutes(e *echo.Group) {
	e.POST("/media", h.PostMedia)
}

func (h *mediaHandler) PostMedia(c echo.Context) error {

	user, ok := c.Get("user").(eUser.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, "invalid user type")
	}
	fmt.Println(user)

	file, err := c.FormFile("media") // "file" is the name of the input field in the HTML form
	if err != nil {
		fmt.Println("err ", err)
		return err
	}

	caption := c.FormValue("caption")
	fmt.Println("caption ", caption)

	src, err := file.Open()
	if err != nil {
		fmt.Println("err ", err)
		return err
	}

	defer src.Close()

	ctx := c.Request().Context()

	err = h.mediaService.CreateAndStoreMedia(ctx, user.ID, file.Filename, src, caption)
	if err != nil {
		fmt.Errorf("This is error", err)
		fmt.Println("err ", err)
		return c.JSON(500, map[string]string{"status": "ERROR"})
	}
	// // Destination
	// dst, err := os.Create(file.Filename) // Save with original filename
	// if err != nil {
	// 	return err
	// }
	// defer dst.Close()

	// // Copy
	// if _, err = io.Copy(dst, src); err != nil {
	// 	return err
	// }

	return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
}
