package controller

import (
	"encoding/json"
	"fms/config"
	"fms/models"
	"fms/utils"
	"github.com/gofiber/fiber/v2"
)

func FetchFolders(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)

	var folder []models.Folders

	res := config.DB.Where("user_id = ?", userId).Find(&folder)
	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch user folders"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "fetched all folders", "folders": folder})
}

func CreateFolder(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)

	params := c.AllParams()

	folderName := params["name"]

	//create a new folder of the given name

	var folder models.Folders

	folder.Name = folderName
	folder.Id = utils.GenerateID()
	folder.UserId = userId

	res := config.DB.Create(folder)
	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create new folder"})
	}

	jsonFolder, _ := json.Marshal(folder)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "created folder", "folder": jsonFolder})
}

type FolderBody struct {
	FolderId string
	Name     string
}

func RenameFolder(c *fiber.Ctx) error {

	// userId = c.Locals("userId").(string)
	var body FolderBody

	c.BodyParser(&body)

	var folder models.Folders
	res := config.DB.Where("id = ?", body.FolderId).First(&folder)
	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch folder"})
	}

	folder.Name = body.Name

	config.DB.Save(&folder)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "folder name has been updated", "folder": folder})
}

func DeleteFolder(c *fiber.Ctx) error {
	params := c.AllParams()
	folderId := params["folderID"]

	res := config.DB.Where("id = ?", folderId).Delete(models.Folders{})
	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to delete folder"})
	}

	return c.SendStatus(200)
}
