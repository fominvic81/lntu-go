package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mailru/easyjson"
	"slices"
	"strconv"

	"github.com/fominvic81/lntu-go/lab_07/notes"
)

func main() {
	app := fiber.New()

	notesArr := notes.Notes{}

	app.Get("/notes", func(c fiber.Ctx) error {
		response, err := easyjson.Marshal(notesArr)
		if err != nil {
			return err
		}

		if err := c.SendString(string(response)); err != nil {
			return err
		}
		return nil
	})

	app.Get("/notes/:id", func(c fiber.Ctx) error {
		id := c.Params("id")
		for _, note := range notesArr {
			if strconv.FormatInt(int64(note.Id), 10) == id {
				response, err := easyjson.Marshal(note)
				if err != nil {
					return err
				}
				if err := c.SendString(string(response)); err != nil {
					return err
				}
				return nil
			}
		}
		if err := c.Status(404).SendString("note not found"); err != nil {
			return err
		}
		return nil
	})

	app.Post("/notes", func(c fiber.Ctx) error {
		note := notes.Note{}
		if err := easyjson.Unmarshal(c.Body(), &note); err != nil {
			return err
		}
		notesArr = append(notesArr, note)

		response, err := easyjson.Marshal(note)
		if err != nil {
			return err
		}

		if err := c.SendString(string(response)); err != nil {
			return err
		}
		return nil
	})

	app.Put("/notes/:id", func(c fiber.Ctx) error {
		id := c.Params("id")
		for i, note := range notesArr {
			if strconv.FormatInt(int64(note.Id), 10) == id {
				newNote := notes.Note{}
				if err := easyjson.Unmarshal(c.Body(), &newNote); err != nil {
					return err
				}
				if newNote.Id != note.Id {
					if err := c.Status(400).SendString("Can not modify note id"); err != nil {
						return err
					}
				}
				notesArr[i] = newNote
				response, err := easyjson.Marshal(newNote)
				if err != nil {
					return err
				}
				if err := c.SendString(string(response)); err != nil {
					return err
				}
				return nil
			}
		}
		if err := c.Status(404).SendString("note not found"); err != nil {
			return err
		}

		return nil
	})

	app.Delete("/notes/:id", func(c fiber.Ctx) error {
		id := c.Params("id")
		notesArr = slices.DeleteFunc(notesArr, func(note *notes.Note) bool {
			return strconv.FormatInt(int64(note.Id), 10) == id
		})
		return nil
	})

	app.Listen(":6969")
}
