package main
import (
	"log"
	"os"
	_ "github.com/shivani-1521/RSS-feeds-Matcher/matchers"
	"github.com/shivani-1521/RSS-feeds-Matcher/search"
)


func init(){
	log.SetOutput(os.Stdout)
}

func main(){
	search.Run("Donald")
}