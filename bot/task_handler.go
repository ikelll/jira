package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (b *MyBot) GetAssignedIssues(token string, username string) string {
	baseURL := "https://test-jira.gorizont-vs.ru"
	client := &http.Client{}

	apiEndpoint := fmt.Sprintf("/rest/api/2/search?jql=assignee=%s", username)

	req, err := http.NewRequest("GET", baseURL+apiEndpoint, nil)
	if err != nil {
		return "Ошибка при создании запроса: " + err.Error()
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return "Ошибка при отправке запроса: " + err.Error()
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var searchData map[string]interface{}
		err := json.NewDecoder(resp.Body).Decode(&searchData)
		if err != nil {
			return "Ошибка при чтении ответа: " + err.Error()
		}

		issues := searchData["issues"].([]interface{})
		if len(issues) == 0 {
			return "Задачи для пользователя " + username + " не найдены."
		}

		var tasksInfo strings.Builder
		tasksInfo.WriteString("Задачи для пользователя " + username + ":\n")

		for _, issue := range issues {
			issueData := issue.(map[string]interface{})
			issueKey := issueData["key"].(string)
			issueSummary := issueData["fields"].(map[string]interface{})["summary"].(string)
			tasksInfo.WriteString(fmt.Sprintf("- %s: %s\n", issueKey, issueSummary))
		}

		return tasksInfo.String()
	} else {
		return "Не удалось получить список задач. Статус ответа: " + resp.Status
	}
}
