package AnekdotProviders

import (
	"context"
	"fmt"
	"io"
	"newbot/internal/domain/ForGeneration"
)

func (d *Domain) GetJoke(ctx context.Context, category string) (string, error) {
	joke, err := d.client.Get(ForGeneration.GenerateURL(category))
	if err != nil {
		fmt.Println("ОШИБКА ЗАПРОСА GET: ", err)
	}
	body, err := io.ReadAll(joke.Body)
	if err != nil {
		fmt.Println("ОШЫБКА ЧТЕНИЯ ГЕТ-ЗАПРОСА: ", err)
	}
	return string(body), err
}
