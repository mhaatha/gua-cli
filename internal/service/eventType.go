package service

import (
	"fmt"
	"log"
)

func commitCommentEvent(rawData map[string]interface{}) {
	var repoName string

	if repo, ok := rawData["repo"].(map[string]interface{}); ok {
		repoName = repo["name"].(string)
	} else {
		log.Println("Error: Cannot fetch repository data.")
		return
	}

	fmt.Printf("- Created commit comment on %s\n", repoName)
}

func createEvent(rawData map[string]interface{}) {
	var repoName, branchOrTagName string

	if payload, ok := rawData["payload"].(map[string]interface{}); ok {
		// If payload.ref_type is repository then payload.ref is nil,
		// but if payload.ref_type is branch or tag, payload.ref is a name of branch or tag itself.
		if payload["ref_type"].(string) == "repository" {

			if repo, ok := rawData["repo"].(map[string]interface{}); ok {
				repoName = repo["name"].(string)
				fmt.Printf("- Created a new repository, named %s\n", repoName)
			}

		} else {
			branchOrTagName = payload["ref"].(string)
			fmt.Printf("- Created a new %s, named %s\n", payload["ref_type"].(string), branchOrTagName)
		}
	}
}

func watchEvent(rawData map[string]interface{}) {
	var repoName string

	if repo, ok := rawData["repo"].(map[string]interface{}); ok {
		repoName = repo["name"].(string)
	} else {
		log.Println("Error: Cannot fetch repository data.")
		return
	}

	fmt.Printf("- Starred %s\n", repoName)
}
