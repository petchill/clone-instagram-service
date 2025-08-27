package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type mediaHandler struct {
}

func NewMediaHandler() *mediaHandler {
	return &mediaHandler{}
}

func (h *mediaHandler) RegisterRoutes(e *echo.Echo) {
	e.POST("/media", h.PostMedia)
}

func (h *mediaHandler) PostMedia(c echo.Context) error {
	file, err := c.FormFile("media") // "file" is the name of the input field in the HTML form
	if err != nil {
		return err
	}
	fmt.Println("media", file)
	src, err := file.Open()
	if err != nil {
		return err
	}
	fmt.Println("src", src)

	defer src.Close()

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
