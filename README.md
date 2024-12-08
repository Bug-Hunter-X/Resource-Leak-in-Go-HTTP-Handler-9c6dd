# Go HTTP Handler Resource Leak

This repository demonstrates a common resource leak in Go HTTP handlers that can occur when processing POST requests. The issue arises from the improper handling of the `http.Request.Body` which needs to be closed explicitly using `defer` to prevent resource exhaustion.

## Bug Description

The `handleRequest` function in `bug.go` fails to properly close the request body in all cases. If an error happens during the body's processing, the connection stays open and resources leak.

## Solution

The `bugSolution.go` file demonstrates the corrected version, ensuring `r.Body.Close()` is executed even if errors occur during request processing using `defer`.

## How to Reproduce

1. Clone the repository.
2. Run the `bug.go` file, send POST requests to the server with a large body.
3. Observe the resource consumption. Repeat this multiple times to see the cumulative effect of the leak.
4. Repeat with the fixed version in `bugSolution.go` to see the difference. Note: You'll need to modify the code to make it functional, as the examples are simplified for demonstration purposes.

This example highlights the importance of carefully handling resources in Go HTTP handlers to prevent resource exhaustion and ensure application stability.