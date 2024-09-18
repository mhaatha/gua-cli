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
	var repoName, branchOrTag, branchOrTagName string

	if payload, ok := rawData["payload"].(map[string]interface{}); ok {
		// If payload.ref_type is repository then payload.ref is nil,
		// but if payload.ref_type is branch or tag, payload.ref is a name of branch or tag itself.
		if payload["ref_type"].(string) == "repository" {

			if repo, ok := rawData["repo"].(map[string]interface{}); ok {
				repoName = repo["name"].(string)
				fmt.Printf("- Created a new repository, named %s\n", repoName)
			}

		} else {
			branchOrTag = payload["ref"].(string)
			branchOrTagName = payload["ref_type"].(string)
			fmt.Printf("- Created a new %s, named %s\n", branchOrTag, branchOrTagName)
		}
	}
}

func deleteEvent(rawData map[string]interface{}) {
	var repoName, branchOrTag, branchOrTagName string

	if payload, ok := rawData["payload"].(map[string]interface{}); ok {
		branchOrTag = payload["ref_type"].(string)
		branchOrTagName = payload["ref"].(string)

		if repo, ok := rawData["repo"].(map[string]interface{}); ok {
			repoName = repo["name"].(string)
		}

	}

	fmt.Printf("- Deleted a %s, named %s from %s repository\n", branchOrTag, branchOrTagName, repoName)
}

func forkEvent(rawData map[string]interface{}) {
	var forkedRepoName string

	if payload, ok := rawData["payload"].(map[string]interface{}); ok {
		if forkee, ok := payload["forkee"].(map[string]interface{}); ok {
			forkedRepoName = forkee["name"].(string)
		}
	}

	fmt.Printf("- Forked %s\n", forkedRepoName)
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
