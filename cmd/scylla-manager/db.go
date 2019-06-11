// Copyright (C) 2017 ScyllaDB

package main

import (
	"bytes"
	"context"
	"text/template"
	"time"

	"github.com/gocql/gocql"
	"github.com/pkg/errors"
	"github.com/scylladb/go-log"
	"github.com/scylladb/gocqlx/migrate"
	"github.com/scylladb/mermaid/schema/cql"
)

func waitForDatabase(ctx context.Context, config *serverConfig, logger log.Logger) error {
	const (
		wait        = 5 * time.Second
		maxAttempts = 60
	)

	for i := 0; i < maxAttempts; i++ {
		if err := tryConnectToDatabase(config); err != nil {
			logger.Info(ctx, "Could not connect to database",
				"sleep", wait,
				"error", err,
			)
			time.Sleep(wait)
		} else {
			return nil
		}
	}

	return errors.New("could not connect to database, max attempts reached")
}

func tryConnectToDatabase(config *serverConfig) error {
	session, err := gocqlClusterConfigForDBInit(config).CreateSession()
	if session != nil {
		session.Close()
	}
	return err
}

func keyspaceExists(config *serverConfig) (bool, error) {
	session, err := gocqlClusterConfigForDBInit(config).CreateSession()
	if err != nil {
		return false, err
	}
	defer session.Close()

	var cnt int
	q := session.Query("SELECT COUNT(keyspace_name) FROM system_schema.keyspaces WHERE keyspace_name = ?").Bind(config.Database.Keyspace)
	return cnt == 1, q.Scan(&cnt)
}

func createKeyspace(config *serverConfig) error {
	session, err := gocqlClusterConfigForDBInit(config).CreateSession()
	if err != nil {
		return err
	}
	defer session.Close()

	// Auto upgrade replication factor if needed. RF=1 with multiple hosts means
	// data loss when one of the nodes is down. This is understood with a single
	// node deployment but must be avoided if we have more nodes.
	if config.Database.ReplicationFactor == 1 {
		var peers int
		q := session.Query("SELECT COUNT(*) FROM system.peers")
		if err := q.Scan(&peers); err != nil {
			return err
		}
		if peers > 0 {
			rf := peers + 1
			if rf > 3 {
				rf = 3
			}
			config.Database.ReplicationFactor = rf
		}
	}

	return session.Query(mustEvaluateCreateKeyspaceStmt(config)).Exec()
}

const createKeyspaceStmt = "CREATE KEYSPACE {{.Keyspace}} WITH replication = {'class': 'SimpleStrategy', 'replication_factor': {{.ReplicationFactor}}}"

func mustEvaluateCreateKeyspaceStmt(config *serverConfig) string {
	t := template.New("")
	if _, err := t.Parse(string(createKeyspaceStmt)); err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	if err := t.Execute(buf, config.Database); err != nil {
		panic(err)
	}

	return buf.String()
}

func migrateSchema(config *serverConfig, logger log.Logger) error {
	c := gocqlClusterConfigForDBInit(config)
	c.Keyspace = config.Database.Keyspace

	session, err := c.CreateSession()
	if err != nil {
		return err
	}
	defer session.Close()

	cql.Logger = logger
	migrate.Callback = cql.MigrateCallback
	return migrate.Migrate(context.Background(), session, config.Database.MigrateDir)
}

func gocqlClusterConfigForDBInit(config *serverConfig) *gocql.ClusterConfig {
	c := gocqlClusterConfig(config)
	c.Keyspace = "system"
	c.Timeout = config.Database.MigrateTimeout
	c.MaxWaitSchemaAgreement = config.Database.MigrateMaxWaitSchemaAgreement
	return c
}

func gocqlClusterConfig(config *serverConfig) *gocql.ClusterConfig {
	c := gocql.NewCluster(config.Database.Hosts...)

	// Chose consistency level, for a single node deployments use ONE, for
	// multi-dc deployments use LOCAL_QUORUM, otherwise use QUORUM.
	switch {
	case config.Database.LocalDC != "":
		c.Consistency = gocql.LocalQuorum
	case config.Database.ReplicationFactor == 1:
		c.Consistency = gocql.One
	default:
		c.Consistency = gocql.Quorum
	}

	c.Keyspace = config.Database.Keyspace
	c.Timeout = config.Database.Timeout

	// SSL
	if config.Database.SSL {
		c.SslOpts = &gocql.SslOptions{
			CaPath:                 config.SSL.CertFile,
			CertPath:               config.SSL.UserCertFile,
			KeyPath:                config.SSL.UserKeyFile,
			EnableHostVerification: config.SSL.Validate,
		}
	}

	// Authentication
	if config.Database.User != "" {
		c.Authenticator = gocql.PasswordAuthenticator{
			Username: config.Database.User,
			Password: config.Database.Password,
		}
	}

	// Enable token aware host selection policy
	fallback := gocql.RoundRobinHostPolicy()
	if config.Database.LocalDC != "" {
		fallback = gocql.DCAwareRoundRobinPolicy(config.Database.LocalDC)
	}
	c.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(fallback)

	return c
}
