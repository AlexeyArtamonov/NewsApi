package newsHandler

import (
	"news_api/database"
	"news_api/model"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/reform.v1"
)

func GetNews(c *fiber.Ctx) error {
	news, err := database.DB.SelectAllFrom(model.NewsTable, "")
	if err != nil {
		c.Context().Logger().Printf("Get News Handler: Failed - %v", err.Error())
		return err
	}

	err = fillCategories(&news)
	if err != nil {
		c.Context().Logger().Printf("Get News Handler: Failed - %v", err.Error())
		return err
	}
	return c.Status(200).JSON(fiber.Map{"News": news, "Success": true})
}

func fillCategories(news *[]reform.Struct) error {
	for _, element := range *news {
		row := element.(*model.News)
		reform_categories, err := database.DB.FindAllFrom(model.NewscategoriesView, "newsid", row.ID)

		if err != nil {
			return err
		}
		row.Categories = make([]int64, 0)
		for _, elemet := range reform_categories {
			row.Categories = append(row.Categories, elemet.(*model.Newscategories).Categoryid)
		}
	}
	return nil
}

func PostNews(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("Id", 0)

	var updatedNews model.JSONNews
	err := c.BodyParser(&updatedNews)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Incorrect input"})
	}

	news, err := database.DB.FindByPrimaryKeyFrom(model.NewsTable, id)

	// Insert
	if id == 0 || news == nil || err == reform.ErrNoRows {
		insertCategories(updatedNews.Id, updatedNews.Categories)

		insertNews(updatedNews)

		return c.SendStatus(200)
	}
	// Update
	if news != nil {
		new_news := news.(*model.News)
		if len(updatedNews.Categories) != 0 {
			updateCategories(new_news.ID, updatedNews.Id, updatedNews.Categories)
		}
		updateNews(new_news, updatedNews)

		return c.SendStatus(200)
	}

	return nil
}

func insertNews(news model.JSONNews) {
	var new_news = new(model.News)

	new_news.ID = news.Id
	new_news.Content = news.Content
	new_news.Title = news.Title

	database.DB.Save(new_news)
}
func updateNews(old *model.News, new model.JSONNews) {
	if len(new.Content) != 0 {
		old.Content = new.Content
	}

	if len(new.Title) != 0 {
		old.Title = new.Title
	}

	if new.Id != 0 {
		old.ID = new.Id
	}
	database.DB.Update(old)
}
func insertCategories(news_id int64, categories []int64) {
	for _, element := range categories {
		var category = new(model.Newscategories)

		category.Categoryid = element
		category.Newsid = news_id

		database.DB.Insert(category)
	}
}
func updateCategories(old_id int64, new_id int64, categories []int64) {
	database.DB.DeleteFrom(model.NewscategoriesView, "newsid", old_id)

	for _, element := range categories {
		var category = new(model.Newscategories)

		category.Categoryid = element
		category.Newsid = new_id

		database.DB.Insert(category)
	}
}
