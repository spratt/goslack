# goslack
A simple tool for scripting slack actions from the command line

Example of usage:

```bash
TOKEN=xoxp-12345 CHANNEL=C012345 go run goslack.go -text "Your message goes here"
```

Make sure you replace xoxp-12345 with your actual authorization token,
and C012345 with your actual channel ID.
