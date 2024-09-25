package service

import (
	"fmt"
	"log"
	"strings"
)

func commitCommentEvent(rawData map[string]interface{}) {
	var repoName string

	if repo, ok := rawData["repo"].(map[string]interface{}); ok {
		if repoName, ok = repo["name"].(string); !ok {
			log.Println("Error: Cannot fetch repository data.")
			return
		}
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
		if refType, ok := payload["ref_type"].(string); refType == "repository" && ok {

			if repo, ok := rawData["repo"].(map[string]interface{}); ok {
				if repoName, ok = repo["name"].(string); ok {
					fmt.Printf("- Created a new repository, named %s\n", repoName)
				} else {
					log.Println("Error: Cannot fetch repository data.")
					return
				}
			} else {
				log.Println("Error: Cannot fetch repository data.")
				return
			}

		} else if !ok {
			log.Println("Error: Cannot fetch repository data.")
			return
		} else {
			branchOrTag = payload["ref"].(string)
			branchOrTagName = payload["ref_type"].(string)
			fmt.Printf("- Created a new %s, named %s\n", branchOrTag, branchOrTagName)
		}
	} else {
		log.Println("Error: Cannot fetch repository data.")
		return
	}
}

func deleteEvent(rawData map[string]interface{}) {
	var repoName, branchOrTag, branchOrTagName string

	if payload, ok := rawData["payload"].(map[string]interface{}); ok {
		if branchOrTag, ok = payload["ref_type"].(string); !ok {
			log.Println("Error: Cannot fetch repository data.")
			return
		}
		if branchOrTagName, ok = payload["ref"].(string); !ok {
			log.Println("Error: Cannot fetch repository data.")
			return
		}

		if repo, ok := rawData["repo"].(map[string]interface{}); ok {
			if repoName, ok = repo["name"].(string); !ok {
				log.Println("Error: Cannot fetch repository data.")
				return
			}
		} else {
			log.Println("Error: Cannot fetch repository data.")
			return
		}

	} else {
		log.Println("Error: Cannot fetch repository data.")
		return
	}

	fmt.Printf("- Deleted a %s, named %s from %s repository\n", branchOrTag, branchOrTagName, repoName)
}

func forkEvent(rawData map[string]interface{}) {
	var forkedRepoName string

	if payload, ok := rawData["payload"].(map[string]interface{}); ok {
		if forkee, ok := payload["forkee"].(map[string]interface{}); ok {
			if forkedRepoName, ok = forkee["name"].(string); !ok {
				log.Println("Error: Cannot fetch repository data.")
				return
			}
		} else {
			log.Println("Error: Cannot fetch repository data.")
			return
		}
	} else {
		log.Println("Error: Cannot fetch repository data.")
		return
	}

	fmt.Printf("- Forked %s\n", forkedRepoName)
}

func gollumEvent(rawData map[string]interface{}) {
	var pageName, title, action string

	if payload, ok := rawData["payload"].(map[string]interface{}); ok {

		if pages, ok := payload["pages"].([]interface{}); ok {
			for _, page := range pages {
				if pageData, ok := page.(map[string]interface{}); ok {
					pageName = pageData["page_name"].(string)
					title = pageData["title"].(string)
					action = pageData["title"].(string)
				} else {
					log.Println("Error: Cannot fetch repository data.")
					return
				}
			}
		} else {
			log.Println("Error: Cannot fetch repository data.")
			return
		}

	} else {
		log.Println("Error: Cannot fetch repository data.")
		return
	}

	fmt.Printf("%s a wiki named %s in %s page\n", strings.ToUpper(string(action[0]))+strings.ToLower(action[1:]), title, pageName)
}

func issueCommentEvent(rawData map[string]interface{}) {
	var actionType string

	if payload, ok := rawData["payload"].(map[string]interface{}); ok {
		if actionType, ok = payload["action"].(string); !ok {
			log.Println("Error: Cannot fetch repository data.")
			return
		}
	} else {
		log.Println("Error: Cannot fetch repository data.")
		return
	}

	fmt.Printf("- %s an issue comment\n", strings.ToUpper(string(actionType[0]))+strings.ToLower(actionType[1:]))
}

func issusesEvent(rawData map[string]interface{}) {
	var actionType string

	if payload, ok := rawData["payload"].(map[string]interface{}); ok {
		if actionType, ok = payload["action"].(string); !ok {
			log.Println("Error: Cannot fetch repository data.")
			return
		}
	} else {
		log.Println("Error: Cannot fetch repository data.")
		return
	}

	fmt.Printf("- %s an issue\n", strings.ToUpper(string(actionType[0]))+strings.ToLower(actionType[1:]))
}

