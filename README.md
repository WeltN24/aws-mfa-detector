# aws-mfa-detector

Fetches all users from an aws account and returns the once which don't have activated MFA.

### Installing
        go get github.com/WeltN24/aws-mfa-detector

### Examples
        aws-mfa-detector -aws.region eu-west-1

exclude user from detection

        aws-mfa-detector -aws.region eu-west-1 -exclude user1.name,user2.name
