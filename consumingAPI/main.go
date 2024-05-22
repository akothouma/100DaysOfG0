
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	req, err := http.NewRequest("GET", "https://parseapi.back4app.com/classes/Colors_Color?count=1&limit=10&keys=name,hexCode", nil)
	if err != nil {
			fmt.Println("Error accessing the database")
			os.Exit(1)
	}
	req.Header.Add("X-Parse-Application-Id", "1mMPHQ4CEx3x3S09uo1c0bXYIgolUFcZ9m3JaXzj") //adding headers
	req.Header.Add("X-Parse-REST-API-Key", "MhUbwZ3vDA91V0AlmUHx6zwq9V3Czb5YBdgOwOnS")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
			fmt.Println("Error adding headers to the request")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
			fmt.Println("Error adding headers to the request")
	}
	fmt.Println(string(body))
}