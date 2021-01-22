# Proxy-Range-Filter
Filters proxies in a txt file based on the numbers before the first dot (e.g. 123.xxx.xxx.xxx will remove all proxies that start with 123)

I noticed that proxies from a specific provider with of a certain range were constantly failing, so I made this script to filter out the failing proxies.

This is my first Go project. It takes in an input for the file name (which should be placed in the same directory as the go program) and the range to delete.
The output is a text file named "test.txt" with the filtered output. There is an example "test.txt" attached that you may use to test the program.
