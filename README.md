# Distance

A very simple Go-Kit based microservice to determine string similarity.

## Example

The service should currently be up at http://dockerer-152911.nitrousapp.com:8080.

It expects a POST request to `/` with the following as a JSON encoded body:

- **source**: The first string to compare.
- **target**: The seconds string to compare.

And will return the following on a successful comparison:

- **distance**: An integer that represents the similarity of the two strings - the lower it is, the more similar the strings are.

### Example request

```bash
⟩ echo '{ "source": "This is a sentence", "target": "This is a similar sentence" }' | http http://dockerer-152911.nitrousapp.com:8080
HTTP/1.1 200 OK
Connection: keep-alive
Content-Length: 15
Content-Type: text/plain; charset=utf-8
Date: Sun, 13 Mar 2016 17:17:23 GMT
Keep-Alive: timeout=30
Server: openresty/1.7.10.1

{"distance":8}
```

## Algorithms

Current algorithms implemented:

- [Levenshtein distance](http://dockerer-152911.nitrousapp.com:8080)

Algorithms that might be worth implementing:

- Longest common substring
- Longest common subsequence
- [Naive Bayes filtering](https://en.wikipedia.org/wiki/Naive_Bayes_spam_filtering)
