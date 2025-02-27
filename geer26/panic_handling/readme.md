# Panic recover example

Lets say se have an API the send a 10 characters long random alphanumerical string upon a GET request.

We intended to save theese strings into a file, but we stating that any string that contains a specific character [eg. "A"], is invalid, and we want to save the collected strings before the invalid string response.

 - start api: sudo docker compose up
 - start collect: cd client && go run .

 The saved string can be found in client/results.txt!

 At some point there can be two case of panic:
 - We get only valid strings, and the buffer fills up. In this case save the buffer, and exit the application
 - We get invalid string prematurelly. In this case do the same.

## !!! The algorythm is for practicing panic recovering, do not use it in production !!!