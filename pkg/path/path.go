package path

import (
	"os"
	"path/filepath"
)

// kafka
var (
	KfkHome             = os.Getenv("KAFKA_HOME")
	KfkConfigDir        = filepath.FromSlash(KfkHome + "/config")
	KfkConfig           = filepath.FromSlash(KfkConfigDir + "/server.properties")
	KfkOriginalConfig   = filepath.FromSlash(KfkConfigDir + "/server_original.properties")
	KRaftConfigDir      = filepath.FromSlash(KfkConfigDir + "/kraft")
	KRaftConfig         = filepath.FromSlash(KRaftConfigDir + "/server.properties")
	KRaftOriginalConfig = filepath.FromSlash(KRaftConfigDir + "/server_original.properties")
)

// mate
var (
	KfkMatePath                  = filepath.FromSlash(KfkHome + "/mate")
	KfkScripts                   = filepath.FromSlash(KfkMatePath + "/scripts")
	KfkStartScript               = filepath.FromSlash(KfkScripts + "/start-kafka.sh")
	KfkStartRaftScript           = filepath.FromSlash(KfkScripts + "/start-kafka-raft.sh")
	KfkStartStandaloneScript     = filepath.FromSlash(KfkScripts + "/start-kafka-standalone.sh")
	KfkStartRaftStandaloneScript = filepath.FromSlash(KfkScripts + "/start-kafka-raft-standalone.sh")
)
