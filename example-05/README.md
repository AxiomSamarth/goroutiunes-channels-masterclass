# Problem statement

Write a Go program to handle the jobs with a worker pool. A worker pool is a limited number of workers that process the jobs in the queue. The worker pool must close itself once all the jobs in the job queue have been processed. 

The worker count in the worker pool and the jobs should be configurable. 

## Example
Imagine there are 10 jobs in the job queue and 3 workers in the worker pool. The 3 workers pick one job after the other from the job queue until all the jobs are processes.

## Significance
There could be scenarios where a number of goroutines could be spawned to perform a certain operations over a set of inputs. Spawning too many goroutines might not be a good idea and it would create a number of processes to be handled by the Go compiler as well as the operating system. Hence, having a check on the number of goroutines is a good idea. When there is indepedent processing of input data and resource constraint for concurrency, designing a worker pool is efficient method to manage resources as well as achieving computation speed. 