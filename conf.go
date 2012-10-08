package main

import (
	"github.com/couchbaselabs/cbfs/config"
)

const configKey = "/@globalConfig"

// Update this config within a bucket.
func StoreConfig(conf cbfsconfig.CBFSConfig) error {
	return couchbase.Set(configKey, 0, &conf)
}

// Update this config from the db.
func RetrieveConfig() (*cbfsconfig.CBFSConfig, error) {
	conf := &cbfsconfig.CBFSConfig{}
	err := couchbase.Get(configKey, conf)
	return conf, err
}