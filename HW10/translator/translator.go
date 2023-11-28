package translator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TranslationRequest struct {
	Text     string `json:"q"`
	FromLang string `json:"source"`
	ToLang   string `json:"target"`
}

type TranslationResponce struct {
	Error          string
	TranslatedText string
}

type Translation struct {
	FromLang   string
	ToLang     string
	InputText  string
	OutputText string
}

type TranslatorClient struct {
}

func (client TranslatorClient) Translate(fromLang, toLang, text string) (*Translation, error) {

	params := TranslationRequest{
		FromLang: fromLang,
		ToLang:   toLang,
		Text:     text,
	}

	jsonBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPost, "https://libretranslate.de/translate", bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed, status code %d", resp.StatusCode)
	}

	var translationResp TranslationResponce
	if err := json.NewDecoder(resp.Body).Decode(&translationResp); err != nil {
		return nil, err
	}

	if translationResp.Error != "" {
		return nil, fmt.Errorf("unable to translate, error: %s", translationResp.Error)
	}

	translation := getTranslationFromTranslationResponce(translationResp, fromLang, toLang, text)
	return &translation, nil
}

func getTranslationFromTranslationResponce(resp TranslationResponce, fromLang, toLang, text string) Translation {
	return Translation{
		FromLang:   fromLang,
		ToLang:     toLang,
		InputText:  text,
		OutputText: resp.TranslatedText,
	}
}
