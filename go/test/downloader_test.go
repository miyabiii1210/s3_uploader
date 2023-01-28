package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/miyabiii1210/s3_uploader/go/pkg/s3opm"
	"github.com/miyabiii1210/s3_uploader/go/pkg/util"
)

func TestS3Downloader(t *testing.T) {
	type args struct {
		ctx                                       context.Context
		s3BucketName, s3ObjectPath, localFilePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "s3 downloader test",
			args: args{
				ctx:           context.TODO(),
				s3BucketName:  "",
				s3ObjectPath:  "media/sample_01_.jpg",
				localFilePath: "sample_01_.jpg",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// /* s3に接続するためのセッション生成 */
			session := s3opm.GenerateSession()

			/* s3の読み込みファイルを生成 */
			file, err := os.Create(tt.args.localFilePath)
			if err != nil {
				t.Errorf("os.Create error: %v\n", err)
				return
			}

			/* ファイルの存在確認 */
			initDate, err := os.Stat(tt.args.localFilePath)
			if err != nil {
				t.Errorf("os.Stat error: %v\n", err)
				return
			}

			/* s3から指定したファイルをダウンロード */
			mng, err := s3opm.S3DownLoader(session, file, tt.args.s3BucketName, tt.args.s3ObjectPath)
			if err != nil {
				t.Errorf("S3DownLoader error: %v\n", err)
				return
			}
			util.Sleep(1)

			afterDate, err := os.Stat(tt.args.localFilePath)
			if err != nil {
				t.Errorf("os.Stat error: %v\n", err)
				return
			}

			/* ダウンロード前後でファイル容量を比較 */
			if afterDate.Size() <= initDate.Size() || afterDate.Size() != mng {
				t.Error("s3 download process may have failed.")
				return
			}

			fmt.Printf("Download process from S3 completed. DownloadedSize: %d byte\n", mng)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}
