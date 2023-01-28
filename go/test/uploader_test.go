package test

import (
	"context"
	"os"
	"testing"

	"github.com/miyabiii1210/s3_uploader/go/pkg/s3opm"
	"github.com/miyabiii1210/s3_uploader/go/pkg/util"
)

func TestS3Uploader(t *testing.T) {
	type args struct {
		ctx                                        context.Context
		s3BucketName, s3ObjectPath, uploadFilePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "s3 uploader test",
			args: args{
				ctx:            context.TODO(),
				s3BucketName:   "",
				s3ObjectPath:   "media/sample_02_.jpg",
				uploadFilePath: "../../media/sample_02_.jpg",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			/* s3に接続するためのセッション生成 */
			session := s3opm.GenerateSession()

			/* s3へアップロードするファイルを指定 */
			file, err := os.Open(tt.args.uploadFilePath)
			if err != nil {
				t.Errorf("os.Open error: %v\n", err)
				return
			}
			defer file.Close()

			/* s3へのアップロード処理を実行 */
			mng, err := s3opm.S3UpLoader(session, tt.args.s3BucketName, tt.args.s3ObjectPath, file)
			if err != nil {
				t.Errorf("Upload error: %v\n", err)
				return
			}
			t.Log("upload success.", mng)

			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}
