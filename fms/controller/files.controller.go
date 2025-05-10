package controller

import (
	"fms/config"
	"fms/models"
	"fms/utils"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	folderId := c.FormValue("folderId")
	if err != nil || folderId == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get user file"})
	}
	var filename string = utils.GenerateID()
	if err = c.SaveFile(file, "./upload/"+filename); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to save file"})
	}

	var newFile models.Files = models.Files{
		Id:       utils.GenerateID(),
		Name:     file.Filename,
		FolderId: folderId,
	}

	res := config.DB.Create(newFile)
	if res.Error != nil {
		os.Remove("./upload/" + filename)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to save file"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "file saved", "filename": filename})
}

func GetAllFiles(c *fiber.Ctx) error {
	params := c.AllParams()
	folderId := params["folderId"]

	var files []models.Files
	res := config.DB.Where("folder_id = ?", folderId).Find(&files)
	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch all files"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "successfully fetch files", "files": files})
}
