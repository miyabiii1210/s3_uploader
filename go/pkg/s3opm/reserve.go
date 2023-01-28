package s3opm

import "github.com/aws/aws-sdk-go/aws/session"

func GenerateSession() *session.Session {
	s := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           "default",
		SharedConfigState: session.SharedConfigEnable,
	}))
	return s
}
