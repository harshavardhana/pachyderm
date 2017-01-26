package server

import (
	log "github.com/Sirupsen/logrus"
	pfsclient "github.com/pachyderm/pachyderm/src/client/pfs"
	"github.com/pachyderm/pachyderm/src/server/pfs/drive"
	"github.com/pachyderm/pachyderm/src/server/pkg/metrics"
	"github.com/pachyderm/pachyderm/src/server/pkg/obj"
)

var (
	blockSize = 8 * 1024 * 1024 // 8 Megabytes
	// MaxMsgSize is used to define the GRPC frame size, which we need to be greater than a block
	MaxMsgSize = 3 * blockSize
)

// Valid object storage backends
const (
	MinioBackendEnvVar     = "MINIO"
	AmazonBackendEnvVar    = "AMAZON"
	GoogleBackendEnvVar    = "GOOGLE"
	MicrosoftBackendEnvVar = "MICROSOFT"
)

// APIServer represents and api server.
type APIServer interface {
	pfsclient.APIServer
}

// NewAPIServer creates an APIServer.
func NewAPIServer(driver drive.Driver, reporter *metrics.Reporter) APIServer {
	return newAPIServer(driver, reporter)
}

// NewLocalBlockAPIServer creates a BlockAPIServer.
func NewLocalBlockAPIServer(dir string) (pfsclient.BlockAPIServer, error) {
	return newLocalBlockAPIServer(dir)
}

// NewObjBlockAPIServer create a BlockAPIServer from an obj.Client.
func NewObjBlockAPIServer(dir string, cacheBytes int64, objClient obj.Client) (pfsclient.BlockAPIServer, error) {
	return newObjBlockAPIServer(dir, cacheBytes, objClient)
}

// NewBlockAPIServer creates a BlockAPIServer using the credentials
// it finds in the environment
func NewBlockAPIServer(dir string, cacheBytes int64, backend string) (pfsclient.BlockAPIServer, error) {
	log.Info("Initializing new blockAPIServer", dir, cacheBytes, backend)
	switch backend {
	case MinioBackendEnvVar:
		// S3 compatible doesn't like leading slashes
		if len(dir) > 0 && dir[0] == '/' {
			dir = dir[1:]
		}
		blockAPIServer, err := newMinioBlockAPIServer(dir, cacheBytes)
		if err != nil {
			return nil, err
		}
		return blockAPIServer, nil
	case AmazonBackendEnvVar:
		// amazon doesn't like leading slashes
		if len(dir) > 0 && dir[0] == '/' {
			dir = dir[1:]
		}
		blockAPIServer, err := newAmazonBlockAPIServer(dir, cacheBytes)
		if err != nil {
			return nil, err
		}
		return blockAPIServer, nil
	case GoogleBackendEnvVar:
		// TODO figure out if google likes leading slashses
		blockAPIServer, err := newGoogleBlockAPIServer(dir, cacheBytes)
		if err != nil {
			return nil, err
		}
		return blockAPIServer, nil
	case MicrosoftBackendEnvVar:
		blockAPIServer, err := newMicrosoftBlockAPIServer(dir, cacheBytes)
		if err != nil {
			return nil, err
		}
		return blockAPIServer, nil
	default:
		return NewLocalBlockAPIServer(dir)
	}
}
