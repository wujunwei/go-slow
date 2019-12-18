# go-slow
a reliable rate limiter which achieved by token bucket and slide windows

## todo
* add debug print
* if the required permits is large than stored ,it should wait until the stored permits is enough or return false 
* cli tool
* status api