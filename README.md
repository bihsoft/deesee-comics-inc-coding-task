# deesee-comics-inc-coding-task

To run the app simply run the go run ./ -port=8090 or go run main.go -port=8090 command
If no port argument is specified it will default to 8080

To run the tests simply run the go test task/deeSeeComics/tests command

Used an in memory database approach to expose the entities to the service

On startup the Migrate function gets called and loads the superheroes into the in memory db

Implementd an simple mapper to map dto to view objects, mapper uses reflection to dynamicly map fields
and their json tags. 

Implemented simple test to test key/string pairs against the DeeSeeChiffre algorithm.

The algorithm is keept verry simple, get an rune and add the key value, if the unicode character is larger then
english alphabet letters count we reverse the encryption. We then return each character from the input value.

