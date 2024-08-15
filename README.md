# Timeout Wrapper

```go get github.com/revaldimijaya/timeout-wrapper@latest```

# Implementation
Medium : [Wrapper Function Handling Timeout](https://medium.com/@revaldimijaya.rev/how-to-set-a-max-timeout-for-your-function-in-go-0cca08090e27)

# Notes
1. This approach is based on my own research, so results may vary 🐛.
2. Ensure that your function returns no more than 2 values. If there are more than 2 values, wrap them in a struct before using this method.