func memberEvent(rawData map[string]interface{}) {
	var actionType string

	if payload, ok := rawData["payload"].(map[string]interface{}); ok {
		if actionType, ok = payload["action"].(string); actionType == "added" && ok {
			if member, ok := payload["member"].(map[string]interface{}); ok {
				if memberName, ok := member["login"].(string); ok {
					fmt.Printf("- Added %s as a collaborator\n", memberName)
				} else {
					log.Println("Error: Cannot fetch repository data.")
					return
				}
			} else {
				log.Println("Error: Cannot fetch repository data.")
				return
			}
		} else if actionType == "edited" && ok {
			if changes, ok := payload["changes"].(map[string]interface{}); ok {
				if member, ok := payload["member"].(map[string]interface{}); ok {
					if memberName, ok := member["login"].(string); ok {
						if newPermission, ok := changes["role"].(map[string]interface{}); ok {
							if newPermissionValue, ok := newPermission["new_value"].(string); ok {
								fmt.Printf("- Changed %s's permission to %s\n", memberName, newPermissionValue)
							} else {
								log.Println("Error: Cannot fetch repository data.")
								return
							}
						} else {
							log.Println("Error: Cannot fetch repository data.")
							return
						}
					} else {
						log.Println("Error: Cannot fetch repository data.")
						return
					}
				} else {
					log.Println("Error: Cannot fetch repository data.")
					return
				}
			} else {
				log.Println("Error: Cannot fetch repository data.")
				return
			}
		} else {
			log.Println("Error: Cannot fetch repository data.")
			return
		}
	} else {
		log.Println("Error: Cannot fetch repository data.")
		return
	}
}

func publicEvent() {
	fmt.Printf("- Published a private repository\n")
}

func pullRequestEvent(rawData map[string]interface{}) {
	var actionType, repoName string

	if repo, ok := rawData["repo"].(map[string]interface{}); ok {
		if repoName, ok = repo["name"].(string); !ok {
			log.Println("Error: Cannot fetch repository data.")
			return
		}
	}
	if payload, ok := rawData["payload"].(map[string]interface{}); ok {
		if actionType, ok = payload["action"].(string); !ok {
			log.Println("Error: Cannot fetch repository data.")
			return
		}
	} else {
		log.Println("Error: Cannot fetch repository data.")
		return
	}

	fmt.Printf("- %s a pull request in %s\n", strings.ToUpper(string(actionType[0]))+strings.ToLower(actionType[1:]), repoName)
}

func pullRequestReviewEvent(rawData map[string]interface{}) {
	var actionType string

	if payload, ok := rawData["payload"].(map[string]interface{}); ok {
		if actionType, ok = payload["action"].(string); !ok {
			log.Println("Error: Cannot fetch repository data.")
			return
		}
	} else {
		log.Println("Error: Cannot fetch repository data.")
		return
	}

	fmt.Printf("- %s a PR review\n", strings.ToUpper(string(actionType[0]))+strings.ToLower(actionType[1:]))
}

func pullRequestReviewCommentEvent(rawData map[string]interface{}) {
	var actionType string

	if payload, ok := rawData["payload"].(map[string]interface{}); ok {
		if actionType, ok = payload["action"].(string); !ok {
			log.Println("Error: Cannot fetch repository data.")
			return
		}
	} else {
		log.Println("Error: Cannot fetch repository data.")
		return
	}

	fmt.Printf("- %s a PR review comment\n", strings.ToUpper(string(actionType[0]))+strings.ToLower(actionType[1:]))
}

func pullRequestReviewThreadEvent(rawData map[string]interface{}) {
	var actionType string

	if payload, ok := rawData["payload"].(map[string]interface{}); ok {
		if actionType, ok = payload["action"].(string); !ok {
			log.Println("Error: Cannot fetch repository data.")
			return
		}
	} else {
		log.Println("Error: Cannot fetch repository data.")
		return
	}

	fmt.Printf("- Marked a PR comment thread to %s\n", actionType)
}

func pushEvent(rawData map[string]interface{}) {
	var size float64
	var repoName string

	// Fetch size of commit data
	if payload, ok := rawData["payload"].(map[string]interface{}); ok {
		if size, ok = payload["size"].(float64); !ok {
			log.Println("Error: Cannot fetch repository data.")
			return
		}
	} else {
		log.Println("Error: Cannot fetch repository data.")
		return
	}

	// Fetch repository name
	if repo, ok := rawData["repo"].(map[string]interface{}); ok {
		if repoName, ok = repo["name"].(string); !ok {
			log.Println("Error: Cannot fetch repository data.")
			return
		}
	} else {
		log.Println("Error: Cannot fetch repository data.")
		return
	}

	fmt.Printf("- Pushed %v commit(s) to %s\n", size, repoName)
}

func watchEvent(rawData map[string]interface{}) {
	var repoName string

	if repo, ok := rawData["repo"].(map[string]interface{}); ok {
		if repoName, ok = repo["name"].(string); !ok {
			log.Println("Error: Cannot fetch repository data.")
			return
		}
	} else {
		log.Println("Error: Cannot fetch repository data.")
		return
	}

	fmt.Printf("- Starred %s\n", repoName)
}
