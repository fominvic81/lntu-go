package main

import (
	"github.com/fominvic81/lntu-go/lab_07/notes"
	"github.com/gofiber/fiber/v3"
	"github.com/mailru/easyjson"
	"maps"
	"net/http"
	"slices"
	"strconv"
)

func main() {
	app := fiber.New()

	notesMap := map[string]notes.Note{}
	lastId := 1

	app.Get("/notes", func(c fiber.Ctx) error {
		response, err := easyjson.Marshal(notes.Notes(slices.Collect(maps.Values(notesMap))))
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
		note, ok := notesMap[id]
		if !ok {
			if err := c.Status(http.StatusNotFound).SendString("note not found"); err != nil {
				return err
			}
			return nil
		}
		response, err := easyjson.Marshal(note)
		if err != nil {
			return err
		}
		if err := c.SendString(string(response)); err != nil {
			return err
		}
		return nil
	})

	app.Post("/notes", func(c fiber.Ctx) error {
		note := notes.Note{}
		if err := easyjson.Unmarshal(c.Body(), &note); err != nil {
			return err
		}
		note.Id = lastId
		lastId += 1

		notesMap[strconv.Itoa(note.Id)] = note

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
		note, ok := notesMap[id]
		if !ok {
			if err := c.Status(http.StatusNotFound).SendString("note not found"); err != nil {
				return err
			}
			return nil
		}
		newNote := notes.Note{}
		if err := easyjson.Unmarshal(c.Body(), &newNote); err != nil {
			return err
		}
		newNote.Id = note.Id

		notesMap[id] = newNote
		response, err := easyjson.Marshal(newNote)
		if err != nil {
			return err
		}
		if err := c.SendString(string(response)); err != nil {
			return err
		}
		return nil
	})

	app.Delete("/notes/:id", func(c fiber.Ctx) error {
		id := c.Params("id")
		_, ok := notesMap[id]
		if !ok {
			if err := c.Status(http.StatusNotFound).SendString("note not found"); err != nil {
				return err
			}
			return nil
		}
		delete(notesMap, id)
		if err := c.SendStatus(http.StatusNoContent); err != nil {
			return err
		}
		return nil
	})

	app.Listen(":6969")
}
