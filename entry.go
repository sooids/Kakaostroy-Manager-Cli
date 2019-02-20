package main

import (
	"github.com/git-soo/Kakaostory-Manager-Cli/core"
)

func main() {
	InitAPIS()
	req := NewKakaoRequest()
	f, err := os.Open(`session_test.json`)
	if err != nil {
		log.Fatal(err)
		return
	}

	req.SetSession(ImportSession(f))

	// profile := req.GetProfile("lproject")
	// feeds := req.GetFeeds()
	// if feeds == nil {
	// 	log.Fatalf("nill return")
	// 	return
	// }
	// tree := treeprint.New()
	// for _, feed := range feeds.Feeds {
	// 	branch_str := fmt.Sprintf("[%s] %s - %s", feed.Actor.DisplayName, feed.Content, feed.CreatedAt.String())
	// 	branch := tree.AddBranch(branch_str)
	// 	for _, comment := range feed.Comments {
	// 		branch.AddNode(fmt.Sprintf("[%s] %s - %s", comment.Writer.DisplayName, comment.Text, comment.CreatedAt.String()))
	// 	}
	// }
	
	// fmt.Printf(tree.String())
}
