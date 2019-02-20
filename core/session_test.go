package core

import (
	"fmt"
	"log"
	"os"

	"github.com/xlab/treeprint"
)

// func TestImportSession(t *testing.T) {
// 	data := ImportSession("/Users/soo/go/src/github.com/soo/kakaostory-manager-cli/testdata/session_test.json")
// 	if data == nil {
// 		fmt.Println("nil returned")
// 		return
// 	}
// 	fmt.Printf("Result: %+v\n", data)
// 	return
// 	// Output: .
// }

func ExampleEE() {
	InitAPIS()
	req := NewKakaoRequest()
	f, err := os.Open(`C:\Users\Dev\go\src\github.com\git-soo\Kakaostory-Manager-Cli\testdata\session_test.json`)
	if err != nil {
		log.Fatal(err)
		return
	}

	req.SetSession(ImportSession(f))
	// profile := req.GetProfile("lproject")
	feeds := req.GetFeeds()
	if feeds == nil {
		log.Fatalf("nill return")
		return
	}
	tree := treeprint.New()
	for _, feed := range feeds.Feeds {
		branch_str := fmt.Sprintf("[%s] %s - %s", feed.Actor.DisplayName, feed.Content, feed.CreatedAt.String())
		branch := tree.AddBranch(branch_str)
		for _, comment := range feed.Comments {
			branch.AddNode(fmt.Sprintf("[%s] %s - %s", comment.Writer.DisplayName, comment.Text, comment.CreatedAt.String()))
		}
	}
	
	fmt.Printf(tree.String())
	// output:.
}
