# crowdsvt
This repo has Go code with the following folders:
- externalapi contains the lambda function that implements the request to obtain a signedurl for S3
- testclient contains the client code to request a signed URL through an API gateway and the send an HTTPS PUT request to S3 to upload a file
